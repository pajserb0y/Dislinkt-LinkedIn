// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: requests-service.proto

package request

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

// RequestsServiceClient is the client API for RequestsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RequestsServiceClient interface {
	GetAllByRecieverId(ctx context.Context, in *GetAllByRecieverIdRequest, opts ...grpc.CallOption) (*GetAllByRecieverIdResponse, error)
	AcceptRequest(ctx context.Context, in *AcceptRequestRequest, opts ...grpc.CallOption) (*AcceptRequestResponse, error)
	DeclineRequest(ctx context.Context, in *DeclineRequestRequest, opts ...grpc.CallOption) (*DeclineRequestResponse, error)
	SendRequest(ctx context.Context, in *SendRequestRequest, opts ...grpc.CallOption) (*SendRequestResponse, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	AreConnected(ctx context.Context, in *AreConnectedRequest, opts ...grpc.CallOption) (*AreConnectedResponse, error)
	FindConnections(ctx context.Context, in *FindConnectionsRequest, opts ...grpc.CallOption) (*FindConnectionsResponse, error)
	FindMessages(ctx context.Context, in *FindMessagesRequest, opts ...grpc.CallOption) (*FindMessagesResponse, error)
	DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionResponse, error)
	GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error)
}

type requestsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestsServiceClient(cc grpc.ClientConnInterface) RequestsServiceClient {
	return &requestsServiceClient{cc}
}

