// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: envoy/extensions/matching/common_inputs/network/v3/network_inputs.proto

package networkv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
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

// Specifies that matching should be performed by the destination IP address.
// [#extension: envoy.matching.inputs.destination_ip]
type DestinationIPInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DestinationIPInput) Reset() {
	*x = DestinationIPInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestinationIPInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationIPInput) ProtoMessage() {}

func (x *DestinationIPInput) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationIPInput.ProtoReflect.Descriptor instead.
func (*DestinationIPInput) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescGZIP(), []int{0}
}

// Specifies that matching should be performed by the destination port.
// [#extension: envoy.matching.inputs.destination_port]
type DestinationPortInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DestinationPortInput) Reset() {
	*x = DestinationPortInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestinationPortInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestinationPortInput) ProtoMessage() {}

func (x *DestinationPortInput) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestinationPortInput.ProtoReflect.Descriptor instead.
func (*DestinationPortInput) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescGZIP(), []int{1}
}

// Specifies that matching should be performed by the source IP address.
// [#extension: envoy.matching.inputs.source_ip]
type SourceIPInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SourceIPInput) Reset() {
	*x = SourceIPInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceIPInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceIPInput) ProtoMessage() {}

func (x *SourceIPInput) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceIPInput.ProtoReflect.Descriptor instead.
func (*SourceIPInput) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescGZIP(), []int{2}
}

// Specifies that matching should be performed by the source port.
// [#extension: envoy.matching.inputs.source_port]
type SourcePortInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SourcePortInput) Reset() {
	*x = SourcePortInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourcePortInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourcePortInput) ProtoMessage() {}

func (x *SourcePortInput) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourcePortInput.ProtoReflect.Descriptor instead.
func (*SourcePortInput) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescGZIP(), []int{3}
}

// Input that matches by the directly connected source IP address (this
// will only be different from the source IP address when using a listener
// filter that overrides the source address, such as the :ref:`Proxy Protocol
// listener filter <config_listener_filters_proxy_protocol>`).
// [#extension: envoy.matching.inputs.direct_source_ip]
type DirectSourceIPInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DirectSourceIPInput) Reset() {
	*x = DirectSourceIPInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DirectSourceIPInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectSourceIPInput) ProtoMessage() {}

func (x *DirectSourceIPInput) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectSourceIPInput.ProtoReflect.Descriptor instead.
func (*DirectSourceIPInput) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescGZIP(), []int{4}
}

// Input that matches by the source IP type.
// Specifies the source IP match type. The values include:
//
// * “local“ - matches a connection originating from the same host,
// [#extension: envoy.matching.inputs.source_type]
type SourceTypeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SourceTypeInput) Reset() {
	*x = SourceTypeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceTypeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceTypeInput) ProtoMessage() {}

func (x *SourceTypeInput) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceTypeInput.ProtoReflect.Descriptor instead.
func (*SourceTypeInput) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescGZIP(), []int{5}
}

// Input that matches by the requested server name (e.g. SNI in TLS).
//
// :ref:`TLS Inspector <config_listener_filters_tls_inspector>` provides the requested server name based on SNI,
// when TLS protocol is detected.
// [#extension: envoy.matching.inputs.server_name]
type ServerNameInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerNameInput) Reset() {
	*x = ServerNameInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerNameInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerNameInput) ProtoMessage() {}

func (x *ServerNameInput) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerNameInput.ProtoReflect.Descriptor instead.
func (*ServerNameInput) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescGZIP(), []int{6}
}

// Input that matches by the transport protocol.
//
// Suggested values include:
//
//   - “raw_buffer“ - default, used when no transport protocol is detected,
//   - “tls“ - set by :ref:`envoy.filters.listener.tls_inspector <config_listener_filters_tls_inspector>`
//     when TLS protocol is detected.
//
// [#extension: envoy.matching.inputs.transport_protocol]
type TransportProtocolInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TransportProtocolInput) Reset() {
	*x = TransportProtocolInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransportProtocolInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransportProtocolInput) ProtoMessage() {}

func (x *TransportProtocolInput) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransportProtocolInput.ProtoReflect.Descriptor instead.
func (*TransportProtocolInput) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescGZIP(), []int{7}
}

