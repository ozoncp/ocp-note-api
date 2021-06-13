// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: ocp-note-api.proto

package ocp_note_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateNoteV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClassroomId int32 `protobuf:"varint,3,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
	DocumentId  int32 `protobuf:"varint,4,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
}

func (x *CreateNoteV1Request) Reset() {
	*x = CreateNoteV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_note_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNoteV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteV1Request) ProtoMessage() {}

func (x *CreateNoteV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_note_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteV1Request.ProtoReflect.Descriptor instead.
func (*CreateNoteV1Request) Descriptor() ([]byte, []int) {
	return file_ocp_note_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNoteV1Request) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateNoteV1Request) GetClassroomId() int32 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *CreateNoteV1Request) GetDocumentId() int32 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

type CreateNoteV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteId uint64 `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
}

func (x *CreateNoteV1Response) Reset() {
	*x = CreateNoteV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_note_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNoteV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteV1Response) ProtoMessage() {}

func (x *CreateNoteV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_note_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteV1Response.ProtoReflect.Descriptor instead.
func (*CreateNoteV1Response) Descriptor() ([]byte, []int) {
	return file_ocp_note_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNoteV1Response) GetNoteId() uint64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

type MultiCreateNotesV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notes []*NewNote `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *MultiCreateNotesV1Request) Reset() {
	*x = MultiCreateNotesV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_note_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiCreateNotesV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiCreateNotesV1Request) ProtoMessage() {}

func (x *MultiCreateNotesV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_note_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiCreateNotesV1Request.ProtoReflect.Descriptor instead.
func (*MultiCreateNotesV1Request) Descriptor() ([]byte, []int) {
	return file_ocp_note_api_proto_rawDescGZIP(), []int{2}
}

func (x *MultiCreateNotesV1Request) GetNotes() []*NewNote {
	if x != nil {
		return x.Notes
	}
	return nil
}

type MultiCreateNotesV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberOfNotesCreated uint64 `protobuf:"varint,1,opt,name=number_of_notes_created,json=numberOfNotesCreated,proto3" json:"number_of_notes_created,omitempty"`
}

func (x *MultiCreateNotesV1Response) Reset() {
	*x = MultiCreateNotesV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_note_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiCreateNotesV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiCreateNotesV1Response) ProtoMessage() {}

func (x *MultiCreateNotesV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_note_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiCreateNotesV1Response.ProtoReflect.Descriptor instead.
func (*MultiCreateNotesV1Response) Descriptor() ([]byte, []int) {
	return file_ocp_note_api_proto_rawDescGZIP(), []int{3}
}

func (x *MultiCreateNotesV1Response) GetNumberOfNotesCreated() uint64 {
	if x != nil {
		return x.NumberOfNotesCreated
	}
	return 0
}

type DescribeNoteV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteId int64 `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
}

func (x *DescribeNoteV1Request) Reset() {
	*x = DescribeNoteV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_note_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeNoteV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeNoteV1Request) ProtoMessage() {}

func (x *DescribeNoteV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_note_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeNoteV1Request.ProtoReflect.Descriptor instead.
func (*DescribeNoteV1Request) Descriptor() ([]byte, []int) {
	return file_ocp_note_api_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeNoteV1Request) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

type DescribeNoteV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Note *Note `protobuf:"bytes,1,opt,name=note,proto3" json:"note,omitempty"`
}

func (x *DescribeNoteV1Response) Reset() {
	*x = DescribeNoteV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_note_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeNoteV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeNoteV1Response) ProtoMessage() {}

func (x *DescribeNoteV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_note_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeNoteV1Response.ProtoReflect.Descriptor instead.
func (*DescribeNoteV1Response) Descriptor() ([]byte, []int) {
	return file_ocp_note_api_proto_rawDescGZIP(), []int{5}
}

func (x *DescribeNoteV1Response) GetNote() *Note {
	if x != nil {
		return x.Note
	}
	return nil
}

type ListNotesV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListNotesV1Request) Reset() {
	*x = ListNotesV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_note_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotesV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesV1Request) ProtoMessage() {}

func (x *ListNotesV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_note_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesV1Request.ProtoReflect.Descriptor instead.
func (*ListNotesV1Request) Descriptor() ([]byte, []int) {
	return file_ocp_note_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListNotesV1Request) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListNotesV1Request) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListNotesV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notes []*Note `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
}

func (x *ListNotesV1Response) Reset() {
	*x = ListNotesV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_note_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotesV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesV1Response) ProtoMessage() {}

