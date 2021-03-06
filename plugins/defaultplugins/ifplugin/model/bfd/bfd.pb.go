// Code generated by protoc-gen-gogo.
// source: bfd.proto
// DO NOT EDIT!

/*
Package bfd is a generated protocol buffer package.

Package bfd provides data model to single-hop UDP transport BFD based on RFC 5880 and RFC 5881.

It is generated from these files:
	bfd.proto

It has these top-level messages:
	SingleHopBFD
*/
package bfd

import "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type SingleHopBFD_Key_AuthenticationType int32

const (
	SingleHopBFD_Key_KEYED_SHA1            SingleHopBFD_Key_AuthenticationType = 0
	SingleHopBFD_Key_METICULOUS_KEYED_SHA1 SingleHopBFD_Key_AuthenticationType = 1
)

var SingleHopBFD_Key_AuthenticationType_name = map[int32]string{
	0: "KEYED_SHA1",
	1: "METICULOUS_KEYED_SHA1",
}
var SingleHopBFD_Key_AuthenticationType_value = map[string]int32{
	"KEYED_SHA1":            0,
	"METICULOUS_KEYED_SHA1": 1,
}

func (x SingleHopBFD_Key_AuthenticationType) String() string {
	return proto.EnumName(SingleHopBFD_Key_AuthenticationType_name, int32(x))
}

type SingleHopBFD struct {
	Sessions     []*SingleHopBFD_Session    `protobuf:"bytes,1,rep,name=sessions" json:"sessions,omitempty"`
	Keys         []*SingleHopBFD_Key        `protobuf:"bytes,2,rep,name=keys" json:"keys,omitempty"`
	EchoFunction *SingleHopBFD_EchoFunction `protobuf:"bytes,3,opt,name=echo_function" json:"echo_function,omitempty"`
}

func (m *SingleHopBFD) Reset()         { *m = SingleHopBFD{} }
func (m *SingleHopBFD) String() string { return proto.CompactTextString(m) }
func (*SingleHopBFD) ProtoMessage()    {}

func (m *SingleHopBFD) GetSessions() []*SingleHopBFD_Session {
	if m != nil {
		return m.Sessions
	}
	return nil
}

func (m *SingleHopBFD) GetKeys() []*SingleHopBFD_Key {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *SingleHopBFD) GetEchoFunction() *SingleHopBFD_EchoFunction {
	if m != nil {
		return m.EchoFunction
	}
	return nil
}

type SingleHopBFD_Session struct {
	Interface             string                               `protobuf:"bytes,3,opt,name=interface,proto3" json:"interface,omitempty"`
	DestinationAddress    string                               `protobuf:"bytes,4,opt,name=destination_address,proto3" json:"destination_address,omitempty"`
	SourceAddress         string                               `protobuf:"bytes,5,opt,name=source_address,proto3" json:"source_address,omitempty"`
	Enabled               bool                                 `protobuf:"varint,7,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DesiredMinTxInterval  uint32                               `protobuf:"varint,8,opt,name=desired_min_tx_interval,proto3" json:"desired_min_tx_interval,omitempty"`
	RequiredMinRxInterval uint32                               `protobuf:"varint,9,opt,name=required_min_rx_interval,proto3" json:"required_min_rx_interval,omitempty"`
	DetectMultiplier      uint32                               `protobuf:"varint,10,opt,name=detect_multiplier,proto3" json:"detect_multiplier,omitempty"`
	Authentication        *SingleHopBFD_Session_Authentication `protobuf:"bytes,11,opt,name=authentication" json:"authentication,omitempty"`
}

func (m *SingleHopBFD_Session) Reset()         { *m = SingleHopBFD_Session{} }
func (m *SingleHopBFD_Session) String() string { return proto.CompactTextString(m) }
func (*SingleHopBFD_Session) ProtoMessage()    {}

func (m *SingleHopBFD_Session) GetAuthentication() *SingleHopBFD_Session_Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

type SingleHopBFD_Session_Authentication struct {
	KeyId           uint32 `protobuf:"varint,1,opt,name=key_id,proto3" json:"key_id,omitempty"`
	AdvertisedKeyId uint32 `protobuf:"varint,2,opt,name=advertised_key_id,proto3" json:"advertised_key_id,omitempty"`
}

func (m *SingleHopBFD_Session_Authentication) Reset()         { *m = SingleHopBFD_Session_Authentication{} }
func (m *SingleHopBFD_Session_Authentication) String() string { return proto.CompactTextString(m) }
func (*SingleHopBFD_Session_Authentication) ProtoMessage()    {}

type SingleHopBFD_Key struct {
	Name               string                              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AuthKeyIndex       uint32                              `protobuf:"varint,2,opt,name=auth_key_index,proto3" json:"auth_key_index,omitempty"`
	Id                 uint32                              `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	AuthenticationType SingleHopBFD_Key_AuthenticationType `protobuf:"varint,4,opt,name=authentication_type,proto3,enum=bfd.SingleHopBFD_Key_AuthenticationType" json:"authentication_type,omitempty"`
	Secret             string                              `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (m *SingleHopBFD_Key) Reset()         { *m = SingleHopBFD_Key{} }
func (m *SingleHopBFD_Key) String() string { return proto.CompactTextString(m) }
func (*SingleHopBFD_Key) ProtoMessage()    {}

type SingleHopBFD_EchoFunction struct {
	Name                string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	EchoSourceInterface string `protobuf:"bytes,2,opt,name=echo_source_interface,proto3" json:"echo_source_interface,omitempty"`
}

func (m *SingleHopBFD_EchoFunction) Reset()         { *m = SingleHopBFD_EchoFunction{} }
func (m *SingleHopBFD_EchoFunction) String() string { return proto.CompactTextString(m) }
func (*SingleHopBFD_EchoFunction) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("bfd.SingleHopBFD_Key_AuthenticationType", SingleHopBFD_Key_AuthenticationType_name, SingleHopBFD_Key_AuthenticationType_value)
}
