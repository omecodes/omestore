// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// HandlerUnitClient is the client API for HandlerUnit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HandlerUnitClient interface {
	CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error)
	GetCollection(ctx context.Context, in *GetCollectionRequest, opts ...grpc.CallOption) (*GetCollectionResponse, error)
	ListCollections(ctx context.Context, in *ListCollectionsRequest, opts ...grpc.CallOption) (*ListCollectionsResponse, error)
	DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*DeleteCollectionResponse, error)
	PutObject(ctx context.Context, in *PutObjectRequest, opts ...grpc.CallOption) (*PutObjectResponse, error)
	PatchObject(ctx context.Context, in *PatchObjectRequest, opts ...grpc.CallOption) (*PatchObjectResponse, error)
	MoveObject(ctx context.Context, in *MoveObjectRequest, opts ...grpc.CallOption) (*MoveObjectResponse, error)
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error)
	DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error)
	ObjectInfo(ctx context.Context, in *ObjectInfoRequest, opts ...grpc.CallOption) (*ObjectInfoResponse, error)
	ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (HandlerUnit_ListObjectsClient, error)
	SearchObjects(ctx context.Context, in *SearchObjectsRequest, opts ...grpc.CallOption) (HandlerUnit_SearchObjectsClient, error)
}

type handlerUnitClient struct {
	cc grpc.ClientConnInterface
}

func NewHandlerUnitClient(cc grpc.ClientConnInterface) HandlerUnitClient {
	return &handlerUnitClient{cc}
}

