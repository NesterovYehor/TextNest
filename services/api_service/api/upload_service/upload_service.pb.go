// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: paste_upload_service/upload_service.proto

package paste_upload

import (
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

// Upload request message
type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key            string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`                                             // Unique paste key
	ExpirationDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"` // Expiration date
	Data           []byte                 `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`                                           // Paste content as bytes
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	mi := &file_paste_upload_service_upload_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paste_upload_service_upload_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_paste_upload_service_upload_service_proto_rawDescGZIP(), []int{0}
}

func (x *UploadRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UploadRequest) GetExpirationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationDate
	}
	return nil
}

func (x *UploadRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// Upload response message
type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message        string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`                                     // Response message (e.g., "Upload successful")
	StatusCode     int32                  `protobuf:"varint,2,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`            // HTTP-like status code (e.g., 200 for success)
	ExpirationDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"` // Echo back the expiration date for confirmation
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	mi := &file_paste_upload_service_upload_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_paste_upload_service_upload_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_paste_upload_service_upload_service_proto_rawDescGZIP(), []int{1}
}

func (x *UploadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UploadResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *UploadResponse) GetExpirationDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpirationDate
	}
	return nil
}

var File_paste_upload_service_upload_service_proto protoreflect.FileDescriptor

var file_paste_upload_service_upload_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x61, 0x73, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x61, 0x73,
	0x74, 0x65, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x0d, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x43, 0x0a, 0x0f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x32, 0x50, 0x0a, 0x0b, 0x50, 0x61, 0x73, 0x74,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x73, 0x74, 0x65, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x61, 0x73, 0x74, 0x65, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f,
	0x3b, 0x70, 0x61, 0x73, 0x74, 0x65, 0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paste_upload_service_upload_service_proto_rawDescOnce sync.Once
	file_paste_upload_service_upload_service_proto_rawDescData = file_paste_upload_service_upload_service_proto_rawDesc
)

func file_paste_upload_service_upload_service_proto_rawDescGZIP() []byte {
	file_paste_upload_service_upload_service_proto_rawDescOnce.Do(func() {
		file_paste_upload_service_upload_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_paste_upload_service_upload_service_proto_rawDescData)
	})
	return file_paste_upload_service_upload_service_proto_rawDescData
}

var file_paste_upload_service_upload_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_paste_upload_service_upload_service_proto_goTypes = []any{
	(*UploadRequest)(nil),         // 0: pasteupload.UploadRequest
	(*UploadResponse)(nil),        // 1: pasteupload.UploadResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_paste_upload_service_upload_service_proto_depIdxs = []int32{
	2, // 0: pasteupload.UploadRequest.expiration_date:type_name -> google.protobuf.Timestamp
	2, // 1: pasteupload.UploadResponse.expiration_date:type_name -> google.protobuf.Timestamp
	0, // 2: pasteupload.PasteUpload.Upload:input_type -> pasteupload.UploadRequest
	1, // 3: pasteupload.PasteUpload.Upload:output_type -> pasteupload.UploadResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_paste_upload_service_upload_service_proto_init() }
func file_paste_upload_service_upload_service_proto_init() {
	if File_paste_upload_service_upload_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_paste_upload_service_upload_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paste_upload_service_upload_service_proto_goTypes,
		DependencyIndexes: file_paste_upload_service_upload_service_proto_depIdxs,
		MessageInfos:      file_paste_upload_service_upload_service_proto_msgTypes,
	}.Build()
	File_paste_upload_service_upload_service_proto = out.File
	file_paste_upload_service_upload_service_proto_rawDesc = nil
	file_paste_upload_service_upload_service_proto_goTypes = nil
	file_paste_upload_service_upload_service_proto_depIdxs = nil
}
