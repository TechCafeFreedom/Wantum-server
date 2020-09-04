// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.4
// source: memory.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

// 思い出情報取得用リクエスト
type GetMemoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 思い出1件を特定するための思い出ID
	MemoryId int64 `protobuf:"varint,1,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
}

func (x *GetMemoryRequest) Reset() {
	*x = GetMemoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoryRequest) ProtoMessage() {}

func (x *GetMemoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoryRequest.ProtoReflect.Descriptor instead.
func (*GetMemoryRequest) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{0}
}

func (x *GetMemoryRequest) GetMemoryId() int64 {
	if x != nil {
		return x.MemoryId
	}
	return 0
}

// 思い出一覧情報
type Memories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 思い出カウント情報
	MemoryCount *MemoryCountInfo `protobuf:"bytes,1,opt,name=memory_count,json=memoryCount,proto3" json:"memory_count,omitempty"`
	// 思い出一覧（配列）
	Memories []*MemoryInfo `protobuf:"bytes,2,rep,name=memories,proto3" json:"memories,omitempty"`
}

func (x *Memories) Reset() {
	*x = Memories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Memories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Memories) ProtoMessage() {}

func (x *Memories) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Memories.ProtoReflect.Descriptor instead.
func (*Memories) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{1}
}

func (x *Memories) GetMemoryCount() *MemoryCountInfo {
	if x != nil {
		return x.MemoryCount
	}
	return nil
}

func (x *Memories) GetMemories() []*MemoryInfo {
	if x != nil {
		return x.Memories
	}
	return nil
}

// 思い出のカウント情報
type MemoryCountInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 過去に投稿した思い出の数
	MemoriesCount int64 `protobuf:"varint,1,opt,name=memories_count,json=memoriesCount,proto3" json:"memories_count,omitempty"`
	// 公開中の思い出の数
	PublishedCount int64 `protobuf:"varint,2,opt,name=published_count,json=publishedCount,proto3" json:"published_count,omitempty"`
}

func (x *MemoryCountInfo) Reset() {
	*x = MemoryCountInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryCountInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryCountInfo) ProtoMessage() {}

func (x *MemoryCountInfo) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryCountInfo.ProtoReflect.Descriptor instead.
func (*MemoryCountInfo) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{2}
}

func (x *MemoryCountInfo) GetMemoriesCount() int64 {
	if x != nil {
		return x.MemoriesCount
	}
	return 0
}

func (x *MemoryCountInfo) GetPublishedCount() int64 {
	if x != nil {
		return x.PublishedCount
	}
	return 0
}

// 思い出詳細情報
type MemoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 思い出ID
	MemoryId int64 `protobuf:"varint,1,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	// やったこと
	Activity string `protobuf:"bytes,2,opt,name=activity,proto3" json:"activity,omitempty"`
	// 日付(UNIX)
	Date int64 `protobuf:"varint,3,opt,name=date,proto3" json:"date,omitempty"`
	// 思い出の説明
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// 場所
	Place string `protobuf:"bytes,5,opt,name=place,proto3" json:"place,omitempty"`
	// 投稿されている画像一覧（配列）
	ImageUrls []string `protobuf:"bytes,6,rep,name=image_urls,json=imageUrls,proto3" json:"image_urls,omitempty"`
	// タグ一覧（配列）
	Tags []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *MemoryInfo) Reset() {
	*x = MemoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memory_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryInfo) ProtoMessage() {}

func (x *MemoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_memory_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryInfo.ProtoReflect.Descriptor instead.
func (*MemoryInfo) Descriptor() ([]byte, []int) {
	return file_memory_proto_rawDescGZIP(), []int{3}
}

func (x *MemoryInfo) GetMemoryId() int64 {
	if x != nil {
		return x.MemoryId
	}
	return 0
}

func (x *MemoryInfo) GetActivity() string {
	if x != nil {
		return x.Activity
	}
	return ""
}

func (x *MemoryInfo) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *MemoryInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MemoryInfo) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *MemoryInfo) GetImageUrls() []string {
	if x != nil {
		return x.ImageUrls
	}
	return nil
}

