// Code generated by protoc-gen-go. DO NOT EDIT.
// source: RequestMsg.proto

/*
Package protocal is a generated protocol buffer package.

It is generated from these files:
	RequestMsg.proto

It has these top-level messages:
	LogBean
	Metrics
	BaseInfo
	Request
	Response
*/
package protocal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// [logType, collectTime, hostName, fileName, appName, domain, ip, topic, fileOffset, fileNode, sortedId]
type LogBean struct {
	LogType          *string `protobuf:"bytes,1,opt,name=LogType" json:"LogType,omitempty"`
	CollectTime      *string `protobuf:"bytes,2,opt,name=CollectTime" json:"CollectTime,omitempty"`
	HostName         *string `protobuf:"bytes,3,opt,name=HostName" json:"HostName,omitempty"`
	FileName         *string `protobuf:"bytes,4,opt,name=FileName" json:"FileName,omitempty"`
	AppName          *string `protobuf:"bytes,5,opt,name=AppName" json:"AppName,omitempty"`
	Domain           *string `protobuf:"bytes,6,opt,name=Domain" json:"Domain,omitempty"`
	Ip               *string `protobuf:"bytes,7,opt,name=Ip" json:"Ip,omitempty"`
	Topic            *string `protobuf:"bytes,8,opt,name=Topic" json:"Topic,omitempty"`
	FileOffset       *string `protobuf:"bytes,9,opt,name=FileOffset" json:"FileOffset,omitempty"`
	FileNode         *string `protobuf:"bytes,10,opt,name=FileNode" json:"FileNode,omitempty"`
	SortedId         *string `protobuf:"bytes,11,opt,name=SortedId" json:"SortedId,omitempty"`
	Body             *string `protobuf:"bytes,12,opt,name=Body" json:"Body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogBean) Reset()                    { *m = LogBean{} }
func (m *LogBean) String() string            { return proto.CompactTextString(m) }
func (*LogBean) ProtoMessage()               {}
func (*LogBean) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogBean) GetLogType() string {
	if m != nil && m.LogType != nil {
		return *m.LogType
	}
	return ""
}

func (m *LogBean) GetCollectTime() string {
	if m != nil && m.CollectTime != nil {
		return *m.CollectTime
	}
	return ""
}

func (m *LogBean) GetHostName() string {
	if m != nil && m.HostName != nil {
		return *m.HostName
	}
	return ""
}

func (m *LogBean) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *LogBean) GetAppName() string {
	if m != nil && m.AppName != nil {
		return *m.AppName
	}
	return ""
}

func (m *LogBean) GetDomain() string {
	if m != nil && m.Domain != nil {
		return *m.Domain
	}
	return ""
}

func (m *LogBean) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *LogBean) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *LogBean) GetFileOffset() string {
	if m != nil && m.FileOffset != nil {
		return *m.FileOffset
	}
	return ""
}

func (m *LogBean) GetFileNode() string {
	if m != nil && m.FileNode != nil {
		return *m.FileNode
	}
	return ""
}

func (m *LogBean) GetSortedId() string {
	if m != nil && m.SortedId != nil {
		return *m.SortedId
	}
	return ""
}

func (m *LogBean) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

type Metrics struct {
	Endpoint         *string `protobuf:"bytes,1,opt,name=Endpoint" json:"Endpoint,omitempty"`
	Metric           *string `protobuf:"bytes,2,opt,name=Metric" json:"Metric,omitempty"`
	Value            *string `protobuf:"bytes,3,opt,name=Value" json:"Value,omitempty"`
	Step             *int64  `protobuf:"varint,4,opt,name=Step" json:"Step,omitempty"`
	Type             *string `protobuf:"bytes,5,opt,name=Type" json:"Type,omitempty"`
	Tags             *string `protobuf:"bytes,6,opt,name=Tags" json:"Tags,omitempty"`
	Timestamp        *int64  `protobuf:"varint,7,opt,name=Timestamp" json:"Timestamp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Metrics) Reset()                    { *m = Metrics{} }
func (m *Metrics) String() string            { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()               {}
func (*Metrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Metrics) GetEndpoint() string {
	if m != nil && m.Endpoint != nil {
		return *m.Endpoint
	}
	return ""
}

func (m *Metrics) GetMetric() string {
	if m != nil && m.Metric != nil {
		return *m.Metric
	}
	return ""
}

func (m *Metrics) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *Metrics) GetStep() int64 {
	if m != nil && m.Step != nil {
		return *m.Step
	}
	return 0
}

func (m *Metrics) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Metrics) GetTags() string {
	if m != nil && m.Tags != nil {
		return *m.Tags
	}
	return ""
}

