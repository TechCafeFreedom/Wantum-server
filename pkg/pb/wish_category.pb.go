// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.4
// source: wish_category.proto

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

// やりたいことカテゴリー新規作成リクエスト
type CreateWishCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// やりたいことリストID
	WishListId int64 `protobuf:"varint,1,opt,name=wish_list_id,json=wishListId,proto3" json:"wish_list_id,omitempty"`
	// やりたいことカテゴリー名
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateWishCategoryRequest) Reset() {
	*x = CreateWishCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wish_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWishCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWishCategoryRequest) ProtoMessage() {}

func (x *CreateWishCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wish_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWishCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateWishCategoryRequest) Descriptor() ([]byte, []int) {
	return file_wish_category_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWishCategoryRequest) GetWishListId() int64 {
	if x != nil {
		return x.WishListId
	}
	return 0
}

func (x *CreateWishCategoryRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// やりたいことカテゴリー名更新リクエスト
type UpdateWishCategoryNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// やりたいことカテゴリーID
	WishCategoryId int64 `protobuf:"varint,1,opt,name=wish_category_id,json=wishCategoryId,proto3" json:"wish_category_id,omitempty"`
	// 新規やりたいことカテゴリー名
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *UpdateWishCategoryNameRequest) Reset() {
	*x = UpdateWishCategoryNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wish_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWishCategoryNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWishCategoryNameRequest) ProtoMessage() {}

func (x *UpdateWishCategoryNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wish_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWishCategoryNameRequest.ProtoReflect.Descriptor instead.
func (*UpdateWishCategoryNameRequest) Descriptor() ([]byte, []int) {
	return file_wish_category_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateWishCategoryNameRequest) GetWishCategoryId() int64 {
	if x != nil {
		return x.WishCategoryId
	}
	return 0
}

func (x *UpdateWishCategoryNameRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// やりたいことカテゴリー情報
type WishCategory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// やりたいことカテゴリーを特定するためのID
	WishCategoryId int64 `protobuf:"varint,1,opt,name=wish_category_id,json=wishCategoryId,proto3" json:"wish_category_id,omitempty"`
	// やりたいことカテゴリー名
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// カテゴリー内のやりたいことカード一覧（配列）
	WishCards []*WishCard `protobuf:"bytes,3,rep,name=wish_cards,json=wishCards,proto3" json:"wish_cards,omitempty"`
}

func (x *WishCategory) Reset() {
	*x = WishCategory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wish_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WishCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishCategory) ProtoMessage() {}

func (x *WishCategory) ProtoReflect() protoreflect.Message {
	mi := &file_wish_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishCategory.ProtoReflect.Descriptor instead.
func (*WishCategory) Descriptor() ([]byte, []int) {
	return file_wish_category_proto_rawDescGZIP(), []int{2}
}

func (x *WishCategory) GetWishCategoryId() int64 {
	if x != nil {
		return x.WishCategoryId
	}
	return 0
}

func (x *WishCategory) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *WishCategory) GetWishCards() []*WishCard {
	if x != nil {
		return x.WishCards
	}
	return nil
}

var File_wish_category_proto protoreflect.FileDescriptor

var file_wish_category_proto_rawDesc = []byte{
	0x0a, 0x13, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69, 0x73,
	0x68, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x0f, 0x77, 0x69, 0x73, 0x68, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x57, 0x69, 0x73, 0x68, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x77, 0x69, 0x73, 0x68,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x5f, 0x0a, 0x1d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x10, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x77, 0x69, 0x73, 0x68, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x87, 0x01,
	0x0a, 0x0c, 0x57, 0x69, 0x73, 0x68, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28,
	0x0a, 0x10, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x77, 0x69, 0x73, 0x68, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x37,
	0x0a, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x63,
	0x61, 0x72, 0x64, 0x2e, 0x57, 0x69, 0x73, 0x68, 0x43, 0x61, 0x72, 0x64, 0x52, 0x09, 0x77, 0x69,
	0x73, 0x68, 0x43, 0x61, 0x72, 0x64, 0x73, 0x32, 0xdb, 0x01, 0x0a, 0x13, 0x57, 0x69, 0x73, 0x68,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69,
	0x73, 0x68, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x57, 0x69, 0x73, 0x68, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x65,
	0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x77, 0x69, 0x73, 0x68, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wish_category_proto_rawDescOnce sync.Once
	file_wish_category_proto_rawDescData = file_wish_category_proto_rawDesc
)

func file_wish_category_proto_rawDescGZIP() []byte {
	file_wish_category_proto_rawDescOnce.Do(func() {
		file_wish_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_wish_category_proto_rawDescData)
	})
	return file_wish_category_proto_rawDescData
}

