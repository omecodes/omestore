package router

import (
	"context"
	"github.com/google/uuid"
	"github.com/omecodes/common/errors"
	"github.com/omecodes/common/utils/log"
	"github.com/omecodes/store/acl"
	"github.com/omecodes/store/objects"
	"github.com/omecodes/store/pb"
)

type ExecHandler struct {
	BaseHandler
}

func (e *ExecHandler) PutObject(ctx context.Context, object *pb.Object, security *pb.PathAccessRules, opts objects.PutDataOptions) (string, error) {
	if object.Header.Id == "" {
		object.Header.Id = uuid.New().String()
	}

	accessStore := acl.GetStore(ctx)
	if accessStore == nil {
		log.Info("exec-handler.PutObject: missing access store in context")
		return "", errors.Internal
	}

	err := accessStore.SaveRules(ctx, object.Header.Id, security)
	if err != nil {
		log.Error("exec-handler.PutObject: failed to save object access security rules", log.Err(err))
		return "", errors.Internal
	}

	storage := objects.Get(ctx)
	if storage == nil {
		log.Error("exec-handler.PutObject: missing storage in context")
		if err2 := accessStore.Delete(ctx, object.Header.Id); err2 != nil {
			log.Error("exec-handler.PutObject: failed to clear access rules", log.Err(err2))
		}
		return "", errors.Internal
	}

	err = storage.Save(ctx, object, opts.Indexes...)
	if err != nil {
		if err2 := accessStore.Delete(ctx, object.Header.Id); err2 != nil {
			log.Error("exec-handler.PutObject: failed to clear access rules", log.Err(err2))
		}
		return "", err
	}

	return object.Header.Id, nil
}

func (e *ExecHandler) PatchObject(ctx context.Context, patch *pb.Patch, opts objects.PatchOptions) error {
	storage := objects.Get(ctx)
	if storage == nil {
		log.Info("missing storage in context")
		return errors.Internal
	}
	return storage.Patch(ctx, patch)
}

func (e *ExecHandler) GetObject(ctx context.Context, objectID string, opts objects.GetObjectOptions) (*pb.Object, error) {
	storage := objects.Get(ctx)
	if storage == nil {
		log.Info("missing DB in context")
		return nil, errors.Internal
	}

	if opts.Path == "" {
		return storage.Get(ctx, objectID)
	} else {
		return storage.GetAt(ctx, objectID, opts.Path)
	}
}

func (e *ExecHandler) GetObjectHeader(ctx context.Context, objectID string) (*pb.Header, error) {
	storage := objects.Get(ctx)
	if storage == nil {
		log.Info("missing DB in context")
		return nil, errors.Internal
	}
	return storage.Info(ctx, objectID)
}

func (e *ExecHandler) DeleteObject(ctx context.Context, objectID string) error {
	storage := objects.Get(ctx)
	if storage == nil {
		log.Info("exec-handler.DeleteObjet: missing DB in context")
		return errors.Internal
	}

	err := storage.Delete(ctx, objectID)
	if err != nil {
		log.Error("exec-handler.DeleteObjet: failed to delete object from storage", log.Err(err))
		return err
	}

	accessStore := acl.GetStore(ctx)
	if accessStore == nil {
		log.Info("exec-handler.DeleteObjet: missing access store in context")
		return errors.Internal
	}

	return accessStore.Delete(ctx, objectID)
}

func (e *ExecHandler) ListObjects(ctx context.Context, opts objects.ListOptions) (*pb.ObjectList, error) {
	storage := objects.Get(ctx)
	if storage == nil {
		log.Info("missing DB in context")
		return nil, errors.Internal
	}

	return storage.List(ctx, opts.Filter, opts)
}
