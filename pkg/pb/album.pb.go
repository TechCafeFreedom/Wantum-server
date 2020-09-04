// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.4
// source: album.proto

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

// アルバム作成リクエスト
type CreateAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// アルバム名
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateAlbumRequest) Reset() {
	*x = CreateAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumRequest) ProtoMessage() {}

func (x *CreateAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumRequest.ProtoReflect.Descriptor instead.
func (*CreateAlbumRequest) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAlbumRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// アルバム1件取得用リクエスト
type GetAlbumInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// アルバムを特定するためのID
	AlbumId int64 `protobuf:"varint,1,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
}

func (x *GetAlbumInfoRequest) Reset() {
	*x = GetAlbumInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumInfoRequest) ProtoMessage() {}

func (x *GetAlbumInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumInfoRequest.ProtoReflect.Descriptor instead.
func (*GetAlbumInfoRequest) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{1}
}

func (x *GetAlbumInfoRequest) GetAlbumId() int64 {
	if x != nil {
		return x.AlbumId
	}
	return 0
}

// アルバム一覧（ID, title)取得用リクエスト
type GetAlbumsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 取得開始位置を知らせるためのアルバムID
	AlbumId int64 `protobuf:"varint,1,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
	// 最大取得件数
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAlbumsRequest) Reset() {
	*x = GetAlbumsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumsRequest) ProtoMessage() {}

func (x *GetAlbumsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumsRequest.ProtoReflect.Descriptor instead.
func (*GetAlbumsRequest) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{2}
}

func (x *GetAlbumsRequest) GetAlbumId() int64 {
	if x != nil {
		return x.AlbumId
	}
	return 0
}

func (x *GetAlbumsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// アルバム名変更用リクエスト
type ChangeAlbumTitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 変更したいアルバムのID
	AlbumId int64 `protobuf:"varint,1,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
	// 変更したいアルバム名
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *ChangeAlbumTitleRequest) Reset() {
	*x = ChangeAlbumTitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAlbumTitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAlbumTitleRequest) ProtoMessage() {}

func (x *ChangeAlbumTitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeAlbumTitleRequest.ProtoReflect.Descriptor instead.
func (*ChangeAlbumTitleRequest) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeAlbumTitleRequest) GetAlbumId() int64 {
	if x != nil {
		return x.AlbumId
	}
	return 0
}

func (x *ChangeAlbumTitleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// アルバム一覧（配列）
type Albums struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// アルバム一覧（配列）
	Albums []*AlbumInfo `protobuf:"bytes,1,rep,name=albums,proto3" json:"albums,omitempty"`
}

func (x *Albums) Reset() {
	*x = Albums{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Albums) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Albums) ProtoMessage() {}

func (x *Albums) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Albums.ProtoReflect.Descriptor instead.
func (*Albums) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{4}
}

func (x *Albums) GetAlbums() []*AlbumInfo {
	if x != nil {
		return x.Albums
	}
	return nil
}

// アルバム情報
type AlbumInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// アルバムを特定するためのID
	AlbumId int64 `protobuf:"varint,1,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
	// アルバムのタイトル
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// アルバムに保存されている思い出一覧（配列）
	Memories *Memories `protobuf:"bytes,3,opt,name=memories,proto3" json:"memories,omitempty"`
}

func (x *AlbumInfo) Reset() {
	*x = AlbumInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumInfo) ProtoMessage() {}

func (x *AlbumInfo) ProtoReflect() protoreflect.Message {
	mi := &file_album_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumInfo.ProtoReflect.Descriptor instead.
func (*AlbumInfo) Descriptor() ([]byte, []int) {
	return file_album_proto_rawDescGZIP(), []int{5}
}

func (x *AlbumInfo) GetAlbumId() int64 {
	if x != nil {
		return x.AlbumId
	}
	return 0
}

func (x *AlbumInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AlbumInfo) GetMemories() *Memories {
	if x != nil {
		return x.Memories
	}
	return nil
}

var File_album_proto protoreflect.FileDescriptor

var file_album_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x1a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x22, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x4a, 0x0a, 0x17, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x38, 0x0a, 0x06, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x73, 0x12, 0x2e, 0x0a, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x73, 0x22, 0x70, 0x0a, 0x09, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x32, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x32, 0xbd, 0x02, 0x0a, 0x0c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x6c, 0x62, 0x75,
	0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x4a,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x4d, 0x79, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x22, 0x00, 0x12,
	0x52, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x6c, 0x62, 0x75,
	0x6d, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_album_proto_rawDescOnce sync.Once
	file_album_proto_rawDescData = file_album_proto_rawDesc
)