func (c *handlerUnitClient) CreateCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CreateCollectionResponse, error) {
	out := new(CreateCollectionResponse)
	err := c.cc.Invoke(ctx, "/HandlerUnit/CreateCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerUnitClient) GetCollection(ctx context.Context, in *GetCollectionRequest, opts ...grpc.CallOption) (*GetCollectionResponse, error) {
	out := new(GetCollectionResponse)
	err := c.cc.Invoke(ctx, "/HandlerUnit/GetCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerUnitClient) ListCollections(ctx context.Context, in *ListCollectionsRequest, opts ...grpc.CallOption) (*ListCollectionsResponse, error) {
	out := new(ListCollectionsResponse)
	err := c.cc.Invoke(ctx, "/HandlerUnit/ListCollections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerUnitClient) DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*DeleteCollectionResponse, error) {
	out := new(DeleteCollectionResponse)
	err := c.cc.Invoke(ctx, "/HandlerUnit/DeleteCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerUnitClient) PutObject(ctx context.Context, in *PutObjectRequest, opts ...grpc.CallOption) (*PutObjectResponse, error) {
	out := new(PutObjectResponse)
	err := c.cc.Invoke(ctx, "/HandlerUnit/PutObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerUnitClient) PatchObject(ctx context.Context, in *PatchObjectRequest, opts ...grpc.CallOption) (*PatchObjectResponse, error) {
	out := new(PatchObjectResponse)
	err := c.cc.Invoke(ctx, "/HandlerUnit/PatchObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerUnitClient) MoveObject(ctx context.Context, in *MoveObjectRequest, opts ...grpc.CallOption) (*MoveObjectResponse, error) {
	out := new(MoveObjectResponse)
	err := c.cc.Invoke(ctx, "/HandlerUnit/MoveObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerUnitClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error) {
	out := new(GetObjectResponse)
	err := c.cc.Invoke(ctx, "/HandlerUnit/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerUnitClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error) {
	out := new(DeleteObjectResponse)
	err := c.cc.Invoke(ctx, "/HandlerUnit/DeleteObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerUnitClient) ObjectInfo(ctx context.Context, in *ObjectInfoRequest, opts ...grpc.CallOption) (*ObjectInfoResponse, error) {
	out := new(ObjectInfoResponse)
	err := c.cc.Invoke(ctx, "/HandlerUnit/ObjectInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *handlerUnitClient) ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (HandlerUnit_ListObjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HandlerUnit_serviceDesc.Streams[0], "/HandlerUnit/ListObjects", opts...)
	if err != nil {
		return nil, err
	}
	x := &handlerUnitListObjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HandlerUnit_ListObjectsClient interface {
	Recv() (*Object, error)
	grpc.ClientStream
}

type handlerUnitListObjectsClient struct {
	grpc.ClientStream
}

func (x *handlerUnitListObjectsClient) Recv() (*Object, error) {
	m := new(Object)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *handlerUnitClient) SearchObjects(ctx context.Context, in *SearchObjectsRequest, opts ...grpc.CallOption) (HandlerUnit_SearchObjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HandlerUnit_serviceDesc.Streams[1], "/HandlerUnit/SearchObjects", opts...)
	if err != nil {
		return nil, err
	}
	x := &handlerUnitSearchObjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HandlerUnit_SearchObjectsClient interface {
	Recv() (*Object, error)
	grpc.ClientStream
}

type handlerUnitSearchObjectsClient struct {
	grpc.ClientStream
}

func (x *handlerUnitSearchObjectsClient) Recv() (*Object, error) {
	m := new(Object)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HandlerUnitServer is the server API for HandlerUnit service.
// All implementations must embed UnimplementedHandlerUnitServer
// for forward compatibility
type HandlerUnitServer interface {
	CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error)
	GetCollection(context.Context, *GetCollectionRequest) (*GetCollectionResponse, error)
	ListCollections(context.Context, *ListCollectionsRequest) (*ListCollectionsResponse, error)
	DeleteCollection(context.Context, *DeleteCollectionRequest) (*DeleteCollectionResponse, error)
	PutObject(context.Context, *PutObjectRequest) (*PutObjectResponse, error)
	PatchObject(context.Context, *PatchObjectRequest) (*PatchObjectResponse, error)
	MoveObject(context.Context, *MoveObjectRequest) (*MoveObjectResponse, error)
	GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error)
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)
	ObjectInfo(context.Context, *ObjectInfoRequest) (*ObjectInfoResponse, error)
	ListObjects(*ListObjectsRequest, HandlerUnit_ListObjectsServer) error
	SearchObjects(*SearchObjectsRequest, HandlerUnit_SearchObjectsServer) error
	mustEmbedUnimplementedHandlerUnitServer()
}

// UnimplementedHandlerUnitServer must be embedded to have forward compatible implementations.
type UnimplementedHandlerUnitServer struct {
}

func (UnimplementedHandlerUnitServer) CreateCollection(context.Context, *CreateCollectionRequest) (*CreateCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (UnimplementedHandlerUnitServer) GetCollection(context.Context, *GetCollectionRequest) (*GetCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCollection not implemented")
}
func (UnimplementedHandlerUnitServer) ListCollections(context.Context, *ListCollectionsRequest) (*ListCollectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollections not implemented")
}
func (UnimplementedHandlerUnitServer) DeleteCollection(context.Context, *DeleteCollectionRequest) (*DeleteCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCollection not implemented")
}
func (UnimplementedHandlerUnitServer) PutObject(context.Context, *PutObjectRequest) (*PutObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutObject not implemented")
}
func (UnimplementedHandlerUnitServer) PatchObject(context.Context, *PatchObjectRequest) (*PatchObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchObject not implemented")
}
func (UnimplementedHandlerUnitServer) MoveObject(context.Context, *MoveObjectRequest) (*MoveObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveObject not implemented")
}
func (UnimplementedHandlerUnitServer) GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (UnimplementedHandlerUnitServer) DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (UnimplementedHandlerUnitServer) ObjectInfo(context.Context, *ObjectInfoRequest) (*ObjectInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObjectInfo not implemented")
}
func (UnimplementedHandlerUnitServer) ListObjects(*ListObjectsRequest, HandlerUnit_ListObjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListObjects not implemented")
}
func (UnimplementedHandlerUnitServer) SearchObjects(*SearchObjectsRequest, HandlerUnit_SearchObjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchObjects not implemented")
}
func (UnimplementedHandlerUnitServer) mustEmbedUnimplementedHandlerUnitServer() {}

// UnsafeHandlerUnitServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HandlerUnitServer will
// result in compilation errors.
type UnsafeHandlerUnitServer interface {
	mustEmbedUnimplementedHandlerUnitServer()
}

func RegisterHandlerUnitServer(s grpc.ServiceRegistrar, srv HandlerUnitServer) {
	s.RegisterService(&_HandlerUnit_serviceDesc, srv)
}

func _HandlerUnit_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerUnitServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HandlerUnit/CreateCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerUnitServer).CreateCollection(ctx, req.(*CreateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerUnit_GetCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerUnitServer).GetCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HandlerUnit/GetCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerUnitServer).GetCollection(ctx, req.(*GetCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerUnit_ListCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerUnitServer).ListCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HandlerUnit/ListCollections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerUnitServer).ListCollections(ctx, req.(*ListCollectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerUnit_DeleteCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerUnitServer).DeleteCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HandlerUnit/DeleteCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerUnitServer).DeleteCollection(ctx, req.(*DeleteCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerUnit_PutObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerUnitServer).PutObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HandlerUnit/PutObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerUnitServer).PutObject(ctx, req.(*PutObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerUnit_PatchObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerUnitServer).PatchObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HandlerUnit/PatchObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerUnitServer).PatchObject(ctx, req.(*PatchObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerUnit_MoveObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerUnitServer).MoveObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HandlerUnit/MoveObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerUnitServer).MoveObject(ctx, req.(*MoveObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerUnit_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerUnitServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HandlerUnit/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerUnitServer).GetObject(ctx, req.(*GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerUnit_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerUnitServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HandlerUnit/DeleteObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerUnitServer).DeleteObject(ctx, req.(*DeleteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerUnit_ObjectInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandlerUnitServer).ObjectInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/HandlerUnit/ObjectInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandlerUnitServer).ObjectInfo(ctx, req.(*ObjectInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HandlerUnit_ListObjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListObjectsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HandlerUnitServer).ListObjects(m, &handlerUnitListObjectsServer{stream})
}

type HandlerUnit_ListObjectsServer interface {
	Send(*Object) error
	grpc.ServerStream
}

type handlerUnitListObjectsServer struct {
	grpc.ServerStream
}

func (x *handlerUnitListObjectsServer) Send(m *Object) error {
	return x.ServerStream.SendMsg(m)
}

func _HandlerUnit_SearchObjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchObjectsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HandlerUnitServer).SearchObjects(m, &handlerUnitSearchObjectsServer{stream})
}

type HandlerUnit_SearchObjectsServer interface {
	Send(*Object) error
	grpc.ServerStream
}

type handlerUnitSearchObjectsServer struct {
	grpc.ServerStream
}

func (x *handlerUnitSearchObjectsServer) Send(m *Object) error {
	return x.ServerStream.SendMsg(m)
}

var _HandlerUnit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HandlerUnit",
	HandlerType: (*HandlerUnitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCollection",
			Handler:    _HandlerUnit_CreateCollection_Handler,
		},
		{
			MethodName: "GetCollection",
			Handler:    _HandlerUnit_GetCollection_Handler,
		},
		{
			MethodName: "ListCollections",
			Handler:    _HandlerUnit_ListCollections_Handler,
		},
		{
			MethodName: "DeleteCollection",
			Handler:    _HandlerUnit_DeleteCollection_Handler,
		},
		{
			MethodName: "PutObject",
			Handler:    _HandlerUnit_PutObject_Handler,
		},
		{
			MethodName: "PatchObject",
			Handler:    _HandlerUnit_PatchObject_Handler,
		},
		{
			MethodName: "MoveObject",
			Handler:    _HandlerUnit_MoveObject_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _HandlerUnit_GetObject_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _HandlerUnit_DeleteObject_Handler,
		},
		{
			MethodName: "ObjectInfo",
			Handler:    _HandlerUnit_ObjectInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListObjects",
			Handler:       _HandlerUnit_ListObjects_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchObjects",
			Handler:       _HandlerUnit_SearchObjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb.proto",
}

// ACLClient is the client API for ACL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ACLClient interface {
	PutRules(ctx context.Context, in *PutRulesRequest, opts ...grpc.CallOption) (*PutRulesResponse, error)
	GetRules(ctx context.Context, in *GetRulesRequest, opts ...grpc.CallOption) (*GetRulesResponse, error)
	GetRulesForPath(ctx context.Context, in *GetRulesForPathRequest, opts ...grpc.CallOption) (*GetRulesForPathResponse, error)
	DeleteRules(ctx context.Context, in *DeleteRulesRequest, opts ...grpc.CallOption) (*DeleteRulesResponse, error)
	DeleteRulesForPath(ctx context.Context, in *DeleteRulesForPathRequest, opts ...grpc.CallOption) (*DeleteRulesForPathResponse, error)
}

type aCLClient struct {
	cc grpc.ClientConnInterface
}

func NewACLClient(cc grpc.ClientConnInterface) ACLClient {
	return &aCLClient{cc}
}

func (c *aCLClient) PutRules(ctx context.Context, in *PutRulesRequest, opts ...grpc.CallOption) (*PutRulesResponse, error) {
	out := new(PutRulesResponse)
	err := c.cc.Invoke(ctx, "/ACL/PutRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCLClient) GetRules(ctx context.Context, in *GetRulesRequest, opts ...grpc.CallOption) (*GetRulesResponse, error) {
	out := new(GetRulesResponse)
	err := c.cc.Invoke(ctx, "/ACL/GetRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCLClient) GetRulesForPath(ctx context.Context, in *GetRulesForPathRequest, opts ...grpc.CallOption) (*GetRulesForPathResponse, error) {
	out := new(GetRulesForPathResponse)
	err := c.cc.Invoke(ctx, "/ACL/GetRulesForPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCLClient) DeleteRules(ctx context.Context, in *DeleteRulesRequest, opts ...grpc.CallOption) (*DeleteRulesResponse, error) {
	out := new(DeleteRulesResponse)
	err := c.cc.Invoke(ctx, "/ACL/DeleteRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aCLClient) DeleteRulesForPath(ctx context.Context, in *DeleteRulesForPathRequest, opts ...grpc.CallOption) (*DeleteRulesForPathResponse, error) {
	out := new(DeleteRulesForPathResponse)
	err := c.cc.Invoke(ctx, "/ACL/DeleteRulesForPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ACLServer is the server API for ACL service.
// All implementations must embed UnimplementedACLServer
// for forward compatibility
type ACLServer interface {
	PutRules(context.Context, *PutRulesRequest) (*PutRulesResponse, error)
	GetRules(context.Context, *GetRulesRequest) (*GetRulesResponse, error)
	GetRulesForPath(context.Context, *GetRulesForPathRequest) (*GetRulesForPathResponse, error)
	DeleteRules(context.Context, *DeleteRulesRequest) (*DeleteRulesResponse, error)
	DeleteRulesForPath(context.Context, *DeleteRulesForPathRequest) (*DeleteRulesForPathResponse, error)
	mustEmbedUnimplementedACLServer()
}

// UnimplementedACLServer must be embedded to have forward compatible implementations.
type UnimplementedACLServer struct {
}

func (UnimplementedACLServer) PutRules(context.Context, *PutRulesRequest) (*PutRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutRules not implemented")
}
func (UnimplementedACLServer) GetRules(context.Context, *GetRulesRequest) (*GetRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRules not implemented")
}
func (UnimplementedACLServer) GetRulesForPath(context.Context, *GetRulesForPathRequest) (*GetRulesForPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRulesForPath not implemented")
}
func (UnimplementedACLServer) DeleteRules(context.Context, *DeleteRulesRequest) (*DeleteRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRules not implemented")
}
func (UnimplementedACLServer) DeleteRulesForPath(context.Context, *DeleteRulesForPathRequest) (*DeleteRulesForPathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRulesForPath not implemented")
}
func (UnimplementedACLServer) mustEmbedUnimplementedACLServer() {}

// UnsafeACLServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ACLServer will
// result in compilation errors.
type UnsafeACLServer interface {
	mustEmbedUnimplementedACLServer()
}

func RegisterACLServer(s grpc.ServiceRegistrar, srv ACLServer) {
	s.RegisterService(&_ACL_serviceDesc, srv)
}

func _ACL_PutRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServer).PutRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ACL/PutRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServer).PutRules(ctx, req.(*PutRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACL_GetRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServer).GetRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ACL/GetRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServer).GetRules(ctx, req.(*GetRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACL_GetRulesForPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRulesForPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServer).GetRulesForPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ACL/GetRulesForPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServer).GetRulesForPath(ctx, req.(*GetRulesForPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACL_DeleteRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServer).DeleteRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ACL/DeleteRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServer).DeleteRules(ctx, req.(*DeleteRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ACL_DeleteRulesForPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRulesForPathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ACLServer).DeleteRulesForPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ACL/DeleteRulesForPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ACLServer).DeleteRulesForPath(ctx, req.(*DeleteRulesForPathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ACL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ACL",
	HandlerType: (*ACLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutRules",
			Handler:    _ACL_PutRules_Handler,
		},
		{
			MethodName: "GetRules",
			Handler:    _ACL_GetRules_Handler,
		},
		{
			MethodName: "GetRulesForPath",
			Handler:    _ACL_GetRulesForPath_Handler,
		},
		{
			MethodName: "DeleteRules",
			Handler:    _ACL_DeleteRules_Handler,
		},
		{
			MethodName: "DeleteRulesForPath",
			Handler:    _ACL_DeleteRulesForPath_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}

// SearchEngineClient is the client API for SearchEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchEngineClient interface {
	Feed(ctx context.Context, opts ...grpc.CallOption) (SearchEngine_FeedClient, error)
	Search(ctx context.Context, in *ResearchRequest, opts ...grpc.CallOption) (SearchEngine_SearchClient, error)
}

type searchEngineClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchEngineClient(cc grpc.ClientConnInterface) SearchEngineClient {
	return &searchEngineClient{cc}
}

func (c *searchEngineClient) Feed(ctx context.Context, opts ...grpc.CallOption) (SearchEngine_FeedClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SearchEngine_serviceDesc.Streams[0], "/SearchEngine/Feed", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchEngineFeedClient{stream}
	return x, nil
}

type SearchEngine_FeedClient interface {
	Send(*MessageFeed) error
	CloseAndRecv() (*FeedResponse, error)
	grpc.ClientStream
}

type searchEngineFeedClient struct {
	grpc.ClientStream
}

func (x *searchEngineFeedClient) Send(m *MessageFeed) error {
	return x.ClientStream.SendMsg(m)
}

func (x *searchEngineFeedClient) CloseAndRecv() (*FeedResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FeedResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *searchEngineClient) Search(ctx context.Context, in *ResearchRequest, opts ...grpc.CallOption) (SearchEngine_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SearchEngine_serviceDesc.Streams[1], "/SearchEngine/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &searchEngineSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SearchEngine_SearchClient interface {
	Recv() (*SearchResult, error)
	grpc.ClientStream
}

type searchEngineSearchClient struct {
	grpc.ClientStream
}

func (x *searchEngineSearchClient) Recv() (*SearchResult, error) {
	m := new(SearchResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SearchEngineServer is the server API for SearchEngine service.
// All implementations must embed UnimplementedSearchEngineServer
// for forward compatibility
type SearchEngineServer interface {
	Feed(SearchEngine_FeedServer) error
	Search(*ResearchRequest, SearchEngine_SearchServer) error
	mustEmbedUnimplementedSearchEngineServer()
}

// UnimplementedSearchEngineServer must be embedded to have forward compatible implementations.
type UnimplementedSearchEngineServer struct {
}

func (UnimplementedSearchEngineServer) Feed(SearchEngine_FeedServer) error {
	return status.Errorf(codes.Unimplemented, "method Feed not implemented")
}
func (UnimplementedSearchEngineServer) Search(*ResearchRequest, SearchEngine_SearchServer) error {
	return status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedSearchEngineServer) mustEmbedUnimplementedSearchEngineServer() {}

// UnsafeSearchEngineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchEngineServer will
// result in compilation errors.
type UnsafeSearchEngineServer interface {
	mustEmbedUnimplementedSearchEngineServer()
}

func RegisterSearchEngineServer(s grpc.ServiceRegistrar, srv SearchEngineServer) {
	s.RegisterService(&_SearchEngine_serviceDesc, srv)
}

func _SearchEngine_Feed_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SearchEngineServer).Feed(&searchEngineFeedServer{stream})
}

type SearchEngine_FeedServer interface {
	SendAndClose(*FeedResponse) error
	Recv() (*MessageFeed, error)
	grpc.ServerStream
}

type searchEngineFeedServer struct {
	grpc.ServerStream
}

func (x *searchEngineFeedServer) SendAndClose(m *FeedResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *searchEngineFeedServer) Recv() (*MessageFeed, error) {
	m := new(MessageFeed)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SearchEngine_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ResearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SearchEngineServer).Search(m, &searchEngineSearchServer{stream})
}

type SearchEngine_SearchServer interface {
	Send(*SearchResult) error
	grpc.ServerStream
}

type searchEngineSearchServer struct {
	grpc.ServerStream
}

func (x *searchEngineSearchServer) Send(m *SearchResult) error {
	return x.ServerStream.SendMsg(m)
}

var _SearchEngine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SearchEngine",
	HandlerType: (*SearchEngineServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Feed",
			Handler:       _SearchEngine_Feed_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Search",
			Handler:       _SearchEngine_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb.proto",
}