func (x *ListNotesV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_note_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesV1Response.ProtoReflect.Descriptor instead.
func (*ListNotesV1Response) Descriptor() ([]byte, []int) {
	return file_ocp_note_api_proto_rawDescGZIP(), []int{7}
}

func (x *ListNotesV1Response) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

type RemoveNoteV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteId int64 `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
}

func (x *RemoveNoteV1Request) Reset() {
	*x = RemoveNoteV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_note_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveNoteV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveNoteV1Request) ProtoMessage() {}

func (x *RemoveNoteV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_note_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveNoteV1Request.ProtoReflect.Descriptor instead.
func (*RemoveNoteV1Request) Descriptor() ([]byte, []int) {
	return file_ocp_note_api_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveNoteV1Request) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

type RemoveNoteV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *RemoveNoteV1Response) Reset() {
	*x = RemoveNoteV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_note_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveNoteV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveNoteV1Response) ProtoMessage() {}

func (x *RemoveNoteV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_note_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveNoteV1Response.ProtoReflect.Descriptor instead.
func (*RemoveNoteV1Response) Descriptor() ([]byte, []int) {
	return file_ocp_note_api_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveNoteV1Response) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClassroomId uint32 `protobuf:"varint,3,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
	DocumentId  uint32 `protobuf:"varint,4,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_note_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_note_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_ocp_note_api_proto_rawDescGZIP(), []int{10}
}

func (x *Note) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Note) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Note) GetClassroomId() uint32 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *Note) GetDocumentId() uint32 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

type NewNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      uint32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClassroomId uint32 `protobuf:"varint,2,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
	DocumentId  uint32 `protobuf:"varint,3,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
}

func (x *NewNote) Reset() {
	*x = NewNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ocp_note_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewNote) ProtoMessage() {}

func (x *NewNote) ProtoReflect() protoreflect.Message {
	mi := &file_ocp_note_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewNote.ProtoReflect.Descriptor instead.
func (*NewNote) Descriptor() ([]byte, []int) {
	return file_ocp_note_api_proto_rawDescGZIP(), []int{11}
}

func (x *NewNote) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *NewNote) GetClassroomId() uint32 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *NewNote) GetDocumentId() uint32 {
	if x != nil {
		return x.DocumentId
	}
	return 0
}

var File_ocp_note_api_proto protoreflect.FileDescriptor

var file_ocp_note_api_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x63, 0x70, 0x2d, 0x6e, 0x6f, 0x74, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x63, 0x70, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02,
	0x20, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x19, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x65, 0x77, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e,
	0x6f, 0x74, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x1a, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x14, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4e, 0x6f, 0x74,
	0x65, 0x73, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x39, 0x0a, 0x15, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06, 0x6e, 0x6f,
	0x74, 0x65, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f,
	0x63, 0x70, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x74, 0x65,
	0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x54, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f,
	0x74, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x20, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x22, 0x02, 0x28, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3f, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x37, 0x0a,
	0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06,
	0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x22, 0x73, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f,
	0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x07, 0x4e, 0x65, 0x77,
	0x4e, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x32, 0xc6, 0x04, 0x0a, 0x0a, 0x4f, 0x63, 0x70, 0x4e, 0x6f, 0x74, 0x65, 0x41, 0x70, 0x69,
	0x12, 0x67, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31,
	0x12, 0x21, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22,
	0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x7a, 0x0a, 0x12, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x56, 0x31, 0x12,
	0x27, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d,
	0x75, 0x6c, 0x74, 0x69, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6e,
	0x6f, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x78, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x12, 0x23, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f,
	0x63, 0x70, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0x65, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x56, 0x31, 0x12, 0x20,
	0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x72, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x4e, 0x6f, 0x74, 0x65, 0x56, 0x31, 0x12, 0x21, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x2f, 0x7b, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f,
	0x6f, 0x63, 0x70, 0x2d, 0x6e, 0x6f, 0x74, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x6e, 0x6f, 0x74, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x63,
	0x70, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ocp_note_api_proto_rawDescOnce sync.Once
	file_ocp_note_api_proto_rawDescData = file_ocp_note_api_proto_rawDesc
)

func file_ocp_note_api_proto_rawDescGZIP() []byte {
	file_ocp_note_api_proto_rawDescOnce.Do(func() {
		file_ocp_note_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_ocp_note_api_proto_rawDescData)
	})
	return file_ocp_note_api_proto_rawDescData
}

var file_ocp_note_api_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_ocp_note_api_proto_goTypes = []interface{}{
	(*CreateNoteV1Request)(nil),        // 0: ocp.note.api.CreateNoteV1Request
	(*CreateNoteV1Response)(nil),       // 1: ocp.note.api.CreateNoteV1Response
	(*MultiCreateNotesV1Request)(nil),  // 2: ocp.note.api.MultiCreateNotesV1Request
	(*MultiCreateNotesV1Response)(nil), // 3: ocp.note.api.MultiCreateNotesV1Response
	(*DescribeNoteV1Request)(nil),      // 4: ocp.note.api.DescribeNoteV1Request
	(*DescribeNoteV1Response)(nil),     // 5: ocp.note.api.DescribeNoteV1Response
	(*ListNotesV1Request)(nil),         // 6: ocp.note.api.ListNotesV1Request
	(*ListNotesV1Response)(nil),        // 7: ocp.note.api.ListNotesV1Response
	(*RemoveNoteV1Request)(nil),        // 8: ocp.note.api.RemoveNoteV1Request
	(*RemoveNoteV1Response)(nil),       // 9: ocp.note.api.RemoveNoteV1Response
	(*Note)(nil),                       // 10: ocp.note.api.Note
	(*NewNote)(nil),                    // 11: ocp.note.api.NewNote
}
var file_ocp_note_api_proto_depIdxs = []int32{
	11, // 0: ocp.note.api.MultiCreateNotesV1Request.notes:type_name -> ocp.note.api.NewNote
	10, // 1: ocp.note.api.DescribeNoteV1Response.note:type_name -> ocp.note.api.Note
	10, // 2: ocp.note.api.ListNotesV1Response.notes:type_name -> ocp.note.api.Note
	0,  // 3: ocp.note.api.OcpNoteApi.CreateNoteV1:input_type -> ocp.note.api.CreateNoteV1Request
	2,  // 4: ocp.note.api.OcpNoteApi.MultiCreateNotesV1:input_type -> ocp.note.api.MultiCreateNotesV1Request
	4,  // 5: ocp.note.api.OcpNoteApi.DescribeNoteV1:input_type -> ocp.note.api.DescribeNoteV1Request
	6,  // 6: ocp.note.api.OcpNoteApi.ListNotesV1:input_type -> ocp.note.api.ListNotesV1Request
	8,  // 7: ocp.note.api.OcpNoteApi.RemoveNoteV1:input_type -> ocp.note.api.RemoveNoteV1Request
	1,  // 8: ocp.note.api.OcpNoteApi.CreateNoteV1:output_type -> ocp.note.api.CreateNoteV1Response
	3,  // 9: ocp.note.api.OcpNoteApi.MultiCreateNotesV1:output_type -> ocp.note.api.MultiCreateNotesV1Response
	5,  // 10: ocp.note.api.OcpNoteApi.DescribeNoteV1:output_type -> ocp.note.api.DescribeNoteV1Response
	7,  // 11: ocp.note.api.OcpNoteApi.ListNotesV1:output_type -> ocp.note.api.ListNotesV1Response
	9,  // 12: ocp.note.api.OcpNoteApi.RemoveNoteV1:output_type -> ocp.note.api.RemoveNoteV1Response
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_ocp_note_api_proto_init() }
func file_ocp_note_api_proto_init() {
	if File_ocp_note_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ocp_note_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNoteV1Request); i {
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
		file_ocp_note_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNoteV1Response); i {
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
		file_ocp_note_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiCreateNotesV1Request); i {
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
		file_ocp_note_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiCreateNotesV1Response); i {
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
		file_ocp_note_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeNoteV1Request); i {
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
		file_ocp_note_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeNoteV1Response); i {
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
		file_ocp_note_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNotesV1Request); i {
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
		file_ocp_note_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNotesV1Response); i {
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
		file_ocp_note_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveNoteV1Request); i {
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
		file_ocp_note_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveNoteV1Response); i {
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
		file_ocp_note_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
		file_ocp_note_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewNote); i {
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
			RawDescriptor: file_ocp_note_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ocp_note_api_proto_goTypes,
		DependencyIndexes: file_ocp_note_api_proto_depIdxs,
		MessageInfos:      file_ocp_note_api_proto_msgTypes,
	}.Build()
	File_ocp_note_api_proto = out.File
	file_ocp_note_api_proto_rawDesc = nil
	file_ocp_note_api_proto_goTypes = nil
	file_ocp_note_api_proto_depIdxs = nil
}
