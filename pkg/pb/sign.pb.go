//
// Copyright (c) 2023 sixwaaaay.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: sign.proto

package pb

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

type JSONRegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Account *User  `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *JSONRegisterReply) Reset() {
	*x = JSONRegisterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sign_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JSONRegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSONRegisterReply) ProtoMessage() {}

func (x *JSONRegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_sign_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSONRegisterReply.ProtoReflect.Descriptor instead.
func (*JSONRegisterReply) Descriptor() ([]byte, []int) {
	return file_sign_proto_rawDescGZIP(), []int{0}
}

func (x *JSONRegisterReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *JSONRegisterReply) GetAccount() *User {
	if x != nil {
		return x.Account
	}
	return nil
}

type JSONLoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Account *User  `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *JSONLoginReply) Reset() {
	*x = JSONLoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sign_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JSONLoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSONLoginReply) ProtoMessage() {}

func (x *JSONLoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_sign_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSONLoginReply.ProtoReflect.Descriptor instead.
func (*JSONLoginReply) Descriptor() ([]byte, []int) {
	return file_sign_proto_rawDescGZIP(), []int{1}
}

func (x *JSONLoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *JSONLoginReply) GetAccount() *User {
	if x != nil {
		return x.Account
	}
	return nil
}

var File_sign_proto protoreflect.FileDescriptor

var file_sign_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x69,
	0x78, 0x77, 0x61, 0x61, 0x61, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x11, 0x4a, 0x53, 0x4f, 0x4e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x69, 0x78, 0x77, 0x61, 0x61, 0x61, 0x61, 0x79,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x56, 0x0a, 0x0e, 0x4a, 0x53, 0x4f, 0x4e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73,
	0x69, 0x78, 0x77, 0x61, 0x61, 0x61, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sign_proto_rawDescOnce sync.Once
	file_sign_proto_rawDescData = file_sign_proto_rawDesc
)

func file_sign_proto_rawDescGZIP() []byte {
	file_sign_proto_rawDescOnce.Do(func() {
		file_sign_proto_rawDescData = protoimpl.X.CompressGZIP(file_sign_proto_rawDescData)
	})
	return file_sign_proto_rawDescData
}

var file_sign_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sign_proto_goTypes = []interface{}{
	(*JSONRegisterReply)(nil), // 0: sixwaaaay.user.JSONRegisterReply
	(*JSONLoginReply)(nil),    // 1: sixwaaaay.user.JSONLoginReply
	(*User)(nil),              // 2: sixwaaaay.user.User
}
var file_sign_proto_depIdxs = []int32{
	2, // 0: sixwaaaay.user.JSONRegisterReply.account:type_name -> sixwaaaay.user.User
	2, // 1: sixwaaaay.user.JSONLoginReply.account:type_name -> sixwaaaay.user.User
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sign_proto_init() }
func file_sign_proto_init() {
	if File_sign_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sign_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JSONRegisterReply); i {
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
		file_sign_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JSONLoginReply); i {
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
			RawDescriptor: file_sign_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sign_proto_goTypes,
		DependencyIndexes: file_sign_proto_depIdxs,
		MessageInfos:      file_sign_proto_msgTypes,
	}.Build()
	File_sign_proto = out.File
	file_sign_proto_rawDesc = nil
	file_sign_proto_goTypes = nil
	file_sign_proto_depIdxs = nil
}
