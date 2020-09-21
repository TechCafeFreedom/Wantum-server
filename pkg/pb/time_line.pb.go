// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.4
// source: time_line.proto

package pb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// TODO: タイムライン情報(仮)
type Timeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// タイムラインに表示する思い出一覧
	Memories []*Memory `protobuf:"bytes,1,rep,name=memories,proto3" json:"memories,omitempty"`
}

func (x *Timeline) Reset() {
	*x = Timeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_line_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timeline) ProtoMessage() {}

func (x *Timeline) ProtoReflect() protoreflect.Message {
	mi := &file_time_line_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timeline.ProtoReflect.Descriptor instead.
func (*Timeline) Descriptor() ([]byte, []int) {
	return file_time_line_proto_rawDescGZIP(), []int{0}
}

func (x *Timeline) GetMemories() []*Memory {
	if x != nil {
		return x.Memories
	}
	return nil
}

var File_time_line_proto protoreflect.FileDescriptor

var file_time_line_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x1a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x08,
	0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x32, 0x54, 0x0a, 0x0f, 0x54, 0x69,
	0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x00,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_time_line_proto_rawDescOnce sync.Once
	file_time_line_proto_rawDescData = file_time_line_proto_rawDesc
)

func file_time_line_proto_rawDescGZIP() []byte {
	file_time_line_proto_rawDescOnce.Do(func() {
		file_time_line_proto_rawDescData = protoimpl.X.CompressGZIP(file_time_line_proto_rawDescData)
	})
	return file_time_line_proto_rawDescData
}

var file_time_line_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_time_line_proto_goTypes = []interface{}{
	(*Timeline)(nil),    // 0: proto_timeline.Timeline
	(*Memory)(nil),      // 1: proto_memory.Memory
	(*empty.Empty)(nil), // 2: google.protobuf.Empty
}
var file_time_line_proto_depIdxs = []int32{
	1, // 0: proto_timeline.Timeline.memories:type_name -> proto_memory.Memory
	2, // 1: proto_timeline.TimelineService.GetTimeline:input_type -> google.protobuf.Empty
	0, // 2: proto_timeline.TimelineService.GetTimeline:output_type -> proto_timeline.Timeline
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_time_line_proto_init() }
func file_time_line_proto_init() {
	if File_time_line_proto != nil {
		return
	}
	file_memory_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_time_line_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timeline); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_time_line_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_time_line_proto_goTypes,
		DependencyIndexes: file_time_line_proto_depIdxs,
		MessageInfos:      file_time_line_proto_msgTypes,
	}.Build()
	File_time_line_proto = out.File
	file_time_line_proto_rawDesc = nil
	file_time_line_proto_goTypes = nil
	file_time_line_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TimelineServiceClient is the client API for TimelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimelineServiceClient interface {
	// タイムライン情報の取得
	GetTimeline(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Timeline, error)
}

type timelineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTimelineServiceClient(cc grpc.ClientConnInterface) TimelineServiceClient {
	return &timelineServiceClient{cc}
}

func (c *timelineServiceClient) GetTimeline(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Timeline, error) {
	out := new(Timeline)
	err := c.cc.Invoke(ctx, "/proto_timeline.TimelineService/GetTimeline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimelineServiceServer is the server API for TimelineService service.
type TimelineServiceServer interface {
	// タイムライン情報の取得
	GetTimeline(context.Context, *empty.Empty) (*Timeline, error)
}

// UnimplementedTimelineServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTimelineServiceServer struct {
}

func (*UnimplementedTimelineServiceServer) GetTimeline(context.Context, *empty.Empty) (*Timeline, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeline not implemented")
}

func RegisterTimelineServiceServer(s *grpc.Server, srv TimelineServiceServer) {
	s.RegisterService(&_TimelineService_serviceDesc, srv)
}

func _TimelineService_GetTimeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimelineServiceServer).GetTimeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_timeline.TimelineService/GetTimeline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimelineServiceServer).GetTimeline(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TimelineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_timeline.TimelineService",
	HandlerType: (*TimelineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTimeline",
			Handler:    _TimelineService_GetTimeline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "time_line.proto",
}
