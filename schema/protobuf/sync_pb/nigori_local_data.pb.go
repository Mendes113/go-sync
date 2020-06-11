// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: nigori_local_data.proto

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

type CryptographerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains all known Nigori keys.
	KeyBag *NigoriKeyBag `protobuf:"bytes,1,opt,name=key_bag,json=keyBag" json:"key_bag,omitempty"`
	// Default key is the key, that should be used for encryption. Can be empty
	// in case we have pending keys (waiting for explicit passphrase, or client
	// didn't received keystore keys).
	DefaultKeyName *string `protobuf:"bytes,2,opt,name=default_key_name,json=defaultKeyName" json:"default_key_name,omitempty"`
}

func (x *CryptographerData) Reset() {
	*x = CryptographerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nigori_local_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptographerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptographerData) ProtoMessage() {}

func (x *CryptographerData) ProtoReflect() protoreflect.Message {
	mi := &file_nigori_local_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptographerData.ProtoReflect.Descriptor instead.
func (*CryptographerData) Descriptor() ([]byte, []int) {
	return file_nigori_local_data_proto_rawDescGZIP(), []int{0}
}

func (x *CryptographerData) GetKeyBag() *NigoriKeyBag {
	if x != nil {
		return x.KeyBag
	}
	return nil
}

func (x *CryptographerData) GetDefaultKeyName() string {
	if x != nil && x.DefaultKeyName != nil {
		return *x.DefaultKeyName
	}
	return ""
}

type CustomPassphraseKeyDerivationParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Method used for deriving keys from custom passphrase.
	CustomPassphraseKeyDerivationMethod *NigoriSpecifics_KeyDerivationMethod `protobuf:"varint,1,opt,name=custom_passphrase_key_derivation_method,json=customPassphraseKeyDerivationMethod,enum=sync_pb.NigoriSpecifics_KeyDerivationMethod" json:"custom_passphrase_key_derivation_method,omitempty"`
	// Salt used for the derivation of the key from the custom passphrase. Should
	// be set iff custom_passphrase_key_derivation_method == SCRYPT_8192_8_11.
	CustomPassphraseKeyDerivationSalt *string `protobuf:"bytes,2,opt,name=custom_passphrase_key_derivation_salt,json=customPassphraseKeyDerivationSalt" json:"custom_passphrase_key_derivation_salt,omitempty"`
}

func (x *CustomPassphraseKeyDerivationParams) Reset() {
	*x = CustomPassphraseKeyDerivationParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nigori_local_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomPassphraseKeyDerivationParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomPassphraseKeyDerivationParams) ProtoMessage() {}

