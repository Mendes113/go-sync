// Copyright 2021 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: webauthn_credential_specifics.proto

package sync_pb

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

// WebauthnCredentialSpecifics is an entity that backs a WebAuthn
// PublicKeyCredential. Since it contains the authenticator’s view of this
// object, it has a private key rather than a public key.
// (https://www.w3.org/TR/webauthn-2/#iface-pkcredential).
//
// Names of fields are taken from WebAuthn where possible. E.g.
// user.displayName in WebAuthn becomes user_display_name here.
//
// All fields are immutable after creation except for user_name and
// user_display_name, which may be updated by a user.
type WebauthnCredentialSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sync's ID for this entity (sometimes called the client unique tag), 16
	// random bytes. This value is used within Sync to identify this entity. The
	// credential ID is not used because the (hashed) sync_id is exposed to the
	// Sync server, and we don’t want Google to be able to map a credential ID to
	// an account. Password entities construct this value from the concatenation
	// of many fields and depend on the fact that the server only sees a hash of
	// it. But the only high-entropy secret here is the private key, which will
	// have different encryption in the future, and private keys are not the sort
	// of data to copy into other fields. Therefore this independent value is
	// provided to form the client's ID.
	SyncId []byte `protobuf:"bytes,1,opt,name=sync_id,json=syncId" json:"sync_id,omitempty"`
	// The credential ID, 16 random bytes. This is a value surfaced in
	// the WebAuthn API (https://www.w3.org/TR/webauthn-2/#credential-id).
	CredentialId []byte `protobuf:"bytes,2,opt,name=credential_id,json=credentialId" json:"credential_id,omitempty"`
	// An RP ID is a WebAuthn concept:
	// https://www.w3.org/TR/webauthn-2/#rp-id. It’s usually a domain name,
	// although in non-Web contexts it can be a URL with a non-Web scheme.
	RpId *string `protobuf:"bytes,3,opt,name=rp_id,json=rpId" json:"rp_id,omitempty"`
	// The user ID, which is also called a “user handle” in WebAuthn
	// (https://www.w3.org/TR/webauthn-2/#user-handle), is an RP-specific
	// identifier that is up to 64-bytes long. An authenticator conceptually only
	// stores a single credential for a given (rp_id, user_id) pair, but there
	// may be several credentials in Sync. They are prioritised using
	// newly_shadowed_credential_ids and creation_time. See below.
	//
	// (We wish to be able to retain several entities for a single (rp_id,
	// user_id) pair because there’s an edge case where we may wish to revert to
	// an older entry and thus need to keep the older entry around in Sync. The
	// revert could happen on a different device too.)
	UserId []byte `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	// The id of credentials with the same (rp_id, user_id) that were
	// shadowed by the creation of this entity.
	//
	// A credential is shadowed if one or more other credentials (from the same
	// account, and with the same (rp_id, user_id)) include its credential_id in
	// their list of shadowed IDs. Shadowed credentials are ignored when finding
	// a credential to sign with. If there is more than one candidate remaining
	// after filtering shadowed credentials then the most recently created (based
	// on creation_time) is used.
	//
	// The reason for all this is that sites can replace a credential by creating
	// another one with the same (rp_id, user_id) pair. However, we don't
	// immediately know whether the WebAuthn response reached the website's
	// server. Consider a user with a poor internet connection. Javascript in the
	// site's origin triggers a credential creation that “overwrites” an existing
	// credential, but the Javascript is unable to send the new public key to the
	// website's server. The user is now locked out: the old credential has been
	// over-written but the website's server doesn't know about the new one.
	//
	// Thus we wish to keep “overwritten” credentials around for a while to allow
	// for some sort of recovery. In the simple case, a new credential shadows
	// the single, previous old credential. We could depend on creation_time, but
	// client clocks aren't always accurate, thus this field.
	//
	// In complicated cases two devices might race to replace a credential, in
	// which case (after mutual syncing) two candidate credentials exist for the
	// same (rp_id, user_id) pair because neither shadows the other. In this case
	// we pick the newest based on |creation_time| but it's quite possible that
	// some recovery will be needed because the website's server thinks the other
	// one is correct.
	//
	// A generation counter isn't used because a single device might replace a
	// series of credentials as it tries to update the website's server. But that
	// doesn't mean that it should dominate a different device that replaced it
	// only once, but later.
	NewlyShadowedCredentialIds [][]byte `protobuf:"bytes,5,rep,name=newly_shadowed_credential_ids,json=newlyShadowedCredentialIds" json:"newly_shadowed_credential_ids,omitempty"`
	// The local time on the device when this credential was created. Given in
	// milliseconds since the UNIX epoch. This is used to break ties between
	// credentials. See newly_shadowed_credential_ids.
	CreationTime *int64 `protobuf:"varint,6,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	// The human-readable account identifier. Usually an email address. This is
	// mutable.
	// https://www.w3.org/TR/webauthn-2/#dom-publickeycredentialentity-name
	UserName *string `protobuf:"bytes,7,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	// The human-readable name. Usually a legal name. This is mutable.
	// https://www.w3.org/TR/webauthn-2/#dom-publickeycredentialuserentity-displayname.
	UserDisplayName *string `protobuf:"bytes,8,opt,name=user_display_name,json=userDisplayName" json:"user_display_name,omitempty"`
	// Credentials may optionally be enabled for Secure Payment Confirmation[1] on
	// third-party sites. This is opt-in at creation time.
	//
	// [1] https://www.w3.org/TR/secure-payment-confirmation/
	ThirdPartyPaymentsSupport *bool `protobuf:"varint,11,opt,name=third_party_payments_support,json=thirdPartyPaymentsSupport" json:"third_party_payments_support,omitempty"`
	// Time when this passkey was last successfully asserted. Number of
	// microseconds since 1601, aka Windows epoch. This mirrors the
	// `date_last_used` field in PasswordSpecificsData.
	LastUsedTimeWindowsEpochMicros *int64 `protobuf:"varint,13,opt,name=last_used_time_windows_epoch_micros,json=lastUsedTimeWindowsEpochMicros" json:"last_used_time_windows_epoch_micros,omitempty"`
	// The "version" (sometimes called the epoch) of the key that `encrypted_data`
	// is encrypted with. This allows trial decryption to be avoided when present.
	KeyVersion *int32 `protobuf:"varint,14,opt,name=key_version,json=keyVersion" json:"key_version,omitempty"`
	// Types that are assignable to EncryptedData:
	//	*WebauthnCredentialSpecifics_PrivateKey
	//	*WebauthnCredentialSpecifics_Encrypted_
	EncryptedData isWebauthnCredentialSpecifics_EncryptedData `protobuf_oneof:"encrypted_data"`
}

func (x *WebauthnCredentialSpecifics) Reset() {
	*x = WebauthnCredentialSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webauthn_credential_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebauthnCredentialSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebauthnCredentialSpecifics) ProtoMessage() {}

func (x *WebauthnCredentialSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_webauthn_credential_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebauthnCredentialSpecifics.ProtoReflect.Descriptor instead.
func (*WebauthnCredentialSpecifics) Descriptor() ([]byte, []int) {
	return file_webauthn_credential_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *WebauthnCredentialSpecifics) GetSyncId() []byte {
	if x != nil {
		return x.SyncId
	}
	return nil
}

func (x *WebauthnCredentialSpecifics) GetCredentialId() []byte {
	if x != nil {
		return x.CredentialId
	}
	return nil
}

func (x *WebauthnCredentialSpecifics) GetRpId() string {
	if x != nil && x.RpId != nil {
		return *x.RpId
	}
	return ""
}

func (x *WebauthnCredentialSpecifics) GetUserId() []byte {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *WebauthnCredentialSpecifics) GetNewlyShadowedCredentialIds() [][]byte {
	if x != nil {
		return x.NewlyShadowedCredentialIds
	}
	return nil
}

func (x *WebauthnCredentialSpecifics) GetCreationTime() int64 {
	if x != nil && x.CreationTime != nil {
		return *x.CreationTime
	}
	return 0
}

func (x *WebauthnCredentialSpecifics) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

func (x *WebauthnCredentialSpecifics) GetUserDisplayName() string {
	if x != nil && x.UserDisplayName != nil {
		return *x.UserDisplayName
	}
	return ""
}

func (x *WebauthnCredentialSpecifics) GetThirdPartyPaymentsSupport() bool {
	if x != nil && x.ThirdPartyPaymentsSupport != nil {
		return *x.ThirdPartyPaymentsSupport
	}
	return false
}

func (x *WebauthnCredentialSpecifics) GetLastUsedTimeWindowsEpochMicros() int64 {
	if x != nil && x.LastUsedTimeWindowsEpochMicros != nil {
		return *x.LastUsedTimeWindowsEpochMicros
	}
	return 0
}

func (x *WebauthnCredentialSpecifics) GetKeyVersion() int32 {
	if x != nil && x.KeyVersion != nil {
		return *x.KeyVersion
	}
	return 0
}

func (m *WebauthnCredentialSpecifics) GetEncryptedData() isWebauthnCredentialSpecifics_EncryptedData {
	if m != nil {
		return m.EncryptedData
	}
	return nil
}

func (x *WebauthnCredentialSpecifics) GetPrivateKey() []byte {
	if x, ok := x.GetEncryptedData().(*WebauthnCredentialSpecifics_PrivateKey); ok {
		return x.PrivateKey
	}
	return nil
}

func (x *WebauthnCredentialSpecifics) GetEncrypted() []byte {
	if x, ok := x.GetEncryptedData().(*WebauthnCredentialSpecifics_Encrypted_); ok {
		return x.Encrypted
	}
	return nil
}

type isWebauthnCredentialSpecifics_EncryptedData interface {
	isWebauthnCredentialSpecifics_EncryptedData()
}

type WebauthnCredentialSpecifics_PrivateKey struct {
	// The bytes of the private key, encrypted with a feature-specific security
	// domain.
	PrivateKey []byte `protobuf:"bytes,9,opt,name=private_key,json=privateKey,oneof"`
}

type WebauthnCredentialSpecifics_Encrypted_ struct {
	// A serialized, `Encrypted` message, encrypted with a feature-specific
	// security domain.
	Encrypted []byte `protobuf:"bytes,12,opt,name=encrypted,oneof"`
}

func (*WebauthnCredentialSpecifics_PrivateKey) isWebauthnCredentialSpecifics_EncryptedData() {}

func (*WebauthnCredentialSpecifics_Encrypted_) isWebauthnCredentialSpecifics_EncryptedData() {}

type WebauthnCredentialSpecifics_Encrypted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bytes of the private key, in PKCS#8 format.
	PrivateKey []byte `protobuf:"bytes,1,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	// The secret for implementing the hmac-secret extension
	// https://fidoalliance.org/specs/fido-v2.1-ps-20210615/fido-client-to-authenticator-protocol-v2.1-ps-errata-20220621.html#sctn-hmac-secret-extension
	HmacSecret []byte `protobuf:"bytes,2,opt,name=hmac_secret,json=hmacSecret" json:"hmac_secret,omitempty"`
	// The contents of the credential's credBlob.
	// https://fidoalliance.org/specs/fido-v2.1-ps-20210615/fido-client-to-authenticator-protocol-v2.1-ps-errata-20220621.html#sctn-credBlob-extension
	CredBlob []byte `protobuf:"bytes,3,opt,name=cred_blob,json=credBlob" json:"cred_blob,omitempty"`
	// The contents of the credential's largeBlob value(*). Unlike with security
	// keys, largeBlob data is not stored in a single lump for all credentials,
	// but as per-credential data. This data is presented to the authenticator
	// over CTAP and thus has already had the required DEFLATE compression
	// applied by the remote platform. The uncompressed size of this data is in
	// the next field.
	//
	// (*) "large" with respect to embedded devices. Maximum length is 2KiB for
	// Google Password Manager.
	//
	// https://fidoalliance.org/specs/fido-v2.1-ps-20210615/fido-client-to-authenticator-protocol-v2.1-ps-errata-20220621.html#authenticatorLargeBlobs
	LargeBlob []byte `protobuf:"bytes,4,opt,name=large_blob,json=largeBlob" json:"large_blob,omitempty"`
	// The claimed uncompressed size of the DEFLATE-compressed data in
	// `large_blob`. This corresponds to the `origSize` field from the spec:
	// https://fidoalliance.org/specs/fido-v2.1-ps-20210615/fido-client-to-authenticator-protocol-v2.1-ps-errata-20220621.html#large-blob
	LargeBlobUncompressedSize *uint64 `protobuf:"varint,5,opt,name=large_blob_uncompressed_size,json=largeBlobUncompressedSize" json:"large_blob_uncompressed_size,omitempty"`
}