func file_album_proto_rawDescGZIP() []byte {
	file_album_proto_rawDescOnce.Do(func() {
		file_album_proto_rawDescData = protoimpl.X.CompressGZIP(file_album_proto_rawDescData)
	})
	return file_album_proto_rawDescData
}

var file_album_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_album_proto_goTypes = []interface{}{
	(*CreateAlbumRequest)(nil),      // 0: proto_album.CreateAlbumRequest
	(*GetAlbumInfoRequest)(nil),     // 1: proto_album.GetAlbumInfoRequest
	(*GetAlbumsRequest)(nil),        // 2: proto_album.GetAlbumsRequest
	(*ChangeAlbumTitleRequest)(nil), // 3: proto_album.ChangeAlbumTitleRequest
	(*Albums)(nil),                  // 4: proto_album.Albums
	(*AlbumInfo)(nil),               // 5: proto_album.AlbumInfo
	(*Memories)(nil),                // 6: proto_memory.Memories
}
var file_album_proto_depIdxs = []int32{
	5, // 0: proto_album.Albums.albums:type_name -> proto_album.AlbumInfo
	6, // 1: proto_album.AlbumInfo.memories:type_name -> proto_memory.Memories
	0, // 2: proto_album.AlbumService.CreateAlbum:input_type -> proto_album.CreateAlbumRequest
	1, // 3: proto_album.AlbumService.GetAlbumInfo:input_type -> proto_album.GetAlbumInfoRequest
	2, // 4: proto_album.AlbumService.GetMyAlbums:input_type -> proto_album.GetAlbumsRequest
	3, // 5: proto_album.AlbumService.ChangeAlbumTitle:input_type -> proto_album.ChangeAlbumTitleRequest
	5, // 6: proto_album.AlbumService.CreateAlbum:output_type -> proto_album.AlbumInfo
	5, // 7: proto_album.AlbumService.GetAlbumInfo:output_type -> proto_album.AlbumInfo
	4, // 8: proto_album.AlbumService.GetMyAlbums:output_type -> proto_album.Albums
	5, // 9: proto_album.AlbumService.ChangeAlbumTitle:output_type -> proto_album.AlbumInfo
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_album_proto_init() }
func file_album_proto_init() {
	if File_album_proto != nil {
		return
	}
	file_memory_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_album_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumRequest); i {
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
		file_album_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumInfoRequest); i {
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
		file_album_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumsRequest); i {
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
		file_album_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeAlbumTitleRequest); i {
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
		file_album_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Albums); i {
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
		file_album_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumInfo); i {
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
			RawDescriptor: file_album_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_album_proto_goTypes,
		DependencyIndexes: file_album_proto_depIdxs,
		MessageInfos:      file_album_proto_msgTypes,
	}.Build()
	File_album_proto = out.File
	file_album_proto_rawDesc = nil
	file_album_proto_goTypes = nil
	file_album_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AlbumServiceClient is the client API for AlbumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlbumServiceClient interface {
	// アルバムの新規作成
	CreateAlbum(ctx context.Context, in *CreateAlbumRequest, opts ...grpc.CallOption) (*AlbumInfo, error)
	// IDをもとにアルバム1件を取得
	GetAlbumInfo(ctx context.Context, in *GetAlbumInfoRequest, opts ...grpc.CallOption) (*AlbumInfo, error)
	// アルバム一覧情報の取得
	GetMyAlbums(ctx context.Context, in *GetAlbumsRequest, opts ...grpc.CallOption) (*Albums, error)
	// アルバム名の変更
	ChangeAlbumTitle(ctx context.Context, in *ChangeAlbumTitleRequest, opts ...grpc.CallOption) (*AlbumInfo, error)
}

type albumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlbumServiceClient(cc grpc.ClientConnInterface) AlbumServiceClient {
	return &albumServiceClient{cc}
}

