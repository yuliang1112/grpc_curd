// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package update_data

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UpdateDataManageClient is the client API for UpdateDataManage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpdateDataManageClient interface {
	UpdateData(ctx context.Context, in *UpdateDataRequest, opts ...grpc.CallOption) (*UpdateDataResponse, error)
}

type updateDataManageClient struct {
	cc grpc.ClientConnInterface
}

func NewUpdateDataManageClient(cc grpc.ClientConnInterface) UpdateDataManageClient {
	return &updateDataManageClient{cc}
}

func (c *updateDataManageClient) UpdateData(ctx context.Context, in *UpdateDataRequest, opts ...grpc.CallOption) (*UpdateDataResponse, error) {
	out := new(UpdateDataResponse)
	err := c.cc.Invoke(ctx, "/update_data.UpdateDataManage/UpdateData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateDataManageServer is the server API for UpdateDataManage service.
// All implementations must embed UnimplementedUpdateDataManageServer
// for forward compatibility
type UpdateDataManageServer interface {
	UpdateData(context.Context, *UpdateDataRequest) (*UpdateDataResponse, error)
	mustEmbedUnimplementedUpdateDataManageServer()
}

// UnimplementedUpdateDataManageServer must be embedded to have forward compatible implementations.
type UnimplementedUpdateDataManageServer struct {
}

func (UnimplementedUpdateDataManageServer) UpdateData(context.Context, *UpdateDataRequest) (*UpdateDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateData not implemented")
}
func (UnimplementedUpdateDataManageServer) mustEmbedUnimplementedUpdateDataManageServer() {}

// UnsafeUpdateDataManageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpdateDataManageServer will
// result in compilation errors.
type UnsafeUpdateDataManageServer interface {
	mustEmbedUnimplementedUpdateDataManageServer()
}

func RegisterUpdateDataManageServer(s grpc.ServiceRegistrar, srv UpdateDataManageServer) {
	s.RegisterService(&UpdateDataManage_ServiceDesc, srv)
}

func _UpdateDataManage_UpdateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateDataManageServer).UpdateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/update_data.UpdateDataManage/UpdateData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateDataManageServer).UpdateData(ctx, req.(*UpdateDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UpdateDataManage_ServiceDesc is the grpc.ServiceDesc for UpdateDataManage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UpdateDataManage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "update_data.UpdateDataManage",
	HandlerType: (*UpdateDataManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateData",
			Handler:    _UpdateDataManage_UpdateData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "updatedata.proto",
}
