// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/gerrithooks/gerrithooks.proto
// DO NOT EDIT!

/*
Package gerrithooks is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/gerrithooks/gerrithooks.proto

It has these top-level messages:
	HookRequest
	HookResponse
	Repository
	Watcher
	Build
	CreateRepoRequest
	CreateRepoResponse
	Commit
	ChangeList
	Change
	TriggerBuildRequest
	OpenChangesRequest
*/
package gerrithooks

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// comment: message pingresponse
type HookRequest struct {
	Hostname    string   `protobuf:"bytes,1,opt,name=Hostname" json:"Hostname,omitempty"`
	Environment []string `protobuf:"bytes,2,rep,name=Environment" json:"Environment,omitempty"`
	Arguments   []string `protobuf:"bytes,3,rep,name=Arguments" json:"Arguments,omitempty"`
	HookName    string   `protobuf:"bytes,4,opt,name=HookName" json:"HookName,omitempty"`
}

func (m *HookRequest) Reset()                    { *m = HookRequest{} }
func (m *HookRequest) String() string            { return proto.CompactTextString(m) }
func (*HookRequest) ProtoMessage()               {}
func (*HookRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HookRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *HookRequest) GetEnvironment() []string {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *HookRequest) GetArguments() []string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func (m *HookRequest) GetHookName() string {
	if m != nil {
		return m.HookName
	}
	return ""
}

type HookResponse struct {
	ExitCode       int32    `protobuf:"varint,1,opt,name=ExitCode" json:"ExitCode,omitempty"`
	StdoutMessages []string `protobuf:"bytes,2,rep,name=StdoutMessages" json:"StdoutMessages,omitempty"`
}

func (m *HookResponse) Reset()                    { *m = HookResponse{} }
func (m *HookResponse) String() string            { return proto.CompactTextString(m) }
func (*HookResponse) ProtoMessage()               {}
func (*HookResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HookResponse) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *HookResponse) GetStdoutMessages() []string {
	if m != nil {
		return m.StdoutMessages
	}
	return nil
}

type Repository struct {
	ID     uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Server string `protobuf:"bytes,2,opt,name=Server" json:"Server,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
}

func (m *Repository) Reset()                    { *m = Repository{} }
func (m *Repository) String() string            { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()               {}
func (*Repository) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Repository) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Repository) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *Repository) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Watcher struct {
	ID       uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	ChangeID string `protobuf:"bytes,2,opt,name=ChangeID" json:"ChangeID,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=Email" json:"Email,omitempty"`
}

func (m *Watcher) Reset()                    { *m = Watcher{} }
func (m *Watcher) String() string            { return proto.CompactTextString(m) }
func (*Watcher) ProtoMessage()               {}
func (*Watcher) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Watcher) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Watcher) GetChangeID() string {
	if m != nil {
		return m.ChangeID
	}
	return ""
}

func (m *Watcher) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Build struct {
	ID        uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	CommitID  string `protobuf:"bytes,2,opt,name=CommitID" json:"CommitID,omitempty"`
	Project   string `protobuf:"bytes,3,opt,name=Project" json:"Project,omitempty"`
	URL       string `protobuf:"bytes,4,opt,name=URL" json:"URL,omitempty"`
	Timestamp uint32 `protobuf:"varint,5,opt,name=Timestamp" json:"Timestamp,omitempty"`
}

func (m *Build) Reset()                    { *m = Build{} }
func (m *Build) String() string            { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()               {}
func (*Build) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Build) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Build) GetCommitID() string {
	if m != nil {
		return m.CommitID
	}
	return ""
}

func (m *Build) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Build) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *Build) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type CreateRepoRequest struct {
	Name        string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description" json:"Description,omitempty"`
}

