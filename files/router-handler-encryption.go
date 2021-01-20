package files

import (
	"context"
	"github.com/omecodes/errors"
	"github.com/omecodes/libome/crypt"
	"io"
)

type EncryptionHandler struct {
	BaseHandler
}

func (h *EncryptionHandler) WriteFileContent(ctx context.Context, filename string, content io.Reader, size int64, opts WriteOptions) error {
	source := GetSource(ctx)
	if source == nil {
		return errors.Create(errors.Internal, "missing source in context")
	}

	if source.Encryption == nil {
		return h.next.WriteFileContent(ctx, filename, content, size, opts)
	}

	encryptStream := crypt.NewEncryptWrapper(nil, crypt.WithBlockSize(4096))
	return h.next.WriteFileContent(ctx, filename, encryptStream.WrapReader(content), size, opts)
}

func (h *EncryptionHandler) ReadFileContent(ctx context.Context, filename string, opts ReadOptions) (io.ReadCloser, int64, error) {
	source := GetSource(ctx)
	if source == nil {
		return nil, 0, errors.Create(errors.Internal, "missing source in context")
	}

	if source.Encryption == nil {
		return h.next.ReadFileContent(ctx, filename, opts)
	}

	readCloser, size, err := h.next.ReadFileContent(ctx, filename, opts)
	if err != nil {
		return nil, 0, err
	}

	decryptStream := crypt.NewDecryptWrapper(nil, crypt.WithLimit(opts.Range.Length), crypt.WithOffset(opts.Range.Offset))
	return decryptStream.WrapReadCloser(readCloser), size, nil
}

func (h *EncryptionHandler) AddContentPart(ctx context.Context, sessionID string, content io.Reader, size int64, info *ContentPartInfo) error {
	source := GetSource(ctx)
	if source == nil {
		return errors.Create(errors.Internal, "missing source in context")
	}

	if source.Encryption == nil {
		return h.next.AddContentPart(ctx, sessionID, content, size, info)
	}

	encryptStream := crypt.NewEncryptWrapper(nil, crypt.WithBlockSize(4096))
	return h.next.AddContentPart(ctx, sessionID, encryptStream.WrapReader(content), size, info)
}
