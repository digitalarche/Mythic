// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: pushC2GRPC.proto

package services

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

// agent sends along the base64 message like normal and c2 profile adds in c2 profile name
// this is all normal c2 profiles pass along to mythic anyway
// C2ProfileName is required
// RemoteIP should be the IP of the remote connection to the c2 profile if known
// Message is a base64 decoded message if you want the c2 profile to base64 decode what the agents are saying
// Base64Message is what an agent would normally send. This is mutually exclusive with Message
// TrackingID is some custom string that the c2 server provides so that it can correlate input with output from the stream
type PushC2MessageFromAgent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C2ProfileName     string `protobuf:"bytes,1,opt,name=C2ProfileName,proto3" json:"C2ProfileName,omitempty"`
	RemoteIP          string `protobuf:"bytes,2,opt,name=RemoteIP,proto3" json:"RemoteIP,omitempty"`
	Message           []byte `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	OuterUUID         string `protobuf:"bytes,4,opt,name=OuterUUID,proto3" json:"OuterUUID,omitempty"`
	Base64Message     []byte `protobuf:"bytes,5,opt,name=Base64Message,proto3" json:"Base64Message,omitempty"`
	TrackingID        string `protobuf:"bytes,6,opt,name=TrackingID,proto3" json:"TrackingID,omitempty"`
	AgentDisconnected bool   `protobuf:"varint,7,opt,name=AgentDisconnected,proto3" json:"AgentDisconnected,omitempty"`
}

func (x *PushC2MessageFromAgent) Reset() {
	*x = PushC2MessageFromAgent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pushC2GRPC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushC2MessageFromAgent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushC2MessageFromAgent) ProtoMessage() {}

func (x *PushC2MessageFromAgent) ProtoReflect() protoreflect.Message {
	mi := &file_pushC2GRPC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushC2MessageFromAgent.ProtoReflect.Descriptor instead.
func (*PushC2MessageFromAgent) Descriptor() ([]byte, []int) {
	return file_pushC2GRPC_proto_rawDescGZIP(), []int{0}
}

func (x *PushC2MessageFromAgent) GetC2ProfileName() string {
	if x != nil {
		return x.C2ProfileName
	}
	return ""
}

func (x *PushC2MessageFromAgent) GetRemoteIP() string {
	if x != nil {
		return x.RemoteIP
	}
	return ""
}

func (x *PushC2MessageFromAgent) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *PushC2MessageFromAgent) GetOuterUUID() string {
	if x != nil {
		return x.OuterUUID
	}
	return ""
}

func (x *PushC2MessageFromAgent) GetBase64Message() []byte {
	if x != nil {
		return x.Base64Message
	}
	return nil
}

func (x *PushC2MessageFromAgent) GetTrackingID() string {
	if x != nil {
		return x.TrackingID
	}
	return ""
}

func (x *PushC2MessageFromAgent) GetAgentDisconnected() bool {
	if x != nil {
		return x.AgentDisconnected
	}
	return false
}

// mythic sends along success/error information if any
// mythic sends along the messages
type PushC2MessageFromMythic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool   `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Error      string `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	Message    []byte `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	TrackingID string `protobuf:"bytes,4,opt,name=TrackingID,proto3" json:"TrackingID,omitempty"`
}

func (x *PushC2MessageFromMythic) Reset() {
	*x = PushC2MessageFromMythic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pushC2GRPC_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushC2MessageFromMythic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushC2MessageFromMythic) ProtoMessage() {}

func (x *PushC2MessageFromMythic) ProtoReflect() protoreflect.Message {
	mi := &file_pushC2GRPC_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushC2MessageFromMythic.ProtoReflect.Descriptor instead.
func (*PushC2MessageFromMythic) Descriptor() ([]byte, []int) {
	return file_pushC2GRPC_proto_rawDescGZIP(), []int{1}
}

func (x *PushC2MessageFromMythic) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PushC2MessageFromMythic) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *PushC2MessageFromMythic) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *PushC2MessageFromMythic) GetTrackingID() string {
	if x != nil {
		return x.TrackingID
	}
	return ""
}