// List of quoted and comma-separated requested application protocols. The list consists of a
// single negotiated application protocol once the network stream is established.
//
// Examples:
//
// * “'h2','http/1.1'“
// * “'h2c'“
//
// Suggested values in the list include:
//
//   - “http/1.1“ - set by :ref:`envoy.filters.listener.tls_inspector
//     <config_listener_filters_tls_inspector>` and :ref:`envoy.filters.listener.http_inspector
//     <config_listener_filters_http_inspector>`,
//   - “h2“ - set by :ref:`envoy.filters.listener.tls_inspector <config_listener_filters_tls_inspector>`
//   - “h2c“ - set by :ref:`envoy.filters.listener.http_inspector <config_listener_filters_http_inspector>`
//
// .. attention::
//
//	Currently, :ref:`TLS Inspector <config_listener_filters_tls_inspector>` provides
//	application protocol detection based on the requested
//	`ALPN <https://en.wikipedia.org/wiki/Application-Layer_Protocol_Negotiation>`_ values.
//
//	However, the use of ALPN is pretty much limited to the HTTP/2 traffic on the Internet,
//	and matching on values other than ``h2`` is going to lead to a lot of false negatives,
//	unless all connecting clients are known to use ALPN.
//
// [#extension: envoy.matching.inputs.application_protocol]
type ApplicationProtocolInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApplicationProtocolInput) Reset() {
	*x = ApplicationProtocolInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationProtocolInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationProtocolInput) ProtoMessage() {}

func (x *ApplicationProtocolInput) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationProtocolInput.ProtoReflect.Descriptor instead.
func (*ApplicationProtocolInput) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescGZIP(), []int{8}
}

var File_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto protoreflect.FileDescriptor

var file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDesc = []byte{
	0x0a, 0x47, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x32, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x50, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x50, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x15,
	0x0a, 0x13, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x50,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x42, 0xc5, 0x01, 0x0a, 0x40, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x76, 0x33, 0x42, 0x12, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x76, 0x33, 0x3b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x76,
	0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescOnce sync.Once
	file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescData = file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDesc
)

func file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescGZIP() []byte {
	file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescData)
	})
	return file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDescData
}

var file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_goTypes = []interface{}{
	(*DestinationIPInput)(nil),       // 0: envoy.extensions.matching.common_inputs.network.v3.DestinationIPInput
	(*DestinationPortInput)(nil),     // 1: envoy.extensions.matching.common_inputs.network.v3.DestinationPortInput
	(*SourceIPInput)(nil),            // 2: envoy.extensions.matching.common_inputs.network.v3.SourceIPInput
	(*SourcePortInput)(nil),          // 3: envoy.extensions.matching.common_inputs.network.v3.SourcePortInput
	(*DirectSourceIPInput)(nil),      // 4: envoy.extensions.matching.common_inputs.network.v3.DirectSourceIPInput
	(*SourceTypeInput)(nil),          // 5: envoy.extensions.matching.common_inputs.network.v3.SourceTypeInput
	(*ServerNameInput)(nil),          // 6: envoy.extensions.matching.common_inputs.network.v3.ServerNameInput
	(*TransportProtocolInput)(nil),   // 7: envoy.extensions.matching.common_inputs.network.v3.TransportProtocolInput
	(*ApplicationProtocolInput)(nil), // 8: envoy.extensions.matching.common_inputs.network.v3.ApplicationProtocolInput
}
var file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_init() }
func file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_init() {
	if File_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestinationIPInput); i {
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
		file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestinationPortInput); i {
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
		file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceIPInput); i {
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
		file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourcePortInput); i {
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
		file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DirectSourceIPInput); i {
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
		file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceTypeInput); i {
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
		file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerNameInput); i {
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
		file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransportProtocolInput); i {
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
		file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationProtocolInput); i {
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
			RawDescriptor: file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_msgTypes,
	}.Build()
	File_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto = out.File
	file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_rawDesc = nil
	file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_goTypes = nil
	file_envoy_extensions_matching_common_inputs_network_v3_network_inputs_proto_depIdxs = nil
}