func (m *CreateRepoRequest) Reset()                    { *m = CreateRepoRequest{} }
func (m *CreateRepoRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRepoRequest) ProtoMessage()               {}
func (*CreateRepoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateRepoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRepoRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateRepoResponse struct {
	ID uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *CreateRepoResponse) Reset()                    { *m = CreateRepoResponse{} }
func (m *CreateRepoResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateRepoResponse) ProtoMessage()               {}
func (*CreateRepoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateRepoResponse) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type Commit struct {
	ID             uint64      `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Repository     *Repository `protobuf:"bytes,2,opt,name=Repository" json:"Repository,omitempty"`
	ChangeID       string      `protobuf:"bytes,3,opt,name=ChangeID" json:"ChangeID,omitempty"`
	CommitID       string      `protobuf:"bytes,4,opt,name=CommitID" json:"CommitID,omitempty"`
	CommitterID    string      `protobuf:"bytes,5,opt,name=CommitterID" json:"CommitterID,omitempty"`
	CommitterEmail string      `protobuf:"bytes,6,opt,name=CommitterEmail" json:"CommitterEmail,omitempty"`
	Committed      uint32      `protobuf:"varint,7,opt,name=Committed" json:"Committed,omitempty"`
	FirstReview    uint32      `protobuf:"varint,8,opt,name=FirstReview" json:"FirstReview,omitempty"`
	ReviewerID     string      `protobuf:"bytes,9,opt,name=ReviewerID" json:"ReviewerID,omitempty"`
	ReviewedByDod  bool        `protobuf:"varint,10,opt,name=ReviewedByDod" json:"ReviewedByDod,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Commit) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Commit) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *Commit) GetChangeID() string {
	if m != nil {
		return m.ChangeID
	}
	return ""
}

func (m *Commit) GetCommitID() string {
	if m != nil {
		return m.CommitID
	}
	return ""
}

func (m *Commit) GetCommitterID() string {
	if m != nil {
		return m.CommitterID
	}
	return ""
}

func (m *Commit) GetCommitterEmail() string {
	if m != nil {
		return m.CommitterEmail
	}
	return ""
}

func (m *Commit) GetCommitted() uint32 {
	if m != nil {
		return m.Committed
	}
	return 0
}

func (m *Commit) GetFirstReview() uint32 {
	if m != nil {
		return m.FirstReview
	}
	return 0
}

func (m *Commit) GetReviewerID() string {
	if m != nil {
		return m.ReviewerID
	}
	return ""
}

func (m *Commit) GetReviewedByDod() bool {
	if m != nil {
		return m.ReviewedByDod
	}
	return false
}

type ChangeList struct {
	Changes []*Change `protobuf:"bytes,1,rep,name=Changes" json:"Changes,omitempty"`
}

func (m *ChangeList) Reset()                    { *m = ChangeList{} }
func (m *ChangeList) String() string            { return proto.CompactTextString(m) }
func (*ChangeList) ProtoMessage()               {}
func (*ChangeList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ChangeList) GetChanges() []*Change {
	if m != nil {
		return m.Changes
	}
	return nil
}

type Change struct {
	Commit *Commit `protobuf:"bytes,1,opt,name=Commit" json:"Commit,omitempty"`
}

func (m *Change) Reset()                    { *m = Change{} }
func (m *Change) String() string            { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()               {}
func (*Change) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Change) GetCommit() *Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

type TriggerBuildRequest struct {
	ChangeID string `protobuf:"bytes,1,opt,name=ChangeID" json:"ChangeID,omitempty"`
}

func (m *TriggerBuildRequest) Reset()                    { *m = TriggerBuildRequest{} }
func (m *TriggerBuildRequest) String() string            { return proto.CompactTextString(m) }
func (*TriggerBuildRequest) ProtoMessage()               {}
func (*TriggerBuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TriggerBuildRequest) GetChangeID() string {
	if m != nil {
		return m.ChangeID
	}
	return ""
}

type OpenChangesRequest struct {
}

func (m *OpenChangesRequest) Reset()                    { *m = OpenChangesRequest{} }
func (m *OpenChangesRequest) String() string            { return proto.CompactTextString(m) }
func (*OpenChangesRequest) ProtoMessage()               {}
func (*OpenChangesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*HookRequest)(nil), "gerrithooks.HookRequest")
	proto.RegisterType((*HookResponse)(nil), "gerrithooks.HookResponse")
	proto.RegisterType((*Repository)(nil), "gerrithooks.Repository")
	proto.RegisterType((*Watcher)(nil), "gerrithooks.Watcher")
	proto.RegisterType((*Build)(nil), "gerrithooks.Build")
	proto.RegisterType((*CreateRepoRequest)(nil), "gerrithooks.CreateRepoRequest")
	proto.RegisterType((*CreateRepoResponse)(nil), "gerrithooks.CreateRepoResponse")
	proto.RegisterType((*Commit)(nil), "gerrithooks.Commit")
	proto.RegisterType((*ChangeList)(nil), "gerrithooks.ChangeList")
	proto.RegisterType((*Change)(nil), "gerrithooks.Change")
	proto.RegisterType((*TriggerBuildRequest)(nil), "gerrithooks.TriggerBuildRequest")
	proto.RegisterType((*OpenChangesRequest)(nil), "gerrithooks.OpenChangesRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GerritHooks service

type GerritHooksClient interface {
	//
	// execute a hook. This is meant to be in a format so that
	// it does not require much parsing or logic on the client.
	// it throws an error if execution failed (not if code tests fail for example)
	GerritHook(ctx context.Context, in *HookRequest, opts ...grpc.CallOption) (*HookResponse, error)
	//
	// create a new repository, useful for 'create module from scratch builder'
	CreateRepository(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error)
	// (re)build a certain change
	TriggerBuild(ctx context.Context, in *TriggerBuildRequest, opts ...grpc.CallOption) (*common.Void, error)
	// get list of open changes
	GetOpenChanges(ctx context.Context, in *OpenChangesRequest, opts ...grpc.CallOption) (*ChangeList, error)
}

type gerritHooksClient struct {
	cc *grpc.ClientConn
}

func NewGerritHooksClient(cc *grpc.ClientConn) GerritHooksClient {
	return &gerritHooksClient{cc}
}

func (c *gerritHooksClient) GerritHook(ctx context.Context, in *HookRequest, opts ...grpc.CallOption) (*HookResponse, error) {
	out := new(HookResponse)
	err := grpc.Invoke(ctx, "/gerrithooks.GerritHooks/GerritHook", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerritHooksClient) CreateRepository(ctx context.Context, in *CreateRepoRequest, opts ...grpc.CallOption) (*CreateRepoResponse, error) {
	out := new(CreateRepoResponse)
	err := grpc.Invoke(ctx, "/gerrithooks.GerritHooks/CreateRepository", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerritHooksClient) TriggerBuild(ctx context.Context, in *TriggerBuildRequest, opts ...grpc.CallOption) (*common.Void, error) {
	out := new(common.Void)
	err := grpc.Invoke(ctx, "/gerrithooks.GerritHooks/TriggerBuild", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gerritHooksClient) GetOpenChanges(ctx context.Context, in *OpenChangesRequest, opts ...grpc.CallOption) (*ChangeList, error) {
	out := new(ChangeList)
	err := grpc.Invoke(ctx, "/gerrithooks.GerritHooks/GetOpenChanges", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GerritHooks service

type GerritHooksServer interface {
	//
	// execute a hook. This is meant to be in a format so that
	// it does not require much parsing or logic on the client.
	// it throws an error if execution failed (not if code tests fail for example)
	GerritHook(context.Context, *HookRequest) (*HookResponse, error)
	//
	// create a new repository, useful for 'create module from scratch builder'
	CreateRepository(context.Context, *CreateRepoRequest) (*CreateRepoResponse, error)
	// (re)build a certain change
	TriggerBuild(context.Context, *TriggerBuildRequest) (*common.Void, error)
	// get list of open changes
	GetOpenChanges(context.Context, *OpenChangesRequest) (*ChangeList, error)
}

func RegisterGerritHooksServer(s *grpc.Server, srv GerritHooksServer) {
	s.RegisterService(&_GerritHooks_serviceDesc, srv)
}

func _GerritHooks_GerritHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerritHooksServer).GerritHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerrithooks.GerritHooks/GerritHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerritHooksServer).GerritHook(ctx, req.(*HookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GerritHooks_CreateRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerritHooksServer).CreateRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerrithooks.GerritHooks/CreateRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerritHooksServer).CreateRepository(ctx, req.(*CreateRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GerritHooks_TriggerBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerritHooksServer).TriggerBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerrithooks.GerritHooks/TriggerBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerritHooksServer).TriggerBuild(ctx, req.(*TriggerBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GerritHooks_GetOpenChanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenChangesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerritHooksServer).GetOpenChanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerrithooks.GerritHooks/GetOpenChanges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerritHooksServer).GetOpenChanges(ctx, req.(*OpenChangesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GerritHooks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gerrithooks.GerritHooks",
	HandlerType: (*GerritHooksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GerritHook",
			Handler:    _GerritHooks_GerritHook_Handler,
		},
		{
			MethodName: "CreateRepository",
			Handler:    _GerritHooks_CreateRepository_Handler,
		},
		{
			MethodName: "TriggerBuild",
			Handler:    _GerritHooks_TriggerBuild_Handler,
		},
		{
			MethodName: "GetOpenChanges",
			Handler:    _GerritHooks_GetOpenChanges_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/gerrithooks/gerrithooks.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/gerrithooks/gerrithooks.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 719 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0x5b, 0x4f, 0x1b, 0x3b,
	0x10, 0xd6, 0xe6, 0x9e, 0x09, 0x20, 0x8e, 0x41, 0x87, 0x25, 0xe7, 0x08, 0xf6, 0xac, 0xd0, 0x51,
	0xa4, 0xaa, 0x41, 0xa5, 0x17, 0x55, 0xaa, 0x54, 0x89, 0x10, 0x0a, 0x69, 0xe9, 0x45, 0x86, 0xb6,
	0xcf, 0xdb, 0xac, 0xb5, 0xb8, 0xb0, 0xeb, 0xd4, 0x76, 0xa0, 0x3c, 0xf4, 0x05, 0xa9, 0x7f, 0xa2,
	0x8f, 0xfd, 0x07, 0xfd, 0x73, 0x7d, 0xae, 0x7c, 0xc9, 0xc6, 0x9b, 0xa4, 0x7d, 0xca, 0x7c, 0xdf,
	0x78, 0x3f, 0xcf, 0xcc, 0xe7, 0x0c, 0x3c, 0x4e, 0xd8, 0x65, 0x94, 0x25, 0xdd, 0x21, 0xcb, 0x78,
	0x14, 0x5f, 0x33, 0x16, 0x77, 0x33, 0x22, 0x77, 0xa3, 0x11, 0x15, 0xbb, 0x09, 0xe1, 0x9c, 0xca,
	0x73, 0xc6, 0x2e, 0x0a, 0x71, 0x77, 0xc4, 0x99, 0x64, 0xa8, 0xe5, 0x50, 0xed, 0xee, 0x1f, 0x64,
	0x86, 0x2c, 0x4d, 0x59, 0x66, 0x7f, 0xcc, 0xc7, 0xe1, 0x57, 0x0f, 0x5a, 0xc7, 0x8c, 0x5d, 0x60,
	0xf2, 0x69, 0x4c, 0x84, 0x44, 0x6d, 0x68, 0x1c, 0x33, 0x21, 0xb3, 0x28, 0x25, 0xbe, 0x17, 0x78,
	0x9d, 0x26, 0xce, 0x31, 0x0a, 0xa0, 0x75, 0x98, 0x5d, 0x51, 0xce, 0xb2, 0x94, 0x64, 0xd2, 0x2f,
	0x05, 0xe5, 0x4e, 0x13, 0xbb, 0x14, 0xfa, 0x17, 0x9a, 0xfb, 0x3c, 0x19, 0xab, 0x58, 0xf8, 0x65,
	0x9d, 0x9f, 0x12, 0x46, 0x9b, 0x5d, 0xbc, 0x52, 0xda, 0x95, 0x89, 0xb6, 0xc1, 0x21, 0x86, 0x25,
	0x53, 0x86, 0x18, 0xb1, 0x4c, 0x10, 0x75, 0xf6, 0xf0, 0x33, 0x95, 0x07, 0x2c, 0x36, 0x75, 0x54,
	0x71, 0x8e, 0xd1, 0xff, 0xb0, 0x72, 0x2a, 0x63, 0x36, 0x96, 0x2f, 0x89, 0x10, 0x51, 0x42, 0x84,
	0x2d, 0x65, 0x86, 0x0d, 0x8f, 0x01, 0x30, 0x19, 0x31, 0x41, 0x25, 0xe3, 0x37, 0x68, 0x05, 0x4a,
	0x83, 0xbe, 0xd6, 0xaa, 0xe0, 0xd2, 0xa0, 0x8f, 0xfe, 0x86, 0xda, 0x29, 0xe1, 0x57, 0x84, 0xfb,
	0x25, 0x5d, 0x8b, 0x45, 0x08, 0x41, 0x45, 0x57, 0x58, 0xd6, 0xac, 0x8e, 0xc3, 0x17, 0x50, 0x7f,
	0x1f, 0xc9, 0xe1, 0x39, 0xe1, 0x73, 0x32, 0x6d, 0x68, 0x1c, 0x9c, 0x47, 0x59, 0x42, 0x06, 0x7d,
	0x2b, 0x94, 0x63, 0xb4, 0x0e, 0xd5, 0xc3, 0x34, 0xa2, 0x97, 0x56, 0xcb, 0x80, 0xf0, 0x0b, 0x54,
	0x7b, 0x63, 0x7a, 0x19, 0x2f, 0x94, 0x62, 0x69, 0x4a, 0xa5, 0x23, 0x65, 0x31, 0xf2, 0xa1, 0xfe,
	0x86, 0xb3, 0x8f, 0x64, 0x28, 0xad, 0xd8, 0x04, 0xa2, 0x55, 0x28, 0xbf, 0xc5, 0x27, 0x76, 0xa0,
	0x2a, 0x54, 0x2e, 0x9c, 0xd1, 0x94, 0x08, 0x19, 0xa5, 0x23, 0xbf, 0x1a, 0x78, 0x9d, 0x65, 0x3c,
	0x25, 0xc2, 0x01, 0xfc, 0x75, 0xc0, 0x49, 0x24, 0x89, 0x9a, 0xcd, 0xc4, 0xf6, 0x49, 0xd3, 0xde,
	0xb4, 0x69, 0x65, 0x77, 0x9f, 0x88, 0x21, 0xa7, 0x23, 0x49, 0x59, 0x66, 0x2b, 0x72, 0xa9, 0x70,
	0x07, 0x90, 0x2b, 0x65, 0xad, 0x9b, 0x69, 0x2b, 0xfc, 0x59, 0x82, 0x9a, 0xe9, 0x63, 0xae, 0xe3,
	0x33, 0xd7, 0x21, 0x7d, 0x43, 0x6b, 0x6f, 0xa3, 0xeb, 0x3e, 0xf1, 0x69, 0xba, 0xf7, 0xcf, 0xb7,
	0xdb, 0xcd, 0xda, 0x98, 0x66, 0xf2, 0xd1, 0x83, 0xef, 0xb7, 0x9b, 0xcb, 0x3c, 0xcf, 0x74, 0x69,
	0x8c, 0x5d, 0xa7, 0x5d, 0x4b, 0xca, 0x33, 0x96, 0xb8, 0x33, 0xae, 0xcc, 0xcc, 0x38, 0x80, 0x96,
	0x89, 0x25, 0xe1, 0x83, 0xbe, 0x9e, 0x5c, 0x13, 0xbb, 0x94, 0x7a, 0x79, 0x39, 0x34, 0xce, 0xd6,
	0xf4, 0xa1, 0x19, 0x56, 0x39, 0x30, 0x61, 0x62, 0xbf, 0x6e, 0x1c, 0xc8, 0x09, 0x75, 0xcf, 0x33,
	0xca, 0x85, 0xc4, 0xe4, 0x8a, 0x92, 0x6b, 0xbf, 0xa1, 0xf3, 0x2e, 0x85, 0xb6, 0xd4, 0x5c, 0x54,
	0xa4, 0x0b, 0x69, 0xea, 0x3b, 0x1c, 0x06, 0xed, 0xc0, 0xb2, 0x45, 0x71, 0xef, 0xa6, 0xcf, 0x62,
	0x1f, 0x02, 0xaf, 0xd3, 0xc0, 0x45, 0x32, 0x7c, 0x02, 0x60, 0xfa, 0x3e, 0xa1, 0x42, 0xa2, 0xbb,
	0x50, 0x37, 0x48, 0xf8, 0x5e, 0x50, 0xee, 0xb4, 0xf6, 0xd6, 0x0a, 0x83, 0x36, 0x39, 0x3c, 0x39,
	0x13, 0x3e, 0x84, 0x9a, 0x09, 0xd1, 0x9d, 0x89, 0x7d, 0xda, 0xb8, 0xb9, 0xef, 0x74, 0x0a, 0xdb,
	0x23, 0xe1, 0x3d, 0x58, 0x3b, 0xe3, 0x34, 0x49, 0x08, 0xd7, 0x6f, 0xdc, 0x59, 0x2b, 0xb9, 0x25,
	0x5e, 0xd1, 0x92, 0x70, 0x1d, 0xd0, 0xeb, 0x11, 0xc9, 0xec, 0xc5, 0xf6, 0x8b, 0xbd, 0x1f, 0x25,
	0x68, 0x1d, 0xe9, 0x7b, 0xd4, 0x5e, 0x10, 0x68, 0x1f, 0x60, 0x0a, 0x91, 0x5f, 0xa8, 0xc1, 0x59,
	0x60, 0xed, 0xcd, 0x05, 0x19, 0xfb, 0x30, 0x4f, 0x61, 0x75, 0xfa, 0x5c, 0xed, 0x5b, 0xd9, 0x2a,
	0x36, 0x33, 0xfb, 0xc7, 0x68, 0x6f, 0xff, 0x36, 0x6f, 0x45, 0x9f, 0xc2, 0x92, 0xdb, 0x30, 0x0a,
	0x0a, 0x1f, 0x2c, 0x98, 0x45, 0x7b, 0xa9, 0x6b, 0x37, 0xf0, 0x3b, 0x46, 0x63, 0xf4, 0x1c, 0x56,
	0x8e, 0x88, 0x74, 0x06, 0x80, 0x8a, 0x57, 0xce, 0x8f, 0xa6, 0xbd, 0xb1, 0xc0, 0x38, 0x65, 0x71,
	0xef, 0x3f, 0xd8, 0xce, 0x88, 0x74, 0x77, 0xbf, 0xda, 0xfb, 0xee, 0xe9, 0x0f, 0x35, 0xbd, 0xf6,
	0xef, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x7b, 0x2b, 0x6b, 0x6f, 0x06, 0x00, 0x00,
}