func (x *MemoryInfo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_memory_proto protoreflect.FileDescriptor

var file_memory_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x2f, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x82, 0x01,
	0x0a, 0x08, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x6d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x08,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x61, 0x0a, 0x0f, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xc4, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x32, 0x58, 0x0a, 0x0d,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_memory_proto_rawDescOnce sync.Once
	file_memory_proto_rawDescData = file_memory_proto_rawDesc
)

func file_memory_proto_rawDescGZIP() []byte {
	file_memory_proto_rawDescOnce.Do(func() {
		file_memory_proto_rawDescData = protoimpl.X.CompressGZIP(file_memory_proto_rawDescData)
	})
	return file_memory_proto_rawDescData
}

var file_memory_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_memory_proto_goTypes = []interface{}{
	(*GetMemoryRequest)(nil), // 0: proto_memory.GetMemoryRequest
	(*Memories)(nil),         // 1: proto_memory.Memories
	(*MemoryCountInfo)(nil),  // 2: proto_memory.MemoryCountInfo
	(*MemoryInfo)(nil),       // 3: proto_memory.MemoryInfo
}
var file_memory_proto_depIdxs = []int32{
	2, // 0: proto_memory.Memories.memory_count:type_name -> proto_memory.MemoryCountInfo
	3, // 1: proto_memory.Memories.memories:type_name -> proto_memory.MemoryInfo
	0, // 2: proto_memory.MemoryService.GetMemory:input_type -> proto_memory.GetMemoryRequest
	3, // 3: proto_memory.MemoryService.GetMemory:output_type -> proto_memory.MemoryInfo
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_memory_proto_init() }
func file_memory_proto_init() {
	if File_memory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_memory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemoryRequest); i {
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
		file_memory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Memories); i {
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
		file_memory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryCountInfo); i {
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
		file_memory_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryInfo); i {
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
			RawDescriptor: file_memory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_memory_proto_goTypes,
		DependencyIndexes: file_memory_proto_depIdxs,
		MessageInfos:      file_memory_proto_msgTypes,
	}.Build()
	File_memory_proto = out.File
	file_memory_proto_rawDesc = nil
	file_memory_proto_goTypes = nil
	file_memory_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MemoryServiceClient is the client API for MemoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MemoryServiceClient interface {
	// IDをもとに思い出1件を取得
	GetMemory(ctx context.Context, in *GetMemoryRequest, opts ...grpc.CallOption) (*MemoryInfo, error)
}

type memoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemoryServiceClient(cc grpc.ClientConnInterface) MemoryServiceClient {
	return &memoryServiceClient{cc}
}

func (c *memoryServiceClient) GetMemory(ctx context.Context, in *GetMemoryRequest, opts ...grpc.CallOption) (*MemoryInfo, error) {
	out := new(MemoryInfo)
	err := c.cc.Invoke(ctx, "/proto_memory.MemoryService/GetMemory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemoryServiceServer is the server API for MemoryService service.
type MemoryServiceServer interface {
	// IDをもとに思い出1件を取得
	GetMemory(context.Context, *GetMemoryRequest) (*MemoryInfo, error)
}

// UnimplementedMemoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMemoryServiceServer struct {
}

func (*UnimplementedMemoryServiceServer) GetMemory(context.Context, *GetMemoryRequest) (*MemoryInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemory not implemented")
}

func RegisterMemoryServiceServer(s *grpc.Server, srv MemoryServiceServer) {
	s.RegisterService(&_MemoryService_serviceDesc, srv)
}

func _MemoryService_GetMemory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemoryServiceServer).GetMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_memory.MemoryService/GetMemory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemoryServiceServer).GetMemory(ctx, req.(*GetMemoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MemoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_memory.MemoryService",
	HandlerType: (*MemoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMemory",
			Handler:    _MemoryService_GetMemory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "memory.proto",
}