func (m *Metrics) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type BaseInfo struct {
	ProtocalVersion  *int32  `protobuf:"varint,1,req,name=ProtocalVersion" json:"ProtocalVersion,omitempty"`
	Cmd              *int32  `protobuf:"varint,2,req,name=Cmd" json:"Cmd,omitempty"`
	ReqId            *int64  `protobuf:"varint,3,opt,name=ReqId" json:"ReqId,omitempty"`
	ExtendParams     *string `protobuf:"bytes,4,opt,name=ExtendParams" json:"ExtendParams,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BaseInfo) Reset()                    { *m = BaseInfo{} }
func (m *BaseInfo) String() string            { return proto.CompactTextString(m) }
func (*BaseInfo) ProtoMessage()               {}
func (*BaseInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BaseInfo) GetProtocalVersion() int32 {
	if m != nil && m.ProtocalVersion != nil {
		return *m.ProtocalVersion
	}
	return 0
}

func (m *BaseInfo) GetCmd() int32 {
	if m != nil && m.Cmd != nil {
		return *m.Cmd
	}
	return 0
}

func (m *BaseInfo) GetReqId() int64 {
	if m != nil && m.ReqId != nil {
		return *m.ReqId
	}
	return 0
}

func (m *BaseInfo) GetExtendParams() string {
	if m != nil && m.ExtendParams != nil {
		return *m.ExtendParams
	}
	return ""
}

type Request struct {
	BaseInfo         *BaseInfo  `protobuf:"bytes,1,req,name=BaseInfo" json:"BaseInfo,omitempty"`
	Logs             []*LogBean `protobuf:"bytes,2,rep,name=Logs" json:"Logs,omitempty"`
	MertricsValue    []*Metrics `protobuf:"bytes,3,rep,name=MertricsValue" json:"MertricsValue,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Request) GetBaseInfo() *BaseInfo {
	if m != nil {
		return m.BaseInfo
	}
	return nil
}

func (m *Request) GetLogs() []*LogBean {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *Request) GetMertricsValue() []*Metrics {
	if m != nil {
		return m.MertricsValue
	}
	return nil
}

