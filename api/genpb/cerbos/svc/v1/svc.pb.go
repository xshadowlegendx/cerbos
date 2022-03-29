// Copyright 2021-2022 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: cerbos/svc/v1/svc.proto

package svcv1

import (
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/request/v1"
	v11 "github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_cerbos_svc_v1_svc_proto protoreflect.FileDescriptor

var file_cerbos_svc_v1_svc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x65, 0x72, 0x62, 0x6f,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc2, 0x07, 0x0a, 0x0d, 0x43,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xf6, 0x01, 0x0a,
	0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65,
	0x74, 0x12, 0x2a, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x87, 0x01, 0x92, 0x41,
	0x6f, 0x12, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x66, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x20,
	0x77, 0x68, 0x65, 0x74, 0x68, 0x65, 0x72, 0x20, 0x61, 0x20, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69,
	0x70, 0x61, 0x6c, 0x20, 0x68, 0x61, 0x73, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x20, 0x6f, 0x6e, 0x20, 0x61, 0x20, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x86, 0x02, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2c, 0x2e, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x91, 0x01, 0x92, 0x41, 0x6a,
	0x12, 0x14, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x20, 0x62, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x52, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x20, 0x61, 0x20,
	0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x27, 0x73, 0x20, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x61, 0x20, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x20, 0x6f, 0x66, 0x20, 0x68, 0x65, 0x74, 0x65, 0x72, 0x6f, 0x67, 0x65, 0x6e, 0x65,
	0x6f, 0x75, 0x73, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x61, 0x6e,
	0x64, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e,
	0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0xc5,
	0x01, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69, 0x92, 0x41, 0x4e,
	0x12, 0x16, 0x47, 0x65, 0x74, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x69, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x34, 0x47, 0x65, 0x74, 0x20, 0x69, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x2e, 0x67, 0x2e, 0x20,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0xc3, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x2c, 0x2e,
	0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x65,
	0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x92, 0x41, 0x2c,
	0x12, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x20, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x20, 0x70, 0x6c, 0x61, 0x6e, 0x1a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x20, 0x71, 0x75, 0x65, 0x72, 0x79, 0x20, 0x70, 0x6c, 0x61, 0x6e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x1a, 0x21, 0x92, 0x41,
	0x1e, 0x12, 0x1c, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x20, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x20, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x32,
	0xb0, 0x0c, 0x0a, 0x12, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc9, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x4f, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2b, 0x2e, 0x63,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x92, 0x41, 0x29, 0x12, 0x16, 0x41,
	0x64, 0x64, 0x20, 0x6f, 0x72, 0x20, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x20, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x22, 0x0d, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x01, 0x2a, 0x5a, 0x12,
	0x1a, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a,
	0x01, 0x2a, 0x12, 0x9c, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x65,
	0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x92, 0x41, 0x20, 0x12, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x12, 0x0f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x8e, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x23, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x92, 0x41, 0x1d,
	0x12, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x0f, 0x0a, 0x0d,
	0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0xc8, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x2e, 0x63, 0x65, 0x72,
	0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f, 0x92, 0x41, 0x29, 0x12,
	0x16, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x61, 0x75, 0x64, 0x69, 0x74, 0x20, 0x6c, 0x6f, 0x67, 0x20,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x6c, 0x6f, 0x67, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x6b, 0x69, 0x6e, 0x64, 0x7d, 0x30, 0x01, 0x12, 0xc7, 0x01,
	0x0a, 0x11, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x2b, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x56, 0x92, 0x41, 0x27, 0x12, 0x14, 0x41, 0x64, 0x64, 0x20, 0x6f, 0x72, 0x20, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x26, 0x22, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x3a, 0x01, 0x2a, 0x5a, 0x12, 0x1a, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x3a, 0x01, 0x2a, 0x12, 0x97, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x25, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x92, 0x41, 0x1f, 0x12, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x12, 0x0e, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x73, 0x12, 0x8e, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x23, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x92, 0x41, 0x1d,
	0x12, 0x0a, 0x47, 0x65, 0x74, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x0f, 0x0a, 0x0d,
	0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x9a, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x26, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x65,
	0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x92, 0x41, 0x20, 0x12, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x2a, 0x0d, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x9c, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12,
	0x25, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x6f,
	0x61, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x3d, 0x92, 0x41, 0x1f, 0x12, 0x0c, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x20, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x62, 0x0f, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x22,
	0x92, 0x41, 0x1f, 0x12, 0x1d, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x20, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x32, 0xe5, 0x04, 0x0a, 0x17, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x50, 0x6c, 0x61,
	0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x97,
	0x01, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x87, 0x01, 0x0a, 0x0e, 0x50, 0x6c, 0x61,
	0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x12, 0x28, 0x2e, 0x63, 0x65,
	0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0x97, 0x01, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x63, 0x65, 0x72, 0x62,
	0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73,
	0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22,
	0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x2f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x8b, 0x01, 0x0a,
	0x0f, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x12, 0x29, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x65,
	0x72, 0x62, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a,
	0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x3a, 0x01, 0x2a, 0x42, 0xe1, 0x01, 0x0a, 0x15, 0x64,
	0x65, 0x76, 0x2e, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x76, 0x63, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x62, 0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2f,
	0x73, 0x76, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x76, 0x63, 0x76, 0x31, 0xaa, 0x02, 0x11, 0x43,
	0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x76, 0x63,
	0x92, 0x41, 0x7b, 0x12, 0x3f, 0x0a, 0x06, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x22, 0x2d, 0x0a,
	0x06, 0x43, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x12, 0x12, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x1a, 0x0f, 0x69, 0x6e, 0x66,
	0x6f, 0x40, 0x63, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x32, 0x06, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x11, 0x0a, 0x0f, 0x0a,
	0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x02, 0x08, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_cerbos_svc_v1_svc_proto_goTypes = []interface{}{
	(*v1.CheckResourceSetRequest)(nil),      // 0: cerbos.request.v1.CheckResourceSetRequest
	(*v1.CheckResourceBatchRequest)(nil),    // 1: cerbos.request.v1.CheckResourceBatchRequest
	(*v1.ServerInfoRequest)(nil),            // 2: cerbos.request.v1.ServerInfoRequest
	(*v1.ResourcesQueryPlanRequest)(nil),    // 3: cerbos.request.v1.ResourcesQueryPlanRequest
	(*v1.AddOrUpdatePolicyRequest)(nil),     // 4: cerbos.request.v1.AddOrUpdatePolicyRequest
	(*v1.ListPoliciesRequest)(nil),          // 5: cerbos.request.v1.ListPoliciesRequest
	(*v1.GetPolicyRequest)(nil),             // 6: cerbos.request.v1.GetPolicyRequest
	(*v1.ListAuditLogEntriesRequest)(nil),   // 7: cerbos.request.v1.ListAuditLogEntriesRequest
	(*v1.AddOrUpdateSchemaRequest)(nil),     // 8: cerbos.request.v1.AddOrUpdateSchemaRequest
	(*v1.ListSchemasRequest)(nil),           // 9: cerbos.request.v1.ListSchemasRequest
	(*v1.GetSchemaRequest)(nil),             // 10: cerbos.request.v1.GetSchemaRequest
	(*v1.DeleteSchemaRequest)(nil),          // 11: cerbos.request.v1.DeleteSchemaRequest
	(*v1.ReloadStoreRequest)(nil),           // 12: cerbos.request.v1.ReloadStoreRequest
	(*v1.PlaygroundValidateRequest)(nil),    // 13: cerbos.request.v1.PlaygroundValidateRequest
	(*v1.PlaygroundTestRequest)(nil),        // 14: cerbos.request.v1.PlaygroundTestRequest
	(*v1.PlaygroundEvaluateRequest)(nil),    // 15: cerbos.request.v1.PlaygroundEvaluateRequest
	(*v1.PlaygroundProxyRequest)(nil),       // 16: cerbos.request.v1.PlaygroundProxyRequest
	(*v11.CheckResourceSetResponse)(nil),    // 17: cerbos.response.v1.CheckResourceSetResponse
	(*v11.CheckResourceBatchResponse)(nil),  // 18: cerbos.response.v1.CheckResourceBatchResponse
	(*v11.ServerInfoResponse)(nil),          // 19: cerbos.response.v1.ServerInfoResponse
	(*v11.ResourcesQueryPlanResponse)(nil),  // 20: cerbos.response.v1.ResourcesQueryPlanResponse
	(*v11.AddOrUpdatePolicyResponse)(nil),   // 21: cerbos.response.v1.AddOrUpdatePolicyResponse
	(*v11.ListPoliciesResponse)(nil),        // 22: cerbos.response.v1.ListPoliciesResponse
	(*v11.GetPolicyResponse)(nil),           // 23: cerbos.response.v1.GetPolicyResponse
	(*v11.ListAuditLogEntriesResponse)(nil), // 24: cerbos.response.v1.ListAuditLogEntriesResponse
	(*v11.AddOrUpdateSchemaResponse)(nil),   // 25: cerbos.response.v1.AddOrUpdateSchemaResponse
	(*v11.ListSchemasResponse)(nil),         // 26: cerbos.response.v1.ListSchemasResponse
	(*v11.GetSchemaResponse)(nil),           // 27: cerbos.response.v1.GetSchemaResponse
	(*v11.DeleteSchemaResponse)(nil),        // 28: cerbos.response.v1.DeleteSchemaResponse
	(*v11.ReloadStoreResponse)(nil),         // 29: cerbos.response.v1.ReloadStoreResponse
	(*v11.PlaygroundValidateResponse)(nil),  // 30: cerbos.response.v1.PlaygroundValidateResponse
	(*v11.PlaygroundTestResponse)(nil),      // 31: cerbos.response.v1.PlaygroundTestResponse
	(*v11.PlaygroundEvaluateResponse)(nil),  // 32: cerbos.response.v1.PlaygroundEvaluateResponse
	(*v11.PlaygroundProxyResponse)(nil),     // 33: cerbos.response.v1.PlaygroundProxyResponse
}
var file_cerbos_svc_v1_svc_proto_depIdxs = []int32{
	0,  // 0: cerbos.svc.v1.CerbosService.CheckResourceSet:input_type -> cerbos.request.v1.CheckResourceSetRequest
	1,  // 1: cerbos.svc.v1.CerbosService.CheckResourceBatch:input_type -> cerbos.request.v1.CheckResourceBatchRequest
	2,  // 2: cerbos.svc.v1.CerbosService.ServerInfo:input_type -> cerbos.request.v1.ServerInfoRequest
	3,  // 3: cerbos.svc.v1.CerbosService.ResourcesQueryPlan:input_type -> cerbos.request.v1.ResourcesQueryPlanRequest
	4,  // 4: cerbos.svc.v1.CerbosAdminService.AddOrUpdatePolicy:input_type -> cerbos.request.v1.AddOrUpdatePolicyRequest
	5,  // 5: cerbos.svc.v1.CerbosAdminService.ListPolicies:input_type -> cerbos.request.v1.ListPoliciesRequest
	6,  // 6: cerbos.svc.v1.CerbosAdminService.GetPolicy:input_type -> cerbos.request.v1.GetPolicyRequest
	7,  // 7: cerbos.svc.v1.CerbosAdminService.ListAuditLogEntries:input_type -> cerbos.request.v1.ListAuditLogEntriesRequest
	8,  // 8: cerbos.svc.v1.CerbosAdminService.AddOrUpdateSchema:input_type -> cerbos.request.v1.AddOrUpdateSchemaRequest
	9,  // 9: cerbos.svc.v1.CerbosAdminService.ListSchemas:input_type -> cerbos.request.v1.ListSchemasRequest
	10, // 10: cerbos.svc.v1.CerbosAdminService.GetSchema:input_type -> cerbos.request.v1.GetSchemaRequest
	11, // 11: cerbos.svc.v1.CerbosAdminService.DeleteSchema:input_type -> cerbos.request.v1.DeleteSchemaRequest
	12, // 12: cerbos.svc.v1.CerbosAdminService.ReloadStore:input_type -> cerbos.request.v1.ReloadStoreRequest
	13, // 13: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundValidate:input_type -> cerbos.request.v1.PlaygroundValidateRequest
	14, // 14: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundTest:input_type -> cerbos.request.v1.PlaygroundTestRequest
	15, // 15: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundEvaluate:input_type -> cerbos.request.v1.PlaygroundEvaluateRequest
	16, // 16: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundProxy:input_type -> cerbos.request.v1.PlaygroundProxyRequest
	17, // 17: cerbos.svc.v1.CerbosService.CheckResourceSet:output_type -> cerbos.response.v1.CheckResourceSetResponse
	18, // 18: cerbos.svc.v1.CerbosService.CheckResourceBatch:output_type -> cerbos.response.v1.CheckResourceBatchResponse
	19, // 19: cerbos.svc.v1.CerbosService.ServerInfo:output_type -> cerbos.response.v1.ServerInfoResponse
	20, // 20: cerbos.svc.v1.CerbosService.ResourcesQueryPlan:output_type -> cerbos.response.v1.ResourcesQueryPlanResponse
	21, // 21: cerbos.svc.v1.CerbosAdminService.AddOrUpdatePolicy:output_type -> cerbos.response.v1.AddOrUpdatePolicyResponse
	22, // 22: cerbos.svc.v1.CerbosAdminService.ListPolicies:output_type -> cerbos.response.v1.ListPoliciesResponse
	23, // 23: cerbos.svc.v1.CerbosAdminService.GetPolicy:output_type -> cerbos.response.v1.GetPolicyResponse
	24, // 24: cerbos.svc.v1.CerbosAdminService.ListAuditLogEntries:output_type -> cerbos.response.v1.ListAuditLogEntriesResponse
	25, // 25: cerbos.svc.v1.CerbosAdminService.AddOrUpdateSchema:output_type -> cerbos.response.v1.AddOrUpdateSchemaResponse
	26, // 26: cerbos.svc.v1.CerbosAdminService.ListSchemas:output_type -> cerbos.response.v1.ListSchemasResponse
	27, // 27: cerbos.svc.v1.CerbosAdminService.GetSchema:output_type -> cerbos.response.v1.GetSchemaResponse
	28, // 28: cerbos.svc.v1.CerbosAdminService.DeleteSchema:output_type -> cerbos.response.v1.DeleteSchemaResponse
	29, // 29: cerbos.svc.v1.CerbosAdminService.ReloadStore:output_type -> cerbos.response.v1.ReloadStoreResponse
	30, // 30: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundValidate:output_type -> cerbos.response.v1.PlaygroundValidateResponse
	31, // 31: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundTest:output_type -> cerbos.response.v1.PlaygroundTestResponse
	32, // 32: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundEvaluate:output_type -> cerbos.response.v1.PlaygroundEvaluateResponse
	33, // 33: cerbos.svc.v1.CerbosPlaygroundService.PlaygroundProxy:output_type -> cerbos.response.v1.PlaygroundProxyResponse
	17, // [17:34] is the sub-list for method output_type
	0,  // [0:17] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_cerbos_svc_v1_svc_proto_init() }
func file_cerbos_svc_v1_svc_proto_init() {
	if File_cerbos_svc_v1_svc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cerbos_svc_v1_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_cerbos_svc_v1_svc_proto_goTypes,
		DependencyIndexes: file_cerbos_svc_v1_svc_proto_depIdxs,
	}.Build()
	File_cerbos_svc_v1_svc_proto = out.File
	file_cerbos_svc_v1_svc_proto_rawDesc = nil
	file_cerbos_svc_v1_svc_proto_goTypes = nil
	file_cerbos_svc_v1_svc_proto_depIdxs = nil
}
