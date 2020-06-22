// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/extensions/filters/udp/dns_filter/v3alpha/dns_filter.proto

package envoy_extensions_filters_udp_dns_filter_v3alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/data/dns/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DnsFilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatPrefix   string                               `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	ServerConfig *DnsFilterConfig_ServerContextConfig `protobuf:"bytes,2,opt,name=server_config,json=serverConfig,proto3" json:"server_config,omitempty"`
	ClientConfig *DnsFilterConfig_ClientContextConfig `protobuf:"bytes,3,opt,name=client_config,json=clientConfig,proto3" json:"client_config,omitempty"`
}

func (x *DnsFilterConfig) Reset() {
	*x = DnsFilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsFilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsFilterConfig) ProtoMessage() {}

func (x *DnsFilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsFilterConfig.ProtoReflect.Descriptor instead.
func (*DnsFilterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDescGZIP(), []int{0}
}

func (x *DnsFilterConfig) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *DnsFilterConfig) GetServerConfig() *DnsFilterConfig_ServerContextConfig {
	if x != nil {
		return x.ServerConfig
	}
	return nil
}

func (x *DnsFilterConfig) GetClientConfig() *DnsFilterConfig_ClientContextConfig {
	if x != nil {
		return x.ClientConfig
	}
	return nil
}

type DnsFilterConfig_ServerContextConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ConfigSource:
	//	*DnsFilterConfig_ServerContextConfig_InlineDnsTable
	//	*DnsFilterConfig_ServerContextConfig_ExternalDnsTable
	ConfigSource isDnsFilterConfig_ServerContextConfig_ConfigSource `protobuf_oneof:"config_source"`
}

func (x *DnsFilterConfig_ServerContextConfig) Reset() {
	*x = DnsFilterConfig_ServerContextConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsFilterConfig_ServerContextConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsFilterConfig_ServerContextConfig) ProtoMessage() {}

func (x *DnsFilterConfig_ServerContextConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsFilterConfig_ServerContextConfig.ProtoReflect.Descriptor instead.
func (*DnsFilterConfig_ServerContextConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDescGZIP(), []int{0, 0}
}

func (m *DnsFilterConfig_ServerContextConfig) GetConfigSource() isDnsFilterConfig_ServerContextConfig_ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func (x *DnsFilterConfig_ServerContextConfig) GetInlineDnsTable() *v3.DnsTable {
	if x, ok := x.GetConfigSource().(*DnsFilterConfig_ServerContextConfig_InlineDnsTable); ok {
		return x.InlineDnsTable
	}
	return nil
}

func (x *DnsFilterConfig_ServerContextConfig) GetExternalDnsTable() *v31.DataSource {
	if x, ok := x.GetConfigSource().(*DnsFilterConfig_ServerContextConfig_ExternalDnsTable); ok {
		return x.ExternalDnsTable
	}
	return nil
}

type isDnsFilterConfig_ServerContextConfig_ConfigSource interface {
	isDnsFilterConfig_ServerContextConfig_ConfigSource()
}

type DnsFilterConfig_ServerContextConfig_InlineDnsTable struct {
	InlineDnsTable *v3.DnsTable `protobuf:"bytes,1,opt,name=inline_dns_table,json=inlineDnsTable,proto3,oneof"`
}

type DnsFilterConfig_ServerContextConfig_ExternalDnsTable struct {
	ExternalDnsTable *v31.DataSource `protobuf:"bytes,2,opt,name=external_dns_table,json=externalDnsTable,proto3,oneof"`
}

func (*DnsFilterConfig_ServerContextConfig_InlineDnsTable) isDnsFilterConfig_ServerContextConfig_ConfigSource() {
}

func (*DnsFilterConfig_ServerContextConfig_ExternalDnsTable) isDnsFilterConfig_ServerContextConfig_ConfigSource() {
}

type DnsFilterConfig_ClientContextConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResolverTimeout   *duration.Duration `protobuf:"bytes,1,opt,name=resolver_timeout,json=resolverTimeout,proto3" json:"resolver_timeout,omitempty"`
	UpstreamResolvers []*v31.Address     `protobuf:"bytes,2,rep,name=upstream_resolvers,json=upstreamResolvers,proto3" json:"upstream_resolvers,omitempty"`
	MaxPendingLookups uint64             `protobuf:"varint,3,opt,name=max_pending_lookups,json=maxPendingLookups,proto3" json:"max_pending_lookups,omitempty"`
}

func (x *DnsFilterConfig_ClientContextConfig) Reset() {
	*x = DnsFilterConfig_ClientContextConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsFilterConfig_ClientContextConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsFilterConfig_ClientContextConfig) ProtoMessage() {}

func (x *DnsFilterConfig_ClientContextConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsFilterConfig_ClientContextConfig.ProtoReflect.Descriptor instead.
func (*DnsFilterConfig_ClientContextConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDescGZIP(), []int{0, 1}
}

func (x *DnsFilterConfig_ClientContextConfig) GetResolverTimeout() *duration.Duration {
	if x != nil {
		return x.ResolverTimeout
	}
	return nil
}

func (x *DnsFilterConfig_ClientContextConfig) GetUpstreamResolvers() []*v31.Address {
	if x != nil {
		return x.UpstreamResolvers
	}
	return nil
}

func (x *DnsFilterConfig_ClientContextConfig) GetMaxPendingLookups() uint64 {
	if x != nil {
		return x.MaxPendingLookups
	}
	return 0
}

var File_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x64, 0x70, 0x2f, 0x64,
	0x6e, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x64, 0x70,
	0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x6e, 0x73, 0x5f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x05, 0x0a, 0x0f, 0x44, 0x6e, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x12, 0x79, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x54, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x6e, 0x73, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x79, 0x0a, 0x0d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x54, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x75,
	0x64, 0x70, 0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x6e, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0xc6, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x47, 0x0a, 0x10, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x64, 0x6e, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x6e,
	0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x50, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x44, 0x6e, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x14, 0x0a, 0x0d, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x03, 0xf8, 0x42, 0x01,
	0x1a, 0xf8, 0x01, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0xaa, 0x01, 0x04, 0x32, 0x02, 0x08, 0x01, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x56, 0x0a, 0x12, 0x75, 0x70,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52,
	0x11, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x73, 0x12, 0x37, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x50, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x73, 0x42, 0x61, 0x0a, 0x3d, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x64, 0x6e, 0x73, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0e, 0x44, 0x6e,
	0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x08, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDescData = file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDesc
)

func file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDescData)
	})
	return file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDescData
}

var file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_goTypes = []interface{}{
	(*DnsFilterConfig)(nil),                     // 0: envoy.extensions.filters.udp.dns_filter.v3alpha.DnsFilterConfig
	(*DnsFilterConfig_ServerContextConfig)(nil), // 1: envoy.extensions.filters.udp.dns_filter.v3alpha.DnsFilterConfig.ServerContextConfig
	(*DnsFilterConfig_ClientContextConfig)(nil), // 2: envoy.extensions.filters.udp.dns_filter.v3alpha.DnsFilterConfig.ClientContextConfig
	(*v3.DnsTable)(nil),                         // 3: envoy.data.dns.v3.DnsTable
	(*v31.DataSource)(nil),                      // 4: envoy.config.core.v3.DataSource
	(*duration.Duration)(nil),                   // 5: google.protobuf.Duration
	(*v31.Address)(nil),                         // 6: envoy.config.core.v3.Address
}
var file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.filters.udp.dns_filter.v3alpha.DnsFilterConfig.server_config:type_name -> envoy.extensions.filters.udp.dns_filter.v3alpha.DnsFilterConfig.ServerContextConfig
	2, // 1: envoy.extensions.filters.udp.dns_filter.v3alpha.DnsFilterConfig.client_config:type_name -> envoy.extensions.filters.udp.dns_filter.v3alpha.DnsFilterConfig.ClientContextConfig
	3, // 2: envoy.extensions.filters.udp.dns_filter.v3alpha.DnsFilterConfig.ServerContextConfig.inline_dns_table:type_name -> envoy.data.dns.v3.DnsTable
	4, // 3: envoy.extensions.filters.udp.dns_filter.v3alpha.DnsFilterConfig.ServerContextConfig.external_dns_table:type_name -> envoy.config.core.v3.DataSource
	5, // 4: envoy.extensions.filters.udp.dns_filter.v3alpha.DnsFilterConfig.ClientContextConfig.resolver_timeout:type_name -> google.protobuf.Duration
	6, // 5: envoy.extensions.filters.udp.dns_filter.v3alpha.DnsFilterConfig.ClientContextConfig.upstream_resolvers:type_name -> envoy.config.core.v3.Address
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_init() }
func file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_init() {
	if File_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsFilterConfig); i {
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
		file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsFilterConfig_ServerContextConfig); i {
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
		file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsFilterConfig_ClientContextConfig); i {
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
	file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DnsFilterConfig_ServerContextConfig_InlineDnsTable)(nil),
		(*DnsFilterConfig_ServerContextConfig_ExternalDnsTable)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto = out.File
	file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_rawDesc = nil
	file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_goTypes = nil
	file_envoy_extensions_filters_udp_dns_filter_v3alpha_dns_filter_proto_depIdxs = nil
}