func (x *WebauthnCredentialSpecifics_Encrypted) Reset() {
	*x = WebauthnCredentialSpecifics_Encrypted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webauthn_credential_specifics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebauthnCredentialSpecifics_Encrypted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebauthnCredentialSpecifics_Encrypted) ProtoMessage() {}

func (x *WebauthnCredentialSpecifics_Encrypted) ProtoReflect() protoreflect.Message {
	mi := &file_webauthn_credential_specifics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebauthnCredentialSpecifics_Encrypted.ProtoReflect.Descriptor instead.
func (*WebauthnCredentialSpecifics_Encrypted) Descriptor() ([]byte, []int) {
	return file_webauthn_credential_specifics_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WebauthnCredentialSpecifics_Encrypted) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *WebauthnCredentialSpecifics_Encrypted) GetHmacSecret() []byte {
	if x != nil {
		return x.HmacSecret
	}
	return nil
}

func (x *WebauthnCredentialSpecifics_Encrypted) GetCredBlob() []byte {
	if x != nil {
		return x.CredBlob
	}
	return nil
}

func (x *WebauthnCredentialSpecifics_Encrypted) GetLargeBlob() []byte {
	if x != nil {
		return x.LargeBlob
	}
	return nil
}

func (x *WebauthnCredentialSpecifics_Encrypted) GetLargeBlobUncompressedSize() uint64 {
	if x != nil && x.LargeBlobUncompressedSize != nil {
		return *x.LargeBlobUncompressedSize
	}
	return 0
}

