// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.2
// source: user/passgenerService.proto

package passenger_service

import (
	base "github.com/Remoulding/12306-go/idl-gen/base"
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

// 根据用户名查询乘车人列表请求
type ListPassengerByUsernameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ListPassengerByUsernameReq) Reset() {
	*x = ListPassengerByUsernameReq{}
	mi := &file_user_passgenerService_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPassengerByUsernameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPassengerByUsernameReq) ProtoMessage() {}

func (x *ListPassengerByUsernameReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPassengerByUsernameReq.ProtoReflect.Descriptor instead.
func (*ListPassengerByUsernameReq) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{0}
}

func (x *ListPassengerByUsernameReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// 根据用户名查询乘车人列表响应
type ListPassengerByUsernameResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PassengerList []*Passenger `protobuf:"bytes,1,rep,name=passenger_list,json=passengerList,proto3" json:"passenger_list,omitempty"`
}

func (x *ListPassengerByUsernameResp) Reset() {
	*x = ListPassengerByUsernameResp{}
	mi := &file_user_passgenerService_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPassengerByUsernameResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPassengerByUsernameResp) ProtoMessage() {}

func (x *ListPassengerByUsernameResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPassengerByUsernameResp.ProtoReflect.Descriptor instead.
func (*ListPassengerByUsernameResp) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{1}
}

func (x *ListPassengerByUsernameResp) GetPassengerList() []*Passenger {
	if x != nil {
		return x.PassengerList
	}
	return nil
}

// 根据乘车人 ID 集合查询乘车人列表请求
type ListPassengerByIdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string  `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Ids      []int64 `protobuf:"varint,2,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ListPassengerByIdsReq) Reset() {
	*x = ListPassengerByIdsReq{}
	mi := &file_user_passgenerService_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPassengerByIdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPassengerByIdsReq) ProtoMessage() {}

func (x *ListPassengerByIdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPassengerByIdsReq.ProtoReflect.Descriptor instead.
func (*ListPassengerByIdsReq) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{2}
}

func (x *ListPassengerByIdsReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ListPassengerByIdsReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

// 根据乘车人 ID 集合查询乘车人列表响应
type ListPassengerByIdsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PassengerList []*PassengerActualResp `protobuf:"bytes,1,rep,name=passenger_list,json=passengerList,proto3" json:"passenger_list,omitempty"`
}

func (x *ListPassengerByIdsResp) Reset() {
	*x = ListPassengerByIdsResp{}
	mi := &file_user_passgenerService_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPassengerByIdsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPassengerByIdsResp) ProtoMessage() {}

func (x *ListPassengerByIdsResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPassengerByIdsResp.ProtoReflect.Descriptor instead.
func (*ListPassengerByIdsResp) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{3}
}

func (x *ListPassengerByIdsResp) GetPassengerList() []*PassengerActualResp {
	if x != nil {
		return x.PassengerList
	}
	return nil
}

// 新增乘车人请求
type SavePassengerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Passenger *PassengerReq `protobuf:"bytes,1,opt,name=passenger,proto3" json:"passenger,omitempty"`
}

func (x *SavePassengerReq) Reset() {
	*x = SavePassengerReq{}
	mi := &file_user_passgenerService_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SavePassengerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePassengerReq) ProtoMessage() {}

func (x *SavePassengerReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePassengerReq.ProtoReflect.Descriptor instead.
func (*SavePassengerReq) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{4}
}

func (x *SavePassengerReq) GetPassenger() *PassengerReq {
	if x != nil {
		return x.Passenger
	}
	return nil
}

// 新增乘车人响应
type SavePassengerResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *base.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SavePassengerResp) Reset() {
	*x = SavePassengerResp{}
	mi := &file_user_passgenerService_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SavePassengerResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePassengerResp) ProtoMessage() {}

func (x *SavePassengerResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePassengerResp.ProtoReflect.Descriptor instead.
func (*SavePassengerResp) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{5}
}

func (x *SavePassengerResp) GetStatus() *base.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// 修改乘车人请求
type UpdatePassengerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Passenger *PassengerReq `protobuf:"bytes,1,opt,name=passenger,proto3" json:"passenger,omitempty"`
}

