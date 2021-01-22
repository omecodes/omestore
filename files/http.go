package files

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/omecodes/common/errors"
	"github.com/omecodes/libome/logs"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const (
	pathVarId = "id"
)

func CreateFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sourceID, filename := Split(r.URL.Path)

	if r.ContentLength > 0 {
		var location *FileLocation
		err := json.NewDecoder(r.Body).Decode(&location)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		route, err := NewRoute(ctx)
		if err != nil {
			w.WriteHeader(errors.HttpStatus(err))
			return
		}

		err = route.CopyFile(ctx, sourceID, filename, strings.TrimPrefix(location.Filename, sourceID))
		if err != nil {
			w.WriteHeader(errors.HttpStatus(err))
			return
		}
	}

	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	err = route.CreateDir(ctx, sourceID, filename)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}
}

func ListDir(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sourceID, filename := Split(r.URL.Path)

	var opts ListDirOptions
	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	content, err := route.ListDir(ctx, sourceID, filename, opts)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}
	_ = json.NewEncoder(w).Encode(content)
}

func GetFileInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sourceID, filename := Split(r.URL.Path)

	var opts GetFileInfoOptions
	opts.WithAttrs = r.URL.Query().Get("attrs") == "true"

	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	info, err := route.GetFileInfo(ctx, sourceID, filename, opts)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}
	_ = json.NewEncoder(w).Encode(info)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sourceID, filename := Split(r.URL.Path)

	var opts DeleteFileOptions
	opts.Recursive = r.URL.Query().Get("recursive") == "true"

	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	err = route.DeleteFile(ctx, sourceID, filename, opts)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}
}

func PatchFileTree(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sourceID, filename := Split(r.URL.Path)

	var patchInfo *TreePatchInfo
	err := json.NewDecoder(r.Body).Decode(&patchInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	if patchInfo.Rename {
		err = route.RenameFile(ctx, sourceID, filename, strings.TrimPrefix(patchInfo.Value, sourceID))
	} else {
		err = route.MoveFile(ctx, sourceID, filename, strings.TrimPrefix(patchInfo.Value, sourceID))
	}
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}
}

func SetFileAttributes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sourceID, filename := Split(r.URL.Path)

	var attributes Attributes
	err := json.NewDecoder(r.Body).Decode(&attributes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	err = route.SetFileMetaData(ctx, sourceID, filename, attributes)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}
}

func GetFileAttributes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sourceID, filename := Split(r.URL.Path)

	name := r.URL.Query().Get("name")

	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	attributes, err := route.GetFileAttributes(ctx, sourceID, filename, name)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	_ = json.NewEncoder(w).Encode(attributes)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sourceID, filename := Split(r.URL.Path)
	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	opts := WriteOptions{}
	opts.Append = r.URL.Query().Get("append") == "true"
	opts.Hash = r.Header.Get("X-Content-Hash")

	err = route.WriteFileContent(ctx, sourceID, filename, r.Body, r.ContentLength, opts)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}
}

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sourceID, filename := Split(r.URL.Path)

	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	opts := ReadOptions{}
	opts.Range.Offset, err = Int64QueryParam(r, "offset")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	opts.Range.Length, err = Int64QueryParam(r, "length")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, size, err := route.ReadFileContent(ctx, sourceID, filename, opts)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	defer func() {
		if cer := file.Close(); cer != nil {
			logs.Error("file stream close", logs.Err(err))
		}
	}()

	if size == 0 {
		return
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", size))
	buf := make([]byte, 1024)
	done := false
	for !done {
		n, err := file.Read(buf)
		if err != nil {
			if done = err == io.EOF; !done {
				logs.Error("failed to read file content", logs.Err(err))
				return
			}
		}

		_, err = w.Write(buf[:n])
		if err != nil {
			logs.Error("failed to write file content", logs.Err(err))
			return
		}
	}
}

func CreateSource(w http.ResponseWriter, r *http.Request) {
	var source *Source
	err := json.NewDecoder(r.Body).Decode(&source)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	err = route.CreateSource(ctx, source)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}
}

func ListSources(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	sources, err := route.ListSources(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}
	_ = json.NewEncoder(w).Encode(sources)
}

func GetSource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceID := vars[pathVarId]

	ctx := r.Context()
	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	source, err := route.GetSource(ctx, sourceID)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}
	_ = json.NewEncoder(w).Encode(source)
}

func DeleteSource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sourceID := vars[pathVarId]

	ctx := r.Context()
	route, err := NewRoute(ctx)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}

	err = route.DeleteSource(ctx, sourceID)
	if err != nil {
		w.WriteHeader(errors.HttpStatus(err))
		return
	}
}

func Int64QueryParam(r *http.Request, name string) (int64, error) {
	param := r.URL.Query().Get(name)
	if param != "" {
		return strconv.ParseInt(param, 10, 64)
	} else {
		return 0, nil
	}
}
