// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.3
// source: proto/rac/v1/rac.proto

package rac

import (
	context "context"
	control_plane "github.com/scionproto/scion/pkg/proto/control_plane"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type InterfaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	IA       uint64 `protobuf:"varint,2,opt,name=IA,proto3" json:"IA,omitempty"`
	LinkType uint32 `protobuf:"varint,3,opt,name=LinkType,proto3" json:"LinkType,omitempty"`
	RemoteID uint32 `protobuf:"varint,4,opt,name=RemoteID,proto3" json:"RemoteID,omitempty"`
	MTU      uint32 `protobuf:"varint,5,opt,name=MTU,proto3" json:"MTU,omitempty"`
}

func (x *InterfaceInfo) Reset() {
	*x = InterfaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rac_v1_rac_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterfaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterfaceInfo) ProtoMessage() {}

func (x *InterfaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rac_v1_rac_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterfaceInfo.ProtoReflect.Descriptor instead.
func (*InterfaceInfo) Descriptor() ([]byte, []int) {
	return file_proto_rac_v1_rac_proto_rawDescGZIP(), []int{0}
}

func (x *InterfaceInfo) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *InterfaceInfo) GetIA() uint64 {
	if x != nil {
		return x.IA
	}
	return 0
}

func (x *InterfaceInfo) GetLinkType() uint32 {
	if x != nil {
		return x.LinkType
	}
	return 0
}

func (x *InterfaceInfo) GetRemoteID() uint32 {
	if x != nil {
		return x.RemoteID
	}
	return 0
}

func (x *InterfaceInfo) GetMTU() uint32 {
	if x != nil {
		return x.MTU
	}
	return 0
}

type ExecutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beacons       []*control_plane.IRECBeacon `protobuf:"bytes,1,rep,name=beacons,proto3" json:"beacons,omitempty"`
	AlgorithmHash []byte                      `protobuf:"bytes,2,opt,name=algorithmHash,proto3" json:"algorithmHash,omitempty"`
	Intfs         []*InterfaceInfo            `protobuf:"bytes,3,rep,name=intfs,proto3" json:"intfs,omitempty"`
	PropIntfs     []uint32                    `protobuf:"varint,4,rep,packed,name=propIntfs,proto3" json:"propIntfs,omitempty"`
}

func (x *ExecutionRequest) Reset() {
	*x = ExecutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rac_v1_rac_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionRequest) ProtoMessage() {}

func (x *ExecutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rac_v1_rac_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionRequest.ProtoReflect.Descriptor instead.
func (*ExecutionRequest) Descriptor() ([]byte, []int) {
	return file_proto_rac_v1_rac_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutionRequest) GetBeacons() []*control_plane.IRECBeacon {
	if x != nil {
		return x.Beacons
	}
	return nil
}

func (x *ExecutionRequest) GetAlgorithmHash() []byte {
	if x != nil {
		return x.AlgorithmHash
	}
	return nil
}

func (x *ExecutionRequest) GetIntfs() []*InterfaceInfo {
	if x != nil {
		return x.Intfs
	}
	return nil
}

func (x *ExecutionRequest) GetPropIntfs() []uint32 {
	if x != nil {
		return x.PropIntfs
	}
	return nil
}

type ExecutionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExecutionResponse) Reset() {
	*x = ExecutionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rac_v1_rac_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionResponse) ProtoMessage() {}

func (x *ExecutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rac_v1_rac_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionResponse.ProtoReflect.Descriptor instead.
func (*ExecutionResponse) Descriptor() ([]byte, []int) {
	return file_proto_rac_v1_rac_proto_rawDescGZIP(), []int{2}
}

type RACRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beacons      []*IRECBeaconSB                  `protobuf:"bytes,1,rep,name=beacons,proto3" json:"beacons,omitempty"`
	EgressIntfs  []uint32                         `protobuf:"varint,2,rep,packed,name=egress_intfs,json=egressIntfs,proto3" json:"egress_intfs,omitempty"`
	BeaconsUnopt []*control_plane.IRECBeaconUnopt `protobuf:"bytes,3,rep,name=beacons_unopt,json=beaconsUnopt,proto3" json:"beacons_unopt,omitempty"`
}

func (x *RACRequest) Reset() {
	*x = RACRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rac_v1_rac_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RACRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RACRequest) ProtoMessage() {}

