package files

import (
	"context"
	pb "github.com/omecodes/store/gen/go/proto"
	"io"

	"github.com/omecodes/errors"
)

type ExecHandler struct {
	BaseHandler
}

func (h *ExecHandler) CreateSource(ctx context.Context, source *pb.Access) error {
	sourceManager := getAccessManager(ctx)
	if sourceManager == nil {
		return errors.Internal("context missing source manager")
	}
	_, err := sourceManager.Save(ctx, source)
	return err
}

func (h *ExecHandler) GetAccess(ctx context.Context, sourceID string) (*pb.Access, error) {
	sourceManager := getAccessManager(ctx)
	if sourceManager == nil {
		return nil, errors.Internal("context missing source manager")
	}
	return sourceManager.Get(ctx, sourceID)
}

func (h *ExecHandler) DeleteAccess(ctx context.Context, sourceID string) error {
	sourceManager := getAccessManager(ctx)
	if sourceManager == nil {
		return errors.Internal("context missing source manager")
	}
	return sourceManager.Delete(ctx, sourceID)
}

func (h *ExecHandler) CreateDir(ctx context.Context, sourceID string, dirname string) error {
	fs, err := getFS(ctx, sourceID)
	if err != nil {
		return err
	}
	return fs.Mkdir(ctx, dirname)
}

func (h *ExecHandler) WriteFileContent(ctx context.Context, sourceID string, filename string, content io.Reader, size int64, opts WriteOptions) error {
	fs, err := getFS(ctx, sourceID)
	if err != nil {
		return err
	}
	return fs.Write(ctx, filename, content, opts.Append)
}

func (h *ExecHandler) ListDir(ctx context.Context, sourceID string, dirname string, opts ListDirOptions) (*DirContent, error) {
	fs, err := getFS(ctx, sourceID)
	if err != nil {
		return nil, err
	}
	return fs.Ls(ctx, dirname, opts.Offset, opts.Count)
}

func (h *ExecHandler) ReadFileContent(ctx context.Context, sourceID string, filename string, opts ReadOptions) (io.ReadCloser, int64, error) {
	fs, err := getFS(ctx, sourceID)
	if err != nil {
		return nil, 0, err
	}
	return fs.Read(ctx, filename, opts.Range.Offset, opts.Range.Length)
}

func (h *ExecHandler) GetFileInfo(ctx context.Context, sourceID string, filename string, opts GetFileOptions) (*pb.File, error) {
	fs, err := getFS(ctx, sourceID)
	if err != nil {
		return nil, err
	}
	return fs.Info(ctx, filename, opts.WithAttrs)
}

func (h *ExecHandler) DeleteFile(ctx context.Context, sourceID string, filename string, opts DeleteFileOptions) error {
	fs, err := getFS(ctx, sourceID)
	if err != nil {
		return err
	}
	return fs.DeleteFile(ctx, filename, opts.Recursive)
}

func (h *ExecHandler) SetFileAttributes(ctx context.Context, sourceID string, filename string, attrs Attributes) error {
	fs, err := getFS(ctx, sourceID)
	if err != nil {
		return err
	}
	return fs.SetAttributes(ctx, filename, attrs)
}

func (h *ExecHandler) GetFileAttributes(ctx context.Context, sourceID string, filename string, names ...string) (Attributes, error) {
	fs, err := getFS(ctx, sourceID)
	if err != nil {
		return nil, err
	}
	return fs.GetAttributes(ctx, filename, names...)
}

func (h *ExecHandler) RenameFile(ctx context.Context, sourceID string, filename string, newName string) error {
	fs, err := getFS(ctx, sourceID)
	if err != nil {
		return err
	}
	return fs.Rename(ctx, filename, newName)
}

func (h *ExecHandler) MoveFile(ctx context.Context, sourceID string, filename string, dirname string) error {
	fs, err := getFS(ctx, sourceID)
	if err != nil {
		return err
	}
	return fs.Rename(ctx, filename, dirname)
}

func (h *ExecHandler) CopyFile(ctx context.Context, sourceID string, filename string, dirname string) error {
	fs, err := getFS(ctx, sourceID)
	if err != nil {
		return err
	}
	return fs.Rename(ctx, filename, dirname)
}

func (h *ExecHandler) OpenMultipartSession(ctx context.Context, sourceID string, filename string, info MultipartSessionInfo) (string, error) {
	panic("implement me")
}

func (h *ExecHandler) WriteFilePart(ctx context.Context, sessionID string, content io.Reader, size int64, info ContentPartInfo) (int64, error) {
	panic("implement me")
}

func (h *ExecHandler) CloseMultipartSession(ctx context.Context, sessionId string) error {
	panic("implement me")
}
