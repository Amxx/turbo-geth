// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package remote

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KVClient interface {
	// open a cursor on given position of given bucket
	// if streaming requested - streams all data: stops if client's buffer is full, resumes when client read enough from buffer
	// if streaming not requested - streams next data only when clients sends message to bi-directional channel
	// no full consistency guarantee - server implementation can close/open underlying db transaction at any time
	Seek(ctx context.Context, opts ...grpc.CallOption) (KV_SeekClient, error)
}

type kVClient struct {
	cc grpc.ClientConnInterface
}

func NewKVClient(cc grpc.ClientConnInterface) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) Seek(ctx context.Context, opts ...grpc.CallOption) (KV_SeekClient, error) {
	stream, err := c.cc.NewStream(ctx, &_KV_serviceDesc.Streams[0], "/remote.KV/Seek", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVSeekClient{stream}
	return x, nil
}

type KV_SeekClient interface {
	Send(*SeekRequest) error
	Recv() (*Pair, error)
	grpc.ClientStream
}

type kVSeekClient struct {
	grpc.ClientStream
}

func (x *kVSeekClient) Send(m *SeekRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kVSeekClient) Recv() (*Pair, error) {
	m := new(Pair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KVServer is the server API for KV service.
// All implementations must embed UnimplementedKVServer
// for forward compatibility
type KVServer interface {
	// open a cursor on given position of given bucket
	// if streaming requested - streams all data: stops if client's buffer is full, resumes when client read enough from buffer
	// if streaming not requested - streams next data only when clients sends message to bi-directional channel
	// no full consistency guarantee - server implementation can close/open underlying db transaction at any time
	Seek(KV_SeekServer) error
	mustEmbedUnimplementedKVServer()
}

// UnimplementedKVServer must be embedded to have forward compatible implementations.
type UnimplementedKVServer struct {
}

func (*UnimplementedKVServer) Seek(KV_SeekServer) error {
	return status.Errorf(codes.Unimplemented, "method Seek not implemented")
}
func (*UnimplementedKVServer) mustEmbedUnimplementedKVServer() {}

func RegisterKVServer(s *grpc.Server, srv KVServer) {
	s.RegisterService(&_KV_serviceDesc, srv)
}

func _KV_Seek_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KVServer).Seek(&kVSeekServer{stream})
}

type KV_SeekServer interface {
	Send(*Pair) error
	Recv() (*SeekRequest, error)
	grpc.ServerStream
}

type kVSeekServer struct {
	grpc.ServerStream
}

func (x *kVSeekServer) Send(m *Pair) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kVSeekServer) Recv() (*SeekRequest, error) {
	m := new(SeekRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _KV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remote.KV",
	HandlerType: (*KVServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Seek",
			Handler:       _KV_Seek_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "remote/kv.proto",
}