type Response struct {
	BaseInfo         *BaseInfo `protobuf:"bytes,1,req,name=BaseInfo" json:"BaseInfo,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Response) GetBaseInfo() *BaseInfo {
	if m != nil {
		return m.BaseInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*LogBean)(nil), "protocal.LogBean")
	proto.RegisterType((*Metrics)(nil), "protocal.Metrics")
	proto.RegisterType((*BaseInfo)(nil), "protocal.BaseInfo")
	proto.RegisterType((*Request)(nil), "protocal.Request")
	proto.RegisterType((*Response)(nil), "protocal.Response")
}

func init() { proto.RegisterFile("RequestMsg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0x5f, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0xd5, 0xfc, 0x69, 0xd2, 0x9b, 0x64, 0xdd, 0x0c, 0x12, 0x7e, 0xa3, 0x8a, 0x78, 0xc8,
	0x53, 0x85, 0xf6, 0x0d, 0xe8, 0x18, 0x22, 0xd2, 0x0a, 0x53, 0x17, 0xf6, 0x6e, 0xc5, 0xb7, 0x51,
	0xa4, 0xc4, 0xf6, 0x62, 0x4f, 0x30, 0xbe, 0x27, 0xdf, 0x07, 0xe5, 0x36, 0xa1, 0x82, 0x27, 0x9e,
	0xec, 0xf3, 0xf3, 0x95, 0x8f, 0xce, 0xb9, 0x70, 0x79, 0xc0, 0xa7, 0x67, 0xb4, 0x6e, 0x6f, 0x9b,
	0xad, 0x19, 0xb4, 0xd3, 0x2c, 0xa6, 0xa3, 0x16, 0x5d, 0xfe, 0x6b, 0x01, 0xd1, 0x9d, 0x6e, 0x76,
	0x28, 0x14, 0x5b, 0xd3, 0xb5, 0x7a, 0x31, 0xc8, 0x17, 0x9b, 0x45, 0xb1, 0x62, 0xaf, 0x20, 0xb9,
	0xd1, 0x5d, 0x87, 0xb5, 0xab, 0xda, 0x1e, 0xb9, 0x47, 0xf0, 0x12, 0xe2, 0xcf, 0xda, 0xba, 0x2f,
	0xa2, 0x47, 0xee, 0xcf, 0xe4, 0x53, 0xdb, 0x21, 0x91, 0x80, 0xc8, 0x1a, 0xa2, 0x0f, 0xc6, 0x10,
	0x08, 0x09, 0x5c, 0xc0, 0xf2, 0xa3, 0xee, 0x45, 0xab, 0xf8, 0x92, 0x34, 0x80, 0x57, 0x1a, 0x1e,
	0xd1, 0x3d, 0x83, 0xb0, 0xd2, 0xa6, 0xad, 0x79, 0x4c, 0x92, 0x01, 0x8c, 0xbf, 0x7d, 0x3d, 0x1e,
	0x2d, 0x3a, 0xbe, 0xfa, 0xcb, 0x41, 0x4b, 0xe4, 0x30, 0x93, 0x07, 0x3d, 0x38, 0x94, 0xa5, 0xe4,
	0x09, 0x91, 0x14, 0x82, 0x9d, 0x96, 0x2f, 0x3c, 0x1d, 0x55, 0xfe, 0x1d, 0xa2, 0x3d, 0xba, 0xa1,
	0xad, 0xed, 0x38, 0x7a, 0xab, 0xa4, 0xd1, 0xad, 0x72, 0x53, 0xae, 0x0b, 0x58, 0x9e, 0x1e, 0xa7,
	0x48, 0x19, 0x84, 0x8f, 0xa2, 0x7b, 0x9e, 0xf3, 0xa4, 0x10, 0x3c, 0x38, 0x34, 0x94, 0xc5, 0x1f,
	0x15, 0x55, 0x12, 0xce, 0x6f, 0x95, 0x68, 0xec, 0x14, 0xe3, 0x0a, 0x56, 0x63, 0x33, 0xd6, 0x89,
	0xfe, 0x94, 0xc6, 0xcf, 0xbf, 0x41, 0xbc, 0x13, 0x16, 0x4b, 0x75, 0xd4, 0xec, 0x0d, 0xac, 0xef,
	0xa7, 0xa2, 0x1f, 0x71, 0xb0, 0xad, 0x56, 0x7c, 0xb1, 0xf1, 0x8a, 0x90, 0x25, 0xe0, 0xdf, 0xf4,
	0x92, 0x7b, 0x24, 0x32, 0x08, 0x0f, 0xf8, 0x54, 0x4a, 0x72, 0xf7, 0xd9, 0x6b, 0x48, 0x6f, 0x7f,
	0x38, 0x54, 0xf2, 0x5e, 0x0c, 0xa2, 0xb7, 0xa7, 0x46, 0xf3, 0x9f, 0x10, 0x4d, 0x5b, 0x64, 0xef,
	0xce, 0x0e, 0xf4, 0x5d, 0x72, 0xcd, 0xb6, 0xf3, 0x3e, 0xb7, 0x7f, 0xbc, 0xdf, 0x42, 0x70, 0xa7,
	0x1b, 0xcb, 0xbd, 0x8d, 0x5f, 0x24, 0xd7, 0x57, 0xe7, 0x89, 0x79, 0xdb, 0x05, 0x64, 0x7b, 0x1c,
	0xa8, 0xa2, 0x39, 0xfc, 0x3f, 0x93, 0x53, 0x81, 0xf9, 0x7b, 0x88, 0x0f, 0x68, 0x8d, 0x56, 0x16,
	0xff, 0xcf, 0xfc, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x33, 0x47, 0x57, 0xa3, 0x72, 0x02, 0x00,
	0x00,
}
