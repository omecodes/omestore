package router

import (
	"context"
	"github.com/google/uuid"
	"github.com/omecodes/bome"
	"github.com/omecodes/common/errors"
	"github.com/omecodes/common/utils/log"
	"github.com/omecodes/omestore/oms"
	"github.com/omecodes/omestore/pb"
)

type execHandler struct {
	BaseHandler
}

func (e *execHandler) SetSettings(ctx context.Context, name string, value string, opts oms.SettingsOptions) error {
	s := Settings(ctx)
	if s == nil {
		log.Info("exec-handler.SetSettings: missing settings database in context")
		return errors.Internal
	}

	entry := &bome.MapEntry{
		Key:   name,
		Value: value,
	}
	return s.Save(entry)
}

func (e *execHandler) DeleteSettings(ctx context.Context, name string) error {
	s := Settings(ctx)
	if s == nil {
		log.Info("exec-handler.DeleteSettings: missing settings database in context")
		return errors.Internal
	}
	return s.Delete(name)
}

func (e *execHandler) ClearSettings(ctx context.Context) error {
	s := Settings(ctx)
	if s == nil {
		log.Info("exec-handler.ClearSettings: missing settings database in context")
		return errors.Internal
	}
	return s.Clear()
}

func (e *execHandler) GetSettings(ctx context.Context, name string) (string, error) {
	s := Settings(ctx)
	if s == nil {
		log.Info("exec-handler.GetSettings: missing settings database in context")
		return "", errors.Internal
	}
	return s.Get(name)
}

func (e *execHandler) PutObject(ctx context.Context, object *oms.Object, security *pb.PathAccessRules, opts oms.PutDataOptions) (string, error) {
	storage := Objects(ctx)
	if storage == nil {
		log.Info("exec-handler.PutObject: missing storage in context")
		return "", errors.Internal
	}

	accessStore := ACLStore(ctx)
	if accessStore == nil {
		log.Info("exec-handler.PutObject: missing access store in context")
		return "", errors.Internal
	}
	id := uuid.New().String()

	err := accessStore.SaveRules(ctx, id, security)
	if err != nil {
		log.Error("exec-handler.PutObject: failed to save object access security rules", log.Err(err))
		return "", errors.Internal
	}

	object.SetID(id)
	return id, storage.Save(ctx, object)
}

func (e *execHandler) PatchObject(ctx context.Context, patch *oms.Patch, opts oms.PatchOptions) error {
	storage := Objects(ctx)
	if storage == nil {
		log.Info("missing storage in context")
		return errors.Internal
	}
	return storage.Patch(ctx, patch)
}

func (e *execHandler) GetObject(ctx context.Context, objectID string, opts oms.GetObjectOptions) (*oms.Object, error) {
	storage := Objects(ctx)
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

func (e *execHandler) GetObjectHeader(ctx context.Context, objectID string) (*pb.Header, error) {
	storage := Objects(ctx)
	if storage == nil {
		log.Info("missing DB in context")
		return nil, errors.Internal
	}
	return storage.Info(ctx, objectID)
}

func (e *execHandler) DeleteObject(ctx context.Context, objectID string) error {
	storage := Objects(ctx)
	if storage == nil {
		log.Info("exec-handler.DeleteObjet: missing DB in context")
		return errors.Internal
	}

	err := storage.Delete(ctx, objectID)
	if err != nil {
		log.Error("exec-handler.DeleteObjet: failed to delete object from storage", log.Err(err))
		return err
	}

	accessStore := ACLStore(ctx)
	if accessStore == nil {
		log.Info("exec-handler.DeleteObjet: missing access store in context")
		return errors.Internal
	}

	return accessStore.Delete(ctx, objectID)
}

func (e *execHandler) ListObjects(ctx context.Context, opts oms.ListOptions) (*oms.ObjectList, error) {
	storage := Objects(ctx)
	if storage == nil {
		log.Info("missing DB in context")
		return nil, errors.Internal
	}
	return storage.List(ctx, opts.Before, opts.Count, opts.Filter)
}
