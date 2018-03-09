// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

/*
Package go_micro_srv_vessel is a generated protocol buffer package.

It is generated from these files:
	proto/vessel/vessel.proto

It has these top-level messages:
	Vessel
	Specification
	Response
*/
package go_micro_srv_vessel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Vessel struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Capacity  int32  `protobuf:"varint,2,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight int32  `protobuf:"varint,3,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Available bool   `protobuf:"varint,5,opt,name=available" json:"available,omitempty"`
	OwnerId   string `protobuf:"bytes,6,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
}

func (m *Vessel) Reset()                    { *m = Vessel{} }
func (m *Vessel) String() string            { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()               {}
func (*Vessel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity  int32 `protobuf:"varint,1,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight int32 `protobuf:"varint,2,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
}

func (m *Specification) Reset()                    { *m = Specification{} }
func (m *Specification) String() string            { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()               {}
func (*Specification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel  *Vessel   `protobuf:"bytes,1,opt,name=vessel" json:"vessel,omitempty"`
	Vessels []*Vessel `protobuf:"bytes,2,rep,name=vessels" json:"vessels,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "go.micro.srv.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "go.micro.srv.vessel.Specification")
	proto.RegisterType((*Response)(nil), "go.micro.srv.vessel.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x65, 0xd3, 0x36, 0x4d, 0x47, 0xea, 0x61, 0xbd, 0x6c, 0xab, 0x85, 0x90, 0x53, 0x4e, 0x2b,
	0xb4, 0xf8, 0x01, 0x5e, 0x04, 0x3d, 0x6e, 0x41, 0x8f, 0x65, 0xbb, 0x19, 0xeb, 0x40, 0x92, 0x0d,
	0xd9, 0x90, 0xd6, 0xbf, 0xf1, 0x53, 0x85, 0x4d, 0x5b, 0xa9, 0x04, 0x3d, 0xed, 0xcc, 0x9b, 0x37,
	0x8f, 0x37, 0x6f, 0x61, 0x56, 0xd5, 0xb6, 0xb1, 0xf7, 0x2d, 0x3a, 0x87, 0xf9, 0xf1, 0x91, 0x1e,
	0xe3, 0x37, 0x3b, 0x2b, 0x0b, 0x32, 0xb5, 0x95, 0xae, 0x6e, 0x65, 0x37, 0x4a, 0xbe, 0x18, 0x84,
	0xaf, 0xbe, 0xe4, 0xd7, 0x10, 0x50, 0x26, 0x58, 0xcc, 0xd2, 0x89, 0x0a, 0x28, 0xe3, 0x73, 0x88,
	0x8c, 0xae, 0xb4, 0xa1, 0xe6, 0x53, 0x04, 0x31, 0x4b, 0x47, 0xea, 0xdc, 0xf3, 0x05, 0x40, 0xa1,
	0x0f, 0x9b, 0x3d, 0xd2, 0xee, 0xa3, 0x11, 0x03, 0x3f, 0x9d, 0x14, 0xfa, 0xf0, 0xe6, 0x01, 0xce,
	0x61, 0x58, 0xea, 0x02, 0xc5, 0xd0, 0x8b, 0xf9, 0x9a, 0xdf, 0xc1, 0x44, 0xb7, 0x9a, 0x72, 0xbd,
	0xcd, 0x51, 0x8c, 0x62, 0x96, 0x46, 0xea, 0x07, 0xe0, 0x33, 0x88, 0xec, 0xbe, 0xc4, 0x7a, 0x43,
	0x99, 0x08, 0xfd, 0xd6, 0xd8, 0xf7, 0xcf, 0x59, 0xf2, 0x02, 0xd3, 0x75, 0x85, 0x86, 0xde, 0xc9,
	0xe8, 0x86, 0x6c, 0x79, 0x61, 0x8c, 0xfd, 0x69, 0x2c, 0xf8, 0x65, 0x2c, 0x69, 0x21, 0x52, 0xe8,
	0x2a, 0x5b, 0x3a, 0xe4, 0x2b, 0x08, 0xbb, 0x10, 0xbc, 0xc8, 0xd5, 0xf2, 0x56, 0xf6, 0x04, 0x24,
	0xbb, 0x70, 0xd4, 0x91, 0xca, 0x1f, 0x60, 0xdc, 0x55, 0x4e, 0x04, 0xf1, 0xe0, 0xbf, 0xad, 0x13,
	0x77, 0x69, 0x60, 0xda, 0x41, 0x6b, 0xac, 0x5b, 0x32, 0xc8, 0x15, 0x4c, 0x9f, 0xa8, 0xcc, 0x1e,
	0xcf, 0x01, 0x24, 0xbd, 0x3a, 0x17, 0x87, 0xcf, 0x17, 0xbd, 0x9c, 0xd3, 0x41, 0xdb, 0xd0, 0xff,
	0xf3, 0xea, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xa6, 0x33, 0x27, 0x04, 0x02, 0x00, 0x00,
}
