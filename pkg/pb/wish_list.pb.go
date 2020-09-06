// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.4
// source: wish_list.proto

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

// やりたいことリスト新規作成リクエスト
type CreateWishListInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// やりたいことリスト名
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *CreateWishListInfoRequest) Reset() {
	*x = CreateWishListInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wish_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWishListInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWishListInfoRequest) ProtoMessage() {}

func (x *CreateWishListInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wish_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWishListInfoRequest.ProtoReflect.Descriptor instead.
func (*CreateWishListInfoRequest) Descriptor() ([]byte, []int) {
	return file_wish_list_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWishListInfoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

// やりたいことリスト（ID, titleのみ）の配列を取得するためのリクエスト
type GetWishListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 取得開始位置を知らせるためのやりたいことリストID
	WishListId int64 `protobuf:"varint,1,opt,name=wish_list_id,json=wishListId,proto3" json:"wish_list_id,omitempty"`
	// 最大取得件数
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetWishListRequest) Reset() {
	*x = GetWishListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wish_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWishListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWishListRequest) ProtoMessage() {}

func (x *GetWishListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wish_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWishListRequest.ProtoReflect.Descriptor instead.
func (*GetWishListRequest) Descriptor() ([]byte, []int) {
	return file_wish_list_proto_rawDescGZIP(), []int{1}
}

func (x *GetWishListRequest) GetWishListId() int64 {
	if x != nil {
		return x.WishListId
	}
	return 0
}

func (x *GetWishListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// IDをもとにやりたいことリストを1件取得するためのリクエスト
type GetWishListInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// やりたいことリストを特定するためのID
	WishListId int64 `protobuf:"varint,1,opt,name=wish_list_id,json=wishListId,proto3" json:"wish_list_id,omitempty"`
}

func (x *GetWishListInfoRequest) Reset() {
	*x = GetWishListInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wish_list_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWishListInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWishListInfoRequest) ProtoMessage() {}

func (x *GetWishListInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wish_list_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWishListInfoRequest.ProtoReflect.Descriptor instead.
func (*GetWishListInfoRequest) Descriptor() ([]byte, []int) {
	return file_wish_list_proto_rawDescGZIP(), []int{2}
}

func (x *GetWishListInfoRequest) GetWishListId() int64 {
	if x != nil {
		return x.WishListId
	}
	return 0
}

// やりたいことリスト削除用リクエスト
type DeleteWishListInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 削除したいやりたいことリストのID
	WishListId int64 `protobuf:"varint,1,opt,name=wish_list_id,json=wishListId,proto3" json:"wish_list_id,omitempty"`
}

func (x *DeleteWishListInfoRequest) Reset() {
	*x = DeleteWishListInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wish_list_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWishListInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWishListInfoRequest) ProtoMessage() {}

func (x *DeleteWishListInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wish_list_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWishListInfoRequest.ProtoReflect.Descriptor instead.
func (*DeleteWishListInfoRequest) Descriptor() ([]byte, []int) {
	return file_wish_list_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteWishListInfoRequest) GetWishListId() int64 {
	if x != nil {
		return x.WishListId
	}
	return 0
}

// やりたいことリストの一覧（配列）
type WishList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// やりたいことリスト一覧（配列）
	WishList []*WishListInfo `protobuf:"bytes,1,rep,name=wish_list,json=wishList,proto3" json:"wish_list,omitempty"`
}

func (x *WishList) Reset() {
	*x = WishList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wish_list_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WishList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishList) ProtoMessage() {}

func (x *WishList) ProtoReflect() protoreflect.Message {
	mi := &file_wish_list_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishList.ProtoReflect.Descriptor instead.
func (*WishList) Descriptor() ([]byte, []int) {
	return file_wish_list_proto_rawDescGZIP(), []int{4}
}

func (x *WishList) GetWishList() []*WishListInfo {
	if x != nil {
		return x.WishList
	}
	return nil
}

// やりたいことリスト情報
type WishListInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// やりたいことリストを特定するためのID
	WishListId int64 `protobuf:"varint,1,opt,name=wish_list_id,json=wishListId,proto3" json:"wish_list_id,omitempty"`
	// やりたいことリストタイトル
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// やりたいことカテゴリー一覧（配列）
	WishCategories []*WishCategoryInfo `protobuf:"bytes,3,rep,name=wish_categories,json=wishCategories,proto3" json:"wish_categories,omitempty"`
	// 編集権限を持っているユーザー一覧（配列）
	Authors []*UserInfo `protobuf:"bytes,4,rep,name=authors,proto3" json:"authors,omitempty"`
	// 招待リンク
	InviteUrl string `protobuf:"bytes,5,opt,name=invite_url,json=inviteUrl,proto3" json:"invite_url,omitempty"`
}

func (x *WishListInfo) Reset() {
	*x = WishListInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wish_list_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WishListInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishListInfo) ProtoMessage() {}

func (x *WishListInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wish_list_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishListInfo.ProtoReflect.Descriptor instead.
func (*WishListInfo) Descriptor() ([]byte, []int) {
	return file_wish_list_proto_rawDescGZIP(), []int{5}
}

func (x *WishListInfo) GetWishListId() int64 {
	if x != nil {
		return x.WishListId
	}
	return 0
}

func (x *WishListInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *WishListInfo) GetWishCategories() []*WishCategoryInfo {
	if x != nil {
		return x.WishCategories
	}
	return nil
}

func (x *WishListInfo) GetAuthors() []*UserInfo {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *WishListInfo) GetInviteUrl() string {
	if x != nil {
		return x.InviteUrl
	}
	return ""
}

// やりたいことカテゴリー情報
type WishCategoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// やりたいことカテゴリーを特定するためのID
	WishCategoryId int64 `protobuf:"varint,1,opt,name=wish_category_id,json=wishCategoryId,proto3" json:"wish_category_id,omitempty"`
	// やりたいことカテゴリー名
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// カテゴリー内のやりたいことカード一覧（配列）
	WishCards []*WishCardInfo `protobuf:"bytes,3,rep,name=wish_cards,json=wishCards,proto3" json:"wish_cards,omitempty"`
}

func (x *WishCategoryInfo) Reset() {
	*x = WishCategoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wish_list_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WishCategoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishCategoryInfo) ProtoMessage() {}

func (x *WishCategoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wish_list_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishCategoryInfo.ProtoReflect.Descriptor instead.
func (*WishCategoryInfo) Descriptor() ([]byte, []int) {
	return file_wish_list_proto_rawDescGZIP(), []int{6}
}

func (x *WishCategoryInfo) GetWishCategoryId() int64 {
	if x != nil {
		return x.WishCategoryId
	}
	return 0
}

func (x *WishCategoryInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *WishCategoryInfo) GetWishCards() []*WishCardInfo {
	if x != nil {
		return x.WishCards
	}
	return nil
}

// やりたいことカード情報
type WishCardInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// wish_cardを特定するためのID
	WishCardId int64 `protobuf:"varint,1,opt,name=wish_card_id,json=wishCardId,proto3" json:"wish_card_id,omitempty"`
	// 何をしたいのか
	Activity string `protobuf:"bytes,2,opt,name=activity,proto3" json:"activity,omitempty"`
	// やりたいことの説明
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// いつやりたいか（UNIX）
	Date int64 `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
	// 完了日時
	DoneAt string `protobuf:"bytes,5,opt,name=done_at,json=doneAt,proto3" json:"done_at,omitempty"`
	// どこでそれをしたいのか
	Place string `protobuf:"bytes,6,opt,name=place,proto3" json:"place,omitempty"`
	// タグ一覧（配列）
	Tags []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *WishCardInfo) Reset() {
	*x = WishCardInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wish_list_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WishCardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WishCardInfo) ProtoMessage() {}

func (x *WishCardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wish_list_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WishCardInfo.ProtoReflect.Descriptor instead.
func (*WishCardInfo) Descriptor() ([]byte, []int) {
	return file_wish_list_proto_rawDescGZIP(), []int{7}
}

func (x *WishCardInfo) GetWishCardId() int64 {
	if x != nil {
		return x.WishCardId
	}
	return 0
}

func (x *WishCardInfo) GetActivity() string {
	if x != nil {
		return x.Activity
	}
	return ""
}

func (x *WishCardInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *WishCardInfo) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *WishCardInfo) GetDoneAt() string {
	if x != nil {
		return x.DoneAt
	}
	return ""
}

func (x *WishCardInfo) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *WishCardInfo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_wish_list_proto protoreflect.FileDescriptor

var file_wish_list_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73,
	0x74, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x19, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x4c, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x3a, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x77, 0x69, 0x73,
	0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x77, 0x69, 0x73, 0x68,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x22, 0x45, 0x0a, 0x08, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69,
	0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x77, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xe0, 0x01,
	0x0a, 0x0c, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20,
	0x0a, 0x0c, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x57, 0x69, 0x73, 0x68, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0e, 0x77, 0x69, 0x73, 0x68, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x55, 0x72, 0x6c,
	0x22, 0x8f, 0x01, 0x0a, 0x10, 0x57, 0x69, 0x73, 0x68, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x77, 0x69, 0x73, 0x68, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x73, 0x68, 0x43,
	0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x77, 0x69, 0x73, 0x68, 0x43, 0x61, 0x72,
	0x64, 0x73, 0x22, 0xc5, 0x01, 0x0a, 0x0c, 0x57, 0x69, 0x73, 0x68, 0x43, 0x61, 0x72, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x77, 0x69, 0x73, 0x68, 0x5f, 0x63, 0x61, 0x72, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x77, 0x69, 0x73, 0x68, 0x43,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x6f, 0x6e, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6e, 0x65, 0x41, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x32, 0xc3, 0x03, 0x0a, 0x0f, 0x57,
	0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69, 0x73,
	0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74,
	0x2e, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x57,
	0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e,
	0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69,
	0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x57,
	0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57,
	0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x77, 0x69, 0x73, 0x68, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x57, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wish_list_proto_rawDescOnce sync.Once
	file_wish_list_proto_rawDescData = file_wish_list_proto_rawDesc
)

func file_wish_list_proto_rawDescGZIP() []byte {
	file_wish_list_proto_rawDescOnce.Do(func() {
		file_wish_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_wish_list_proto_rawDescData)
	})
	return file_wish_list_proto_rawDescData
}

var file_wish_list_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_wish_list_proto_goTypes = []interface{}{
	(*CreateWishListInfoRequest)(nil), // 0: proto_wishlist.CreateWishListInfoRequest
	(*GetWishListRequest)(nil),        // 1: proto_wishlist.GetWishListRequest
	(*GetWishListInfoRequest)(nil),    // 2: proto_wishlist.GetWishListInfoRequest
	(*DeleteWishListInfoRequest)(nil), // 3: proto_wishlist.DeleteWishListInfoRequest
	(*WishList)(nil),                  // 4: proto_wishlist.WishList
	(*WishListInfo)(nil),              // 5: proto_wishlist.WishListInfo
	(*WishCategoryInfo)(nil),          // 6: proto_wishlist.WishCategoryInfo
	(*WishCardInfo)(nil),              // 7: proto_wishlist.WishCardInfo
	(*UserInfo)(nil),                  // 8: proto_user.UserInfo
	(*empty.Empty)(nil),               // 9: google.protobuf.Empty
}
var file_wish_list_proto_depIdxs = []int32{
	5, // 0: proto_wishlist.WishList.wish_list:type_name -> proto_wishlist.WishListInfo
	6, // 1: proto_wishlist.WishListInfo.wish_categories:type_name -> proto_wishlist.WishCategoryInfo
	8, // 2: proto_wishlist.WishListInfo.authors:type_name -> proto_user.UserInfo
	7, // 3: proto_wishlist.WishCategoryInfo.wish_cards:type_name -> proto_wishlist.WishCardInfo
	0, // 4: proto_wishlist.WishListService.CreateWishListInfo:input_type -> proto_wishlist.CreateWishListInfoRequest
	1, // 5: proto_wishlist.WishListService.GetWishList:input_type -> proto_wishlist.GetWishListRequest
	2, // 6: proto_wishlist.WishListService.GetWishListInfo:input_type -> proto_wishlist.GetWishListInfoRequest
	5, // 7: proto_wishlist.WishListService.UpdateWishListInfo:input_type -> proto_wishlist.WishListInfo
	3, // 8: proto_wishlist.WishListService.DeleteWishListInfo:input_type -> proto_wishlist.DeleteWishListInfoRequest
	4, // 9: proto_wishlist.WishListService.CreateWishListInfo:output_type -> proto_wishlist.WishList
	4, // 10: proto_wishlist.WishListService.GetWishList:output_type -> proto_wishlist.WishList
	5, // 11: proto_wishlist.WishListService.GetWishListInfo:output_type -> proto_wishlist.WishListInfo
	9, // 12: proto_wishlist.WishListService.UpdateWishListInfo:output_type -> google.protobuf.Empty
	9, // 13: proto_wishlist.WishListService.DeleteWishListInfo:output_type -> google.protobuf.Empty
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_wish_list_proto_init() }
func file_wish_list_proto_init() {
	if File_wish_list_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wish_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWishListInfoRequest); i {
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
		file_wish_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWishListRequest); i {
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
		file_wish_list_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWishListInfoRequest); i {
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
		file_wish_list_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWishListInfoRequest); i {
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
		file_wish_list_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WishList); i {
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
		file_wish_list_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WishListInfo); i {
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
		file_wish_list_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WishCategoryInfo); i {
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
		file_wish_list_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WishCardInfo); i {
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
			RawDescriptor: file_wish_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wish_list_proto_goTypes,
		DependencyIndexes: file_wish_list_proto_depIdxs,
		MessageInfos:      file_wish_list_proto_msgTypes,
	}.Build()
	File_wish_list_proto = out.File
	file_wish_list_proto_rawDesc = nil
	file_wish_list_proto_goTypes = nil
	file_wish_list_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WishListServiceClient is the client API for WishListService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WishListServiceClient interface {
	// やりたいことリストの新規作成
	CreateWishListInfo(ctx context.Context, in *CreateWishListInfoRequest, opts ...grpc.CallOption) (*WishList, error)
	// やりたいことリスト一覧取得
	GetWishList(ctx context.Context, in *GetWishListRequest, opts ...grpc.CallOption) (*WishList, error)
	// やりたいことリスト単体取得
	GetWishListInfo(ctx context.Context, in *GetWishListInfoRequest, opts ...grpc.CallOption) (WishListService_GetWishListInfoClient, error)
	// やりたいことリスト更新（カードの移動やタイトルの変更など：ユーザーアクション）
	UpdateWishListInfo(ctx context.Context, in *WishListInfo, opts ...grpc.CallOption) (*empty.Empty, error)
	// やりたいことリスト自体の削除
	DeleteWishListInfo(ctx context.Context, in *DeleteWishListInfoRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type wishListServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWishListServiceClient(cc grpc.ClientConnInterface) WishListServiceClient {
	return &wishListServiceClient{cc}
}

func (c *wishListServiceClient) CreateWishListInfo(ctx context.Context, in *CreateWishListInfoRequest, opts ...grpc.CallOption) (*WishList, error) {
	out := new(WishList)
	err := c.cc.Invoke(ctx, "/proto_wishlist.WishListService/CreateWishListInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishListServiceClient) GetWishList(ctx context.Context, in *GetWishListRequest, opts ...grpc.CallOption) (*WishList, error) {
	out := new(WishList)
	err := c.cc.Invoke(ctx, "/proto_wishlist.WishListService/GetWishList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishListServiceClient) GetWishListInfo(ctx context.Context, in *GetWishListInfoRequest, opts ...grpc.CallOption) (WishListService_GetWishListInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WishListService_serviceDesc.Streams[0], "/proto_wishlist.WishListService/GetWishListInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &wishListServiceGetWishListInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WishListService_GetWishListInfoClient interface {
	Recv() (*WishListInfo, error)
	grpc.ClientStream
}

type wishListServiceGetWishListInfoClient struct {
	grpc.ClientStream
}

func (x *wishListServiceGetWishListInfoClient) Recv() (*WishListInfo, error) {
	m := new(WishListInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *wishListServiceClient) UpdateWishListInfo(ctx context.Context, in *WishListInfo, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto_wishlist.WishListService/UpdateWishListInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishListServiceClient) DeleteWishListInfo(ctx context.Context, in *DeleteWishListInfoRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto_wishlist.WishListService/DeleteWishListInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WishListServiceServer is the server API for WishListService service.
type WishListServiceServer interface {
	// やりたいことリストの新規作成
	CreateWishListInfo(context.Context, *CreateWishListInfoRequest) (*WishList, error)
	// やりたいことリスト一覧取得
	GetWishList(context.Context, *GetWishListRequest) (*WishList, error)
	// やりたいことリスト単体取得
	GetWishListInfo(*GetWishListInfoRequest, WishListService_GetWishListInfoServer) error
	// やりたいことリスト更新（カードの移動やタイトルの変更など：ユーザーアクション）
	UpdateWishListInfo(context.Context, *WishListInfo) (*empty.Empty, error)
	// やりたいことリスト自体の削除
	DeleteWishListInfo(context.Context, *DeleteWishListInfoRequest) (*empty.Empty, error)
}

// UnimplementedWishListServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWishListServiceServer struct {
}

func (*UnimplementedWishListServiceServer) CreateWishListInfo(context.Context, *CreateWishListInfoRequest) (*WishList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWishListInfo not implemented")
}
func (*UnimplementedWishListServiceServer) GetWishList(context.Context, *GetWishListRequest) (*WishList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWishList not implemented")
}
func (*UnimplementedWishListServiceServer) GetWishListInfo(*GetWishListInfoRequest, WishListService_GetWishListInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetWishListInfo not implemented")
}
func (*UnimplementedWishListServiceServer) UpdateWishListInfo(context.Context, *WishListInfo) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWishListInfo not implemented")
}
func (*UnimplementedWishListServiceServer) DeleteWishListInfo(context.Context, *DeleteWishListInfoRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWishListInfo not implemented")
}

func RegisterWishListServiceServer(s *grpc.Server, srv WishListServiceServer) {
	s.RegisterService(&_WishListService_serviceDesc, srv)
}

func _WishListService_CreateWishListInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWishListInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishListServiceServer).CreateWishListInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_wishlist.WishListService/CreateWishListInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishListServiceServer).CreateWishListInfo(ctx, req.(*CreateWishListInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WishListService_GetWishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWishListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishListServiceServer).GetWishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_wishlist.WishListService/GetWishList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishListServiceServer).GetWishList(ctx, req.(*GetWishListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WishListService_GetWishListInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetWishListInfoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WishListServiceServer).GetWishListInfo(m, &wishListServiceGetWishListInfoServer{stream})
}

type WishListService_GetWishListInfoServer interface {
	Send(*WishListInfo) error
	grpc.ServerStream
}

type wishListServiceGetWishListInfoServer struct {
	grpc.ServerStream
}

func (x *wishListServiceGetWishListInfoServer) Send(m *WishListInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _WishListService_UpdateWishListInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WishListInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishListServiceServer).UpdateWishListInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_wishlist.WishListService/UpdateWishListInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishListServiceServer).UpdateWishListInfo(ctx, req.(*WishListInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _WishListService_DeleteWishListInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWishListInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishListServiceServer).DeleteWishListInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_wishlist.WishListService/DeleteWishListInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishListServiceServer).DeleteWishListInfo(ctx, req.(*DeleteWishListInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WishListService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_wishlist.WishListService",
	HandlerType: (*WishListServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWishListInfo",
			Handler:    _WishListService_CreateWishListInfo_Handler,
		},
		{
			MethodName: "GetWishList",
			Handler:    _WishListService_GetWishList_Handler,
		},
		{
			MethodName: "UpdateWishListInfo",
			Handler:    _WishListService_UpdateWishListInfo_Handler,
		},
		{
			MethodName: "DeleteWishListInfo",
			Handler:    _WishListService_DeleteWishListInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetWishListInfo",
			Handler:       _WishListService_GetWishListInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "wish_list.proto",
}