func (c *albumServiceClient) CreateAlbum(ctx context.Context, in *CreateAlbumRequest, opts ...grpc.CallOption) (*AlbumInfo, error) {
	out := new(AlbumInfo)
	err := c.cc.Invoke(ctx, "/proto_album.AlbumService/CreateAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *albumServiceClient) GetAlbumInfo(ctx context.Context, in *GetAlbumInfoRequest, opts ...grpc.CallOption) (*AlbumInfo, error) {
	out := new(AlbumInfo)
	err := c.cc.Invoke(ctx, "/proto_album.AlbumService/GetAlbumInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *albumServiceClient) GetMyAlbums(ctx context.Context, in *GetAlbumsRequest, opts ...grpc.CallOption) (*Albums, error) {
	out := new(Albums)
	err := c.cc.Invoke(ctx, "/proto_album.AlbumService/GetMyAlbums", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *albumServiceClient) ChangeAlbumTitle(ctx context.Context, in *ChangeAlbumTitleRequest, opts ...grpc.CallOption) (*AlbumInfo, error) {
	out := new(AlbumInfo)
	err := c.cc.Invoke(ctx, "/proto_album.AlbumService/ChangeAlbumTitle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlbumServiceServer is the server API for AlbumService service.
type AlbumServiceServer interface {
	// アルバムの新規作成
	CreateAlbum(context.Context, *CreateAlbumRequest) (*AlbumInfo, error)
	// IDをもとにアルバム1件を取得
	GetAlbumInfo(context.Context, *GetAlbumInfoRequest) (*AlbumInfo, error)
	// アルバム一覧情報の取得
	GetMyAlbums(context.Context, *GetAlbumsRequest) (*Albums, error)
	// アルバム名の変更
	ChangeAlbumTitle(context.Context, *ChangeAlbumTitleRequest) (*AlbumInfo, error)
}

// UnimplementedAlbumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAlbumServiceServer struct {
}

func (*UnimplementedAlbumServiceServer) CreateAlbum(context.Context, *CreateAlbumRequest) (*AlbumInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlbum not implemented")
}
func (*UnimplementedAlbumServiceServer) GetAlbumInfo(context.Context, *GetAlbumInfoRequest) (*AlbumInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlbumInfo not implemented")
}
func (*UnimplementedAlbumServiceServer) GetMyAlbums(context.Context, *GetAlbumsRequest) (*Albums, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyAlbums not implemented")
}
func (*UnimplementedAlbumServiceServer) ChangeAlbumTitle(context.Context, *ChangeAlbumTitleRequest) (*AlbumInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAlbumTitle not implemented")
}

func RegisterAlbumServiceServer(s *grpc.Server, srv AlbumServiceServer) {
	s.RegisterService(&_AlbumService_serviceDesc, srv)
}

func _AlbumService_CreateAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlbumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlbumServiceServer).CreateAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_album.AlbumService/CreateAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlbumServiceServer).CreateAlbum(ctx, req.(*CreateAlbumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlbumService_GetAlbumInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlbumInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlbumServiceServer).GetAlbumInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_album.AlbumService/GetAlbumInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlbumServiceServer).GetAlbumInfo(ctx, req.(*GetAlbumInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlbumService_GetMyAlbums_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlbumsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlbumServiceServer).GetMyAlbums(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_album.AlbumService/GetMyAlbums",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlbumServiceServer).GetMyAlbums(ctx, req.(*GetAlbumsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlbumService_ChangeAlbumTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeAlbumTitleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlbumServiceServer).ChangeAlbumTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_album.AlbumService/ChangeAlbumTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlbumServiceServer).ChangeAlbumTitle(ctx, req.(*ChangeAlbumTitleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlbumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_album.AlbumService",
	HandlerType: (*AlbumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAlbum",
			Handler:    _AlbumService_CreateAlbum_Handler,
		},
		{
			MethodName: "GetAlbumInfo",
			Handler:    _AlbumService_GetAlbumInfo_Handler,
		},
		{
			MethodName: "GetMyAlbums",
			Handler:    _AlbumService_GetMyAlbums_Handler,
		},
		{
			MethodName: "ChangeAlbumTitle",
			Handler:    _AlbumService_ChangeAlbumTitle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "album.proto",
}
