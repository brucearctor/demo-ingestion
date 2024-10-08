// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/demo.proto

package _go

import (
	_ "github.com/GoogleCloudPlatform/protoc-gen-bq-schema/protos"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PostFlightStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightId         string  `protobuf:"bytes,1,opt,name=flight_id,json=flightId,proto3" json:"flight_id,omitempty"`
	ScheduledTakeoff int64   `protobuf:"varint,2,opt,name=scheduled_takeoff,json=scheduledTakeoff,proto3" json:"scheduled_takeoff,omitempty"` // maybe timestamp type
	ActualTakeoff    int64   `protobuf:"varint,3,opt,name=actual_takeoff,json=actualTakeoff,proto3" json:"actual_takeoff,omitempty"`
	CurrentTimestamp int64   `protobuf:"varint,4,opt,name=current_timestamp,json=currentTimestamp,proto3" json:"current_timestamp,omitempty"`
	Altitude         int64   `protobuf:"varint,5,opt,name=altitude,proto3" json:"altitude,omitempty"`
	InAir            bool    `protobuf:"varint,6,opt,name=in_air,json=inAir,proto3" json:"in_air,omitempty"` // TBD whether want this included
	Landed           bool    `protobuf:"varint,7,opt,name=landed,proto3" json:"landed,omitempty"`
	Latitude         float64 `protobuf:"fixed64,8,opt,name=latitude,proto3" json:"latitude,omitempty"`   // maybe lat/long type?
	Longitude        float64 `protobuf:"fixed64,9,opt,name=longitude,proto3" json:"longitude,omitempty"` // same?
	Direction        int32   `protobuf:"varint,10,opt,name=direction,proto3" json:"direction,omitempty"` // ex: 1-360 for degrees.  Else enum
	Airline          string  `protobuf:"bytes,11,opt,name=airline,proto3" json:"airline,omitempty"`
	FlightNumber     int32   `protobuf:"varint,12,opt,name=flight_number,json=flightNumber,proto3" json:"flight_number,omitempty"`
	DepartingAirport string  `protobuf:"bytes,13,opt,name=departing_airport,json=departingAirport,proto3" json:"departing_airport,omitempty"` // ENUM?
	ArrivingAirport  string  `protobuf:"bytes,14,opt,name=arriving_airport,json=arrivingAirport,proto3" json:"arriving_airport,omitempty"`    // ENUM?
}

func (x *PostFlightStatusRequest) Reset() {
	*x = PostFlightStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostFlightStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostFlightStatusRequest) ProtoMessage() {}

func (x *PostFlightStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostFlightStatusRequest.ProtoReflect.Descriptor instead.
func (*PostFlightStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_demo_proto_rawDescGZIP(), []int{0}
}

func (x *PostFlightStatusRequest) GetFlightId() string {
	if x != nil {
		return x.FlightId
	}
	return ""
}

func (x *PostFlightStatusRequest) GetScheduledTakeoff() int64 {
	if x != nil {
		return x.ScheduledTakeoff
	}
	return 0
}

func (x *PostFlightStatusRequest) GetActualTakeoff() int64 {
	if x != nil {
		return x.ActualTakeoff
	}
	return 0
}

func (x *PostFlightStatusRequest) GetCurrentTimestamp() int64 {
	if x != nil {
		return x.CurrentTimestamp
	}
	return 0
}

func (x *PostFlightStatusRequest) GetAltitude() int64 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

func (x *PostFlightStatusRequest) GetInAir() bool {
	if x != nil {
		return x.InAir
	}
	return false
}

func (x *PostFlightStatusRequest) GetLanded() bool {
	if x != nil {
		return x.Landed
	}
	return false
}

func (x *PostFlightStatusRequest) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *PostFlightStatusRequest) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *PostFlightStatusRequest) GetDirection() int32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *PostFlightStatusRequest) GetAirline() string {
	if x != nil {
		return x.Airline
	}
	return ""
}

func (x *PostFlightStatusRequest) GetFlightNumber() int32 {
	if x != nil {
		return x.FlightNumber
	}
	return 0
}

func (x *PostFlightStatusRequest) GetDepartingAirport() string {
	if x != nil {
		return x.DepartingAirport
	}
	return ""
}

