// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proto/Auth/service.proto

package auth

import (
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

var File_proto_Auth_service_proto protoreflect.FileDescriptor

var file_proto_Auth_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x41, 0x75, 0x74, 0x68, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xfd, 0x01, 0x0a,
	0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0c,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x46, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x65,
	0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x6c, 0x68, 0x61, 0x6d,
	0x6b, 0x61, 0x77, 0x65, 0x2f, 0x63, 0x72, 0x6f, 0x77, 0x64, 0x66, 0x75, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_Auth_service_proto_goTypes = []any{
	(*RegisterUserRequest)(nil),   // 0: auth.RegisterUserRequest
	(*LoginRequest)(nil),          // 1: auth.LoginRequest
	(*UpdateInfoUserRequest)(nil), // 2: auth.UpdateInfoUserRequest
	(*UploadAvatarRequest)(nil),   // 3: auth.UploadAvatarRequest
	(*User)(nil),                  // 4: auth.User
	(*BooleanResponse)(nil),       // 5: auth.BooleanResponse
}
var file_proto_Auth_service_proto_depIdxs = []int32{
	0, // 0: auth.AuthService.RegisterUser:input_type -> auth.RegisterUserRequest
	1, // 1: auth.AuthService.Login:input_type -> auth.LoginRequest
	2, // 2: auth.AuthService.UpdateUserInfo:input_type -> auth.UpdateInfoUserRequest
	3, // 3: auth.AuthService.UploadAvatar:input_type -> auth.UploadAvatarRequest
	4, // 4: auth.AuthService.RegisterUser:output_type -> auth.User
	4, // 5: auth.AuthService.Login:output_type -> auth.User
	5, // 6: auth.AuthService.UpdateUserInfo:output_type -> auth.BooleanResponse
	5, // 7: auth.AuthService.UploadAvatar:output_type -> auth.BooleanResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_Auth_service_proto_init() }
func file_proto_Auth_service_proto_init() {
	if File_proto_Auth_service_proto != nil {
		return
	}
	file_proto_Auth_type_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_Auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_Auth_service_proto_goTypes,
		DependencyIndexes: file_proto_Auth_service_proto_depIdxs,
	}.Build()
	File_proto_Auth_service_proto = out.File
	file_proto_Auth_service_proto_rawDesc = nil
	file_proto_Auth_service_proto_goTypes = nil
	file_proto_Auth_service_proto_depIdxs = nil
}
