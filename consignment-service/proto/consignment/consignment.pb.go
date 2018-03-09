// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

/*
Package go_micro_srv_consignment is a generated protocol buffer package.

It is generated from these files:
	proto/consignment/consignment.proto

It has these top-level messages:
	Consignment
	Container
	GetRequest
	Response
*/
package go_micro_srv_consignment

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

type Consignment struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Containers  []*Container `protobuf:"bytes,4,rep,name=containers" json:"containers,omitempty"`
	VesselId    string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId" json:"vessel_id,omitempty"`
}

func (m *Consignment) Reset()                    { *m = Consignment{} }
func (m *Consignment) String() string            { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()               {}
func (*Consignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin     string `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserUid    string `protobuf:"bytes,4,opt,name=user_uid,json=userUid" json:"user_uid,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserUid() string {
	if m != nil {
		return m.UserUid
	}
	return ""
}

type GetRequest struct {
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Response struct {
	Created      bool           `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Consignment  *Consignment   `protobuf:"bytes,2,opt,name=consignment" json:"consignment,omitempty"`
	Consignments []*Consignment `protobuf:"bytes,3,rep,name=consignments" json:"consignments,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.consignment"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	CreateConsignment(context.Context, *Consignment, *Response) error
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ShippingServiceHandler.GetConsignments(ctx, in, out)
}

func init() { proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x8a, 0xdb, 0x30,
	0x14, 0x85, 0x71, 0x7e, 0xed, 0xeb, 0xd0, 0x50, 0x2d, 0x8a, 0xda, 0x2e, 0x6a, 0x9c, 0x16, 0xb2,
	0x72, 0x21, 0x7d, 0x04, 0x2f, 0x42, 0xb6, 0x0a, 0x5d, 0x15, 0x1a, 0x52, 0xeb, 0xe2, 0x5c, 0x68,
	0x24, 0x57, 0x92, 0xd3, 0x57, 0x9b, 0x79, 0x89, 0x79, 0xa6, 0x21, 0x4a, 0x9c, 0x68, 0x18, 0x32,
	0x64, 0xa7, 0x73, 0x8f, 0xce, 0xf5, 0xe7, 0x83, 0x60, 0xd6, 0x18, 0xed, 0xf4, 0xf7, 0x4a, 0x2b,
	0x4b, 0xb5, 0xda, 0xa3, 0x72, 0xe1, 0xb9, 0xf0, 0x2e, 0xe3, 0xb5, 0x2e, 0xf6, 0x54, 0x19, 0x5d,
	0x58, 0x73, 0x28, 0x02, 0x3f, 0x7f, 0x8c, 0x20, 0x2d, 0xaf, 0x9a, 0xbd, 0x83, 0x1e, 0x49, 0x1e,
	0x65, 0xd1, 0x3c, 0x11, 0x3d, 0x92, 0x2c, 0x83, 0x54, 0xa2, 0xad, 0x0c, 0x35, 0x8e, 0xb4, 0xe2,
	0x3d, 0x6f, 0x84, 0x23, 0xf6, 0x01, 0x46, 0xff, 0x91, 0xea, 0x9d, 0xe3, 0xfd, 0x2c, 0x9a, 0x0f,
	0xc5, 0x59, 0xb1, 0x12, 0xa0, 0xd2, 0xca, 0x6d, 0x49, 0xa1, 0xb1, 0x7c, 0x90, 0xf5, 0xe7, 0xe9,
	0x62, 0x56, 0xdc, 0x02, 0x29, 0xca, 0xee, 0xae, 0x08, 0x62, 0xec, 0x33, 0x24, 0x07, 0xb4, 0x16,
	0xff, 0x6e, 0x48, 0xf2, 0xa1, 0xff, 0x78, 0x7c, 0x1a, 0xac, 0x64, 0xae, 0x21, 0xb9, 0xa4, 0x5e,
	0x81, 0x7f, 0x81, 0xb4, 0x6a, 0xad, 0xd3, 0x7b, 0x34, 0xc7, 0xec, 0x09, 0x1c, 0xba, 0xd1, 0x4a,
	0x1e, 0xb9, 0xb5, 0xa1, 0x9a, 0x94, 0xe7, 0x4e, 0xc4, 0x59, 0xb1, 0x8f, 0x10, 0xb7, 0x16, 0xcd,
	0xa6, 0x25, 0xc9, 0x07, 0xde, 0x19, 0x1f, 0xf5, 0x4f, 0x92, 0xf9, 0x04, 0x60, 0x89, 0x4e, 0xe0,
	0xbf, 0x16, 0xad, 0xcb, 0x1f, 0x22, 0x88, 0x05, 0xda, 0x46, 0x2b, 0x8b, 0x8c, 0xc3, 0xb8, 0x32,
	0xb8, 0x75, 0x78, 0x62, 0x88, 0x45, 0x27, 0xd9, 0x12, 0xd2, 0xe0, 0x3f, 0x3d, 0x48, 0xba, 0xf8,
	0xf6, 0x66, 0x11, 0xdd, 0x59, 0x84, 0x49, 0xb6, 0x82, 0x49, 0x20, 0x2d, 0xef, 0xfb, 0x4a, 0xef,
	0xdc, 0xf4, 0x22, 0xba, 0x78, 0x8a, 0x60, 0xba, 0xde, 0x51, 0xd3, 0x90, 0xaa, 0xd7, 0x68, 0x0e,
	0x54, 0x21, 0xfb, 0x0d, 0xef, 0x4b, 0x8f, 0x1c, 0x3e, 0x87, 0xfb, 0xb6, 0x7f, 0xca, 0x6f, 0x5f,
	0xbb, 0x34, 0xf4, 0x0b, 0xa6, 0x4b, 0x74, 0x41, 0xca, 0xb2, 0xaf, 0xb7, 0x63, 0xd7, 0x9e, 0xef,
	0x59, 0xfe, 0x67, 0xe4, 0xdf, 0xf9, 0x8f, 0xe7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0x8a, 0x3a,
	0x78, 0x0e, 0x03, 0x00, 0x00,
}