func (x *CustomPassphraseKeyDerivationParams) ProtoReflect() protoreflect.Message {
	mi := &file_nigori_local_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomPassphraseKeyDerivationParams.ProtoReflect.Descriptor instead.
func (*CustomPassphraseKeyDerivationParams) Descriptor() ([]byte, []int) {
	return file_nigori_local_data_proto_rawDescGZIP(), []int{1}
}

func (x *CustomPassphraseKeyDerivationParams) GetCustomPassphraseKeyDerivationMethod() NigoriSpecifics_KeyDerivationMethod {
	if x != nil && x.CustomPassphraseKeyDerivationMethod != nil {
		return *x.CustomPassphraseKeyDerivationMethod
	}
	return NigoriSpecifics_UNSPECIFIED
}

func (x *CustomPassphraseKeyDerivationParams) GetCustomPassphraseKeyDerivationSalt() string {
	if x != nil && x.CustomPassphraseKeyDerivationSalt != nil {
		return *x.CustomPassphraseKeyDerivationSalt
	}
	return ""
}

type NigoriModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents state of the cryptographer.
	CryptographerData *CryptographerData `protobuf:"bytes,1,opt,name=cryptographer_data,json=cryptographerData" json:"cryptographer_data,omitempty"`
	// Contains the name of the key, used for encryption of
	// NigoriSpecifics.keystore_decryptor_token, should always be the name of the
	// latest keystore key. Equals to cryptographer_data.default_key_name iff it
	// is Nigori in full keystore mode (not in backward compatible keystore
	// mode).
	CurrentKeystoreKeyName *string `protobuf:"bytes,2,opt,name=current_keystore_key_name,json=currentKeystoreKeyName" json:"current_keystore_key_name,omitempty"`
	// Contains keybag in encrypted form, should be empty once we were able to
	// decrypt keybag from specifics.
	PendingKeys *EncryptedData `protobuf:"bytes,3,opt,name=pending_keys,json=pendingKeys" json:"pending_keys,omitempty"`
	// PassphraseType used for encryption. IMPLICIT_PASSPRHASE shouldn't be used
	// here.
	PassphraseType *NigoriSpecifics_PassphraseType `protobuf:"varint,4,opt,name=passphrase_type,json=passphraseType,enum=sync_pb.NigoriSpecifics_PassphraseType,def=0" json:"passphrase_type,omitempty"`
	// The time (in UNIX epoch milliseconds) at which the keystore migration was
	// performed.
	KeystoreMigrationTime *int64 `protobuf:"varint,5,opt,name=keystore_migration_time,json=keystoreMigrationTime" json:"keystore_migration_time,omitempty"`
	// The time (in UNIX epoch milliseconds) at which a custom passphrase was
	// set.
	// Note: this field may not be set if the custom passphrase was applied before
	// corresponding field in NigoriSpecifics was introduced.
	CustomPassphraseTime *int64 `protobuf:"varint,6,opt,name=custom_passphrase_time,json=customPassphraseTime" json:"custom_passphrase_time,omitempty"`
	// Params used for deriving keys from custom passphrase. Should be set iff
	// |passphrase_type| is CUSTOM_PASSPHRASE.
	CustomPassphraseKeyDerivationParams *CustomPassphraseKeyDerivationParams `protobuf:"bytes,7,opt,name=custom_passphrase_key_derivation_params,json=customPassphraseKeyDerivationParams" json:"custom_passphrase_key_derivation_params,omitempty"`
	// Indicates whether we need to encrypt all encryptable user types.
	EncryptEverything *bool `protobuf:"varint,8,opt,name=encrypt_everything,json=encryptEverything" json:"encrypt_everything,omitempty"`
	// The list of encrypted UserEncryptableTypes, represented by their specifics
	// field number.
	EncryptedTypesSpecificsFieldNumber []int32 `protobuf:"varint,9,rep,name=encrypted_types_specifics_field_number,json=encryptedTypesSpecificsFieldNumber" json:"encrypted_types_specifics_field_number,omitempty"`
	// Keystore keys are used to decrypt keystore-based Nigori. Should be
	// persisted in order to not ask the keystore server for them during every
	// browser startup. Due to backward compatibility requirements keys are
	// always Base64 encoded.
	KeystoreKey []string `protobuf:"bytes,10,rep,name=keystore_key,json=keystoreKey" json:"keystore_key,omitempty"`
	// Encryptor keystore decryptor token. Used for decryption of keystore Nigori
	// in case keystore keys arrived after NigoriSpecifics.
	PendingKeystoreDecryptorToken *EncryptedData `protobuf:"bytes,11,opt,name=pending_keystore_decryptor_token,json=pendingKeystoreDecryptorToken" json:"pending_keystore_decryptor_token,omitempty"`
	// Contains the name of the latest available trusted vault key that was used
	// as the default encryption key. Resets when the client go out of pending
	// decryption state and transits to other passphrase types.
	LastDefaultTrustedVaultKeyName *string `protobuf:"bytes,12,opt,name=last_default_trusted_vault_key_name,json=lastDefaultTrustedVaultKeyName" json:"last_default_trusted_vault_key_name,omitempty"`
}

// Default values for NigoriModel fields.
const (
	Default_NigoriModel_PassphraseType = NigoriSpecifics_UNKNOWN
)

func (x *NigoriModel) Reset() {
	*x = NigoriModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nigori_local_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NigoriModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NigoriModel) ProtoMessage() {}