var File_webauthn_credential_specifics_proto protoreflect.FileDescriptor

var file_webauthn_credential_specifics_proto_rawDesc = []byte{
	0x0a, 0x23, 0x77, 0x65, 0x62, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0x91,
	0x06, 0x0a, 0x1b, 0x57, 0x65, 0x62, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x73, 0x79, 0x6e, 0x63, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05,
	0x72, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x70, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x1d, 0x6e, 0x65,
	0x77, 0x6c, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x1a, 0x6e, 0x65, 0x77, 0x6c, 0x79, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x65, 0x64,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x1c, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x19, 0x74, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x4b, 0x0a, 0x23,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x77,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x5f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1e, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x65, 0x79,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x6b, 0x65, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0b, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a,
	0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x00, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x1a, 0xca, 0x01,
	0x0a, 0x09, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b,
	0x68, 0x6d, 0x61, 0x63, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x68, 0x6d, 0x61, 0x63, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x63, 0x72, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61,
	0x72, 0x67, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x6c, 0x61, 0x72, 0x67, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x3f, 0x0a, 0x1c, 0x6c, 0x61, 0x72,
	0x67, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x75, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x19, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x55, 0x6e, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x4a, 0x04, 0x08, 0x0a,
	0x10, 0x0b, 0x42, 0x36, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01, 0x5a,
	0x09, 0x2e, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
}

