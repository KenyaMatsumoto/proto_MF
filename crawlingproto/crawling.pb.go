// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: crawlingproto/crawling.proto

package crawlingproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInput *UserInput `protobuf:"bytes,1,opt,name=user_input,json=userInput,proto3" json:"user_input,omitempty"`
	Pass      string     `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
	SiteKind  int32      `protobuf:"varint,3,opt,name=site_kind,json=siteKind,proto3" json:"site_kind,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawlingproto_crawling_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crawlingproto_crawling_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_crawlingproto_crawling_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetUserInput() *UserInput {
	if x != nil {
		return x.UserInput
	}
	return nil
}

func (x *UserRequest) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *UserRequest) GetSiteKind() int32 {
	if x != nil {
		return x.SiteKind
	}
	return 0
}

type UserInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserInput) Reset() {
	*x = UserInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawlingproto_crawling_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInput) ProtoMessage() {}

func (x *UserInput) ProtoReflect() protoreflect.Message {
	mi := &file_crawlingproto_crawling_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInput.ProtoReflect.Descriptor instead.
func (*UserInput) Descriptor() ([]byte, []int) {
	return file_crawlingproto_crawling_proto_rawDescGZIP(), []int{1}
}

func (x *UserInput) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawlingproto_crawling_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crawlingproto_crawling_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_crawlingproto_crawling_proto_rawDescGZIP(), []int{2}
}

func (x *UserResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

type MfResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Office *Office `protobuf:"bytes,1,opt,name=office,proto3" json:"office,omitempty"`
}

func (x *MfResponse) Reset() {
	*x = MfResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawlingproto_crawling_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MfResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MfResponse) ProtoMessage() {}

func (x *MfResponse) ProtoReflect() protoreflect.Message {
	mi := &file_crawlingproto_crawling_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MfResponse.ProtoReflect.Descriptor instead.
func (*MfResponse) Descriptor() ([]byte, []int) {
	return file_crawlingproto_crawling_proto_rawDescGZIP(), []int{3}
}

func (x *MfResponse) GetOffice() *Office {
	if x != nil {
		return x.Office
	}
	return nil
}

type Office struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OfficeName string                 `protobuf:"bytes,1,opt,name=office_name,json=officeName,proto3" json:"office_name,omitempty"`
	Crawling   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=crawling,proto3" json:"crawling,omitempty"`
	Banks      []*Bank                `protobuf:"bytes,3,rep,name=banks,proto3" json:"banks,omitempty"`
	Cards      []*Card                `protobuf:"bytes,4,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *Office) Reset() {
	*x = Office{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawlingproto_crawling_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Office) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Office) ProtoMessage() {}

func (x *Office) ProtoReflect() protoreflect.Message {
	mi := &file_crawlingproto_crawling_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Office.ProtoReflect.Descriptor instead.
func (*Office) Descriptor() ([]byte, []int) {
	return file_crawlingproto_crawling_proto_rawDescGZIP(), []int{4}
}

func (x *Office) GetOfficeName() string {
	if x != nil {
		return x.OfficeName
	}
	return ""
}

func (x *Office) GetCrawling() *timestamppb.Timestamp {
	if x != nil {
		return x.Crawling
	}
	return nil
}

func (x *Office) GetBanks() []*Bank {
	if x != nil {
		return x.Banks
	}
	return nil
}

func (x *Office) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

type Bank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BankName       string    `protobuf:"bytes,1,opt,name=bank_name,json=bankName,proto3" json:"bank_name,omitempty"`
	BankAmount     string    `protobuf:"bytes,2,opt,name=bank_amount,json=bankAmount,proto3" json:"bank_amount,omitempty"`
	BankLastCommit string    `protobuf:"bytes,3,opt,name=bank_last_commit,json=bankLastCommit,proto3" json:"bank_last_commit,omitempty"`
	BankStatus     string    `protobuf:"bytes,4,opt,name=bank_status,json=bankStatus,proto3" json:"bank_status,omitempty"`
	Details        []*Detail `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Bank) Reset() {
	*x = Bank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawlingproto_crawling_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bank) ProtoMessage() {}

