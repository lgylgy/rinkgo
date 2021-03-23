// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.0.0
// source: pkg/services/pstat/proto/pstat.proto

package pstat

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Players struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*History `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *Players) Reset() {
	*x = Players{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_pstat_proto_pstat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Players) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Players) ProtoMessage() {}

func (x *Players) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_pstat_proto_pstat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Players.ProtoReflect.Descriptor instead.
func (*Players) Descriptor() ([]byte, []int) {
	return file_pkg_services_pstat_proto_pstat_proto_rawDescGZIP(), []int{0}
}

func (x *Players) GetPlayers() []*History {
	if x != nil {
		return x.Players
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID int32 `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_pstat_proto_pstat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_pstat_proto_pstat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_pkg_services_pstat_proto_pstat_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetPlayerID() int32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type History struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID int32    `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Entries  []*Entry `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *History) Reset() {
	*x = History{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_pstat_proto_pstat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *History) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*History) ProtoMessage() {}

func (x *History) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_pstat_proto_pstat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use History.ProtoReflect.Descriptor instead.
func (*History) Descriptor() ([]byte, []int) {
	return file_pkg_services_pstat_proto_pstat_proto_rawDescGZIP(), []int{2}
}

func (x *History) GetPlayerID() int32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *History) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *History) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Season string `protobuf:"bytes,1,opt,name=season,proto3" json:"season,omitempty"`
	Team   string `protobuf:"bytes,2,opt,name=team,proto3" json:"team,omitempty"`
	Event  string `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	Matchs int32  `protobuf:"varint,4,opt,name=matchs,proto3" json:"matchs,omitempty"`
	Goals  int32  `protobuf:"varint,6,opt,name=goals,proto3" json:"goals,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_services_pstat_proto_pstat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_services_pstat_proto_pstat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_pkg_services_pstat_proto_pstat_proto_rawDescGZIP(), []int{3}
}

func (x *Entry) GetSeason() string {
	if x != nil {
		return x.Season
	}
	return ""
}

func (x *Entry) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *Entry) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *Entry) GetMatchs() int32 {
	if x != nil {
		return x.Matchs
	}
	return 0
}

func (x *Entry) GetGoals() int32 {
	if x != nil {
		return x.Goals
	}
	return 0
}

var File_pkg_services_pstat_proto_pstat_proto protoreflect.FileDescriptor

var file_pkg_services_pstat_proto_pstat_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x70,
	0x73, 0x74, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x73, 0x74, 0x61, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x73, 0x74, 0x61, 0x74, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x07, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22,
	0x25, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0x61, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x77, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x6f, 0x61, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x67, 0x6f, 0x61,
	0x6c, 0x73, 0x32, 0x73, 0x0a, 0x0c, 0x50, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x70, 0x73, 0x74, 0x61,
	0x74, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x2c, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x2e, 0x70, 0x73, 0x74, 0x61, 0x74, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x73, 0x74, 0x61, 0x74, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x67, 0x79, 0x6c, 0x67, 0x79, 0x2f, 0x72, 0x69, 0x6e,
	0x6b, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x73, 0x74, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_services_pstat_proto_pstat_proto_rawDescOnce sync.Once
	file_pkg_services_pstat_proto_pstat_proto_rawDescData = file_pkg_services_pstat_proto_pstat_proto_rawDesc
)

func file_pkg_services_pstat_proto_pstat_proto_rawDescGZIP() []byte {
	file_pkg_services_pstat_proto_pstat_proto_rawDescOnce.Do(func() {
		file_pkg_services_pstat_proto_pstat_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_services_pstat_proto_pstat_proto_rawDescData)
	})
	return file_pkg_services_pstat_proto_pstat_proto_rawDescData
}

var file_pkg_services_pstat_proto_pstat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_services_pstat_proto_pstat_proto_goTypes = []interface{}{
	(*Players)(nil),       // 0: pstat.Players
	(*Request)(nil),       // 1: pstat.Request
	(*History)(nil),       // 2: pstat.History
	(*Entry)(nil),         // 3: pstat.Entry
	(*emptypb.Empty)(nil), // 4: google.protobuf.Empty
}
var file_pkg_services_pstat_proto_pstat_proto_depIdxs = []int32{
	2, // 0: pstat.Players.players:type_name -> pstat.History
	3, // 1: pstat.History.entries:type_name -> pstat.Entry
	4, // 2: pstat.PStatService.ListPlayers:input_type -> google.protobuf.Empty
	1, // 3: pstat.PStatService.GetHistory:input_type -> pstat.Request
	0, // 4: pstat.PStatService.ListPlayers:output_type -> pstat.Players
	2, // 5: pstat.PStatService.GetHistory:output_type -> pstat.History
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_services_pstat_proto_pstat_proto_init() }
func file_pkg_services_pstat_proto_pstat_proto_init() {
	if File_pkg_services_pstat_proto_pstat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_services_pstat_proto_pstat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Players); i {
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
		file_pkg_services_pstat_proto_pstat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_pkg_services_pstat_proto_pstat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*History); i {
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
		file_pkg_services_pstat_proto_pstat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
			RawDescriptor: file_pkg_services_pstat_proto_pstat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_services_pstat_proto_pstat_proto_goTypes,
		DependencyIndexes: file_pkg_services_pstat_proto_pstat_proto_depIdxs,
		MessageInfos:      file_pkg_services_pstat_proto_pstat_proto_msgTypes,
	}.Build()
	File_pkg_services_pstat_proto_pstat_proto = out.File
	file_pkg_services_pstat_proto_pstat_proto_rawDesc = nil
	file_pkg_services_pstat_proto_pstat_proto_goTypes = nil
	file_pkg_services_pstat_proto_pstat_proto_depIdxs = nil
}