var file_wish_category_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_wish_category_proto_goTypes = []interface{}{
	(*CreateWishCategoryRequest)(nil),     // 0: proto_wishcategory.CreateWishCategoryRequest
	(*UpdateWishCategoryNameRequest)(nil), // 1: proto_wishcategory.UpdateWishCategoryNameRequest
	(*WishCategory)(nil),                  // 2: proto_wishcategory.WishCategory
	(*WishCard)(nil),                      // 3: proto_wishcard.WishCard
	(*empty.Empty)(nil),                   // 4: google.protobuf.Empty
}
var file_wish_category_proto_depIdxs = []int32{
	3, // 0: proto_wishcategory.WishCategory.wish_cards:type_name -> proto_wishcard.WishCard
	0, // 1: proto_wishcategory.WishCategoryService.CreateWishCategory:input_type -> proto_wishcategory.CreateWishCategoryRequest
	1, // 2: proto_wishcategory.WishCategoryService.UpdateWishCategoryName:input_type -> proto_wishcategory.UpdateWishCategoryNameRequest
	4, // 3: proto_wishcategory.WishCategoryService.CreateWishCategory:output_type -> google.protobuf.Empty
	4, // 4: proto_wishcategory.WishCategoryService.UpdateWishCategoryName:output_type -> google.protobuf.Empty
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_wish_category_proto_init() }
func file_wish_category_proto_init() {
	if File_wish_category_proto != nil {
		return
	}
	file_wish_card_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wish_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWishCategoryRequest); i {
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
		file_wish_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWishCategoryNameRequest); i {
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
		file_wish_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WishCategory); i {
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
			RawDescriptor: file_wish_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wish_category_proto_goTypes,
		DependencyIndexes: file_wish_category_proto_depIdxs,
		MessageInfos:      file_wish_category_proto_msgTypes,
	}.Build()
	File_wish_category_proto = out.File
	file_wish_category_proto_rawDesc = nil
	file_wish_category_proto_goTypes = nil
	file_wish_category_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WishCategoryServiceClient is the client API for WishCategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WishCategoryServiceClient interface {
	// やりたいことカテゴリーの新規作成
	CreateWishCategory(ctx context.Context, in *CreateWishCategoryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// やりたいことカテゴリー名更新
	UpdateWishCategoryName(ctx context.Context, in *UpdateWishCategoryNameRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type wishCategoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWishCategoryServiceClient(cc grpc.ClientConnInterface) WishCategoryServiceClient {
	return &wishCategoryServiceClient{cc}
}

func (c *wishCategoryServiceClient) CreateWishCategory(ctx context.Context, in *CreateWishCategoryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto_wishcategory.WishCategoryService/CreateWishCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishCategoryServiceClient) UpdateWishCategoryName(ctx context.Context, in *UpdateWishCategoryNameRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto_wishcategory.WishCategoryService/UpdateWishCategoryName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WishCategoryServiceServer is the server API for WishCategoryService service.
type WishCategoryServiceServer interface {
	// やりたいことカテゴリーの新規作成
	CreateWishCategory(context.Context, *CreateWishCategoryRequest) (*empty.Empty, error)
	// やりたいことカテゴリー名更新
	UpdateWishCategoryName(context.Context, *UpdateWishCategoryNameRequest) (*empty.Empty, error)
}

// UnimplementedWishCategoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWishCategoryServiceServer struct {
}

func (*UnimplementedWishCategoryServiceServer) CreateWishCategory(context.Context, *CreateWishCategoryRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWishCategory not implemented")
}
func (*UnimplementedWishCategoryServiceServer) UpdateWishCategoryName(context.Context, *UpdateWishCategoryNameRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWishCategoryName not implemented")
}

func RegisterWishCategoryServiceServer(s *grpc.Server, srv WishCategoryServiceServer) {
	s.RegisterService(&_WishCategoryService_serviceDesc, srv)
}

func _WishCategoryService_CreateWishCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWishCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishCategoryServiceServer).CreateWishCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_wishcategory.WishCategoryService/CreateWishCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishCategoryServiceServer).CreateWishCategory(ctx, req.(*CreateWishCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WishCategoryService_UpdateWishCategoryName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWishCategoryNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishCategoryServiceServer).UpdateWishCategoryName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_wishcategory.WishCategoryService/UpdateWishCategoryName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishCategoryServiceServer).UpdateWishCategoryName(ctx, req.(*UpdateWishCategoryNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WishCategoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_wishcategory.WishCategoryService",
	HandlerType: (*WishCategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWishCategory",
			Handler:    _WishCategoryService_CreateWishCategory_Handler,
		},
		{
			MethodName: "UpdateWishCategoryName",
			Handler:    _WishCategoryService_UpdateWishCategoryName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wish_category.proto",
}