func (x *UpdatePassengerReq) Reset() {
	*x = UpdatePassengerReq{}
	mi := &file_user_passgenerService_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePassengerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePassengerReq) ProtoMessage() {}

func (x *UpdatePassengerReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePassengerReq.ProtoReflect.Descriptor instead.
func (*UpdatePassengerReq) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePassengerReq) GetPassenger() *PassengerReq {
	if x != nil {
		return x.Passenger
	}
	return nil
}

// 修改乘车人响应
type UpdatePassengerResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *base.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdatePassengerResp) Reset() {
	*x = UpdatePassengerResp{}
	mi := &file_user_passgenerService_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePassengerResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePassengerResp) ProtoMessage() {}

func (x *UpdatePassengerResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePassengerResp.ProtoReflect.Descriptor instead.
func (*UpdatePassengerResp) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePassengerResp) GetStatus() *base.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// 移除乘车人请求
type RemovePassengerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PassengerIds []int64 `protobuf:"varint,1,rep,packed,name=passenger_ids,json=passengerIds,proto3" json:"passenger_ids,omitempty"`
}

func (x *RemovePassengerReq) Reset() {
	*x = RemovePassengerReq{}
	mi := &file_user_passgenerService_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemovePassengerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePassengerReq) ProtoMessage() {}

func (x *RemovePassengerReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePassengerReq.ProtoReflect.Descriptor instead.
func (*RemovePassengerReq) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{8}
}

func (x *RemovePassengerReq) GetPassengerIds() []int64 {
	if x != nil {
		return x.PassengerIds
	}
	return nil
}

// 移除乘车人响应
type RemovePassengerResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *base.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RemovePassengerResp) Reset() {
	*x = RemovePassengerResp{}
	mi := &file_user_passgenerService_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemovePassengerResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePassengerResp) ProtoMessage() {}

func (x *RemovePassengerResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePassengerResp.ProtoReflect.Descriptor instead.
func (*RemovePassengerResp) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{9}
}

func (x *RemovePassengerResp) GetStatus() *base.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

// 通用乘车人信息
type PassengerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 乘车人 ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 真实姓名
	RealName string `protobuf:"bytes,2,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	// 证件类型
	IdType int32 `protobuf:"varint,3,opt,name=id_type,json=idType,proto3" json:"id_type,omitempty"`
	// 证件号
	IdCard string `protobuf:"bytes,4,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`
	// 优惠类型
	DiscountType int32 `protobuf:"varint,5,opt,name=discount_type,json=discountType,proto3" json:"discount_type,omitempty"`
	// 手机号
	Phone string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *PassengerReq) Reset() {
	*x = PassengerReq{}
	mi := &file_user_passgenerService_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PassengerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassengerReq) ProtoMessage() {}

func (x *PassengerReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassengerReq.ProtoReflect.Descriptor instead.
func (*PassengerReq) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{10}
}

func (x *PassengerReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PassengerReq) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *PassengerReq) GetIdType() int32 {
	if x != nil {
		return x.IdType
	}
	return 0
}

func (x *PassengerReq) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *PassengerReq) GetDiscountType() int32 {
	if x != nil {
		return x.DiscountType
	}
	return 0
}