func (x *NigoriModel) ProtoReflect() protoreflect.Message {
	mi := &file_nigori_local_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NigoriModel.ProtoReflect.Descriptor instead.
func (*NigoriModel) Descriptor() ([]byte, []int) {
	return file_nigori_local_data_proto_rawDescGZIP(), []int{2}
}

func (x *NigoriModel) GetCryptographerData() *CryptographerData {
	if x != nil {
		return x.CryptographerData
	}
	return nil
}

func (x *NigoriModel) GetCurrentKeystoreKeyName() string {
	if x != nil && x.CurrentKeystoreKeyName != nil {
		return *x.CurrentKeystoreKeyName
	}
	return ""
}

func (x *NigoriModel) GetPendingKeys() *EncryptedData {
	if x != nil {
		return x.PendingKeys
	}
	return nil
}

func (x *NigoriModel) GetPassphraseType() NigoriSpecifics_PassphraseType {
	if x != nil && x.PassphraseType != nil {
		return *x.PassphraseType
	}
	return Default_NigoriModel_PassphraseType
}

func (x *NigoriModel) GetKeystoreMigrationTime() int64 {
	if x != nil && x.KeystoreMigrationTime != nil {
		return *x.KeystoreMigrationTime
	}
	return 0
}

func (x *NigoriModel) GetCustomPassphraseTime() int64 {
	if x != nil && x.CustomPassphraseTime != nil {
		return *x.CustomPassphraseTime
	}
	return 0
}

func (x *NigoriModel) GetCustomPassphraseKeyDerivationParams() *CustomPassphraseKeyDerivationParams {
	if x != nil {
		return x.CustomPassphraseKeyDerivationParams
	}
	return nil
}

func (x *NigoriModel) GetEncryptEverything() bool {
	if x != nil && x.EncryptEverything != nil {
		return *x.EncryptEverything
	}
	return false
}

func (x *NigoriModel) GetEncryptedTypesSpecificsFieldNumber() []int32 {
	if x != nil {
		return x.EncryptedTypesSpecificsFieldNumber
	}
	return nil
}

func (x *NigoriModel) GetKeystoreKey() []string {
	if x != nil {
		return x.KeystoreKey
	}
	return nil
}

func (x *NigoriModel) GetPendingKeystoreDecryptorToken() *EncryptedData {
	if x != nil {
		return x.PendingKeystoreDecryptorToken
	}
	return nil
}

func (x *NigoriModel) GetLastDefaultTrustedVaultKeyName() string {
	if x != nil && x.LastDefaultTrustedVaultKeyName != nil {
		return *x.LastDefaultTrustedVaultKeyName
	}
	return ""
}

// Sync proto to store Nigori data in storage. Proto should be encrypted with
// os_crypt before storing it somewhere, because it contains sensitive data (
// nigori_model.cryptographer_data.key_bag and nigori_model.keystore_keys).
type NigoriLocalData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Global metadata.
	ModelTypeState *ModelTypeState `protobuf:"bytes,1,opt,name=model_type_state,json=modelTypeState" json:"model_type_state,omitempty"`
	// Metadata for Nigori entity.
	EntityMetadata *EntityMetadata `protobuf:"bytes,2,opt,name=entity_metadata,json=entityMetadata" json:"entity_metadata,omitempty"`
	// Nigori model state.
	NigoriModel *NigoriModel `protobuf:"bytes,3,opt,name=nigori_model,json=nigoriModel" json:"nigori_model,omitempty"`
}

func (x *NigoriLocalData) Reset() {
	*x = NigoriLocalData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nigori_local_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NigoriLocalData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NigoriLocalData) ProtoMessage() {}

func (x *NigoriLocalData) ProtoReflect() protoreflect.Message {
	mi := &file_nigori_local_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NigoriLocalData.ProtoReflect.Descriptor instead.
func (*NigoriLocalData) Descriptor() ([]byte, []int) {
	return file_nigori_local_data_proto_rawDescGZIP(), []int{3}
}

func (x *NigoriLocalData) GetModelTypeState() *ModelTypeState {
	if x != nil {
		return x.ModelTypeState
	}
	return nil
}

func (x *NigoriLocalData) GetEntityMetadata() *EntityMetadata {
	if x != nil {
		return x.EntityMetadata
	}
	return nil
}

func (x *NigoriLocalData) GetNigoriModel() *NigoriModel {
	if x != nil {
		return x.NigoriModel
	}
	return nil
}

var File_nigori_local_data_proto protoreflect.FileDescriptor

var file_nigori_local_data_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6e, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x62, 0x1a, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6e, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x11, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x2e, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x5f, 0x62, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x69, 0x67, 0x6f,
	0x72, 0x69, 0x4b, 0x65, 0x79, 0x42, 0x61, 0x67, 0x52, 0x06, 0x6b, 0x65, 0x79, 0x42, 0x61, 0x67,
	0x12, 0x28, 0x0a, 0x10, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xfc, 0x01, 0x0a, 0x23, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x4b,
	0x65, 0x79, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x82, 0x01, 0x0a, 0x27, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x4e,
	0x69, 0x67, 0x6f, 0x72, 0x69, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x4b,
	0x65, 0x79, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x23, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68,
	0x72, 0x61, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x50, 0x0a, 0x25, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x5f, 0x64, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x61, 0x6c, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x21, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x61,
	0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x61, 0x6c, 0x74, 0x22, 0xf0, 0x06, 0x0a, 0x0b, 0x4e, 0x69,
	0x67, 0x6f, 0x72, 0x69, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x49, 0x0a, 0x12, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x11, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x19, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x4b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x39, 0x0a, 0x0c, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x70,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x59, 0x0a, 0x0f, 0x70, 0x61,
	0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x69,
	0x67, 0x6f, 0x72, 0x69, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x50, 0x61,
	0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x52, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x16, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61,
	0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x82, 0x01, 0x0a, 0x27, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x64, 0x65,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65,
	0x4b, 0x65, 0x79, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x52, 0x23, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x70,
	0x68, 0x72, 0x61, 0x73, 0x65, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x45, 0x76, 0x65,
	0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x52, 0x0a, 0x26, 0x65, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x09, 0x20, 0x03, 0x28, 0x05, 0x52, 0x22, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6b,
	0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x5f,
	0x0a, 0x20, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x62, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x1d, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x4b, 0x0a, 0x23, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1e, 0x6c, 0x61,
	0x73, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x72, 0x75, 0x73, 0x74, 0x65, 0x64,
	0x56, 0x61, 0x75, 0x6c, 0x74, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xcf, 0x01, 0x0a,
	0x0f, 0x4e, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x41, 0x0a, 0x10, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x0e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x0c, 0x6e, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x4e, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x0b, 0x6e, 0x69, 0x67, 0x6f, 0x72, 0x69, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x2b,
	0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_nigori_local_data_proto_rawDescOnce sync.Once
	file_nigori_local_data_proto_rawDescData = file_nigori_local_data_proto_rawDesc
)