func (x *Bank) ProtoReflect() protoreflect.Message {
	mi := &file_crawlingproto_crawling_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bank.ProtoReflect.Descriptor instead.
func (*Bank) Descriptor() ([]byte, []int) {
	return file_crawlingproto_crawling_proto_rawDescGZIP(), []int{5}
}

func (x *Bank) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

func (x *Bank) GetBankAmount() string {
	if x != nil {
		return x.BankAmount
	}
	return ""
}

func (x *Bank) GetBankLastCommit() string {
	if x != nil {
		return x.BankLastCommit
	}
	return ""
}

func (x *Bank) GetBankStatus() string {
	if x != nil {
		return x.BankStatus
	}
	return ""
}

func (x *Bank) GetDetails() []*Detail {
	if x != nil {
		return x.Details
	}
	return nil
}

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardName       string    `protobuf:"bytes,1,opt,name=card_name,json=cardName,proto3" json:"card_name,omitempty"`
	CardAmount     string    `protobuf:"bytes,2,opt,name=card_amount,json=cardAmount,proto3" json:"card_amount,omitempty"`
	CardLastCommit string    `protobuf:"bytes,3,opt,name=card_last_commit,json=cardLastCommit,proto3" json:"card_last_commit,omitempty"`
	CardStatus     string    `protobuf:"bytes,4,opt,name=card_status,json=cardStatus,proto3" json:"card_status,omitempty"`
	Details        []*Detail `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawlingproto_crawling_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_crawlingproto_crawling_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_crawlingproto_crawling_proto_rawDescGZIP(), []int{6}
}

func (x *Card) GetCardName() string {
	if x != nil {
		return x.CardName
	}
	return ""
}

func (x *Card) GetCardAmount() string {
	if x != nil {
		return x.CardAmount
	}
	return ""
}

func (x *Card) GetCardLastCommit() string {
	if x != nil {
		return x.CardLastCommit
	}
	return ""
}

func (x *Card) GetCardStatus() string {
	if x != nil {
		return x.CardStatus
	}
	return ""
}

func (x *Card) GetDetails() []*Detail {
	if x != nil {
		return x.Details
	}
	return nil
}

type Detail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DetailName string                 `protobuf:"bytes,1,opt,name=detail_name,json=detailName,proto3" json:"detail_name,omitempty"`
	Contents   string                 `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	Amount     int64                  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Balance    int64                  `protobuf:"varint,4,opt,name=balance,proto3" json:"balance,omitempty"`
	Status     string                 `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	DetailDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=detail_date,json=detailDate,proto3" json:"detail_date,omitempty"`
}

func (x *Detail) Reset() {
	*x = Detail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawlingproto_crawling_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Detail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Detail) ProtoMessage() {}

func (x *Detail) ProtoReflect() protoreflect.Message {
	mi := &file_crawlingproto_crawling_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Detail.ProtoReflect.Descriptor instead.
func (*Detail) Descriptor() ([]byte, []int) {
	return file_crawlingproto_crawling_proto_rawDescGZIP(), []int{7}
}

func (x *Detail) GetDetailName() string {
	if x != nil {
		return x.DetailName
	}
	return ""
}

func (x *Detail) GetContents() string {
	if x != nil {
		return x.Contents
	}
	return ""
}

func (x *Detail) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Detail) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Detail) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Detail) GetDetailDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DetailDate
	}
	return nil
}

type MfRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInput *UserInput `protobuf:"bytes,1,opt,name=user_input,json=userInput,proto3" json:"user_input,omitempty"`
	Pass      string     `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
	StartDay  string     `protobuf:"bytes,3,opt,name=start_day,json=startDay,proto3" json:"start_day,omitempty"`
	LastDay   string     `protobuf:"bytes,4,opt,name=last_day,json=lastDay,proto3" json:"last_day,omitempty"`
}

func (x *MfRequest) Reset() {
	*x = MfRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crawlingproto_crawling_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MfRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MfRequest) ProtoMessage() {}

func (x *MfRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crawlingproto_crawling_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MfRequest.ProtoReflect.Descriptor instead.
func (*MfRequest) Descriptor() ([]byte, []int) {
	return file_crawlingproto_crawling_proto_rawDescGZIP(), []int{8}
}

func (x *MfRequest) GetUserInput() *UserInput {
	if x != nil {
		return x.UserInput
	}
	return nil
}

func (x *MfRequest) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *MfRequest) GetStartDay() string {
	if x != nil {
		return x.StartDay
	}
	return ""
}

func (x *MfRequest) GetLastDay() string {
	if x != nil {
		return x.LastDay
	}
	return ""
}

var File_crawlingproto_crawling_proto protoreflect.FileDescriptor

var file_crawlingproto_crawling_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x63, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69,
	0x74, 0x65, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73,
	0x69, 0x74, 0x65, 0x4b, 0x69, 0x6e, 0x64, 0x22, 0x24, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x3b, 0x0a, 0x0a,
	0x4d, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x72, 0x61,
	0x77, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x06, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a,
	0x05, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61, 0x6e,
	0x6b, 0x52, 0x05, 0x62, 0x61, 0x6e, 0x6b, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x69,
	0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05, 0x63, 0x61,
	0x72, 0x64, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x04, 0x42, 0x61, 0x6e, 0x6b, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x6e,
	0x6b, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x62, 0x61, 0x6e, 0x6b, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x61,
	0x6e, 0x6b, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x61, 0x6e, 0x6b, 0x4c, 0x61, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e,
	0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xc0, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a,
	0x10, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x61, 0x72, 0x64, 0x4c, 0x61, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61,
	0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x72, 0x61, 0x77,
	0x6c, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x06, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x09, 0x4d, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x72, 0x61,
	0x77, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x61, 0x79, 0x32, 0x9c, 0x01, 0x0a, 0x0f,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x48, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x1a,
	0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x72, 0x61,
	0x77, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x4d, 0x66, 0x52,
	0x65, 0x61, 0x64, 0x12, 0x18, 0x2e, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x63, 0x72, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x66,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x63, 0x72,
	0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crawlingproto_crawling_proto_rawDescOnce sync.Once
	file_crawlingproto_crawling_proto_rawDescData = file_crawlingproto_crawling_proto_rawDesc
)

func file_crawlingproto_crawling_proto_rawDescGZIP() []byte {
	file_crawlingproto_crawling_proto_rawDescOnce.Do(func() {
		file_crawlingproto_crawling_proto_rawDescData = protoimpl.X.CompressGZIP(file_crawlingproto_crawling_proto_rawDescData)
	})
	return file_crawlingproto_crawling_proto_rawDescData
}

var file_crawlingproto_crawling_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_crawlingproto_crawling_proto_goTypes = []interface{}{
	(*UserRequest)(nil),           // 0: crawlingproto.UserRequest
	(*UserInput)(nil),             // 1: crawlingproto.UserInput
	(*UserResponse)(nil),          // 2: crawlingproto.UserResponse
	(*MfResponse)(nil),            // 3: crawlingproto.MfResponse
	(*Office)(nil),                // 4: crawlingproto.Office
	(*Bank)(nil),                  // 5: crawlingproto.Bank
	(*Card)(nil),                  // 6: crawlingproto.Card
	(*Detail)(nil),                // 7: crawlingproto.Detail
	(*MfRequest)(nil),             // 8: crawlingproto.MfRequest
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_crawlingproto_crawling_proto_depIdxs = []int32{
	1,  // 0: crawlingproto.UserRequest.user_input:type_name -> crawlingproto.UserInput
	4,  // 1: crawlingproto.MfResponse.office:type_name -> crawlingproto.Office
	9,  // 2: crawlingproto.Office.crawling:type_name -> google.protobuf.Timestamp
	5,  // 3: crawlingproto.Office.banks:type_name -> crawlingproto.Bank
	6,  // 4: crawlingproto.Office.cards:type_name -> crawlingproto.Card
	7,  // 5: crawlingproto.Bank.details:type_name -> crawlingproto.Detail
	7,  // 6: crawlingproto.Card.details:type_name -> crawlingproto.Detail
	9,  // 7: crawlingproto.Detail.detail_date:type_name -> google.protobuf.Timestamp
	1,  // 8: crawlingproto.MfRequest.user_input:type_name -> crawlingproto.UserInput
	0,  // 9: crawlingproto.CrawlingService.UserHandler:input_type -> crawlingproto.UserRequest
	8,  // 10: crawlingproto.CrawlingService.MfRead:input_type -> crawlingproto.MfRequest
	2,  // 11: crawlingproto.CrawlingService.UserHandler:output_type -> crawlingproto.UserResponse
	3,  // 12: crawlingproto.CrawlingService.MfRead:output_type -> crawlingproto.MfResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_crawlingproto_crawling_proto_init() }
func file_crawlingproto_crawling_proto_init() {
	if File_crawlingproto_crawling_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crawlingproto_crawling_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_crawlingproto_crawling_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInput); i {
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
		file_crawlingproto_crawling_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
		file_crawlingproto_crawling_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MfResponse); i {
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
		file_crawlingproto_crawling_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Office); i {
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
		file_crawlingproto_crawling_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bank); i {
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
		file_crawlingproto_crawling_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_crawlingproto_crawling_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Detail); i {
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
		file_crawlingproto_crawling_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MfRequest); i {
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
			RawDescriptor: file_crawlingproto_crawling_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crawlingproto_crawling_proto_goTypes,
		DependencyIndexes: file_crawlingproto_crawling_proto_depIdxs,
		MessageInfos:      file_crawlingproto_crawling_proto_msgTypes,
	}.Build()
	File_crawlingproto_crawling_proto = out.File
	file_crawlingproto_crawling_proto_rawDesc = nil
	file_crawlingproto_crawling_proto_goTypes = nil
	file_crawlingproto_crawling_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CrawlingServiceClient is the client API for CrawlingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CrawlingServiceClient interface {
	UserHandler(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	MfRead(ctx context.Context, in *MfRequest, opts ...grpc.CallOption) (*MfResponse, error)
}

type crawlingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCrawlingServiceClient(cc grpc.ClientConnInterface) CrawlingServiceClient {
	return &crawlingServiceClient{cc}
}

func (c *crawlingServiceClient) UserHandler(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/crawlingproto.CrawlingService/UserHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *crawlingServiceClient) MfRead(ctx context.Context, in *MfRequest, opts ...grpc.CallOption) (*MfResponse, error) {
	out := new(MfResponse)
	err := c.cc.Invoke(ctx, "/crawlingproto.CrawlingService/MfRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CrawlingServiceServer is the server API for CrawlingService service.
type CrawlingServiceServer interface {
	UserHandler(context.Context, *UserRequest) (*UserResponse, error)
	MfRead(context.Context, *MfRequest) (*MfResponse, error)
}

// UnimplementedCrawlingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCrawlingServiceServer struct {
}

func (*UnimplementedCrawlingServiceServer) UserHandler(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserHandler not implemented")
}
func (*UnimplementedCrawlingServiceServer) MfRead(context.Context, *MfRequest) (*MfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MfRead not implemented")
}

func RegisterCrawlingServiceServer(s *grpc.Server, srv CrawlingServiceServer) {
	s.RegisterService(&_CrawlingService_serviceDesc, srv)
}

func _CrawlingService_UserHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlingServiceServer).UserHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crawlingproto.CrawlingService/UserHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlingServiceServer).UserHandler(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CrawlingService_MfRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CrawlingServiceServer).MfRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crawlingproto.CrawlingService/MfRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CrawlingServiceServer).MfRead(ctx, req.(*MfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CrawlingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crawlingproto.CrawlingService",
	HandlerType: (*CrawlingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserHandler",
			Handler:    _CrawlingService_UserHandler_Handler,
		},
		{
			MethodName: "MfRead",
			Handler:    _CrawlingService_MfRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crawlingproto/crawling.proto",
}
