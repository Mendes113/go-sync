// Copyright 2016 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: model_type_state.proto

package sync_pb

import (
	proto "github.com/golang/protobuf/proto"
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

// Sync proto to store data type global metadata in model type storage.
type ModelTypeState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The latest progress markers received from the server.
	ProgressMarker *DataTypeProgressMarker `protobuf:"bytes,1,opt,name=progress_marker,json=progressMarker" json:"progress_marker,omitempty"`
	// A data type context.  Sent to the server in every commit or update
	// request.  May be updated by either responses from the server or requests
	// made on the model thread.  The interpretation of this value may be
	// data-type specific.  Many data types ignore it.
	TypeContext *DataTypeContext `protobuf:"bytes,2,opt,name=type_context,json=typeContext" json:"type_context,omitempty"`
	// This value is set if this type's data should be encrypted on the server.
	// If this key changes, the client will need to re-commit all of its local
	// data to the server using the new encryption key.
	EncryptionKeyName *string `protobuf:"bytes,3,opt,name=encryption_key_name,json=encryptionKeyName" json:"encryption_key_name,omitempty"`
	// This flag is set to true when the first download cycle is complete.  The
	// ModelTypeProcessor should not attempt to commit any items until this
	// flag is set.
	InitialSyncDone *bool `protobuf:"varint,4,opt,name=initial_sync_done,json=initialSyncDone" json:"initial_sync_done,omitempty"`
	// A GUID that identifies the committing sync client. It's persisted within
	// the sync metadata and should be used to check the integrity of the
	// metadata. Mismatches with the guid of the running client indicates invalid
	// persisted sync metadata, because cache_guid is reset when sync is disabled,
	// and disabling sync is supposed to clear sync metadata.
	CacheGuid *string `protobuf:"bytes,5,opt,name=cache_guid,json=cacheGuid" json:"cache_guid,omitempty"`
	// Syncing account ID, representing the user.
	AuthenticatedAccountId *string `protobuf:"bytes,6,opt,name=authenticated_account_id,json=authenticatedAccountId" json:"authenticated_account_id,omitempty"`
}

func (x *ModelTypeState) Reset() {
	*x = ModelTypeState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_type_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelTypeState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelTypeState) ProtoMessage() {}

func (x *ModelTypeState) ProtoReflect() protoreflect.Message {
	mi := &file_model_type_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelTypeState.ProtoReflect.Descriptor instead.
func (*ModelTypeState) Descriptor() ([]byte, []int) {
	return file_model_type_state_proto_rawDescGZIP(), []int{0}
}

func (x *ModelTypeState) GetProgressMarker() *DataTypeProgressMarker {
	if x != nil {
		return x.ProgressMarker
	}
	return nil
}

func (x *ModelTypeState) GetTypeContext() *DataTypeContext {
	if x != nil {
		return x.TypeContext
	}
	return nil
}

func (x *ModelTypeState) GetEncryptionKeyName() string {
	if x != nil && x.EncryptionKeyName != nil {
		return *x.EncryptionKeyName
	}
	return ""
}

func (x *ModelTypeState) GetInitialSyncDone() bool {
	if x != nil && x.InitialSyncDone != nil {
		return *x.InitialSyncDone
	}
	return false
}

func (x *ModelTypeState) GetCacheGuid() string {
	if x != nil && x.CacheGuid != nil {
		return *x.CacheGuid
	}
	return ""
}

func (x *ModelTypeState) GetAuthenticatedAccountId() string {
	if x != nil && x.AuthenticatedAccountId != nil {
		return *x.AuthenticatedAccountId
	}
	return ""
}

var File_model_type_state_proto protoreflect.FileDescriptor

var file_model_type_state_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70,
	0x62, 0x1a, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02,
	0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x48, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x79, 0x6e, 0x63, 0x44,
	0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x67, 0x75, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x47, 0x75,
	0x69, 0x64, 0x12, 0x38, 0x0a, 0x18, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x2b, 0x0a, 0x25,
	0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_model_type_state_proto_rawDescOnce sync.Once
	file_model_type_state_proto_rawDescData = file_model_type_state_proto_rawDesc
)

func file_model_type_state_proto_rawDescGZIP() []byte {
	file_model_type_state_proto_rawDescOnce.Do(func() {
		file_model_type_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_type_state_proto_rawDescData)
	})
	return file_model_type_state_proto_rawDescData
}

var file_model_type_state_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_type_state_proto_goTypes = []interface{}{
	(*ModelTypeState)(nil),         // 0: sync_pb.ModelTypeState
	(*DataTypeProgressMarker)(nil), // 1: sync_pb.DataTypeProgressMarker
	(*DataTypeContext)(nil),        // 2: sync_pb.DataTypeContext
}
var file_model_type_state_proto_depIdxs = []int32{
	1, // 0: sync_pb.ModelTypeState.progress_marker:type_name -> sync_pb.DataTypeProgressMarker
	2, // 1: sync_pb.ModelTypeState.type_context:type_name -> sync_pb.DataTypeContext
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_model_type_state_proto_init() }
func file_model_type_state_proto_init() {
	if File_model_type_state_proto != nil {
		return
	}
	file_sync_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_model_type_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelTypeState); i {
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
			RawDescriptor: file_model_type_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_type_state_proto_goTypes,
		DependencyIndexes: file_model_type_state_proto_depIdxs,
		MessageInfos:      file_model_type_state_proto_msgTypes,
	}.Build()
	File_model_type_state_proto = out.File
	file_model_type_state_proto_rawDesc = nil
	file_model_type_state_proto_goTypes = nil
	file_model_type_state_proto_depIdxs = nil
}
