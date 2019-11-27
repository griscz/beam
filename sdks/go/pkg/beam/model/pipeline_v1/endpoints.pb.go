// Code generated by protoc-gen-go. DO NOT EDIT.
// source: endpoints.proto

package pipeline_v1

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

type ApiServiceDescriptor struct {
	// (Required) The URL to connect to.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// (Optional) The method for authentication. If unspecified, access to the
	// url is already being performed in a trusted context (e.g. localhost,
	// private network).
	//
	// Types that are valid to be assigned to Authentication:
	//	*ApiServiceDescriptor_Oauth2ClientCredentialsGrant
	Authentication       isApiServiceDescriptor_Authentication `protobuf_oneof:"authentication"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ApiServiceDescriptor) Reset()         { *m = ApiServiceDescriptor{} }
func (m *ApiServiceDescriptor) String() string { return proto.CompactTextString(m) }
func (*ApiServiceDescriptor) ProtoMessage()    {}
func (*ApiServiceDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_endpoints_09141d845209a85b, []int{0}
}
func (m *ApiServiceDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiServiceDescriptor.Unmarshal(m, b)
}
func (m *ApiServiceDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiServiceDescriptor.Marshal(b, m, deterministic)
}
func (dst *ApiServiceDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiServiceDescriptor.Merge(dst, src)
}
func (m *ApiServiceDescriptor) XXX_Size() int {
	return xxx_messageInfo_ApiServiceDescriptor.Size(m)
}
func (m *ApiServiceDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiServiceDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_ApiServiceDescriptor proto.InternalMessageInfo

type isApiServiceDescriptor_Authentication interface {
	isApiServiceDescriptor_Authentication()
}

type ApiServiceDescriptor_Oauth2ClientCredentialsGrant struct {
	Oauth2ClientCredentialsGrant *OAuth2ClientCredentialsGrant `protobuf:"bytes,3,opt,name=oauth2_client_credentials_grant,json=oauth2ClientCredentialsGrant,proto3,oneof"`
}

func (*ApiServiceDescriptor_Oauth2ClientCredentialsGrant) isApiServiceDescriptor_Authentication() {}

func (m *ApiServiceDescriptor) GetAuthentication() isApiServiceDescriptor_Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *ApiServiceDescriptor) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ApiServiceDescriptor) GetOauth2ClientCredentialsGrant() *OAuth2ClientCredentialsGrant {
	if x, ok := m.GetAuthentication().(*ApiServiceDescriptor_Oauth2ClientCredentialsGrant); ok {
		return x.Oauth2ClientCredentialsGrant
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ApiServiceDescriptor) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ApiServiceDescriptor_OneofMarshaler, _ApiServiceDescriptor_OneofUnmarshaler, _ApiServiceDescriptor_OneofSizer, []interface{}{
		(*ApiServiceDescriptor_Oauth2ClientCredentialsGrant)(nil),
	}
}

func _ApiServiceDescriptor_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ApiServiceDescriptor)
	// authentication
	switch x := m.Authentication.(type) {
	case *ApiServiceDescriptor_Oauth2ClientCredentialsGrant:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Oauth2ClientCredentialsGrant); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ApiServiceDescriptor.Authentication has unexpected type %T", x)
	}
	return nil
}

func _ApiServiceDescriptor_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ApiServiceDescriptor)
	switch tag {
	case 3: // authentication.oauth2_client_credentials_grant
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OAuth2ClientCredentialsGrant)
		err := b.DecodeMessage(msg)
		m.Authentication = &ApiServiceDescriptor_Oauth2ClientCredentialsGrant{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ApiServiceDescriptor_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ApiServiceDescriptor)
	// authentication
	switch x := m.Authentication.(type) {
	case *ApiServiceDescriptor_Oauth2ClientCredentialsGrant:
		s := proto.Size(x.Oauth2ClientCredentialsGrant)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type OAuth2ClientCredentialsGrant struct {
	// (Required) The URL to submit a "client_credentials" grant type request for
	// an OAuth access token which will be used as a bearer token for requests.
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuth2ClientCredentialsGrant) Reset()         { *m = OAuth2ClientCredentialsGrant{} }
func (m *OAuth2ClientCredentialsGrant) String() string { return proto.CompactTextString(m) }
func (*OAuth2ClientCredentialsGrant) ProtoMessage()    {}
func (*OAuth2ClientCredentialsGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_endpoints_09141d845209a85b, []int{1}
}
func (m *OAuth2ClientCredentialsGrant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuth2ClientCredentialsGrant.Unmarshal(m, b)
}
func (m *OAuth2ClientCredentialsGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuth2ClientCredentialsGrant.Marshal(b, m, deterministic)
}
func (dst *OAuth2ClientCredentialsGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuth2ClientCredentialsGrant.Merge(dst, src)
}
func (m *OAuth2ClientCredentialsGrant) XXX_Size() int {
	return xxx_messageInfo_OAuth2ClientCredentialsGrant.Size(m)
}
func (m *OAuth2ClientCredentialsGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuth2ClientCredentialsGrant.DiscardUnknown(m)
}

var xxx_messageInfo_OAuth2ClientCredentialsGrant proto.InternalMessageInfo

func (m *OAuth2ClientCredentialsGrant) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*ApiServiceDescriptor)(nil), "org.apache.beam.model.pipeline.v1.ApiServiceDescriptor")
	proto.RegisterType((*OAuth2ClientCredentialsGrant)(nil), "org.apache.beam.model.pipeline.v1.OAuth2ClientCredentialsGrant")
}

func init() { proto.RegisterFile("endpoints.proto", fileDescriptor_endpoints_09141d845209a85b) }

var fileDescriptor_endpoints_09141d845209a85b = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0x5d, 0x03, 0x42, 0x36, 0xa0, 0xe1, 0xb0, 0x48, 0x11, 0x30, 0xa6, 0x4a, 0xb5, 0x98,
	0x58, 0x5a, 0x48, 0x2e, 0x8a, 0x76, 0x42, 0xec, 0x6c, 0x96, 0xcd, 0xde, 0x90, 0x0c, 0x6c, 0x76,
	0x96, 0xb9, 0xc9, 0x3d, 0x83, 0x2f, 0xe6, 0x7b, 0xc9, 0x46, 0x4e, 0x1b, 0xc9, 0x75, 0xc3, 0xfc,
	0xf0, 0xfd, 0xfc, 0x9f, 0xbe, 0x82, 0x58, 0x25, 0xc2, 0x28, 0xb5, 0x49, 0x4c, 0x42, 0xc5, 0x2d,
	0xf1, 0xd6, 0xb8, 0xe4, 0xfc, 0x0e, 0xcc, 0x06, 0xdc, 0xde, 0xec, 0xa9, 0x82, 0x60, 0x12, 0x26,
	0x08, 0x18, 0xc1, 0x34, 0xf3, 0xe9, 0x97, 0xd2, 0xd7, 0xcb, 0x84, 0xef, 0xc0, 0x0d, 0x7a, 0x78,
	0x82, 0xda, 0x33, 0x26, 0x21, 0x2e, 0x86, 0xba, 0x77, 0xe0, 0x30, 0x3a, 0x9f, 0xa8, 0x59, 0x7f,
	0x9d, 0xcf, 0xe2, 0x53, 0xe9, 0x1b, 0x72, 0x07, 0xd9, 0x2d, 0xac, 0x0f, 0x08, 0x51, 0xac, 0x67,
	0xa8, 0x20, 0x0a, 0xba, 0x50, 0xdb, 0x2d, 0xbb, 0x28, 0xa3, 0xde, 0x44, 0xcd, 0x06, 0x8b, 0x47,
	0xd3, 0x59, 0x6c, 0xde, 0x96, 0x99, 0xb4, 0x3a, 0x82, 0x56, 0x7f, 0x9c, 0x97, 0x8c, 0x79, 0x3d,
	0x5b, 0x8f, 0x7f, 0x9a, 0xfe, 0xcf, 0xcb, 0xa1, 0xbe, 0xcc, 0x71, 0xfe, 0x79, 0x27, 0x48, 0x71,
	0x7a, 0xa7, 0xc7, 0xa7, 0x88, 0xed, 0x1c, 0xf5, 0x3b, 0xa7, 0x7c, 0xd0, 0xdd, 0x7a, 0xca, 0xfe,
	0x73, 0xab, 0xf4, 0x63, 0xd0, 0xfe, 0x6d, 0x33, 0xdf, 0x5c, 0x1c, 0x05, 0xdf, 0x7f, 0x07, 0x00,
	0x00, 0xff, 0xff, 0x98, 0xdf, 0xf8, 0x2f, 0x73, 0x01, 0x00, 0x00,
}