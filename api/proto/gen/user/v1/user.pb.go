// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: api/proto/user/v1/user.proto

package userv1

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// 登录
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_api_proto_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_api_proto_user_v1_user_proto_rawDescGZIP(), []int{1}
}

// 注册
type SignupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SignupReq) Reset() {
	*x = SignupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupReq) ProtoMessage() {}

func (x *SignupReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_v1_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupReq.ProtoReflect.Descriptor instead.
func (*SignupReq) Descriptor() ([]byte, []int) {
	return file_api_proto_user_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *SignupReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type SignupResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignupResp) Reset() {
	*x = SignupResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_v1_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupResp) ProtoMessage() {}

func (x *SignupResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_v1_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupResp.ProtoReflect.Descriptor instead.
func (*SignupResp) Descriptor() ([]byte, []int) {
	return file_api_proto_user_v1_user_proto_rawDescGZIP(), []int{3}
}

// 登出
type LogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *LogoutReq) Reset() {
	*x = LogoutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_v1_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReq) ProtoMessage() {}

func (x *LogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_v1_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReq.ProtoReflect.Descriptor instead.
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return file_api_proto_user_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *LogoutReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type LogoutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutResp) Reset() {
	*x = LogoutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_v1_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResp) ProtoMessage() {}

func (x *LogoutResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_v1_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResp.ProtoReflect.Descriptor instead.
func (*LogoutResp) Descriptor() ([]byte, []int) {
	return file_api_proto_user_v1_user_proto_rawDescGZIP(), []int{5}
}

// 描述
type ProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProfileReq) Reset() {
	*x = ProfileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_v1_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileReq) ProtoMessage() {}

func (x *ProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_v1_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileReq.ProtoReflect.Descriptor instead.
func (*ProfileReq) Descriptor() ([]byte, []int) {
	return file_api_proto_user_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *ProfileReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProfileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *ProfileResp) Reset() {
	*x = ProfileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_v1_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileResp) ProtoMessage() {}

func (x *ProfileResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_v1_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileResp.ProtoReflect.Descriptor instead.
func (*ProfileResp) Descriptor() ([]byte, []int) {
	return file_api_proto_user_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *ProfileResp) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// 更新
type UpdateNonSensitiveInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateNonSensitiveInfoReq) Reset() {
	*x = UpdateNonSensitiveInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_v1_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNonSensitiveInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNonSensitiveInfoReq) ProtoMessage() {}

func (x *UpdateNonSensitiveInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_v1_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNonSensitiveInfoReq.ProtoReflect.Descriptor instead.
func (*UpdateNonSensitiveInfoReq) Descriptor() ([]byte, []int) {
	return file_api_proto_user_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateNonSensitiveInfoReq) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateNonSensitiveInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateNonSensitiveInfoResp) Reset() {
	*x = UpdateNonSensitiveInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_v1_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNonSensitiveInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNonSensitiveInfoResp) ProtoMessage() {}

func (x *UpdateNonSensitiveInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_v1_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNonSensitiveInfoResp.ProtoReflect.Descriptor instead.
func (*UpdateNonSensitiveInfoResp) Descriptor() ([]byte, []int) {
	return file_api_proto_user_v1_user_proto_rawDescGZIP(), []int{9}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email      string               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Nickname   string               `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Password   string               `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Phone      string               `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	AboutMe    string               `protobuf:"bytes,6,opt,name=about_me,json=aboutMe,proto3" json:"about_me,omitempty"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Birthday   *timestamp.Timestamp `protobuf:"bytes,8,opt,name=birthday,proto3" json:"birthday,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_user_v1_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_user_v1_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_proto_user_v1_user_proto_rawDescGZIP(), []int{10}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetAboutMe() string {
	if x != nil {
		return x.AboutMe
	}
	return ""
}

func (x *User) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *User) GetBirthday() *timestamp.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

var File_api_proto_user_v1_user_proto protoreflect.FileDescriptor

var file_api_proto_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x0b, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x2c, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12, 0x1f,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22,
	0x0c, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x22, 0x21, 0x0a,
	0x09, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x0c, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1c,
	0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x0b,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3c, 0x0a, 0x19,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x6e, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x1c, 0x0a, 0x1a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x6e, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x8a, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x5f, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x4d, 0x65,
	0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a,
	0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x64, 0x61, 0x79, 0x32, 0xa8, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2d, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x10, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x2d, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x10, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x30, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x5d, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x6e, 0x53, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x6e, 0x53, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x6e, 0x53,
	0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x10, 0x5a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_user_v1_user_proto_rawDescOnce sync.Once
	file_api_proto_user_v1_user_proto_rawDescData = file_api_proto_user_v1_user_proto_rawDesc
)

func file_api_proto_user_v1_user_proto_rawDescGZIP() []byte {
	file_api_proto_user_v1_user_proto_rawDescOnce.Do(func() {
		file_api_proto_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_user_v1_user_proto_rawDescData)
	})
	return file_api_proto_user_v1_user_proto_rawDescData
}

var file_api_proto_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_proto_user_v1_user_proto_goTypes = []interface{}{
	(*LoginReq)(nil),                   // 0: users.LoginReq
	(*LoginResp)(nil),                  // 1: users.LoginResp
	(*SignupReq)(nil),                  // 2: users.SignupReq
	(*SignupResp)(nil),                 // 3: users.SignupResp
	(*LogoutReq)(nil),                  // 4: users.LogoutReq
	(*LogoutResp)(nil),                 // 5: users.LogoutResp
	(*ProfileReq)(nil),                 // 6: users.ProfileReq
	(*ProfileResp)(nil),                // 7: users.ProfileResp
	(*UpdateNonSensitiveInfoReq)(nil),  // 8: users.UpdateNonSensitiveInfoReq
	(*UpdateNonSensitiveInfoResp)(nil), // 9: users.UpdateNonSensitiveInfoResp
	(*User)(nil),                       // 10: users.User
	(*timestamp.Timestamp)(nil),        // 11: google.protobuf.Timestamp
}
var file_api_proto_user_v1_user_proto_depIdxs = []int32{
	10, // 0: users.SignupReq.user:type_name -> users.User
	10, // 1: users.ProfileResp.user:type_name -> users.User
	10, // 2: users.UpdateNonSensitiveInfoReq.user:type_name -> users.User
	11, // 3: users.User.create_time:type_name -> google.protobuf.Timestamp
	11, // 4: users.User.birthday:type_name -> google.protobuf.Timestamp
	0,  // 5: users.UserService.Login:input_type -> users.LoginReq
	2,  // 6: users.UserService.Signup:input_type -> users.SignupReq
	4,  // 7: users.UserService.Logout:input_type -> users.LogoutReq
	6,  // 8: users.UserService.Profile:input_type -> users.ProfileReq
	8,  // 9: users.UserService.UpdateNonSensitiveInfo:input_type -> users.UpdateNonSensitiveInfoReq
	1,  // 10: users.UserService.Login:output_type -> users.LoginResp
	3,  // 11: users.UserService.Signup:output_type -> users.SignupResp
	5,  // 12: users.UserService.Logout:output_type -> users.LogoutResp
	7,  // 13: users.UserService.Profile:output_type -> users.ProfileResp
	9,  // 14: users.UserService.UpdateNonSensitiveInfo:output_type -> users.UpdateNonSensitiveInfoResp
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_proto_user_v1_user_proto_init() }
func file_api_proto_user_v1_user_proto_init() {
	if File_api_proto_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_api_proto_user_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
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
		file_api_proto_user_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupReq); i {
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
		file_api_proto_user_v1_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupResp); i {
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
		file_api_proto_user_v1_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutReq); i {
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
		file_api_proto_user_v1_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResp); i {
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
		file_api_proto_user_v1_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileReq); i {
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
		file_api_proto_user_v1_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileResp); i {
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
		file_api_proto_user_v1_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNonSensitiveInfoReq); i {
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
		file_api_proto_user_v1_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNonSensitiveInfoResp); i {
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
		file_api_proto_user_v1_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_api_proto_user_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_user_v1_user_proto_goTypes,
		DependencyIndexes: file_api_proto_user_v1_user_proto_depIdxs,
		MessageInfos:      file_api_proto_user_v1_user_proto_msgTypes,
	}.Build()
	File_api_proto_user_v1_user_proto = out.File
	file_api_proto_user_v1_user_proto_rawDesc = nil
	file_api_proto_user_v1_user_proto_goTypes = nil
	file_api_proto_user_v1_user_proto_depIdxs = nil
}
