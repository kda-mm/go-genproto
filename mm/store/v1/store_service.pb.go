// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: mm/store/v1/store_service.proto

package store

import (
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

var File_mm_store_v1_store_service_proto protoreflect.FileDescriptor

var file_mm_store_v1_store_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x6d, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x6d,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xea,
	0x03, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x6d, 0x6d, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x56, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x6d, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x6d, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x47, 0x65, 0x74, 0x12, 0x62, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1b, 0x2e, 0x6d, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x6d, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x5e, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x6d, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x5e, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x6d, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x64, 0x61, 0x2d, 0x6d, 0x6d,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6d, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mm_store_v1_store_service_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),   // 0: mm.store.v1.CreateRequest
	(*GetRequest)(nil),      // 1: mm.store.v1.GetRequest
	(*GetListRequest)(nil),  // 2: mm.store.v1.GetListRequest
	(*UpdateRequest)(nil),   // 3: mm.store.v1.UpdateRequest
	(*DeleteRequest)(nil),   // 4: mm.store.v1.DeleteRequest
	(*CreateResponse)(nil),  // 5: mm.store.v1.CreateResponse
	(*GetListResponse)(nil), // 6: mm.store.v1.GetListResponse
	(*UpdateResponse)(nil),  // 7: mm.store.v1.UpdateResponse
	(*DeleteResponse)(nil),  // 8: mm.store.v1.DeleteResponse
}
var file_mm_store_v1_store_service_proto_depIdxs = []int32{
	0, // 0: mm.store.v1.StoreService.Create:input_type -> mm.store.v1.CreateRequest
	1, // 1: mm.store.v1.StoreService.Get:input_type -> mm.store.v1.GetRequest
	2, // 2: mm.store.v1.StoreService.GetList:input_type -> mm.store.v1.GetListRequest
	3, // 3: mm.store.v1.StoreService.Update:input_type -> mm.store.v1.UpdateRequest
	4, // 4: mm.store.v1.StoreService.Delete:input_type -> mm.store.v1.DeleteRequest
	5, // 5: mm.store.v1.StoreService.Create:output_type -> mm.store.v1.CreateResponse
	6, // 6: mm.store.v1.StoreService.Get:output_type -> mm.store.v1.GetListResponse
	6, // 7: mm.store.v1.StoreService.GetList:output_type -> mm.store.v1.GetListResponse
	7, // 8: mm.store.v1.StoreService.Update:output_type -> mm.store.v1.UpdateResponse
	8, // 9: mm.store.v1.StoreService.Delete:output_type -> mm.store.v1.DeleteResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mm_store_v1_store_service_proto_init() }
func file_mm_store_v1_store_service_proto_init() {
	if File_mm_store_v1_store_service_proto != nil {
		return
	}
	file_mm_store_v1_store_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mm_store_v1_store_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mm_store_v1_store_service_proto_goTypes,
		DependencyIndexes: file_mm_store_v1_store_service_proto_depIdxs,
	}.Build()
	File_mm_store_v1_store_service_proto = out.File
	file_mm_store_v1_store_service_proto_rawDesc = nil
	file_mm_store_v1_store_service_proto_goTypes = nil
	file_mm_store_v1_store_service_proto_depIdxs = nil
}