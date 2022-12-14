// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: thingmodel/thingmodel.proto

package thingmodel

import (
	drivercommon "github.com/winc-link/edge-driver-proto/drivercommon"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ThingModelOperationType int32

const (
	ThingModelOperationType_UNSPECIFIED                      ThingModelOperationType = 0
	ThingModelOperationType_MODEL_GET                        ThingModelOperationType = 1  // 设备上报向云端获取物模型查询请求
	ThingModelOperationType_MODEL_GET_RESPONSE               ThingModelOperationType = 2  // 云端响应设备上报向云端获取物模型查询请求
	ThingModelOperationType_PROPERTY_REPORT                  ThingModelOperationType = 3  // 设备上报向云端上报设备属性值
	ThingModelOperationType_PROPERTY_REPORT_RESPONSE         ThingModelOperationType = 4  // 云端响应设备上报向云端上报设备属性值
	ThingModelOperationType_PROPERTY_SET                     ThingModelOperationType = 5  // 云端下发属性设置
	ThingModelOperationType_PROPERTY_SET_RESPONSE            ThingModelOperationType = 6  // 设备响应云端下发属性设置
	ThingModelOperationType_PROPERTY_GET                     ThingModelOperationType = 7  // 云端下发属性查询请求
	ThingModelOperationType_PROPERTY_GET_RESPONSE            ThingModelOperationType = 8  // 设备响应云端下发属性查询请求
	ThingModelOperationType_ACTION_EXECUTE                   ThingModelOperationType = 9  // 云端下发设备动作执行消息
	ThingModelOperationType_ACTION_EXECUTE_RESPONSE          ThingModelOperationType = 10 // 设备响应云端下发设备动作执行消息
	ThingModelOperationType_EVENT_TRIGGER                    ThingModelOperationType = 11 // 设备上报事件触发消息
	ThingModelOperationType_EVENT_TRIGGER_RESPONSE           ThingModelOperationType = 12 // 云端响应设备上报事件触发消息
	ThingModelOperationType_DATA_BATCH_REPORT                ThingModelOperationType = 13 // 设备上报 批量上报属性或事件数据
	ThingModelOperationType_DATA_BATCH_REPORT_RESPONSE       ThingModelOperationType = 14 // 云端响应设备上报 批量上报属性或事件数据
	ThingModelOperationType_PROPERTY_DESIRED_GET             ThingModelOperationType = 15 // 设备上报向云端获取属性期望值请求
	ThingModelOperationType_PROPERTY_DESIRED_GET_RESPONSE    ThingModelOperationType = 16 // 云端响应设备上报向云端获取属性期望值请求
	ThingModelOperationType_PROPERTY_DESIRED_DELETE          ThingModelOperationType = 17 // 设备上报向云端发送清除属性期望值请求
	ThingModelOperationType_PROPERTY_DESIRED_DELETE_RESPONSE ThingModelOperationType = 18 // 云端响应设备上报向云端发送清除属性期望值请求
)

// Enum value maps for ThingModelOperationType.
var (
	ThingModelOperationType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "MODEL_GET",
		2:  "MODEL_GET_RESPONSE",
		3:  "PROPERTY_REPORT",
		4:  "PROPERTY_REPORT_RESPONSE",
		5:  "PROPERTY_SET",
		6:  "PROPERTY_SET_RESPONSE",
		7:  "PROPERTY_GET",
		8:  "PROPERTY_GET_RESPONSE",
		9:  "ACTION_EXECUTE",
		10: "ACTION_EXECUTE_RESPONSE",
		11: "EVENT_TRIGGER",
		12: "EVENT_TRIGGER_RESPONSE",
		13: "DATA_BATCH_REPORT",
		14: "DATA_BATCH_REPORT_RESPONSE",
		15: "PROPERTY_DESIRED_GET",
		16: "PROPERTY_DESIRED_GET_RESPONSE",
		17: "PROPERTY_DESIRED_DELETE",
		18: "PROPERTY_DESIRED_DELETE_RESPONSE",
	}
	ThingModelOperationType_value = map[string]int32{
		"UNSPECIFIED":                      0,
		"MODEL_GET":                        1,
		"MODEL_GET_RESPONSE":               2,
		"PROPERTY_REPORT":                  3,
		"PROPERTY_REPORT_RESPONSE":         4,
		"PROPERTY_SET":                     5,
		"PROPERTY_SET_RESPONSE":            6,
		"PROPERTY_GET":                     7,
		"PROPERTY_GET_RESPONSE":            8,
		"ACTION_EXECUTE":                   9,
		"ACTION_EXECUTE_RESPONSE":          10,
		"EVENT_TRIGGER":                    11,
		"EVENT_TRIGGER_RESPONSE":           12,
		"DATA_BATCH_REPORT":                13,
		"DATA_BATCH_REPORT_RESPONSE":       14,
		"PROPERTY_DESIRED_GET":             15,
		"PROPERTY_DESIRED_GET_RESPONSE":    16,
		"PROPERTY_DESIRED_DELETE":          17,
		"PROPERTY_DESIRED_DELETE_RESPONSE": 18,
	}
)

