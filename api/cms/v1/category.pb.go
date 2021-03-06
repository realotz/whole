// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: api/cms/v1/category.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//创建请求
type CreateCategoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Category `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateCategoryReply) Reset() {
	*x = CreateCategoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cms_v1_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryReply) ProtoMessage() {}

func (x *CreateCategoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cms_v1_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryReply.ProtoReflect.Descriptor instead.
func (*CreateCategoryReply) Descriptor() ([]byte, []int) {
	return file_api_cms_v1_category_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCategoryReply) GetData() *Category {
	if x != nil {
		return x.Data
	}
	return nil
}

//更新请求
type UpdateCategoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Category `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *UpdateCategoryReply) Reset() {
	*x = UpdateCategoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cms_v1_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCategoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCategoryReply) ProtoMessage() {}

func (x *UpdateCategoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cms_v1_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCategoryReply.ProtoReflect.Descriptor instead.
func (*UpdateCategoryReply) Descriptor() ([]byte, []int) {
	return file_api_cms_v1_category_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCategoryReply) GetData() *Category {
	if x != nil {
		return x.Data
	}
	return nil
}

//批量删除请求
type DeleteCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteCategoryRequest) Reset() {
	*x = DeleteCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cms_v1_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryRequest) ProtoMessage() {}

func (x *DeleteCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cms_v1_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_cms_v1_category_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCategoryRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

//删除结果
type DeleteCategoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCategoryReply) Reset() {
	*x = DeleteCategoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cms_v1_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCategoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCategoryReply) ProtoMessage() {}

func (x *DeleteCategoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cms_v1_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCategoryReply.ProtoReflect.Descriptor instead.
func (*DeleteCategoryReply) Descriptor() ([]byte, []int) {
	return file_api_cms_v1_category_proto_rawDescGZIP(), []int{3}
}

//查询单个数据
type GetCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCategoryRequest) Reset() {
	*x = GetCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cms_v1_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoryRequest) ProtoMessage() {}

func (x *GetCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cms_v1_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_cms_v1_category_proto_rawDescGZIP(), []int{4}
}

