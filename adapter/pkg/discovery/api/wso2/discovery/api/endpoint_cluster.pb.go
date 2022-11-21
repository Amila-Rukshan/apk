//  Copyright (c) 2021, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
//
//  WSO2 LLC. licenses this file to you under the Apache License,
//  Version 2.0 (the "License"); you may not use this file except
//  in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing,
//  software distributed under the License is distributed on an
//  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
//  KIND, either express or implied.  See the License for the
//  specific language governing permissions and limitations
//  under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: wso2/discovery/api/endpoint_cluster.proto

package api

import (
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

// Holds endpoint URLs and the config
type EndpointCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urls   []*Endpoint            `protobuf:"bytes,1,rep,name=urls,proto3" json:"urls,omitempty"`
	Config *EndpointClusterConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *EndpointCluster) Reset() {
	*x = EndpointCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_endpoint_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointCluster) ProtoMessage() {}

func (x *EndpointCluster) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_endpoint_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointCluster.ProtoReflect.Descriptor instead.
func (*EndpointCluster) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_endpoint_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *EndpointCluster) GetUrls() []*Endpoint {
	if x != nil {
		return x.Urls
	}
	return nil
}

func (x *EndpointCluster) GetConfig() *EndpointClusterConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type EndpointClusterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetryConfig   *RetryConfig   `protobuf:"bytes,1,opt,name=retryConfig,proto3" json:"retryConfig,omitempty"`
	TimeoutConfig *TimeoutConfig `protobuf:"bytes,2,opt,name=timeoutConfig,proto3" json:"timeoutConfig,omitempty"`
}

func (x *EndpointClusterConfig) Reset() {
	*x = EndpointClusterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_endpoint_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointClusterConfig) ProtoMessage() {}

func (x *EndpointClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_endpoint_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointClusterConfig.ProtoReflect.Descriptor instead.
func (*EndpointClusterConfig) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_endpoint_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *EndpointClusterConfig) GetRetryConfig() *RetryConfig {
	if x != nil {
		return x.RetryConfig
	}
	return nil
}

func (x *EndpointClusterConfig) GetTimeoutConfig() *TimeoutConfig {
	if x != nil {
		return x.TimeoutConfig
	}
	return nil
}

type TimeoutConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteTimeoutInMillis uint32 `protobuf:"varint,1,opt,name=routeTimeoutInMillis,proto3" json:"routeTimeoutInMillis,omitempty"`
}

func (x *TimeoutConfig) Reset() {
	*x = TimeoutConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_endpoint_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeoutConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeoutConfig) ProtoMessage() {}

func (x *TimeoutConfig) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_endpoint_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeoutConfig.ProtoReflect.Descriptor instead.
func (*TimeoutConfig) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_endpoint_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *TimeoutConfig) GetRouteTimeoutInMillis() uint32 {
	if x != nil {
		return x.RouteTimeoutInMillis
	}
	return 0
}

type RetryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count       uint32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	StatusCodes []uint32 `protobuf:"varint,2,rep,packed,name=statusCodes,proto3" json:"statusCodes,omitempty"`
}

func (x *RetryConfig) Reset() {
	*x = RetryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wso2_discovery_api_endpoint_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetryConfig) ProtoMessage() {}

func (x *RetryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_wso2_discovery_api_endpoint_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetryConfig.ProtoReflect.Descriptor instead.
func (*RetryConfig) Descriptor() ([]byte, []int) {
	return file_wso2_discovery_api_endpoint_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *RetryConfig) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *RetryConfig) GetStatusCodes() []uint32 {
	if x != nil {
		return x.StatusCodes
	}
	return nil
}

var File_wso2_discovery_api_endpoint_cluster_proto protoreflect.FileDescriptor

var file_wso2_discovery_api_endpoint_cluster_proto_rawDesc = []byte{
	0x0a, 0x29, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x77, 0x73, 0x6f,
	0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x1a,
	0x21, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x0f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x41, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xa3, 0x01, 0x0a, 0x15,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x41, 0x0a, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x77, 0x73, 0x6f,
	0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x72, 0x65, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x47, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x43, 0x0a, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x32, 0x0a, 0x14, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x49, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x14, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x49, 0x6e,
	0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x22, 0x45, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x72, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x42, 0x7e, 0x0a,
	0x25, 0x6f, 0x72, 0x67, 0x2e, 0x77, 0x73, 0x6f, 0x32, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x65, 0x6f,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x14, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x77, 0x73, 0x6f, 0x32, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wso2_discovery_api_endpoint_cluster_proto_rawDescOnce sync.Once
	file_wso2_discovery_api_endpoint_cluster_proto_rawDescData = file_wso2_discovery_api_endpoint_cluster_proto_rawDesc
)

func file_wso2_discovery_api_endpoint_cluster_proto_rawDescGZIP() []byte {
	file_wso2_discovery_api_endpoint_cluster_proto_rawDescOnce.Do(func() {
		file_wso2_discovery_api_endpoint_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_wso2_discovery_api_endpoint_cluster_proto_rawDescData)
	})
	return file_wso2_discovery_api_endpoint_cluster_proto_rawDescData
}

var file_wso2_discovery_api_endpoint_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_wso2_discovery_api_endpoint_cluster_proto_goTypes = []interface{}{
	(*EndpointCluster)(nil),       // 0: wso2.discovery.api.EndpointCluster
	(*EndpointClusterConfig)(nil), // 1: wso2.discovery.api.EndpointClusterConfig
	(*TimeoutConfig)(nil),         // 2: wso2.discovery.api.TimeoutConfig
	(*RetryConfig)(nil),           // 3: wso2.discovery.api.RetryConfig
	(*Endpoint)(nil),              // 4: wso2.discovery.api.Endpoint
}
var file_wso2_discovery_api_endpoint_cluster_proto_depIdxs = []int32{
	4, // 0: wso2.discovery.api.EndpointCluster.urls:type_name -> wso2.discovery.api.Endpoint
	1, // 1: wso2.discovery.api.EndpointCluster.config:type_name -> wso2.discovery.api.EndpointClusterConfig
	3, // 2: wso2.discovery.api.EndpointClusterConfig.retryConfig:type_name -> wso2.discovery.api.RetryConfig
	2, // 3: wso2.discovery.api.EndpointClusterConfig.timeoutConfig:type_name -> wso2.discovery.api.TimeoutConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_wso2_discovery_api_endpoint_cluster_proto_init() }
func file_wso2_discovery_api_endpoint_cluster_proto_init() {
	if File_wso2_discovery_api_endpoint_cluster_proto != nil {
		return
	}
	file_wso2_discovery_api_Endpoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_wso2_discovery_api_endpoint_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointCluster); i {
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
		file_wso2_discovery_api_endpoint_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointClusterConfig); i {
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
		file_wso2_discovery_api_endpoint_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeoutConfig); i {
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
		file_wso2_discovery_api_endpoint_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetryConfig); i {
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
			RawDescriptor: file_wso2_discovery_api_endpoint_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wso2_discovery_api_endpoint_cluster_proto_goTypes,
		DependencyIndexes: file_wso2_discovery_api_endpoint_cluster_proto_depIdxs,
		MessageInfos:      file_wso2_discovery_api_endpoint_cluster_proto_msgTypes,
	}.Build()
	File_wso2_discovery_api_endpoint_cluster_proto = out.File
	file_wso2_discovery_api_endpoint_cluster_proto_rawDesc = nil
	file_wso2_discovery_api_endpoint_cluster_proto_goTypes = nil
	file_wso2_discovery_api_endpoint_cluster_proto_depIdxs = nil
}