func file_nigori_local_data_proto_rawDescGZIP() []byte {
	file_nigori_local_data_proto_rawDescOnce.Do(func() {
		file_nigori_local_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_nigori_local_data_proto_rawDescData)
	})
	return file_nigori_local_data_proto_rawDescData
}

var file_nigori_local_data_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nigori_local_data_proto_goTypes = []interface{}{
	(*CryptographerData)(nil),                   // 0: sync_pb.CryptographerData
	(*CustomPassphraseKeyDerivationParams)(nil), // 1: sync_pb.CustomPassphraseKeyDerivationParams
	(*NigoriModel)(nil),                         // 2: sync_pb.NigoriModel
	(*NigoriLocalData)(nil),                     // 3: sync_pb.NigoriLocalData
	(*NigoriKeyBag)(nil),                        // 4: sync_pb.NigoriKeyBag
	(NigoriSpecifics_KeyDerivationMethod)(0),    // 5: sync_pb.NigoriSpecifics.KeyDerivationMethod
	(*EncryptedData)(nil),                       // 6: sync_pb.EncryptedData
	(NigoriSpecifics_PassphraseType)(0),         // 7: sync_pb.NigoriSpecifics.PassphraseType
	(*ModelTypeState)(nil),                      // 8: sync_pb.ModelTypeState
	(*EntityMetadata)(nil),                      // 9: sync_pb.EntityMetadata
}
var file_nigori_local_data_proto_depIdxs = []int32{
	4,  // 0: sync_pb.CryptographerData.key_bag:type_name -> sync_pb.NigoriKeyBag
	5,  // 1: sync_pb.CustomPassphraseKeyDerivationParams.custom_passphrase_key_derivation_method:type_name -> sync_pb.NigoriSpecifics.KeyDerivationMethod
	0,  // 2: sync_pb.NigoriModel.cryptographer_data:type_name -> sync_pb.CryptographerData
	6,  // 3: sync_pb.NigoriModel.pending_keys:type_name -> sync_pb.EncryptedData
	7,  // 4: sync_pb.NigoriModel.passphrase_type:type_name -> sync_pb.NigoriSpecifics.PassphraseType
	1,  // 5: sync_pb.NigoriModel.custom_passphrase_key_derivation_params:type_name -> sync_pb.CustomPassphraseKeyDerivationParams
	6,  // 6: sync_pb.NigoriModel.pending_keystore_decryptor_token:type_name -> sync_pb.EncryptedData
	8,  // 7: sync_pb.NigoriLocalData.model_type_state:type_name -> sync_pb.ModelTypeState
	9,  // 8: sync_pb.NigoriLocalData.entity_metadata:type_name -> sync_pb.EntityMetadata
	2,  // 9: sync_pb.NigoriLocalData.nigori_model:type_name -> sync_pb.NigoriModel
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_nigori_local_data_proto_init() }
func file_nigori_local_data_proto_init() {
	if File_nigori_local_data_proto != nil {
		return
	}
	file_encryption_proto_init()
	file_entity_metadata_proto_init()
	file_model_type_state_proto_init()
	file_nigori_specifics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_nigori_local_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptographerData); i {
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
		file_nigori_local_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomPassphraseKeyDerivationParams); i {
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
		file_nigori_local_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NigoriModel); i {
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
		file_nigori_local_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NigoriLocalData); i {
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
			RawDescriptor: file_nigori_local_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nigori_local_data_proto_goTypes,
		DependencyIndexes: file_nigori_local_data_proto_depIdxs,
		MessageInfos:      file_nigori_local_data_proto_msgTypes,
	}.Build()
	File_nigori_local_data_proto = out.File
	file_nigori_local_data_proto_rawDesc = nil
	file_nigori_local_data_proto_goTypes = nil
	file_nigori_local_data_proto_depIdxs = nil
}