func (c *requestsServiceClient) GetAllByRecieverId(ctx context.Context, in *GetAllByRecieverIdRequest, opts ...grpc.CallOption) (*GetAllByRecieverIdResponse, error) {
	out := new(GetAllByRecieverIdResponse)
	err := c.cc.Invoke(ctx, "/requests.RequestsService/GetAllByRecieverId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestsServiceClient) AcceptRequest(ctx context.Context, in *AcceptRequestRequest, opts ...grpc.CallOption) (*AcceptRequestResponse, error) {
	out := new(AcceptRequestResponse)
	err := c.cc.Invoke(ctx, "/requests.RequestsService/AcceptRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestsServiceClient) DeclineRequest(ctx context.Context, in *DeclineRequestRequest, opts ...grpc.CallOption) (*DeclineRequestResponse, error) {
	out := new(DeclineRequestResponse)
	err := c.cc.Invoke(ctx, "/requests.RequestsService/DeclineRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestsServiceClient) SendRequest(ctx context.Context, in *SendRequestRequest, opts ...grpc.CallOption) (*SendRequestResponse, error) {
	out := new(SendRequestResponse)
	err := c.cc.Invoke(ctx, "/requests.RequestsService/SendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestsServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/requests.RequestsService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestsServiceClient) AreConnected(ctx context.Context, in *AreConnectedRequest, opts ...grpc.CallOption) (*AreConnectedResponse, error) {
	out := new(AreConnectedResponse)
	err := c.cc.Invoke(ctx, "/requests.RequestsService/AreConnected", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestsServiceClient) FindConnections(ctx context.Context, in *FindConnectionsRequest, opts ...grpc.CallOption) (*FindConnectionsResponse, error) {
	out := new(FindConnectionsResponse)
	err := c.cc.Invoke(ctx, "/requests.RequestsService/FindConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestsServiceClient) FindMessages(ctx context.Context, in *FindMessagesRequest, opts ...grpc.CallOption) (*FindMessagesResponse, error) {
	out := new(FindMessagesResponse)
	err := c.cc.Invoke(ctx, "/requests.RequestsService/FindMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestsServiceClient) DeleteConnection(ctx context.Context, in *DeleteConnectionRequest, opts ...grpc.CallOption) (*DeleteConnectionResponse, error) {
	out := new(DeleteConnectionResponse)
	err := c.cc.Invoke(ctx, "/requests.RequestsService/DeleteConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestsServiceClient) GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error) {
	out := new(GetNotificationsResponse)
	err := c.cc.Invoke(ctx, "/requests.RequestsService/GetNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RequestsServiceServer is the server API for RequestsService service.
// All implementations must embed UnimplementedRequestsServiceServer
// for forward compatibility
type RequestsServiceServer interface {
	GetAllByRecieverId(context.Context, *GetAllByRecieverIdRequest) (*GetAllByRecieverIdResponse, error)
	AcceptRequest(context.Context, *AcceptRequestRequest) (*AcceptRequestResponse, error)
	DeclineRequest(context.Context, *DeclineRequestRequest) (*DeclineRequestResponse, error)
	SendRequest(context.Context, *SendRequestRequest) (*SendRequestResponse, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	AreConnected(context.Context, *AreConnectedRequest) (*AreConnectedResponse, error)
	FindConnections(context.Context, *FindConnectionsRequest) (*FindConnectionsResponse, error)
	FindMessages(context.Context, *FindMessagesRequest) (*FindMessagesResponse, error)
	DeleteConnection(context.Context, *DeleteConnectionRequest) (*DeleteConnectionResponse, error)
	GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error)
	mustEmbedUnimplementedRequestsServiceServer()
}

// UnimplementedRequestsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRequestsServiceServer struct {
}

func (UnimplementedRequestsServiceServer) GetAllByRecieverId(context.Context, *GetAllByRecieverIdRequest) (*GetAllByRecieverIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByRecieverId not implemented")
}
func (UnimplementedRequestsServiceServer) AcceptRequest(context.Context, *AcceptRequestRequest) (*AcceptRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptRequest not implemented")
}
func (UnimplementedRequestsServiceServer) DeclineRequest(context.Context, *DeclineRequestRequest) (*DeclineRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeclineRequest not implemented")
}
func (UnimplementedRequestsServiceServer) SendRequest(context.Context, *SendRequestRequest) (*SendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}
func (UnimplementedRequestsServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedRequestsServiceServer) AreConnected(context.Context, *AreConnectedRequest) (*AreConnectedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AreConnected not implemented")
}
func (UnimplementedRequestsServiceServer) FindConnections(context.Context, *FindConnectionsRequest) (*FindConnectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindConnections not implemented")
}
func (UnimplementedRequestsServiceServer) FindMessages(context.Context, *FindMessagesRequest) (*FindMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMessages not implemented")
}
func (UnimplementedRequestsServiceServer) DeleteConnection(context.Context, *DeleteConnectionRequest) (*DeleteConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConnection not implemented")
}
func (UnimplementedRequestsServiceServer) GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}
func (UnimplementedRequestsServiceServer) mustEmbedUnimplementedRequestsServiceServer() {}

// UnsafeRequestsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RequestsServiceServer will
// result in compilation errors.
type UnsafeRequestsServiceServer interface {
	mustEmbedUnimplementedRequestsServiceServer()
}

func RegisterRequestsServiceServer(s grpc.ServiceRegistrar, srv RequestsServiceServer) {
	s.RegisterService(&RequestsService_ServiceDesc, srv)
}

func _RequestsService_GetAllByRecieverId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllByRecieverIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestsServiceServer).GetAllByRecieverId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/requests.RequestsService/GetAllByRecieverId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestsServiceServer).GetAllByRecieverId(ctx, req.(*GetAllByRecieverIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestsService_AcceptRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestsServiceServer).AcceptRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/requests.RequestsService/AcceptRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestsServiceServer).AcceptRequest(ctx, req.(*AcceptRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestsService_DeclineRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclineRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestsServiceServer).DeclineRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/requests.RequestsService/DeclineRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestsServiceServer).DeclineRequest(ctx, req.(*DeclineRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestsService_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestsServiceServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/requests.RequestsService/SendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestsServiceServer).SendRequest(ctx, req.(*SendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestsService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestsServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/requests.RequestsService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestsServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestsService_AreConnected_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AreConnectedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestsServiceServer).AreConnected(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/requests.RequestsService/AreConnected",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestsServiceServer).AreConnected(ctx, req.(*AreConnectedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestsService_FindConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindConnectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestsServiceServer).FindConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/requests.RequestsService/FindConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestsServiceServer).FindConnections(ctx, req.(*FindConnectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestsService_FindMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestsServiceServer).FindMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/requests.RequestsService/FindMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestsServiceServer).FindMessages(ctx, req.(*FindMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestsService_DeleteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestsServiceServer).DeleteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/requests.RequestsService/DeleteConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestsServiceServer).DeleteConnection(ctx, req.(*DeleteConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestsService_GetNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestsServiceServer).GetNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/requests.RequestsService/GetNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestsServiceServer).GetNotifications(ctx, req.(*GetNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RequestsService_ServiceDesc is the grpc.ServiceDesc for RequestsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RequestsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "requests.RequestsService",
	HandlerType: (*RequestsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllByRecieverId",
			Handler:    _RequestsService_GetAllByRecieverId_Handler,
		},
		{
			MethodName: "AcceptRequest",
			Handler:    _RequestsService_AcceptRequest_Handler,
		},
		{
			MethodName: "DeclineRequest",
			Handler:    _RequestsService_DeclineRequest_Handler,
		},
		{
			MethodName: "SendRequest",
			Handler:    _RequestsService_SendRequest_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _RequestsService_SendMessage_Handler,
		},
		{
			MethodName: "AreConnected",
			Handler:    _RequestsService_AreConnected_Handler,
		},
		{
			MethodName: "FindConnections",
			Handler:    _RequestsService_FindConnections_Handler,
		},
		{
			MethodName: "FindMessages",
			Handler:    _RequestsService_FindMessages_Handler,
		},
		{
			MethodName: "DeleteConnection",
			Handler:    _RequestsService_DeleteConnection_Handler,
		},
		{
			MethodName: "GetNotifications",
			Handler:    _RequestsService_GetNotifications_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "requests-service.proto",
}