func (x *PostFlightStatusRequest) GetArrivingAirport() string {
	if x != nil {
		return x.ArrivingAirport
	}
	return ""
}

type PostFlightStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *PostFlightStatusResponse) Reset() {
	*x = PostFlightStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostFlightStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostFlightStatusResponse) ProtoMessage() {}

func (x *PostFlightStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostFlightStatusResponse.ProtoReflect.Descriptor instead.
func (*PostFlightStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_demo_proto_rawDescGZIP(), []int{1}
}

func (x *PostFlightStatusResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_proto_demo_proto protoreflect.FileDescriptor

var file_proto_demo_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x62, 0x71, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x04,
	0x0a, 0x17, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6c,
	0x69, 0x67, 0x68, 0x74, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6f, 0x66, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x61, 0x6b, 0x65,
	0x6f, 0x66, 0x66, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x74, 0x61,
	0x6b, 0x65, 0x6f, 0x66, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61, 0x63, 0x74,
	0x75, 0x61, 0x6c, 0x54, 0x61, 0x6b, 0x65, 0x6f, 0x66, 0x66, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e, 0x5f, 0x61, 0x69, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x6e, 0x41, 0x69, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61,
	0x6e, 0x64, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6c, 0x61, 0x6e, 0x64,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x69,
	0x72, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x69, 0x72,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x66, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x41,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x72, 0x72, 0x69, 0x76, 0x69,
	0x6e, 0x67, 0x5f, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x61, 0x72, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72,
	0x74, 0x3a, 0x12, 0xea, 0x3f, 0x0f, 0x0a, 0x0d, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x36, 0x0a, 0x18, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8e, 0x01,
	0x0a, 0x0e, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x7c, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x69, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x65, 0x6d,
	0x6f, 0x5f, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0xa4,
	0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x69, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0x42, 0x09, 0x44, 0x65, 0x6d, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3e, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x62, 0x72, 0x75, 0x63, 0x65, 0x61, 0x72, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x64, 0x65,
	0x6d, 0x6f, 0x2d, 0x69, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x6f, 0xa2,
	0x02, 0x03, 0x44, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x44, 0x65, 0x6d, 0x6f, 0x49, 0x6e, 0x67, 0x65,
	0x73, 0x74, 0xca, 0x02, 0x0a, 0x44, 0x65, 0x6d, 0x6f, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0xe2,
	0x02, 0x16, 0x44, 0x65, 0x6d, 0x6f, 0x49, 0x6e, 0x67, 0x65, 0x73, 0x74, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x44, 0x65, 0x6d, 0x6f, 0x49,
	0x6e, 0x67, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_demo_proto_rawDescOnce sync.Once
	file_proto_demo_proto_rawDescData = file_proto_demo_proto_rawDesc
)

func file_proto_demo_proto_rawDescGZIP() []byte {
	file_proto_demo_proto_rawDescOnce.Do(func() {
		file_proto_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_demo_proto_rawDescData)
	})
	return file_proto_demo_proto_rawDescData
}

var file_proto_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_demo_proto_goTypes = []interface{}{
	(*PostFlightStatusRequest)(nil),  // 0: demo_ingest.PostFlightStatusRequest
	(*PostFlightStatusResponse)(nil), // 1: demo_ingest.PostFlightStatusResponse
}
var file_proto_demo_proto_depIdxs = []int32{
	0, // 0: demo_ingest.FlightsService.PostFlightStatus:input_type -> demo_ingest.PostFlightStatusRequest
	1, // 1: demo_ingest.FlightsService.PostFlightStatus:output_type -> demo_ingest.PostFlightStatusResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_demo_proto_init() }
func file_proto_demo_proto_init() {
	if File_proto_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_demo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostFlightStatusRequest); i {
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
		file_proto_demo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostFlightStatusResponse); i {
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
			RawDescriptor: file_proto_demo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_demo_proto_goTypes,
		DependencyIndexes: file_proto_demo_proto_depIdxs,
		MessageInfos:      file_proto_demo_proto_msgTypes,
	}.Build()
	File_proto_demo_proto = out.File
	file_proto_demo_proto_rawDesc = nil
	file_proto_demo_proto_goTypes = nil
	file_proto_demo_proto_depIdxs = nil
}