var (
	file_webauthn_credential_specifics_proto_rawDescOnce sync.Once
	file_webauthn_credential_specifics_proto_rawDescData = file_webauthn_credential_specifics_proto_rawDesc
)

func file_webauthn_credential_specifics_proto_rawDescGZIP() []byte {
	file_webauthn_credential_specifics_proto_rawDescOnce.Do(func() {
		file_webauthn_credential_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_webauthn_credential_specifics_proto_rawDescData)
	})
	return file_webauthn_credential_specifics_proto_rawDescData
}

var file_webauthn_credential_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_webauthn_credential_specifics_proto_goTypes = []interface{}{
	(*WebauthnCredentialSpecifics)(nil),           // 0: sync_pb.WebauthnCredentialSpecifics
	(*WebauthnCredentialSpecifics_Encrypted)(nil), // 1: sync_pb.WebauthnCredentialSpecifics.Encrypted
}
var file_webauthn_credential_specifics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_webauthn_credential_specifics_proto_init() }
func file_webauthn_credential_specifics_proto_init() {
	if File_webauthn_credential_specifics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_webauthn_credential_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebauthnCredentialSpecifics); i {
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
		file_webauthn_credential_specifics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebauthnCredentialSpecifics_Encrypted); i {
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
	file_webauthn_credential_specifics_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*WebauthnCredentialSpecifics_PrivateKey)(nil),
		(*WebauthnCredentialSpecifics_Encrypted_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_webauthn_credential_specifics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_webauthn_credential_specifics_proto_goTypes,
		DependencyIndexes: file_webauthn_credential_specifics_proto_depIdxs,
		MessageInfos:      file_webauthn_credential_specifics_proto_msgTypes,
	}.Build()
	File_webauthn_credential_specifics_proto = out.File
	file_webauthn_credential_specifics_proto_rawDesc = nil
	file_webauthn_credential_specifics_proto_goTypes = nil
	file_webauthn_credential_specifics_proto_depIdxs = nil
}
