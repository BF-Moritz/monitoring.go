// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: grpc_api/monitoring.proto

package grpc_api

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

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Client    string     `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Body      *StatsType `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	TimeStamp string     `protobuf:"bytes,4,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_api_monitoring_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_api_monitoring_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_grpc_api_monitoring_proto_rawDescGZIP(), []int{0}
}

func (x *Ping) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Ping) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *Ping) GetBody() *StatsType {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Ping) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

type StatsType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Memory *MemoryStats `protobuf:"bytes,1,opt,name=Memory,proto3" json:"Memory,omitempty"`
	CPU    *CpuStats    `protobuf:"bytes,2,opt,name=CPU,proto3" json:"CPU,omitempty"`
}

func (x *StatsType) Reset() {
	*x = StatsType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_api_monitoring_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsType) ProtoMessage() {}

func (x *StatsType) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_api_monitoring_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsType.ProtoReflect.Descriptor instead.
func (*StatsType) Descriptor() ([]byte, []int) {
	return file_grpc_api_monitoring_proto_rawDescGZIP(), []int{1}
}

func (x *StatsType) GetMemory() *MemoryStats {
	if x != nil {
		return x.Memory
	}
	return nil
}

func (x *StatsType) GetCPU() *CpuStats {
	if x != nil {
		return x.CPU
	}
	return nil
}

type MemoryStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total  uint64  `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Used   float32 `protobuf:"fixed32,2,opt,name=Used,proto3" json:"Used,omitempty"`
	Cached float32 `protobuf:"fixed32,3,opt,name=Cached,proto3" json:"Cached,omitempty"`
	Free   float32 `protobuf:"fixed32,4,opt,name=Free,proto3" json:"Free,omitempty"`
}

func (x *MemoryStats) Reset() {
	*x = MemoryStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_api_monitoring_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemoryStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryStats) ProtoMessage() {}

func (x *MemoryStats) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_api_monitoring_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryStats.ProtoReflect.Descriptor instead.
func (*MemoryStats) Descriptor() ([]byte, []int) {
	return file_grpc_api_monitoring_proto_rawDescGZIP(), []int{2}
}

func (x *MemoryStats) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MemoryStats) GetUsed() float32 {
	if x != nil {
		return x.Used
	}
	return 0
}

func (x *MemoryStats) GetCached() float32 {
	if x != nil {
		return x.Cached
	}
	return 0
}

func (x *MemoryStats) GetFree() float32 {
	if x != nil {
		return x.Free
	}
	return 0
}

type CpuStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usage float32 `protobuf:"fixed32,1,opt,name=Usage,proto3" json:"Usage,omitempty"`
}

func (x *CpuStats) Reset() {
	*x = CpuStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_api_monitoring_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpuStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpuStats) ProtoMessage() {}

func (x *CpuStats) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_api_monitoring_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpuStats.ProtoReflect.Descriptor instead.
func (*CpuStats) Descriptor() ([]byte, []int) {
	return file_grpc_api_monitoring_proto_rawDescGZIP(), []int{3}
}

func (x *CpuStats) GetUsage() float32 {
	if x != nil {
		return x.Usage
	}
	return 0
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_api_monitoring_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_api_monitoring_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_grpc_api_monitoring_proto_rawDescGZIP(), []int{4}
}

func (x *Pong) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_grpc_api_monitoring_proto protoreflect.FileDescriptor

var file_grpc_api_monitoring_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x77, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x64, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a,
	0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x26,
	0x0a, 0x03, 0x43, 0x50, 0x55, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x70, 0x75, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x03, 0x43, 0x50, 0x55, 0x22, 0x63, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x55, 0x73, 0x65, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x06, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x65, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x46, 0x72, 0x65, 0x65, 0x22, 0x20, 0x0a, 0x08, 0x43,
	0x70, 0x75, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0x16, 0x0a,
	0x04, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0x41, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x1a, 0x10, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x42, 0x46, 0x2d, 0x4d, 0x6f, 0x72, 0x69, 0x74, 0x7a,
	0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_api_monitoring_proto_rawDescOnce sync.Once
	file_grpc_api_monitoring_proto_rawDescData = file_grpc_api_monitoring_proto_rawDesc
)

func file_grpc_api_monitoring_proto_rawDescGZIP() []byte {
	file_grpc_api_monitoring_proto_rawDescOnce.Do(func() {
		file_grpc_api_monitoring_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_api_monitoring_proto_rawDescData)
	})
	return file_grpc_api_monitoring_proto_rawDescData
}

var file_grpc_api_monitoring_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_api_monitoring_proto_goTypes = []interface{}{
	(*Ping)(nil),        // 0: monitoring.Ping
	(*StatsType)(nil),   // 1: monitoring.StatsType
	(*MemoryStats)(nil), // 2: monitoring.MemoryStats
	(*CpuStats)(nil),    // 3: monitoring.CpuStats
	(*Pong)(nil),        // 4: monitoring.Pong
}
var file_grpc_api_monitoring_proto_depIdxs = []int32{
	1, // 0: monitoring.Ping.body:type_name -> monitoring.StatsType
	2, // 1: monitoring.StatsType.Memory:type_name -> monitoring.MemoryStats
	3, // 2: monitoring.StatsType.CPU:type_name -> monitoring.CpuStats
	0, // 3: monitoring.ChatService.SendStatus:input_type -> monitoring.Ping
	4, // 4: monitoring.ChatService.SendStatus:output_type -> monitoring.Pong
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_api_monitoring_proto_init() }
func file_grpc_api_monitoring_proto_init() {
	if File_grpc_api_monitoring_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_api_monitoring_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_grpc_api_monitoring_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsType); i {
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
		file_grpc_api_monitoring_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemoryStats); i {
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
		file_grpc_api_monitoring_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpuStats); i {
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
		file_grpc_api_monitoring_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
			RawDescriptor: file_grpc_api_monitoring_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_api_monitoring_proto_goTypes,
		DependencyIndexes: file_grpc_api_monitoring_proto_depIdxs,
		MessageInfos:      file_grpc_api_monitoring_proto_msgTypes,
	}.Build()
	File_grpc_api_monitoring_proto = out.File
	file_grpc_api_monitoring_proto_rawDesc = nil
	file_grpc_api_monitoring_proto_goTypes = nil
	file_grpc_api_monitoring_proto_depIdxs = nil
}