var File_pushC2GRPC_proto protoreflect.FileDescriptor

var file_pushC2GRPC_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x75, 0x73, 0x68, 0x43, 0x32, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x70, 0x75, 0x73, 0x68, 0x43, 0x32, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x22, 0x86, 0x02, 0x0a, 0x16, 0x50, 0x75, 0x73, 0x68, 0x43, 0x32, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x43, 0x32, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x32, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x50, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x50, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x75, 0x74,
	0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4f, 0x75,
	0x74, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x61, 0x73, 0x65, 0x36,
	0x34, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d,
	0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x2c, 0x0a,
	0x11, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x17,
	0x50, 0x75, 0x73, 0x68, 0x43, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x4d, 0x79, 0x74, 0x68, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49,
	0x44, 0x32, 0xef, 0x01, 0x0a, 0x06, 0x50, 0x75, 0x73, 0x68, 0x43, 0x32, 0x12, 0x6d, 0x0a, 0x14,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x75, 0x73, 0x68, 0x43, 0x32, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x12, 0x26, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x43, 0x32, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x43, 0x32, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x27, 0x2e, 0x70,
	0x75, 0x73, 0x68, 0x43, 0x32, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x43, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d,
	0x79, 0x74, 0x68, 0x69, 0x63, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x76, 0x0a, 0x1d, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x50, 0x75, 0x73, 0x68, 0x43, 0x32, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x4f, 0x6e, 0x65, 0x54, 0x6f, 0x4d, 0x61, 0x6e, 0x79, 0x12, 0x26, 0x2e, 0x70,
	0x75, 0x73, 0x68, 0x43, 0x32, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x43, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x43, 0x32, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x43, 0x32, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4d, 0x79, 0x74, 0x68, 0x69, 0x63, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x74, 0x73, 0x2d, 0x61, 0x2d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x4d, 0x79, 0x74, 0x68, 0x69, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pushC2GRPC_proto_rawDescOnce sync.Once
	file_pushC2GRPC_proto_rawDescData = file_pushC2GRPC_proto_rawDesc
)

func file_pushC2GRPC_proto_rawDescGZIP() []byte {
	file_pushC2GRPC_proto_rawDescOnce.Do(func() {
		file_pushC2GRPC_proto_rawDescData = protoimpl.X.CompressGZIP(file_pushC2GRPC_proto_rawDescData)
	})
	return file_pushC2GRPC_proto_rawDescData
}

var file_pushC2GRPC_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pushC2GRPC_proto_goTypes = []interface{}{
	(*PushC2MessageFromAgent)(nil),  // 0: pushC2Services.PushC2MessageFromAgent
	(*PushC2MessageFromMythic)(nil), // 1: pushC2Services.PushC2MessageFromMythic
}
var file_pushC2GRPC_proto_depIdxs = []int32{
	0, // 0: pushC2Services.PushC2.StartPushC2Streaming:input_type -> pushC2Services.PushC2MessageFromAgent
	0, // 1: pushC2Services.PushC2.StartPushC2StreamingOneToMany:input_type -> pushC2Services.PushC2MessageFromAgent
	1, // 2: pushC2Services.PushC2.StartPushC2Streaming:output_type -> pushC2Services.PushC2MessageFromMythic
	1, // 3: pushC2Services.PushC2.StartPushC2StreamingOneToMany:output_type -> pushC2Services.PushC2MessageFromMythic
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pushC2GRPC_proto_init() }
func file_pushC2GRPC_proto_init() {
	if File_pushC2GRPC_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pushC2GRPC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushC2MessageFromAgent); i {
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
		file_pushC2GRPC_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushC2MessageFromMythic); i {
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
			RawDescriptor: file_pushC2GRPC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pushC2GRPC_proto_goTypes,
		DependencyIndexes: file_pushC2GRPC_proto_depIdxs,
		MessageInfos:      file_pushC2GRPC_proto_msgTypes,
	}.Build()
	File_pushC2GRPC_proto = out.File
	file_pushC2GRPC_proto_rawDesc = nil
	file_pushC2GRPC_proto_goTypes = nil
	file_pushC2GRPC_proto_depIdxs = nil
}
