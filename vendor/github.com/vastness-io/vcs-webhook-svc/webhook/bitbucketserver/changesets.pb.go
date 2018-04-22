// Code generated by protoc-gen-go. DO NOT EDIT.
// source: changesets.proto

/*
Package bitbucketserver is a generated protocol buffer package.

It is generated from these files:
	changesets.proto
	commit.proto
	commit_author.proto
	link.proto
	parent.proto
	post_webhook.proto
	project.proto
	ref_change.proto
	repository.proto
	value.proto
	webhook.proto

It has these top-level messages:
	ChangeSets
	ToCommit
	FromCommit
	CommitAuthor
	Link
	Parent
	PostWebhook
	Project
	RefChange
	Repository
	Value
*/
package bitbucketserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ChangeSets struct {
	Filter     *google_protobuf.Struct `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	IsLastPage bool                    `protobuf:"varint,2,opt,name=isLastPage" json:"isLastPage,omitempty"`
	Limit      int64                   `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Size       int64                   `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
	Start      int64                   `protobuf:"varint,5,opt,name=start" json:"start,omitempty"`
	Values     []*Value                `protobuf:"bytes,6,rep,name=values" json:"values,omitempty"`
}

func (m *ChangeSets) Reset()                    { *m = ChangeSets{} }
func (m *ChangeSets) String() string            { return proto.CompactTextString(m) }
func (*ChangeSets) ProtoMessage()               {}
func (*ChangeSets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ChangeSets) GetFilter() *google_protobuf.Struct {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ChangeSets) GetIsLastPage() bool {
	if m != nil {
		return m.IsLastPage
	}
	return false
}

func (m *ChangeSets) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ChangeSets) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ChangeSets) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ChangeSets) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeSets)(nil), "bitbucketserver.ChangeSets")
}

func init() { proto.RegisterFile("changesets.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x89, 0xdb, 0x2e, 0x32, 0x7b, 0x50, 0x82, 0x68, 0x28, 0x22, 0x8b, 0xa7, 0x3d, 0x65,
	0xa1, 0x3e, 0x82, 0x57, 0x0f, 0xb2, 0x05, 0xef, 0xd9, 0x65, 0x1a, 0x83, 0xd1, 0x48, 0x66, 0xd2,
	0x83, 0xef, 0xe8, 0x3b, 0x49, 0x27, 0x15, 0xc4, 0x5b, 0xfe, 0xf9, 0x3f, 0x3e, 0xfe, 0xc0, 0xe5,
	0xf2, 0xea, 0x3e, 0x3c, 0x12, 0x32, 0xd9, 0xcf, 0x9c, 0x38, 0xe9, 0x8b, 0x39, 0xf0, 0x5c, 0x96,
	0x37, 0x64, 0xc2, 0x7c, 0xc0, 0xbc, 0xb9, 0xf5, 0x29, 0xf9, 0x88, 0xa3, 0xd4, 0x73, 0xd9, 0x8f,
	0xc4, 0xb9, 0x2c, 0x5c, 0xf1, 0x4d, 0x77, 0x70, 0xb1, 0x60, 0x0d, 0xf7, 0xdf, 0x0a, 0xe0, 0x51,
	0x84, 0x3b, 0x64, 0xd2, 0x23, 0xb4, 0xfb, 0x10, 0x19, 0xb3, 0x51, 0xbd, 0x1a, 0xba, 0xed, 0x8d,
	0xad, 0x2a, 0xfb, 0xab, 0xb2, 0x3b, 0x51, 0x4d, 0x27, 0x4c, 0xdf, 0x01, 0x04, 0x7a, 0x72, 0xc4,
	0xcf, 0xce, 0xa3, 0x39, 0xeb, 0xd5, 0x70, 0x3e, 0xfd, 0xb9, 0xe8, 0x2b, 0x58, 0xc7, 0xf0, 0x1e,
	0xd8, 0x34, 0xbd, 0x1a, 0x9a, 0xa9, 0x06, 0xad, 0x61, 0x45, 0xe1, 0x0b, 0xcd, 0x4a, 0x8e, 0xf2,
	0x3e, 0x92, 0xc4, 0x2e, 0xb3, 0x59, 0x57, 0x52, 0x82, 0xb6, 0xd0, 0xca, 0x5c, 0x32, 0x6d, 0xdf,
	0x0c, 0xdd, 0xf6, 0xda, 0xfe, 0xfb, 0xac, 0x7d, 0x39, 0xd6, 0xd3, 0x89, 0x9a, 0x5b, 0x19, 0xfa,
	0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0x59, 0x5d, 0xa2, 0x7c, 0x26, 0x01, 0x00, 0x00,
}