// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: rusprofile.proto

package proto

import (
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

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inn string `protobuf:"bytes,1,opt,name=inn,proto3" json:"inn,omitempty"`
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rusprofile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rusprofile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_rusprofile_proto_rawDescGZIP(), []int{0}
}

func (x *GetInfoRequest) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

type InfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inn     string `protobuf:"bytes,1,opt,name=inn,proto3" json:"inn,omitempty"`
	Kpp     string `protobuf:"bytes,2,opt,name=kpp,proto3" json:"kpp,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CeoName string `protobuf:"bytes,4,opt,name=ceo_name,json=ceoName,proto3" json:"ceo_name,omitempty"`
}

func (x *InfoResponse) Reset() {
	*x = InfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rusprofile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfoResponse) ProtoMessage() {}

func (x *InfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rusprofile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfoResponse.ProtoReflect.Descriptor instead.
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return file_rusprofile_proto_rawDescGZIP(), []int{1}
}

func (x *InfoResponse) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

func (x *InfoResponse) GetKpp() string {
	if x != nil {
		return x.Kpp
	}
	return ""
}

func (x *InfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InfoResponse) GetCeoName() string {
	if x != nil {
		return x.CeoName
	}
	return ""
}

var File_rusprofile_proto protoreflect.FileDescriptor

var file_rusprofile_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x72, 0x75, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e,
	0x22, 0x61, 0x0a, 0x0c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x6e, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x70, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x65, 0x6f, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x65, 0x6f, 0x4e,
	0x61, 0x6d, 0x65, 0x32, 0x68, 0x0a, 0x11, 0x52, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x72, 0x75, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x72, 0x75, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0c, 0x12, 0x0a, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x69, 0x6e, 0x6e, 0x7d, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rusprofile_proto_rawDescOnce sync.Once
	file_rusprofile_proto_rawDescData = file_rusprofile_proto_rawDesc
)

func file_rusprofile_proto_rawDescGZIP() []byte {
	file_rusprofile_proto_rawDescOnce.Do(func() {
		file_rusprofile_proto_rawDescData = protoimpl.X.CompressGZIP(file_rusprofile_proto_rawDescData)
	})
	return file_rusprofile_proto_rawDescData
}

var file_rusprofile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rusprofile_proto_goTypes = []interface{}{
	(*GetInfoRequest)(nil), // 0: russervice.GetInfoRequest
	(*InfoResponse)(nil),   // 1: russervice.InfoResponse
}
var file_rusprofile_proto_depIdxs = []int32{
	0, // 0: russervice.RusProfileService.GetInfo:input_type -> russervice.GetInfoRequest
	1, // 1: russervice.RusProfileService.GetInfo:output_type -> russervice.InfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rusprofile_proto_init() }
func file_rusprofile_proto_init() {
	if File_rusprofile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rusprofile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoRequest); i {
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
		file_rusprofile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfoResponse); i {
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
			RawDescriptor: file_rusprofile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rusprofile_proto_goTypes,
		DependencyIndexes: file_rusprofile_proto_depIdxs,
		MessageInfos:      file_rusprofile_proto_msgTypes,
	}.Build()
	File_rusprofile_proto = out.File
	file_rusprofile_proto_rawDesc = nil
	file_rusprofile_proto_goTypes = nil
	file_rusprofile_proto_depIdxs = nil
}