func (x ThingModelOperationType) Enum() *ThingModelOperationType {
	p := new(ThingModelOperationType)
	*p = x
	return p
}

func (x ThingModelOperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ThingModelOperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_thingmodel_thingmodel_proto_enumTypes[0].Descriptor()
}

func (ThingModelOperationType) Type() protoreflect.EnumType {
	return &file_thingmodel_thingmodel_proto_enumTypes[0]
}

func (x ThingModelOperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ThingModelOperationType.Descriptor instead.
func (ThingModelOperationType) EnumDescriptor() ([]byte, []int) {
	return file_thingmodel_thingmodel_proto_rawDescGZIP(), []int{0}
}

// 物模型
type ThingModelMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *drivercommon.BaseRequestMessage `protobuf:"bytes,5,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	Cid         string                           `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"` // 下发的目标设备id
	OpType      ThingModelOperationType          `protobuf:"varint,2,opt,name=op_type,json=opType,proto3,enum=thingmodel.ThingModelOperationType" json:"op_type,omitempty"`
	Data        string                           `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ThingModelMsg) Reset() {
	*x = ThingModelMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thingmodel_thingmodel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThingModelMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThingModelMsg) ProtoMessage() {}

func (x *ThingModelMsg) ProtoReflect() protoreflect.Message {
	mi := &file_thingmodel_thingmodel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThingModelMsg.ProtoReflect.Descriptor instead.
func (*ThingModelMsg) Descriptor() ([]byte, []int) {
	return file_thingmodel_thingmodel_proto_rawDescGZIP(), []int{0}
}

func (x *ThingModelMsg) GetBaseRequest() *drivercommon.BaseRequestMessage {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *ThingModelMsg) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

func (x *ThingModelMsg) GetOpType() ThingModelOperationType {
	if x != nil {
		return x.OpType
	}
	return ThingModelOperationType_UNSPECIFIED
}

func (x *ThingModelMsg) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_thingmodel_thingmodel_proto protoreflect.FileDescriptor

var file_thingmodel_thingmodel_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x0d, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x4d, 0x73, 0x67, 0x12, 0x42, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x12, 0x3c, 0x0a, 0x07, 0x6f, 0x70, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x06, 0x6f, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0xf5, 0x03, 0x0a, 0x17,
	0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x4f, 0x44, 0x45,
	0x4c, 0x5f, 0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x4f, 0x44, 0x45, 0x4c,
	0x5f, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x02, 0x12,
	0x13, 0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x52, 0x45, 0x50, 0x4f,
	0x52, 0x54, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59,
	0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45,
	0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x53,
	0x45, 0x54, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59,
	0x5f, 0x53, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x06, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x47, 0x45, 0x54, 0x10,
	0x07, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x47, 0x45,
	0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x08, 0x12, 0x12, 0x0a, 0x0e,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x10, 0x09,
	0x12, 0x1b, 0x0a, 0x17, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55,
	0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x0a, 0x12, 0x11, 0x0a,
	0x0d, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x10, 0x0b,
	0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45,
	0x52, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x0c, 0x12, 0x15, 0x0a, 0x11,
	0x44, 0x41, 0x54, 0x41, 0x5f, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52,
	0x54, 0x10, 0x0d, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x42, 0x41, 0x54, 0x43,
	0x48, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53,
	0x45, 0x10, 0x0e, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f,
	0x44, 0x45, 0x53, 0x49, 0x52, 0x45, 0x44, 0x5f, 0x47, 0x45, 0x54, 0x10, 0x0f, 0x12, 0x21, 0x0a,
	0x1d, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x53, 0x49, 0x52, 0x45,
	0x44, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x10,
	0x12, 0x1b, 0x0a, 0x17, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x53,
	0x49, 0x52, 0x45, 0x44, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x11, 0x12, 0x24, 0x0a,
	0x20, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x59, 0x5f, 0x44, 0x45, 0x53, 0x49, 0x52, 0x45,
	0x44, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53,
	0x45, 0x10, 0x12, 0x32, 0x61, 0x0a, 0x13, 0x54, 0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x55, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x13, 0x54, 0x68,
	0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x19, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54,
	0x68, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x73, 0x67, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x6e, 0x63, 0x2d, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x65,
	0x64, 0x67, 0x65, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_thingmodel_thingmodel_proto_rawDescOnce sync.Once
	file_thingmodel_thingmodel_proto_rawDescData = file_thingmodel_thingmodel_proto_rawDesc
)

func file_thingmodel_thingmodel_proto_rawDescGZIP() []byte {
	file_thingmodel_thingmodel_proto_rawDescOnce.Do(func() {
		file_thingmodel_thingmodel_proto_rawDescData = protoimpl.X.CompressGZIP(file_thingmodel_thingmodel_proto_rawDescData)
	})
	return file_thingmodel_thingmodel_proto_rawDescData
}

var file_thingmodel_thingmodel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_thingmodel_thingmodel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_thingmodel_thingmodel_proto_goTypes = []interface{}{
	(ThingModelOperationType)(0),            // 0: thingmodel.ThingModelOperationType
	(*ThingModelMsg)(nil),                   // 1: thingmodel.ThingModelMsg
	(*drivercommon.BaseRequestMessage)(nil), // 2: drivercommon.BaseRequestMessage
	(*emptypb.Empty)(nil),                   // 3: google.protobuf.Empty
}
var file_thingmodel_thingmodel_proto_depIdxs = []int32{
	2, // 0: thingmodel.ThingModelMsg.baseRequest:type_name -> drivercommon.BaseRequestMessage
	0, // 1: thingmodel.ThingModelMsg.op_type:type_name -> thingmodel.ThingModelOperationType
	1, // 2: thingmodel.ThingModelUpService.ThingModelMsgReport:input_type -> thingmodel.ThingModelMsg
	3, // 3: thingmodel.ThingModelUpService.ThingModelMsgReport:output_type -> google.protobuf.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_thingmodel_thingmodel_proto_init() }
func file_thingmodel_thingmodel_proto_init() {
	if File_thingmodel_thingmodel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thingmodel_thingmodel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThingModelMsg); i {
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
			RawDescriptor: file_thingmodel_thingmodel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thingmodel_thingmodel_proto_goTypes,
		DependencyIndexes: file_thingmodel_thingmodel_proto_depIdxs,
		EnumInfos:         file_thingmodel_thingmodel_proto_enumTypes,
		MessageInfos:      file_thingmodel_thingmodel_proto_msgTypes,
	}.Build()
	File_thingmodel_thingmodel_proto = out.File
	file_thingmodel_thingmodel_proto_rawDesc = nil
	file_thingmodel_thingmodel_proto_goTypes = nil
	file_thingmodel_thingmodel_proto_depIdxs = nil
}