func (x *PassengerReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

// 返回的乘车人响应信息
type Passenger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username     string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	RealName     string `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	IdType       int32  `protobuf:"varint,4,opt,name=id_type,json=idType,proto3" json:"id_type,omitempty"`
	IdCard       string `protobuf:"bytes,5,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`
	ActualIdCard string `protobuf:"bytes,6,opt,name=actual_id_card,json=actualIdCard,proto3" json:"actual_id_card,omitempty"`
	DiscountType int32  `protobuf:"varint,7,opt,name=discount_type,json=discountType,proto3" json:"discount_type,omitempty"`
	Phone        string `protobuf:"bytes,8,opt,name=phone,proto3" json:"phone,omitempty"`
	ActualPhone  string `protobuf:"bytes,9,opt,name=actual_phone,json=actualPhone,proto3" json:"actual_phone,omitempty"`
	CreateDate   string `protobuf:"bytes,10,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	VerifyStatus int32  `protobuf:"varint,11,opt,name=verify_status,json=verifyStatus,proto3" json:"verify_status,omitempty"`
}

func (x *Passenger) Reset() {
	*x = Passenger{}
	mi := &file_user_passgenerService_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Passenger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Passenger) ProtoMessage() {}

func (x *Passenger) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Passenger.ProtoReflect.Descriptor instead.
func (*Passenger) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{11}
}

func (x *Passenger) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Passenger) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Passenger) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *Passenger) GetIdType() int32 {
	if x != nil {
		return x.IdType
	}
	return 0
}

func (x *Passenger) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *Passenger) GetActualIdCard() string {
	if x != nil {
		return x.ActualIdCard
	}
	return ""
}

func (x *Passenger) GetDiscountType() int32 {
	if x != nil {
		return x.DiscountType
	}
	return 0
}

func (x *Passenger) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Passenger) GetActualPhone() string {
	if x != nil {
		return x.ActualPhone
	}
	return ""
}

func (x *Passenger) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

func (x *Passenger) GetVerifyStatus() int32 {
	if x != nil {
		return x.VerifyStatus
	}
	return 0
}

// 实际乘车人响应信息
type PassengerActualResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username     string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	RealName     string `protobuf:"bytes,3,opt,name=real_name,json=realName,proto3" json:"real_name,omitempty"`
	IdType       int32  `protobuf:"varint,4,opt,name=id_type,json=idType,proto3" json:"id_type,omitempty"`
	IdCard       string `protobuf:"bytes,5,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`
	DiscountType int32  `protobuf:"varint,6,opt,name=discount_type,json=discountType,proto3" json:"discount_type,omitempty"`
	Phone        string `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	CreateDate   string `protobuf:"bytes,8,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	VerifyStatus int32  `protobuf:"varint,9,opt,name=verify_status,json=verifyStatus,proto3" json:"verify_status,omitempty"`
}

func (x *PassengerActualResp) Reset() {
	*x = PassengerActualResp{}
	mi := &file_user_passgenerService_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PassengerActualResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassengerActualResp) ProtoMessage() {}

func (x *PassengerActualResp) ProtoReflect() protoreflect.Message {
	mi := &file_user_passgenerService_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassengerActualResp.ProtoReflect.Descriptor instead.
func (*PassengerActualResp) Descriptor() ([]byte, []int) {
	return file_user_passgenerService_proto_rawDescGZIP(), []int{12}
}

func (x *PassengerActualResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PassengerActualResp) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PassengerActualResp) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *PassengerActualResp) GetIdType() int32 {
	if x != nil {
		return x.IdType
	}
	return 0
}

func (x *PassengerActualResp) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *PassengerActualResp) GetDiscountType() int32 {
	if x != nil {
		return x.DiscountType
	}
	return 0
}

func (x *PassengerActualResp) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *PassengerActualResp) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

func (x *PassengerActualResp) GetVerifyStatus() int32 {
	if x != nil {
		return x.VerifyStatus
	}
	return 0
}

var File_user_passgenerService_proto protoreflect.FileDescriptor

var file_user_passgenerService_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61,
	0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x55, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x36, 0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50,
	0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e,
	0x67, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x5a,
	0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0d, 0x70, 0x61, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x10, 0x53, 0x61,
	0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x30,
	0x0a, 0x09, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x52, 0x09, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x22, 0x39, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x46, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x30, 0x0a, 0x09, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x52, 0x09, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e,
	0x67, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x39, 0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e,
	0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0c, 0x70,
	0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x3b, 0x0a, 0x13, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61,
	0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x22, 0xd0, 0x02, 0x0a, 0x09, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x72, 0x65, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x91, 0x02, 0x0a, 0x13, 0x50, 0x61, 0x73, 0x73, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x41, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x69, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x9f, 0x03, 0x0a, 0x10, 0x50,
	0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x63, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x54, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12,
	0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x0d, 0x53, 0x61,
	0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x12,
	0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42, 0x2c, 0x5a, 0x2a,
	0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_passgenerService_proto_rawDescOnce sync.Once
	file_user_passgenerService_proto_rawDescData = file_user_passgenerService_proto_rawDesc
)

