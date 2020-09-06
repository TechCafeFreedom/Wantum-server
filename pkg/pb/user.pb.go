// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.4
// source: user.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// 新規ユーザ作成リクエスト
type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 表示名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ユーザ名
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// アイコンのURL
	Thumbnail string `protobuf:"bytes,3,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	// ひとこと
	Bio string `protobuf:"bytes,4,opt,name=bio,proto3" json:"bio,omitempty"`
	// 性別（1: man, 2: woman, 3: unknown）
	Gender int32 `protobuf:"varint,5,opt,name=gender,proto3" json:"gender,omitempty"`
	// 電話番号
	Phone string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	// 現在地的な。Twitterの真似
	Place string `protobuf:"bytes,7,opt,name=place,proto3" json:"place,omitempty"`
	// 生年月日(UNIX)
	Birth int64 `protobuf:"varint,8,opt,name=birth,proto3" json:"birth,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *CreateUserRequest) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *CreateUserRequest) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *CreateUserRequest) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *CreateUserRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateUserRequest) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *CreateUserRequest) GetBirth() int64 {
	if x != nil {
		return x.Birth
	}
	return 0
}

// ユーザー情報取得リクエスト（ユーザーを特定するためのID情報）
type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ユーザー名
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

// ユーザー情報更新リクエスト
type UpdateUserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 更新したいユーザーのID
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 表示名
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// ユーザ名
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// アイコンのURL
	Thumbnail string `protobuf:"bytes,4,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	// ひとこと
	Bio string `protobuf:"bytes,5,opt,name=bio,proto3" json:"bio,omitempty"`
	// 性別（1: man, 2: woman, 3: unknown）
	Gender int32 `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`
	// 電話番号
	Phone string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	// 現在地的な。Twitterの真似
	Place string `protobuf:"bytes,8,opt,name=place,proto3" json:"place,omitempty"`
	// 生年月日(UNIX)
	Birth int64 `protobuf:"varint,9,opt,name=birth,proto3" json:"birth,omitempty"`
}

func (x *UpdateUserProfileRequest) Reset() {
	*x = UpdateUserProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserProfileRequest) ProtoMessage() {}

func (x *UpdateUserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserProfileRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserProfileRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateUserProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserProfileRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UpdateUserProfileRequest) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *UpdateUserProfileRequest) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *UpdateUserProfileRequest) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UpdateUserProfileRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateUserProfileRequest) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *UpdateUserProfileRequest) GetBirth() int64 {
	if x != nil {
		return x.Birth
	}
	return 0
}

// ユーザ情報
type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ユーザーID
	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// 表示名
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// ユーザ名
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// アイコンのURL
	Thumbnail string `protobuf:"bytes,4,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	// ひとこと
	Bio string `protobuf:"bytes,5,opt,name=bio,proto3" json:"bio,omitempty"`
	// 性別（1: man, 2: woman, 3: unknown）
	Gender int32 `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`
	// 電話番号
	Phone string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	// 現在地的な。Twitterの真似
	Place string `protobuf:"bytes,8,opt,name=place,proto3" json:"place,omitempty"`
	// 生年月日(UNIX)
	Birth int64 `protobuf:"varint,9,opt,name=birth,proto3" json:"birth,omitempty"`
	// ユーザーが過去に投稿した思い出一覧
	Memories *Memories `protobuf:"bytes,10,opt,name=memories,proto3" json:"memories,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserInfo) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserInfo) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *UserInfo) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *UserInfo) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *UserInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserInfo) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *UserInfo) GetBirth() int64 {
	if x != nil {
		return x.Birth
	}
	return 0
}

func (x *UserInfo) GetMemories() *Memories {
	if x != nil {
		return x.Memories
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x22, 0x2d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0xee, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x22, 0x92, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x62, 0x69, 0x72, 0x74, 0x68, 0x12, 0x32, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x32, 0xab, 0x02, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x3e,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),        // 0: proto_user.CreateUserRequest
	(*GetUserRequest)(nil),           // 1: proto_user.GetUserRequest
	(*UpdateUserProfileRequest)(nil), // 2: proto_user.UpdateUserProfileRequest
	(*UserInfo)(nil),                 // 3: proto_user.UserInfo
	(*Memories)(nil),                 // 4: proto_memory.Memories
	(*empty.Empty)(nil),              // 5: google.protobuf.Empty
}
var file_user_proto_depIdxs = []int32{
	4, // 0: proto_user.UserInfo.memories:type_name -> proto_memory.Memories
	0, // 1: proto_user.UserService.CreateUser:input_type -> proto_user.CreateUserRequest
	5, // 2: proto_user.UserService.GetMyProfile:input_type -> google.protobuf.Empty
	1, // 3: proto_user.UserService.GetUserProfile:input_type -> proto_user.GetUserRequest
	2, // 4: proto_user.UserService.UpdateUserProfile:input_type -> proto_user.UpdateUserProfileRequest
	3, // 5: proto_user.UserService.CreateUser:output_type -> proto_user.UserInfo
	3, // 6: proto_user.UserService.GetMyProfile:output_type -> proto_user.UserInfo
	3, // 7: proto_user.UserService.GetUserProfile:output_type -> proto_user.UserInfo
	3, // 8: proto_user.UserService.UpdateUserProfile:output_type -> proto_user.UserInfo
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_memory_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserProfileRequest); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	// 新規ユーザ作成
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserInfo, error)
	// ユーザー情報取得（マイページ）
	GetMyProfile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserInfo, error)
	// ユーザー情報取得（他の人のページ）
	GetUserProfile(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserInfo, error)
	// ユーザー情報更新
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UserInfo, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/proto_user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetMyProfile(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/proto_user.UserService/GetMyProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserProfile(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/proto_user.UserService/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/proto_user.UserService/UpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	// 新規ユーザ作成
	CreateUser(context.Context, *CreateUserRequest) (*UserInfo, error)
	// ユーザー情報取得（マイページ）
	GetMyProfile(context.Context, *empty.Empty) (*UserInfo, error)
	// ユーザー情報取得（他の人のページ）
	GetUserProfile(context.Context, *GetUserRequest) (*UserInfo, error)
	// ユーザー情報更新
	UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UserInfo, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserServiceServer) GetMyProfile(context.Context, *empty.Empty) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyProfile not implemented")
}
func (*UnimplementedUserServiceServer) GetUserProfile(context.Context, *GetUserRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetMyProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetMyProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_user.UserService/GetMyProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetMyProfile(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_user.UserService/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserProfile(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_user.UserService/UpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserProfile(ctx, req.(*UpdateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetMyProfile",
			Handler:    _UserService_GetMyProfile_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _UserService_GetUserProfile_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserService_UpdateUserProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