func (x *GetCategoryRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

//列表查询条件
type ListCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//父id查询
	Pid int64 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	//名称搜索
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//页码
	Page uint32 `protobuf:"varint,10,opt,name=page,proto3" json:"page,omitempty"`
	//分页大小
	PageSize uint32 `protobuf:"varint,11,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListCategoryRequest) Reset() {
	*x = ListCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cms_v1_category_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoryRequest) ProtoMessage() {}

func (x *ListCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cms_v1_category_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoryRequest.ProtoReflect.Descriptor instead.
func (*ListCategoryRequest) Descriptor() ([]byte, []int) {
	return file_api_cms_v1_category_proto_rawDescGZIP(), []int{5}
}

func (x *ListCategoryRequest) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ListCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListCategoryRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListCategoryRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

//列表查询返回
type ListCategoryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Category `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	//数据总计
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	//页码
	Page uint32 `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	//分页大小
	PageSize uint32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListCategoryReply) Reset() {
	*x = ListCategoryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cms_v1_category_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCategoryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCategoryReply) ProtoMessage() {}

func (x *ListCategoryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_cms_v1_category_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCategoryReply.ProtoReflect.Descriptor instead.
func (*ListCategoryReply) Descriptor() ([]byte, []int) {
	return file_api_cms_v1_category_proto_rawDescGZIP(), []int{6}
}

func (x *ListCategoryReply) GetList() []*Category {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListCategoryReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListCategoryReply) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListCategoryReply) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// 文章类目结构体(请在此处定义数据结构)
type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//分类名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//父级分类id
	Pid int64 `protobuf:"varint,3,opt,name=pid,proto3" json:"pid,omitempty"`
	//分类图标
	Icon string `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	//分类简介
	Desc       string                 `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,23,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_cms_v1_category_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_api_cms_v1_category_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_api_cms_v1_category_proto_rawDescGZIP(), []int{7}
}

func (x *Category) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Category) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Category) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Category) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Category) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Category) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

var File_api_cms_v1_category_proto protoreflect.FileDescriptor

var file_api_cms_v1_category_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3f, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x6c, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x84,
	0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xe2, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x8e, 0x03, 0x0a, 0x0f, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x47, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x1a,
	0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x54, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x4e, 0x0a, 0x0c, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x36, 0x0a, 0x0a, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6f, 0x74, 0x7a, 0x2f,
	0x77, 0x68, 0x6f, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_cms_v1_category_proto_rawDescOnce sync.Once
	file_api_cms_v1_category_proto_rawDescData = file_api_cms_v1_category_proto_rawDesc
)

func file_api_cms_v1_category_proto_rawDescGZIP() []byte {
	file_api_cms_v1_category_proto_rawDescOnce.Do(func() {
		file_api_cms_v1_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_cms_v1_category_proto_rawDescData)
	})
	return file_api_cms_v1_category_proto_rawDescData
}

var file_api_cms_v1_category_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_cms_v1_category_proto_goTypes = []interface{}{
	(*CreateCategoryReply)(nil),   // 0: api.cms.v1.CreateCategoryReply
	(*UpdateCategoryReply)(nil),   // 1: api.cms.v1.UpdateCategoryReply
	(*DeleteCategoryRequest)(nil), // 2: api.cms.v1.DeleteCategoryRequest
	(*DeleteCategoryReply)(nil),   // 3: api.cms.v1.DeleteCategoryReply
	(*GetCategoryRequest)(nil),    // 4: api.cms.v1.GetCategoryRequest
	(*ListCategoryRequest)(nil),   // 5: api.cms.v1.ListCategoryRequest
	(*ListCategoryReply)(nil),     // 6: api.cms.v1.ListCategoryReply
	(*Category)(nil),              // 7: api.cms.v1.Category
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_api_cms_v1_category_proto_depIdxs = []int32{
	7,  // 0: api.cms.v1.CreateCategoryReply.data:type_name -> api.cms.v1.Category
	7,  // 1: api.cms.v1.UpdateCategoryReply.data:type_name -> api.cms.v1.Category
	7,  // 2: api.cms.v1.ListCategoryReply.list:type_name -> api.cms.v1.Category
	8,  // 3: api.cms.v1.Category.update_time:type_name -> google.protobuf.Timestamp
	8,  // 4: api.cms.v1.Category.create_time:type_name -> google.protobuf.Timestamp
	7,  // 5: api.cms.v1.CategoryService.CreateCategory:input_type -> api.cms.v1.Category
	7,  // 6: api.cms.v1.CategoryService.UpdateCategory:input_type -> api.cms.v1.Category
	2,  // 7: api.cms.v1.CategoryService.DeleteCategory:input_type -> api.cms.v1.DeleteCategoryRequest
	4,  // 8: api.cms.v1.CategoryService.GetCategory:input_type -> api.cms.v1.GetCategoryRequest
	5,  // 9: api.cms.v1.CategoryService.ListCategory:input_type -> api.cms.v1.ListCategoryRequest
	0,  // 10: api.cms.v1.CategoryService.CreateCategory:output_type -> api.cms.v1.CreateCategoryReply
	1,  // 11: api.cms.v1.CategoryService.UpdateCategory:output_type -> api.cms.v1.UpdateCategoryReply
	3,  // 12: api.cms.v1.CategoryService.DeleteCategory:output_type -> api.cms.v1.DeleteCategoryReply
	7,  // 13: api.cms.v1.CategoryService.GetCategory:output_type -> api.cms.v1.Category
	6,  // 14: api.cms.v1.CategoryService.ListCategory:output_type -> api.cms.v1.ListCategoryReply
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_cms_v1_category_proto_init() }
func file_api_cms_v1_category_proto_init() {
	if File_api_cms_v1_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_cms_v1_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryReply); i {
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
		file_api_cms_v1_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCategoryReply); i {
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
		file_api_cms_v1_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCategoryRequest); i {
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
		file_api_cms_v1_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCategoryReply); i {
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
		file_api_cms_v1_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoryRequest); i {
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
		file_api_cms_v1_category_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoryRequest); i {
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
		file_api_cms_v1_category_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCategoryReply); i {
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
		file_api_cms_v1_category_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
			RawDescriptor: file_api_cms_v1_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_cms_v1_category_proto_goTypes,
		DependencyIndexes: file_api_cms_v1_category_proto_depIdxs,
		MessageInfos:      file_api_cms_v1_category_proto_msgTypes,
	}.Build()
	File_api_cms_v1_category_proto = out.File
	file_api_cms_v1_category_proto_rawDesc = nil
	file_api_cms_v1_category_proto_goTypes = nil
	file_api_cms_v1_category_proto_depIdxs = nil
}