func file_user_passgenerService_proto_rawDescGZIP() []byte {
	file_user_passgenerService_proto_rawDescOnce.Do(func() {
		file_user_passgenerService_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_passgenerService_proto_rawDescData)
	})
	return file_user_passgenerService_proto_rawDescData
}

var file_user_passgenerService_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_user_passgenerService_proto_goTypes = []any{
	(*ListPassengerByUsernameReq)(nil),  // 0: user.ListPassengerByUsernameReq
	(*ListPassengerByUsernameResp)(nil), // 1: user.ListPassengerByUsernameResp
	(*ListPassengerByIdsReq)(nil),       // 2: user.ListPassengerByIdsReq
	(*ListPassengerByIdsResp)(nil),      // 3: user.ListPassengerByIdsResp
	(*SavePassengerReq)(nil),            // 4: user.SavePassengerReq
	(*SavePassengerResp)(nil),           // 5: user.SavePassengerResp
	(*UpdatePassengerReq)(nil),          // 6: user.UpdatePassengerReq
	(*UpdatePassengerResp)(nil),         // 7: user.UpdatePassengerResp
	(*RemovePassengerReq)(nil),          // 8: user.RemovePassengerReq
	(*RemovePassengerResp)(nil),         // 9: user.RemovePassengerResp
	(*PassengerReq)(nil),                // 10: user.PassengerReq
	(*Passenger)(nil),                   // 11: user.Passenger
	(*PassengerActualResp)(nil),         // 12: user.PassengerActualResp
	(*base.Status)(nil),                 // 13: base.Status
}
var file_user_passgenerService_proto_depIdxs = []int32{
	11, // 0: user.ListPassengerByUsernameResp.passenger_list:type_name -> user.Passenger
	12, // 1: user.ListPassengerByIdsResp.passenger_list:type_name -> user.PassengerActualResp
	10, // 2: user.SavePassengerReq.passenger:type_name -> user.PassengerReq
	13, // 3: user.SavePassengerResp.status:type_name -> base.Status
	10, // 4: user.UpdatePassengerReq.passenger:type_name -> user.PassengerReq
	13, // 5: user.UpdatePassengerResp.status:type_name -> base.Status
	13, // 6: user.RemovePassengerResp.status:type_name -> base.Status
	0,  // 7: user.PassengerService.ListPassengerQueryByUsername:input_type -> user.ListPassengerByUsernameReq
	2,  // 8: user.PassengerService.ListPassengerQueryByIds:input_type -> user.ListPassengerByIdsReq
	4,  // 9: user.PassengerService.SavePassenger:input_type -> user.SavePassengerReq
	6,  // 10: user.PassengerService.UpdatePassenger:input_type -> user.UpdatePassengerReq
	8,  // 11: user.PassengerService.RemovePassenger:input_type -> user.RemovePassengerReq
	1,  // 12: user.PassengerService.ListPassengerQueryByUsername:output_type -> user.ListPassengerByUsernameResp
	3,  // 13: user.PassengerService.ListPassengerQueryByIds:output_type -> user.ListPassengerByIdsResp
	5,  // 14: user.PassengerService.SavePassenger:output_type -> user.SavePassengerResp
	7,  // 15: user.PassengerService.UpdatePassenger:output_type -> user.UpdatePassengerResp
	9,  // 16: user.PassengerService.RemovePassenger:output_type -> user.RemovePassengerResp
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_user_passgenerService_proto_init() }
func file_user_passgenerService_proto_init() {
	if File_user_passgenerService_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_passgenerService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_passgenerService_proto_goTypes,
		DependencyIndexes: file_user_passgenerService_proto_depIdxs,
		MessageInfos:      file_user_passgenerService_proto_msgTypes,
	}.Build()
	File_user_passgenerService_proto = out.File
	file_user_passgenerService_proto_rawDesc = nil
	file_user_passgenerService_proto_goTypes = nil
	file_user_passgenerService_proto_depIdxs = nil
}