func (x *RACRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rac_v1_rac_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RACRequest.ProtoReflect.Descriptor instead.
func (*RACRequest) Descriptor() ([]byte, []int) {
	return file_proto_rac_v1_rac_proto_rawDescGZIP(), []int{3}
}

func (x *RACRequest) GetBeacons() []*IRECBeaconSB {
	if x != nil {
		return x.Beacons
	}
	return nil
}

func (x *RACRequest) GetEgressIntfs() []uint32 {
	if x != nil {
		return x.EgressIntfs
	}
	return nil
}

func (x *RACRequest) GetBeaconsUnopt() []*control_plane.IRECBeaconUnopt {
	if x != nil {
		return x.BeaconsUnopt
	}
	return nil
}

type IRECBeaconSB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Segment *IRECPathSegment `protobuf:"bytes,1,opt,name=segment,proto3" json:"segment,omitempty"`
	Id      uint32           `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IRECBeaconSB) Reset() {
	*x = IRECBeaconSB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rac_v1_rac_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IRECBeaconSB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IRECBeaconSB) ProtoMessage() {}

func (x *IRECBeaconSB) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rac_v1_rac_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IRECBeaconSB.ProtoReflect.Descriptor instead.
func (*IRECBeaconSB) Descriptor() ([]byte, []int) {
	return file_proto_rac_v1_rac_proto_rawDescGZIP(), []int{4}
}

func (x *IRECBeaconSB) GetSegment() *IRECPathSegment {
	if x != nil {
		return x.Segment
	}
	return nil
}

func (x *IRECBeaconSB) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RACResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Selected []*RACResponseInner `protobuf:"bytes,1,rep,name=selected,proto3" json:"selected,omitempty"`
}

func (x *RACResponse) Reset() {
	*x = RACResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rac_v1_rac_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RACResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RACResponse) ProtoMessage() {}

func (x *RACResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rac_v1_rac_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RACResponse.ProtoReflect.Descriptor instead.
func (*RACResponse) Descriptor() ([]byte, []int) {
	return file_proto_rac_v1_rac_proto_rawDescGZIP(), []int{5}
}

func (x *RACResponse) GetSelected() []*RACResponseInner {
	if x != nil {
		return x.Selected
	}
	return nil
}

type RACResponseInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EgressIntfs []uint32 `protobuf:"varint,2,rep,packed,name=egress_intfs,json=egressIntfs,proto3" json:"egress_intfs,omitempty"`
}

func (x *RACResponseInner) Reset() {
	*x = RACResponseInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rac_v1_rac_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RACResponseInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RACResponseInner) ProtoMessage() {}

func (x *RACResponseInner) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rac_v1_rac_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RACResponseInner.ProtoReflect.Descriptor instead.
func (*RACResponseInner) Descriptor() ([]byte, []int) {
	return file_proto_rac_v1_rac_proto_rawDescGZIP(), []int{6}
}

func (x *RACResponseInner) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RACResponseInner) GetEgressIntfs() []uint32 {
	if x != nil {
		return x.EgressIntfs
	}
	return nil
}

type IRECPathSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SegmentInfo []byte         `protobuf:"bytes,1,opt,name=segment_info,json=segmentInfo,proto3" json:"segment_info,omitempty"`
	AsEntries   []*IRECASEntry `protobuf:"bytes,2,rep,name=as_entries,json=asEntries,proto3" json:"as_entries,omitempty"`
}

func (x *IRECPathSegment) Reset() {
	*x = IRECPathSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rac_v1_rac_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IRECPathSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IRECPathSegment) ProtoMessage() {}

func (x *IRECPathSegment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rac_v1_rac_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IRECPathSegment.ProtoReflect.Descriptor instead.
func (*IRECPathSegment) Descriptor() ([]byte, []int) {
	return file_proto_rac_v1_rac_proto_rawDescGZIP(), []int{7}
}

func (x *IRECPathSegment) GetSegmentInfo() []byte {
	if x != nil {
		return x.SegmentInfo
	}
	return nil
}

func (x *IRECPathSegment) GetAsEntries() []*IRECASEntry {
	if x != nil {
		return x.AsEntries
	}
	return nil
}

type IRECASEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signed   *control_plane.ASEntrySignedBody             `protobuf:"bytes,3,opt,name=signed,proto3" json:"signed,omitempty"`
	Unsigned *control_plane.PathSegmentUnsignedExtensions `protobuf:"bytes,2,opt,name=unsigned,proto3" json:"unsigned,omitempty"`
}

func (x *IRECASEntry) Reset() {
	*x = IRECASEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_rac_v1_rac_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IRECASEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IRECASEntry) ProtoMessage() {}

func (x *IRECASEntry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_rac_v1_rac_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IRECASEntry.ProtoReflect.Descriptor instead.
func (*IRECASEntry) Descriptor() ([]byte, []int) {
	return file_proto_rac_v1_rac_proto_rawDescGZIP(), []int{8}
}

func (x *IRECASEntry) GetSigned() *control_plane.ASEntrySignedBody {
	if x != nil {
		return x.Signed
	}
	return nil
}

func (x *IRECASEntry) GetUnsigned() *control_plane.PathSegmentUnsignedExtensions {
	if x != nil {
		return x.Unsigned
	}
	return nil
}

var File_proto_rac_v1_rac_proto protoreflect.FileDescriptor

var file_proto_rac_v1_rac_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x61, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x72, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x72, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x67, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x41, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x41, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x69, 0x6e,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x4c, 0x69, 0x6e,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49,
	0x44, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x54, 0x55, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x4d, 0x54, 0x55, 0x22, 0xc7, 0x01, 0x0a, 0x10, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x07, 0x62, 0x65, 0x61, 0x63,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x52, 0x45, 0x43, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x52, 0x07, 0x62,
	0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x61,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x48, 0x61, 0x73, 0x68, 0x12, 0x31, 0x0a, 0x05,
	0x69, 0x6e, 0x74, 0x66, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x66, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x49, 0x6e, 0x74, 0x66, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x49, 0x6e, 0x74, 0x66, 0x73, 0x22, 0x13, 0x0a,
	0x11, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x0a, 0x52, 0x41, 0x43, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x34, 0x0a, 0x07, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x61, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x52, 0x45, 0x43, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x53, 0x42, 0x52, 0x07,
	0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x69, 0x6e, 0x74, 0x66, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x65,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x74, 0x66, 0x73, 0x12, 0x4d, 0x0a, 0x0d, 0x62, 0x65,
	0x61, 0x63, 0x6f, 0x6e, 0x73, 0x5f, 0x75, 0x6e, 0x6f, 0x70, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x52, 0x45, 0x43, 0x42,
	0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x75, 0x6e, 0x6f, 0x70, 0x74, 0x52, 0x0c, 0x62, 0x65, 0x61,
	0x63, 0x6f, 0x6e, 0x73, 0x55, 0x6e, 0x6f, 0x70, 0x74, 0x22, 0x57, 0x0a, 0x0c, 0x49, 0x52, 0x45,
	0x43, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x53, 0x42, 0x12, 0x37, 0x0a, 0x07, 0x73, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x72, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x52, 0x45, 0x43, 0x50, 0x61,
	0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x49, 0x0a, 0x0b, 0x52, 0x41, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x61, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x41, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x45, 0x0a,
	0x10, 0x52, 0x41, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x74, 0x66,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x74, 0x66, 0x73, 0x22, 0x6e, 0x0a, 0x0f, 0x49, 0x52, 0x45, 0x43, 0x50, 0x61, 0x74, 0x68,
	0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x73,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x0a, 0x61, 0x73,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x61, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x52,
	0x45, 0x43, 0x41, 0x53, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x61, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x0b, 0x49, 0x52, 0x45, 0x43, 0x41, 0x53, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x41, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x53,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x42, 0x6f, 0x64, 0x79, 0x52,
	0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x12, 0x51, 0x0a, 0x08, 0x75, 0x6e, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x6e,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x08, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x32, 0x5a, 0x0a, 0x0a, 0x52, 0x41,
	0x43, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x61, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x61, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x69, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x63, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x72, 0x61, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_rac_v1_rac_proto_rawDescOnce sync.Once
	file_proto_rac_v1_rac_proto_rawDescData = file_proto_rac_v1_rac_proto_rawDesc
)

func file_proto_rac_v1_rac_proto_rawDescGZIP() []byte {
	file_proto_rac_v1_rac_proto_rawDescOnce.Do(func() {
		file_proto_rac_v1_rac_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_rac_v1_rac_proto_rawDescData)
	})
	return file_proto_rac_v1_rac_proto_rawDescData
}

var file_proto_rac_v1_rac_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_rac_v1_rac_proto_goTypes = []interface{}{
	(*InterfaceInfo)(nil),                               // 0: proto.rac.v1.InterfaceInfo
	(*ExecutionRequest)(nil),                            // 1: proto.rac.v1.ExecutionRequest
	(*ExecutionResponse)(nil),                           // 2: proto.rac.v1.ExecutionResponse
	(*RACRequest)(nil),                                  // 3: proto.rac.v1.RACRequest
	(*IRECBeaconSB)(nil),                                // 4: proto.rac.v1.IRECBeaconSB
	(*RACResponse)(nil),                                 // 5: proto.rac.v1.RACResponse
	(*RACResponseInner)(nil),                            // 6: proto.rac.v1.RACResponseInner
	(*IRECPathSegment)(nil),                             // 7: proto.rac.v1.IRECPathSegment
	(*IRECASEntry)(nil),                                 // 8: proto.rac.v1.IRECASEntry
	(*control_plane.IRECBeacon)(nil),                    // 9: proto.control_plane.v1.IRECBeacon
	(*control_plane.IRECBeaconUnopt)(nil),               // 10: proto.control_plane.v1.IRECBeacon_unopt
	(*control_plane.ASEntrySignedBody)(nil),             // 11: proto.control_plane.v1.ASEntrySignedBody
	(*control_plane.PathSegmentUnsignedExtensions)(nil), // 12: proto.control_plane.v1.PathSegmentUnsignedExtensions
}
var file_proto_rac_v1_rac_proto_depIdxs = []int32{
	9,  // 0: proto.rac.v1.ExecutionRequest.beacons:type_name -> proto.control_plane.v1.IRECBeacon
	0,  // 1: proto.rac.v1.ExecutionRequest.intfs:type_name -> proto.rac.v1.InterfaceInfo
	4,  // 2: proto.rac.v1.RACRequest.beacons:type_name -> proto.rac.v1.IRECBeaconSB
	10, // 3: proto.rac.v1.RACRequest.beacons_unopt:type_name -> proto.control_plane.v1.IRECBeacon_unopt
	7,  // 4: proto.rac.v1.IRECBeaconSB.segment:type_name -> proto.rac.v1.IRECPathSegment
	6,  // 5: proto.rac.v1.RACResponse.selected:type_name -> proto.rac.v1.RACResponseInner
	8,  // 6: proto.rac.v1.IRECPathSegment.as_entries:type_name -> proto.rac.v1.IRECASEntry
	11, // 7: proto.rac.v1.IRECASEntry.signed:type_name -> proto.control_plane.v1.ASEntrySignedBody
	12, // 8: proto.rac.v1.IRECASEntry.unsigned:type_name -> proto.control_plane.v1.PathSegmentUnsignedExtensions
	1,  // 9: proto.rac.v1.RACService.Execute:input_type -> proto.rac.v1.ExecutionRequest
	2,  // 10: proto.rac.v1.RACService.Execute:output_type -> proto.rac.v1.ExecutionResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_rac_v1_rac_proto_init() }
func file_proto_rac_v1_rac_proto_init() {
	if File_proto_rac_v1_rac_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_rac_v1_rac_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterfaceInfo); i {
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
		file_proto_rac_v1_rac_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionRequest); i {
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
		file_proto_rac_v1_rac_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionResponse); i {
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
		file_proto_rac_v1_rac_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RACRequest); i {
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
		file_proto_rac_v1_rac_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IRECBeaconSB); i {
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
		file_proto_rac_v1_rac_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RACResponse); i {
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
		file_proto_rac_v1_rac_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RACResponseInner); i {
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
		file_proto_rac_v1_rac_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IRECPathSegment); i {
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
		file_proto_rac_v1_rac_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IRECASEntry); i {
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
			RawDescriptor: file_proto_rac_v1_rac_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_rac_v1_rac_proto_goTypes,
		DependencyIndexes: file_proto_rac_v1_rac_proto_depIdxs,
		MessageInfos:      file_proto_rac_v1_rac_proto_msgTypes,
	}.Build()
	File_proto_rac_v1_rac_proto = out.File
	file_proto_rac_v1_rac_proto_rawDesc = nil
	file_proto_rac_v1_rac_proto_goTypes = nil
	file_proto_rac_v1_rac_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RACServiceClient is the client API for RACService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RACServiceClient interface {
	Execute(ctx context.Context, in *ExecutionRequest, opts ...grpc.CallOption) (*ExecutionResponse, error)
}

type rACServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRACServiceClient(cc grpc.ClientConnInterface) RACServiceClient {
	return &rACServiceClient{cc}
}

func (c *rACServiceClient) Execute(ctx context.Context, in *ExecutionRequest, opts ...grpc.CallOption) (*ExecutionResponse, error) {
	out := new(ExecutionResponse)
	err := c.cc.Invoke(ctx, "/proto.rac.v1.RACService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RACServiceServer is the server API for RACService service.
type RACServiceServer interface {
	Execute(context.Context, *ExecutionRequest) (*ExecutionResponse, error)
}

// UnimplementedRACServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRACServiceServer struct {
}

func (*UnimplementedRACServiceServer) Execute(context.Context, *ExecutionRequest) (*ExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}

func RegisterRACServiceServer(s *grpc.Server, srv RACServiceServer) {
	s.RegisterService(&_RACService_serviceDesc, srv)
}

func _RACService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RACServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.rac.v1.RACService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RACServiceServer).Execute(ctx, req.(*ExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RACService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.rac.v1.RACService",
	HandlerType: (*RACServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _RACService_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/rac/v1/rac.proto",
}
