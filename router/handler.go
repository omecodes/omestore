package router

import (
	"context"
	"github.com/omecodes/store/oms"
	"github.com/omecodes/store/pb"
)

type Handler interface {
	PutObject(ctx context.Context, object *pb.Object, security *pb.PathAccessRules, opts oms.PutDataOptions) (string, error)
	PatchObject(ctx context.Context, patch *oms.Patch, opts oms.PatchOptions) error
	GetObject(ctx context.Context, id string, opts oms.GetObjectOptions) (*pb.Object, error)
	GetObjectHeader(ctx context.Context, id string) (*pb.Header, error)
	DeleteObject(ctx context.Context, id string) error
	ListObjects(ctx context.Context, opts oms.ListOptions) (*pb.ObjectList, error)
	SearchObjects(ctx context.Context, params oms.SearchParams, opts oms.SearchOptions) (*pb.ObjectList, error)
}
