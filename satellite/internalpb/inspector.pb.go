// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: inspector.proto

package internalpb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"

	pb "storj.io/common/pb"
	drpc "storj.io/drpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CountNodesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountNodesRequest) Reset()         { *m = CountNodesRequest{} }
func (m *CountNodesRequest) String() string { return proto.CompactTextString(m) }
func (*CountNodesRequest) ProtoMessage()    {}
func (*CountNodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07d9034b2dd9d26, []int{0}
}
func (m *CountNodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountNodesRequest.Unmarshal(m, b)
}
func (m *CountNodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountNodesRequest.Marshal(b, m, deterministic)
}
func (m *CountNodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountNodesRequest.Merge(m, src)
}
func (m *CountNodesRequest) XXX_Size() int {
	return xxx_messageInfo_CountNodesRequest.Size(m)
}
func (m *CountNodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CountNodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CountNodesRequest proto.InternalMessageInfo

type CountNodesResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountNodesResponse) Reset()         { *m = CountNodesResponse{} }
func (m *CountNodesResponse) String() string { return proto.CompactTextString(m) }
func (*CountNodesResponse) ProtoMessage()    {}
func (*CountNodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07d9034b2dd9d26, []int{1}
}
func (m *CountNodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountNodesResponse.Unmarshal(m, b)
}
func (m *CountNodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountNodesResponse.Marshal(b, m, deterministic)
}
func (m *CountNodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountNodesResponse.Merge(m, src)
}
func (m *CountNodesResponse) XXX_Size() int {
	return xxx_messageInfo_CountNodesResponse.Size(m)
}
func (m *CountNodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CountNodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CountNodesResponse proto.InternalMessageInfo

func (m *CountNodesResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type DumpNodesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpNodesRequest) Reset()         { *m = DumpNodesRequest{} }
func (m *DumpNodesRequest) String() string { return proto.CompactTextString(m) }
func (*DumpNodesRequest) ProtoMessage()    {}
func (*DumpNodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07d9034b2dd9d26, []int{2}
}
func (m *DumpNodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpNodesRequest.Unmarshal(m, b)
}
func (m *DumpNodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpNodesRequest.Marshal(b, m, deterministic)
}
func (m *DumpNodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpNodesRequest.Merge(m, src)
}
func (m *DumpNodesRequest) XXX_Size() int {
	return xxx_messageInfo_DumpNodesRequest.Size(m)
}
func (m *DumpNodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpNodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DumpNodesRequest proto.InternalMessageInfo

type DumpNodesResponse struct {
	Nodes                []*pb.Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DumpNodesResponse) Reset()         { *m = DumpNodesResponse{} }
func (m *DumpNodesResponse) String() string { return proto.CompactTextString(m) }
func (*DumpNodesResponse) ProtoMessage()    {}
func (*DumpNodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07d9034b2dd9d26, []int{3}
}
func (m *DumpNodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpNodesResponse.Unmarshal(m, b)
}
func (m *DumpNodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpNodesResponse.Marshal(b, m, deterministic)
}
func (m *DumpNodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpNodesResponse.Merge(m, src)
}
func (m *DumpNodesResponse) XXX_Size() int {
	return xxx_messageInfo_DumpNodesResponse.Size(m)
}
func (m *DumpNodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpNodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DumpNodesResponse proto.InternalMessageInfo

func (m *DumpNodesResponse) GetNodes() []*pb.Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type ListIrreparableSegmentsRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	LastSeenSegmentPath  []byte   `protobuf:"bytes,2,opt,name=last_seen_segment_path,json=lastSeenSegmentPath,proto3" json:"last_seen_segment_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListIrreparableSegmentsRequest) Reset()         { *m = ListIrreparableSegmentsRequest{} }
func (m *ListIrreparableSegmentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListIrreparableSegmentsRequest) ProtoMessage()    {}
func (*ListIrreparableSegmentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07d9034b2dd9d26, []int{4}
}
func (m *ListIrreparableSegmentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIrreparableSegmentsRequest.Unmarshal(m, b)
}
func (m *ListIrreparableSegmentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIrreparableSegmentsRequest.Marshal(b, m, deterministic)
}
func (m *ListIrreparableSegmentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIrreparableSegmentsRequest.Merge(m, src)
}
func (m *ListIrreparableSegmentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListIrreparableSegmentsRequest.Size(m)
}
func (m *ListIrreparableSegmentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIrreparableSegmentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListIrreparableSegmentsRequest proto.InternalMessageInfo

func (m *ListIrreparableSegmentsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListIrreparableSegmentsRequest) GetLastSeenSegmentPath() []byte {
	if m != nil {
		return m.LastSeenSegmentPath
	}
	return nil
}

type ListIrreparableSegmentsResponse struct {
	Segments             []*IrreparableSegment `protobuf:"bytes,1,rep,name=segments,proto3" json:"segments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListIrreparableSegmentsResponse) Reset()         { *m = ListIrreparableSegmentsResponse{} }
func (m *ListIrreparableSegmentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListIrreparableSegmentsResponse) ProtoMessage()    {}
func (*ListIrreparableSegmentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07d9034b2dd9d26, []int{5}
}
func (m *ListIrreparableSegmentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIrreparableSegmentsResponse.Unmarshal(m, b)
}
func (m *ListIrreparableSegmentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIrreparableSegmentsResponse.Marshal(b, m, deterministic)
}
func (m *ListIrreparableSegmentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIrreparableSegmentsResponse.Merge(m, src)
}
func (m *ListIrreparableSegmentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListIrreparableSegmentsResponse.Size(m)
}
func (m *ListIrreparableSegmentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIrreparableSegmentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListIrreparableSegmentsResponse proto.InternalMessageInfo

func (m *ListIrreparableSegmentsResponse) GetSegments() []*IrreparableSegment {
	if m != nil {
		return m.Segments
	}
	return nil
}

type IrreparableSegment struct {
	Path                 []byte      `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	SegmentDetail        *pb.Pointer `protobuf:"bytes,2,opt,name=segment_detail,json=segmentDetail,proto3" json:"segment_detail,omitempty"`
	LostPieces           int32       `protobuf:"varint,3,opt,name=lost_pieces,json=lostPieces,proto3" json:"lost_pieces,omitempty"`
	LastRepairAttempt    int64       `protobuf:"varint,4,opt,name=last_repair_attempt,json=lastRepairAttempt,proto3" json:"last_repair_attempt,omitempty"`
	RepairAttemptCount   int64       `protobuf:"varint,5,opt,name=repair_attempt_count,json=repairAttemptCount,proto3" json:"repair_attempt_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *IrreparableSegment) Reset()         { *m = IrreparableSegment{} }
func (m *IrreparableSegment) String() string { return proto.CompactTextString(m) }
func (*IrreparableSegment) ProtoMessage()    {}
func (*IrreparableSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07d9034b2dd9d26, []int{6}
}
func (m *IrreparableSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IrreparableSegment.Unmarshal(m, b)
}
func (m *IrreparableSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IrreparableSegment.Marshal(b, m, deterministic)
}
func (m *IrreparableSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IrreparableSegment.Merge(m, src)
}
func (m *IrreparableSegment) XXX_Size() int {
	return xxx_messageInfo_IrreparableSegment.Size(m)
}
func (m *IrreparableSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_IrreparableSegment.DiscardUnknown(m)
}

var xxx_messageInfo_IrreparableSegment proto.InternalMessageInfo

func (m *IrreparableSegment) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *IrreparableSegment) GetSegmentDetail() *pb.Pointer {
	if m != nil {
		return m.SegmentDetail
	}
	return nil
}

func (m *IrreparableSegment) GetLostPieces() int32 {
	if m != nil {
		return m.LostPieces
	}
	return 0
}

func (m *IrreparableSegment) GetLastRepairAttempt() int64 {
	if m != nil {
		return m.LastRepairAttempt
	}
	return 0
}

func (m *IrreparableSegment) GetRepairAttemptCount() int64 {
	if m != nil {
		return m.RepairAttemptCount
	}
	return 0
}

type ObjectHealthRequest struct {
	EncryptedPath        []byte   `protobuf:"bytes,1,opt,name=encrypted_path,json=encryptedPath,proto3" json:"encrypted_path,omitempty"`
	Bucket               []byte   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	ProjectId            []byte   `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	StartAfterSegment    int64    `protobuf:"varint,4,opt,name=start_after_segment,json=startAfterSegment,proto3" json:"start_after_segment,omitempty"`
	EndBeforeSegment     int64    `protobuf:"varint,5,opt,name=end_before_segment,json=endBeforeSegment,proto3" json:"end_before_segment,omitempty"`
	Limit                int32    `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectHealthRequest) Reset()         { *m = ObjectHealthRequest{} }
func (m *ObjectHealthRequest) String() string { return proto.CompactTextString(m) }
func (*ObjectHealthRequest) ProtoMessage()    {}
func (*ObjectHealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07d9034b2dd9d26, []int{7}
}
func (m *ObjectHealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectHealthRequest.Unmarshal(m, b)
}
func (m *ObjectHealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectHealthRequest.Marshal(b, m, deterministic)
}
func (m *ObjectHealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectHealthRequest.Merge(m, src)
}
func (m *ObjectHealthRequest) XXX_Size() int {
	return xxx_messageInfo_ObjectHealthRequest.Size(m)
}
func (m *ObjectHealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectHealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectHealthRequest proto.InternalMessageInfo

func (m *ObjectHealthRequest) GetEncryptedPath() []byte {
	if m != nil {
		return m.EncryptedPath
	}
	return nil
}

func (m *ObjectHealthRequest) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *ObjectHealthRequest) GetProjectId() []byte {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *ObjectHealthRequest) GetStartAfterSegment() int64 {
	if m != nil {
		return m.StartAfterSegment
	}
	return 0
}

func (m *ObjectHealthRequest) GetEndBeforeSegment() int64 {
	if m != nil {
		return m.EndBeforeSegment
	}
	return 0
}

func (m *ObjectHealthRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ObjectHealthResponse struct {
	Segments             []*SegmentHealth     `protobuf:"bytes,1,rep,name=segments,proto3" json:"segments,omitempty"`
	Redundancy           *pb.RedundancyScheme `protobuf:"bytes,2,opt,name=redundancy,proto3" json:"redundancy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ObjectHealthResponse) Reset()         { *m = ObjectHealthResponse{} }
func (m *ObjectHealthResponse) String() string { return proto.CompactTextString(m) }
func (*ObjectHealthResponse) ProtoMessage()    {}
func (*ObjectHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07d9034b2dd9d26, []int{8}
}
func (m *ObjectHealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectHealthResponse.Unmarshal(m, b)
}
func (m *ObjectHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectHealthResponse.Marshal(b, m, deterministic)
}
func (m *ObjectHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectHealthResponse.Merge(m, src)
}
func (m *ObjectHealthResponse) XXX_Size() int {
	return xxx_messageInfo_ObjectHealthResponse.Size(m)
}
func (m *ObjectHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectHealthResponse proto.InternalMessageInfo

func (m *ObjectHealthResponse) GetSegments() []*SegmentHealth {
	if m != nil {
		return m.Segments
	}
	return nil
}

func (m *ObjectHealthResponse) GetRedundancy() *pb.RedundancyScheme {
	if m != nil {
		return m.Redundancy
	}
	return nil
}

type SegmentHealthRequest struct {
	Bucket               []byte   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	EncryptedPath        []byte   `protobuf:"bytes,2,opt,name=encrypted_path,json=encryptedPath,proto3" json:"encrypted_path,omitempty"`
	SegmentIndex         int64    `protobuf:"varint,3,opt,name=segment_index,json=segmentIndex,proto3" json:"segment_index,omitempty"`
	ProjectId            []byte   `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentHealthRequest) Reset()         { *m = SegmentHealthRequest{} }
func (m *SegmentHealthRequest) String() string { return proto.CompactTextString(m) }
func (*SegmentHealthRequest) ProtoMessage()    {}
func (*SegmentHealthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07d9034b2dd9d26, []int{9}
}
func (m *SegmentHealthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentHealthRequest.Unmarshal(m, b)
}
func (m *SegmentHealthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentHealthRequest.Marshal(b, m, deterministic)
}
func (m *SegmentHealthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentHealthRequest.Merge(m, src)
}
func (m *SegmentHealthRequest) XXX_Size() int {
	return xxx_messageInfo_SegmentHealthRequest.Size(m)
}
func (m *SegmentHealthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentHealthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentHealthRequest proto.InternalMessageInfo

func (m *SegmentHealthRequest) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *SegmentHealthRequest) GetEncryptedPath() []byte {
	if m != nil {
		return m.EncryptedPath
	}
	return nil
}

func (m *SegmentHealthRequest) GetSegmentIndex() int64 {
	if m != nil {
		return m.SegmentIndex
	}
	return 0
}

func (m *SegmentHealthRequest) GetProjectId() []byte {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

type SegmentHealthResponse struct {
	Health               *SegmentHealth       `protobuf:"bytes,1,opt,name=health,proto3" json:"health,omitempty"`
	Redundancy           *pb.RedundancyScheme `protobuf:"bytes,2,opt,name=redundancy,proto3" json:"redundancy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SegmentHealthResponse) Reset()         { *m = SegmentHealthResponse{} }
func (m *SegmentHealthResponse) String() string { return proto.CompactTextString(m) }
func (*SegmentHealthResponse) ProtoMessage()    {}
func (*SegmentHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07d9034b2dd9d26, []int{10}
}
func (m *SegmentHealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentHealthResponse.Unmarshal(m, b)
}
func (m *SegmentHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentHealthResponse.Marshal(b, m, deterministic)
}
func (m *SegmentHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentHealthResponse.Merge(m, src)
}
func (m *SegmentHealthResponse) XXX_Size() int {
	return xxx_messageInfo_SegmentHealthResponse.Size(m)
}
func (m *SegmentHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentHealthResponse proto.InternalMessageInfo

func (m *SegmentHealthResponse) GetHealth() *SegmentHealth {
	if m != nil {
		return m.Health
	}
	return nil
}

func (m *SegmentHealthResponse) GetRedundancy() *pb.RedundancyScheme {
	if m != nil {
		return m.Redundancy
	}
	return nil
}

type SegmentHealth struct {
	HealthyIds           []NodeID `protobuf:"bytes,1,rep,name=healthy_ids,json=healthyIds,proto3,customtype=NodeID" json:"healthy_ids,omitempty"`
	UnhealthyIds         []NodeID `protobuf:"bytes,2,rep,name=unhealthy_ids,json=unhealthyIds,proto3,customtype=NodeID" json:"unhealthy_ids,omitempty"`
	OfflineIds           []NodeID `protobuf:"bytes,3,rep,name=offline_ids,json=offlineIds,proto3,customtype=NodeID" json:"offline_ids,omitempty"`
	Segment              []byte   `protobuf:"bytes,4,opt,name=segment,proto3" json:"segment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SegmentHealth) Reset()         { *m = SegmentHealth{} }
func (m *SegmentHealth) String() string { return proto.CompactTextString(m) }
func (*SegmentHealth) ProtoMessage()    {}
func (*SegmentHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_a07d9034b2dd9d26, []int{11}
}
func (m *SegmentHealth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentHealth.Unmarshal(m, b)
}
func (m *SegmentHealth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentHealth.Marshal(b, m, deterministic)
}
func (m *SegmentHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentHealth.Merge(m, src)
}
func (m *SegmentHealth) XXX_Size() int {
	return xxx_messageInfo_SegmentHealth.Size(m)
}
func (m *SegmentHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentHealth.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentHealth proto.InternalMessageInfo

func (m *SegmentHealth) GetSegment() []byte {
	if m != nil {
		return m.Segment
	}
	return nil
}

func init() {
	proto.RegisterType((*CountNodesRequest)(nil), "inspector.CountNodesRequest")
	proto.RegisterType((*CountNodesResponse)(nil), "inspector.CountNodesResponse")
	proto.RegisterType((*DumpNodesRequest)(nil), "inspector.DumpNodesRequest")
	proto.RegisterType((*DumpNodesResponse)(nil), "inspector.DumpNodesResponse")
	proto.RegisterType((*ListIrreparableSegmentsRequest)(nil), "inspector.ListIrreparableSegmentsRequest")
	proto.RegisterType((*ListIrreparableSegmentsResponse)(nil), "inspector.ListIrreparableSegmentsResponse")
	proto.RegisterType((*IrreparableSegment)(nil), "inspector.IrreparableSegment")
	proto.RegisterType((*ObjectHealthRequest)(nil), "inspector.ObjectHealthRequest")
	proto.RegisterType((*ObjectHealthResponse)(nil), "inspector.ObjectHealthResponse")
	proto.RegisterType((*SegmentHealthRequest)(nil), "inspector.SegmentHealthRequest")
	proto.RegisterType((*SegmentHealthResponse)(nil), "inspector.SegmentHealthResponse")
	proto.RegisterType((*SegmentHealth)(nil), "inspector.SegmentHealth")
}

func init() { proto.RegisterFile("inspector.proto", fileDescriptor_a07d9034b2dd9d26) }

var fileDescriptor_a07d9034b2dd9d26 = []byte{
	// 819 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0x66, 0xe2, 0xd8, 0x90, 0xb2, 0xbd, 0x49, 0x3a, 0x66, 0xb1, 0xbc, 0x9b, 0xb5, 0x35, 0xab,
	0x95, 0xc2, 0x2e, 0xb2, 0x57, 0x0e, 0x1c, 0x22, 0x4e, 0x09, 0x11, 0xc2, 0x12, 0x22, 0x61, 0xc2,
	0x09, 0x21, 0x8d, 0xc6, 0x33, 0xe5, 0x78, 0x92, 0x71, 0xf7, 0xd0, 0xdd, 0x46, 0xf8, 0x05, 0x10,
	0x3c, 0x02, 0x0f, 0xc0, 0x1b, 0x70, 0xe2, 0x51, 0x38, 0x70, 0x84, 0xd7, 0x40, 0xfd, 0x33, 0xe3,
	0xb6, 0x1d, 0x47, 0x48, 0xdc, 0xba, 0xab, 0xbe, 0xfe, 0xe6, 0xab, 0xfa, 0xca, 0x65, 0xd8, 0x4f,
	0xa9, 0xc8, 0x31, 0x96, 0x8c, 0xf7, 0x73, 0xce, 0x24, 0x23, 0x7b, 0x65, 0xa0, 0x03, 0xb7, 0xec,
	0x96, 0x99, 0x70, 0x07, 0x28, 0x4b, 0xd0, 0x9e, 0xf7, 0x73, 0x96, 0x52, 0x89, 0x3c, 0x19, 0x9b,
	0x80, 0x7f, 0x04, 0x87, 0x9f, 0xb1, 0x39, 0x95, 0x5f, 0xb1, 0x04, 0x45, 0x80, 0xdf, 0xcf, 0x51,
	0x48, 0xff, 0x35, 0x10, 0x37, 0x28, 0x72, 0x46, 0x05, 0x92, 0x16, 0x54, 0x63, 0x15, 0x6d, 0x7b,
	0x3d, 0xef, 0xa4, 0x12, 0x98, 0x8b, 0x4f, 0xe0, 0xe0, 0x72, 0x3e, 0xcb, 0x57, 0xde, 0x7f, 0x02,
	0x87, 0x4e, 0xcc, 0x3e, 0xef, 0x41, 0x55, 0x09, 0x11, 0x6d, 0xaf, 0x57, 0x39, 0xa9, 0x0f, 0xa1,
	0xaf, 0x65, 0x29, 0x4c, 0x60, 0x12, 0xfe, 0x3d, 0xbc, 0xf8, 0x32, 0x15, 0x72, 0xc4, 0x39, 0xe6,
	0x11, 0x8f, 0xc6, 0x19, 0xde, 0xe0, 0xed, 0x0c, 0xa9, 0x2c, 0x88, 0x95, 0x84, 0x2c, 0x9d, 0xa5,
	0x46, 0x42, 0x35, 0x30, 0x17, 0x72, 0x0a, 0x4f, 0xb3, 0x48, 0xc8, 0x50, 0x20, 0xd2, 0x50, 0x98,
	0x27, 0x61, 0x1e, 0xc9, 0x69, 0x7b, 0xa7, 0xe7, 0x9d, 0x34, 0x82, 0x23, 0x95, 0xbd, 0x41, 0xa4,
	0x96, 0xee, 0x3a, 0x92, 0x53, 0xff, 0x3b, 0xe8, 0x6e, 0xfd, 0x98, 0x55, 0x7c, 0x06, 0xef, 0x59,
	0xb6, 0x42, 0xf4, 0x71, 0x7f, 0xd9, 0xf3, 0xcd, 0x97, 0x41, 0x09, 0xf7, 0xff, 0xf6, 0x80, 0x6c,
	0x02, 0x08, 0x81, 0x5d, 0xad, 0xcb, 0xd3, 0xba, 0xf4, 0x99, 0x9c, 0xc1, 0x93, 0x42, 0x73, 0x82,
	0x32, 0x4a, 0x33, 0xad, 0xba, 0x3e, 0x24, 0xfd, 0xa5, 0x57, 0xd7, 0xe6, 0x14, 0x34, 0x2d, 0xf2,
	0x52, 0x03, 0x49, 0x17, 0xea, 0x19, 0x13, 0x32, 0xcc, 0x53, 0x8c, 0x51, 0xb4, 0x2b, 0xba, 0x29,
	0xa0, 0x42, 0xd7, 0x3a, 0x42, 0xfa, 0xa0, 0x6b, 0x0f, 0x95, 0x90, 0x94, 0x87, 0x91, 0x94, 0x38,
	0xcb, 0x65, 0x7b, 0x57, 0x1b, 0x78, 0xa8, 0x52, 0x81, 0xce, 0x9c, 0x9b, 0x04, 0x79, 0x0b, 0xad,
	0x55, 0x68, 0x68, 0x1c, 0xaf, 0xea, 0x07, 0x84, 0xbb, 0x60, 0x3d, 0x21, 0xfe, 0x3f, 0x1e, 0x1c,
	0x5d, 0x8d, 0xef, 0x30, 0x96, 0x5f, 0x60, 0x94, 0xc9, 0x69, 0xe1, 0xd4, 0x2b, 0x78, 0x82, 0x34,
	0xe6, 0x8b, 0x5c, 0x62, 0x12, 0x3a, 0x35, 0x37, 0xcb, 0xa8, 0x72, 0x81, 0x3c, 0x85, 0xda, 0x78,
	0x1e, 0xdf, 0xa3, 0xb4, 0x56, 0xd9, 0x1b, 0x39, 0x06, 0xc8, 0x39, 0x53, 0xb4, 0x61, 0x9a, 0xe8,
	0xc2, 0x1a, 0xc1, 0x9e, 0x8d, 0x8c, 0x12, 0x55, 0x97, 0x90, 0x11, 0x97, 0x61, 0x34, 0x91, 0xc8,
	0x0b, 0xcf, 0x8b, 0xba, 0x74, 0xea, 0x5c, 0x65, 0x8a, 0xbe, 0x7f, 0x04, 0x04, 0x69, 0x12, 0x8e,
	0x71, 0xc2, 0x38, 0x96, 0x70, 0x53, 0xd5, 0x01, 0xd2, 0xe4, 0x42, 0x27, 0x0a, 0x74, 0x39, 0x65,
	0x35, 0x67, 0xca, 0xfc, 0x5f, 0x3c, 0x68, 0xad, 0x56, 0x6a, 0xc7, 0xe4, 0xe3, 0x8d, 0x31, 0x69,
	0x3b, 0x63, 0x62, 0x49, 0xed, 0x9b, 0x12, 0x49, 0x3e, 0x05, 0xe0, 0x98, 0xcc, 0x69, 0x12, 0xd1,
	0x78, 0x61, 0x2d, 0x7f, 0xe6, 0x58, 0x1e, 0x94, 0xc9, 0x9b, 0x78, 0x8a, 0x33, 0x0c, 0x1c, 0xb8,
	0xff, 0xab, 0x07, 0xad, 0x55, 0x62, 0xdb, 0xf6, 0x65, 0x3f, 0xbd, 0x95, 0x7e, 0x6e, 0xda, 0xb1,
	0xf3, 0x90, 0x1d, 0x2f, 0xa1, 0x98, 0xb0, 0x30, 0xa5, 0x09, 0xfe, 0xa8, 0x3b, 0x5f, 0x09, 0x1a,
	0x36, 0x38, 0x52, 0xb1, 0x35, 0x6f, 0x76, 0xd7, 0xbc, 0xf1, 0x7f, 0xf2, 0xe0, 0xfd, 0x35, 0x6d,
	0xb6, 0x51, 0x6f, 0xa1, 0x36, 0xd5, 0x11, 0x2d, 0xee, 0xb1, 0x36, 0x59, 0xdc, 0xff, 0x6b, 0xd2,
	0xef, 0x1e, 0x34, 0x57, 0x68, 0xc9, 0x1b, 0xa8, 0x1b, 0xe2, 0x45, 0x98, 0x26, 0xc6, 0xac, 0xc6,
	0x05, 0xfc, 0xf9, 0x57, 0xb7, 0xa6, 0xd6, 0xd0, 0xe8, 0x32, 0x00, 0x9b, 0x1e, 0x25, 0x82, 0x0c,
	0xa0, 0x39, 0xa7, 0x2e, 0x7c, 0x67, 0x03, 0xde, 0x28, 0x01, 0xea, 0xc1, 0x1b, 0xa8, 0xb3, 0xc9,
	0x24, 0x4b, 0x29, 0x6a, 0x78, 0x65, 0x93, 0xdd, 0xa6, 0x15, 0xb8, 0x0d, 0xef, 0xba, 0x53, 0xdb,
	0x08, 0x8a, 0xeb, 0xf0, 0x37, 0x0f, 0x0e, 0xae, 0x7e, 0x40, 0x9e, 0x45, 0x8b, 0x51, 0xd1, 0x1e,
	0x32, 0x02, 0x58, 0x6e, 0x64, 0xf2, 0xdc, 0x69, 0xdc, 0xc6, 0xf6, 0xee, 0x1c, 0x6f, 0xc9, 0x5a,
	0x17, 0x3e, 0x87, 0xbd, 0x72, 0x39, 0x93, 0x67, 0x0e, 0x76, 0x7d, 0x8d, 0x77, 0x9e, 0x3f, 0x9c,
	0x34, 0x3c, 0xc3, 0x9f, 0x3d, 0x68, 0x39, 0x2b, 0x6e, 0xa9, 0x35, 0x87, 0x0f, 0xb6, 0x6c, 0x56,
	0xf2, 0xa1, 0xc3, 0xf8, 0xf8, 0xaa, 0xef, 0xbc, 0xfe, 0x2f, 0x50, 0x2b, 0xe5, 0x0f, 0x0f, 0xf6,
	0x8d, 0xc5, 0x4b, 0x15, 0x5f, 0x43, 0xc3, 0xfd, 0xb5, 0x92, 0x17, 0x0e, 0xdf, 0x03, 0x0b, 0xab,
	0xd3, 0xdd, 0x9a, 0x37, 0x1f, 0xf1, 0xdf, 0x21, 0xdf, 0xac, 0xcf, 0x53, 0x77, 0xeb, 0x00, 0x5b,
	0xd2, 0xde, 0x76, 0x40, 0xc1, 0x7a, 0xf1, 0xea, 0xdb, 0x97, 0x42, 0x32, 0x7e, 0xd7, 0x4f, 0xd9,
	0x40, 0x1f, 0x06, 0x22, 0x92, 0x98, 0x65, 0xa9, 0xc4, 0x81, 0x9e, 0x73, 0x1a, 0x65, 0xf9, 0x78,
	0x5c, 0xd3, 0xff, 0xd7, 0xa7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x65, 0xe3, 0x13, 0xf6,
	0x07, 0x00, 0x00,
}

// --- DRPC BEGIN ---

type DRPCOverlayInspectorClient interface {
	DRPCConn() drpc.Conn

	// CountNodes returns the number of nodes in the cache
	CountNodes(ctx context.Context, in *CountNodesRequest) (*CountNodesResponse, error)
	// DumpNodes returns all the nodes in the cache
	DumpNodes(ctx context.Context, in *DumpNodesRequest) (*DumpNodesResponse, error)
}

type drpcOverlayInspectorClient struct {
	cc drpc.Conn
}

func NewDRPCOverlayInspectorClient(cc drpc.Conn) DRPCOverlayInspectorClient {
	return &drpcOverlayInspectorClient{cc}
}

func (c *drpcOverlayInspectorClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcOverlayInspectorClient) CountNodes(ctx context.Context, in *CountNodesRequest) (*CountNodesResponse, error) {
	out := new(CountNodesResponse)
	err := c.cc.Invoke(ctx, "/inspector.OverlayInspector/CountNodes", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcOverlayInspectorClient) DumpNodes(ctx context.Context, in *DumpNodesRequest) (*DumpNodesResponse, error) {
	out := new(DumpNodesResponse)
	err := c.cc.Invoke(ctx, "/inspector.OverlayInspector/DumpNodes", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCOverlayInspectorServer interface {
	// CountNodes returns the number of nodes in the cache
	CountNodes(context.Context, *CountNodesRequest) (*CountNodesResponse, error)
	// DumpNodes returns all the nodes in the cache
	DumpNodes(context.Context, *DumpNodesRequest) (*DumpNodesResponse, error)
}

type DRPCOverlayInspectorDescription struct{}

func (DRPCOverlayInspectorDescription) NumMethods() int { return 2 }

func (DRPCOverlayInspectorDescription) Method(n int) (string, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/inspector.OverlayInspector/CountNodes",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCOverlayInspectorServer).
					CountNodes(
						ctx,
						in1.(*CountNodesRequest),
					)
			}, DRPCOverlayInspectorServer.CountNodes, true
	case 1:
		return "/inspector.OverlayInspector/DumpNodes",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCOverlayInspectorServer).
					DumpNodes(
						ctx,
						in1.(*DumpNodesRequest),
					)
			}, DRPCOverlayInspectorServer.DumpNodes, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterOverlayInspector(mux drpc.Mux, impl DRPCOverlayInspectorServer) error {
	return mux.Register(impl, DRPCOverlayInspectorDescription{})
}

type DRPCOverlayInspector_CountNodesStream interface {
	drpc.Stream
	SendAndClose(*CountNodesResponse) error
}

type drpcOverlayInspectorCountNodesStream struct {
	drpc.Stream
}

func (x *drpcOverlayInspectorCountNodesStream) SendAndClose(m *CountNodesResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCOverlayInspector_DumpNodesStream interface {
	drpc.Stream
	SendAndClose(*DumpNodesResponse) error
}

type drpcOverlayInspectorDumpNodesStream struct {
	drpc.Stream
}

func (x *drpcOverlayInspectorDumpNodesStream) SendAndClose(m *DumpNodesResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCIrreparableInspectorClient interface {
	DRPCConn() drpc.Conn

	// ListIrreparableSegments returns damaged segments
	ListIrreparableSegments(ctx context.Context, in *ListIrreparableSegmentsRequest) (*ListIrreparableSegmentsResponse, error)
}

type drpcIrreparableInspectorClient struct {
	cc drpc.Conn
}

func NewDRPCIrreparableInspectorClient(cc drpc.Conn) DRPCIrreparableInspectorClient {
	return &drpcIrreparableInspectorClient{cc}
}

func (c *drpcIrreparableInspectorClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcIrreparableInspectorClient) ListIrreparableSegments(ctx context.Context, in *ListIrreparableSegmentsRequest) (*ListIrreparableSegmentsResponse, error) {
	out := new(ListIrreparableSegmentsResponse)
	err := c.cc.Invoke(ctx, "/inspector.IrreparableInspector/ListIrreparableSegments", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCIrreparableInspectorServer interface {
	// ListIrreparableSegments returns damaged segments
	ListIrreparableSegments(context.Context, *ListIrreparableSegmentsRequest) (*ListIrreparableSegmentsResponse, error)
}

type DRPCIrreparableInspectorDescription struct{}

func (DRPCIrreparableInspectorDescription) NumMethods() int { return 1 }

func (DRPCIrreparableInspectorDescription) Method(n int) (string, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/inspector.IrreparableInspector/ListIrreparableSegments",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCIrreparableInspectorServer).
					ListIrreparableSegments(
						ctx,
						in1.(*ListIrreparableSegmentsRequest),
					)
			}, DRPCIrreparableInspectorServer.ListIrreparableSegments, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterIrreparableInspector(mux drpc.Mux, impl DRPCIrreparableInspectorServer) error {
	return mux.Register(impl, DRPCIrreparableInspectorDescription{})
}

type DRPCIrreparableInspector_ListIrreparableSegmentsStream interface {
	drpc.Stream
	SendAndClose(*ListIrreparableSegmentsResponse) error
}

type drpcIrreparableInspectorListIrreparableSegmentsStream struct {
	drpc.Stream
}

func (x *drpcIrreparableInspectorListIrreparableSegmentsStream) SendAndClose(m *ListIrreparableSegmentsResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCHealthInspectorClient interface {
	DRPCConn() drpc.Conn

	// ObjectHealth will return stats about the health of an object
	ObjectHealth(ctx context.Context, in *ObjectHealthRequest) (*ObjectHealthResponse, error)
	// SegmentHealth will return stats about the health of a segment
	SegmentHealth(ctx context.Context, in *SegmentHealthRequest) (*SegmentHealthResponse, error)
}

type drpcHealthInspectorClient struct {
	cc drpc.Conn
}

func NewDRPCHealthInspectorClient(cc drpc.Conn) DRPCHealthInspectorClient {
	return &drpcHealthInspectorClient{cc}
}

func (c *drpcHealthInspectorClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcHealthInspectorClient) ObjectHealth(ctx context.Context, in *ObjectHealthRequest) (*ObjectHealthResponse, error) {
	out := new(ObjectHealthResponse)
	err := c.cc.Invoke(ctx, "/inspector.HealthInspector/ObjectHealth", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcHealthInspectorClient) SegmentHealth(ctx context.Context, in *SegmentHealthRequest) (*SegmentHealthResponse, error) {
	out := new(SegmentHealthResponse)
	err := c.cc.Invoke(ctx, "/inspector.HealthInspector/SegmentHealth", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCHealthInspectorServer interface {
	// ObjectHealth will return stats about the health of an object
	ObjectHealth(context.Context, *ObjectHealthRequest) (*ObjectHealthResponse, error)
	// SegmentHealth will return stats about the health of a segment
	SegmentHealth(context.Context, *SegmentHealthRequest) (*SegmentHealthResponse, error)
}

type DRPCHealthInspectorDescription struct{}

func (DRPCHealthInspectorDescription) NumMethods() int { return 2 }

func (DRPCHealthInspectorDescription) Method(n int) (string, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/inspector.HealthInspector/ObjectHealth",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCHealthInspectorServer).
					ObjectHealth(
						ctx,
						in1.(*ObjectHealthRequest),
					)
			}, DRPCHealthInspectorServer.ObjectHealth, true
	case 1:
		return "/inspector.HealthInspector/SegmentHealth",
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCHealthInspectorServer).
					SegmentHealth(
						ctx,
						in1.(*SegmentHealthRequest),
					)
			}, DRPCHealthInspectorServer.SegmentHealth, true
	default:
		return "", nil, nil, false
	}
}

func DRPCRegisterHealthInspector(mux drpc.Mux, impl DRPCHealthInspectorServer) error {
	return mux.Register(impl, DRPCHealthInspectorDescription{})
}

type DRPCHealthInspector_ObjectHealthStream interface {
	drpc.Stream
	SendAndClose(*ObjectHealthResponse) error
}

type drpcHealthInspectorObjectHealthStream struct {
	drpc.Stream
}

func (x *drpcHealthInspectorObjectHealthStream) SendAndClose(m *ObjectHealthResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCHealthInspector_SegmentHealthStream interface {
	drpc.Stream
	SendAndClose(*SegmentHealthResponse) error
}

type drpcHealthInspectorSegmentHealthStream struct {
	drpc.Stream
}

func (x *drpcHealthInspectorSegmentHealthStream) SendAndClose(m *SegmentHealthResponse) error {
	if err := x.MsgSend(m); err != nil {
		return err
	}
	return x.CloseSend()
}

// --- DRPC END ---
