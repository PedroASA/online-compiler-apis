// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: ProtoBuf/code.proto

package apis

import (
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

type CodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Input string `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *CodeRequest) Reset() {
	*x = CodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBuf_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeRequest) ProtoMessage() {}

func (x *CodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBuf_code_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeRequest.ProtoReflect.Descriptor instead.
func (*CodeRequest) Descriptor() ([]byte, []int) {
	return file_ProtoBuf_code_proto_rawDescGZIP(), []int{0}
}

func (x *CodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CodeRequest) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type CodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutMessage string `protobuf:"bytes,1,opt,name=out_message,json=outMessage,proto3" json:"out_message,omitempty"`
	ErrMessage string `protobuf:"bytes,2,opt,name=err_message,json=errMessage,proto3" json:"err_message,omitempty"`
}

func (x *CodeReply) Reset() {
	*x = CodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProtoBuf_code_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeReply) ProtoMessage() {}

func (x *CodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_ProtoBuf_code_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeReply.ProtoReflect.Descriptor instead.
func (*CodeReply) Descriptor() ([]byte, []int) {
	return file_ProtoBuf_code_proto_rawDescGZIP(), []int{1}
}

func (x *CodeReply) GetOutMessage() string {
	if x != nil {
		return x.OutMessage
	}
	return ""
}

func (x *CodeReply) GetErrMessage() string {
	if x != nil {
		return x.ErrMessage
	}
	return ""
}

var File_ProtoBuf_code_proto protoreflect.FileDescriptor

var file_ProtoBuf_code_proto_rawDesc = []byte{
	0x0a, 0x13, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x37, 0x0a, 0x0b, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x22, 0x4d, 0x0a, 0x09, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x72, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x34, 0x0a, 0x05, 0x52, 0x75, 0x6e, 0x4a, 0x53, 0x12, 0x2b, 0x0a, 0x03,
	0x52, 0x75, 0x6e, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x65, 0x0a, 0x05, 0x52, 0x75, 0x6e,
	0x48, 0x53, 0x12, 0x2f, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x12, 0x11, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x03, 0x52, 0x75, 0x6e, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x32, 0x66, 0x0a, 0x06, 0x52, 0x75, 0x6e, 0x43, 0x70, 0x70, 0x12, 0x2f, 0x0a, 0x07, 0x43, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x03, 0x52,
	0x75, 0x6e, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ProtoBuf_code_proto_rawDescOnce sync.Once
	file_ProtoBuf_code_proto_rawDescData = file_ProtoBuf_code_proto_rawDesc
)

func file_ProtoBuf_code_proto_rawDescGZIP() []byte {
	file_ProtoBuf_code_proto_rawDescOnce.Do(func() {
		file_ProtoBuf_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_ProtoBuf_code_proto_rawDescData)
	})
	return file_ProtoBuf_code_proto_rawDescData
}

var file_ProtoBuf_code_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ProtoBuf_code_proto_goTypes = []interface{}{
	(*CodeRequest)(nil), // 0: main.CodeRequest
	(*CodeReply)(nil),   // 1: main.CodeReply
}
var file_ProtoBuf_code_proto_depIdxs = []int32{
	0, // 0: main.RunJS.Run:input_type -> main.CodeRequest
	0, // 1: main.RunHS.Compile:input_type -> main.CodeRequest
	0, // 2: main.RunHS.Run:input_type -> main.CodeRequest
	0, // 3: main.RunCpp.Compile:input_type -> main.CodeRequest
	0, // 4: main.RunCpp.Run:input_type -> main.CodeRequest
	1, // 5: main.RunJS.Run:output_type -> main.CodeReply
	1, // 6: main.RunHS.Compile:output_type -> main.CodeReply
	1, // 7: main.RunHS.Run:output_type -> main.CodeReply
	1, // 8: main.RunCpp.Compile:output_type -> main.CodeReply
	1, // 9: main.RunCpp.Run:output_type -> main.CodeReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ProtoBuf_code_proto_init() }
func file_ProtoBuf_code_proto_init() {
	if File_ProtoBuf_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ProtoBuf_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeRequest); i {
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
		file_ProtoBuf_code_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeReply); i {
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
			RawDescriptor: file_ProtoBuf_code_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_ProtoBuf_code_proto_goTypes,
		DependencyIndexes: file_ProtoBuf_code_proto_depIdxs,
		MessageInfos:      file_ProtoBuf_code_proto_msgTypes,
	}.Build()
	File_ProtoBuf_code_proto = out.File
	file_ProtoBuf_code_proto_rawDesc = nil
	file_ProtoBuf_code_proto_goTypes = nil
	file_ProtoBuf_code_proto_depIdxs = nil
}
