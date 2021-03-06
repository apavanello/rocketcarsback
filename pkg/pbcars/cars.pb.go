// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/cars.proto

package pbcars

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Length    float64 `protobuf:"fixed64,2,opt,name=length,proto3" json:"length,omitempty"`
	Width     float64 `protobuf:"fixed64,3,opt,name=width,proto3" json:"width,omitempty"`
	Height    float64 `protobuf:"fixed64,4,opt,name=height,proto3" json:"height,omitempty"`
	SA        float64 `protobuf:"fixed64,5,opt,name=SA,proto3" json:"SA,omitempty"`
	Elevation float64 `protobuf:"fixed64,6,opt,name=elevation,proto3" json:"elevation,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Car) GetLength() float64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Car) GetWidth() float64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Car) GetHeight() float64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Car) GetSA() float64 {
	if x != nil {
		return x.SA
	}
	return 0
}

func (x *Car) GetElevation() float64 {
	if x != nil {
		return x.Elevation
	}
	return 0
}

type CarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *CarResponse) Reset() {
	*x = CarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarResponse) ProtoMessage() {}

func (x *CarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarResponse.ProtoReflect.Descriptor instead.
func (*CarResponse) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{1}
}

func (x *CarResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type CarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CarRequest) Reset() {
	*x = CarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarRequest) ProtoMessage() {}

func (x *CarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarRequest.ProtoReflect.Descriptor instead.
func (*CarRequest) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{2}
}

func (x *CarRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CarsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cars *Car `protobuf:"bytes,1,opt,name=cars,proto3" json:"cars,omitempty"`
}

func (x *CarsListResponse) Reset() {
	*x = CarsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarsListResponse) ProtoMessage() {}

func (x *CarsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarsListResponse.ProtoReflect.Descriptor instead.
func (*CarsListResponse) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{3}
}

func (x *CarsListResponse) GetCars() *Car {
	if x != nil {
		return x.Cars
	}
	return nil
}

type CarListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CarListRequest) Reset() {
	*x = CarListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarListRequest) ProtoMessage() {}

func (x *CarListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarListRequest.ProtoReflect.Descriptor instead.
func (*CarListRequest) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{4}
}

type ServerVersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerVersionRequest) Reset() {
	*x = ServerVersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerVersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerVersionRequest) ProtoMessage() {}

func (x *ServerVersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerVersionRequest.ProtoReflect.Descriptor instead.
func (*ServerVersionRequest) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{5}
}

type ServerVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp string `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Version   string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ServerVersionResponse) Reset() {
	*x = ServerVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cars_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerVersionResponse) ProtoMessage() {}

func (x *ServerVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cars_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerVersionResponse.ProtoReflect.Descriptor instead.
func (*ServerVersionResponse) Descriptor() ([]byte, []int) {
	return file_proto_cars_proto_rawDescGZIP(), []int{6}
}

func (x *ServerVersionResponse) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *ServerVersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_proto_cars_proto protoreflect.FileDescriptor

var file_proto_cars_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x72, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x03, 0x43, 0x61, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x53, 0x41,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x53, 0x41, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6c,
	0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x65,
	0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x52,
	0x03, 0x63, 0x61, 0x72, 0x22, 0x20, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x63, 0x61,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e,
	0x43, 0x61, 0x72, 0x52, 0x04, 0x63, 0x61, 0x72, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x61, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x32, 0xce, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x12, 0x10,
	0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72,
	0x73, 0x2e, 0x43, 0x61, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x72,
	0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x73, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x63,
	0x61, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cars_proto_rawDescOnce sync.Once
	file_proto_cars_proto_rawDescData = file_proto_cars_proto_rawDesc
)

func file_proto_cars_proto_rawDescGZIP() []byte {
	file_proto_cars_proto_rawDescOnce.Do(func() {
		file_proto_cars_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cars_proto_rawDescData)
	})
	return file_proto_cars_proto_rawDescData
}

var file_proto_cars_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_cars_proto_goTypes = []interface{}{
	(*Car)(nil),                   // 0: cars.Car
	(*CarResponse)(nil),           // 1: cars.CarResponse
	(*CarRequest)(nil),            // 2: cars.CarRequest
	(*CarsListResponse)(nil),      // 3: cars.CarsListResponse
	(*CarListRequest)(nil),        // 4: cars.CarListRequest
	(*ServerVersionRequest)(nil),  // 5: cars.ServerVersionRequest
	(*ServerVersionResponse)(nil), // 6: cars.ServerVersionResponse
}
var file_proto_cars_proto_depIdxs = []int32{
	0, // 0: cars.CarResponse.car:type_name -> cars.Car
	0, // 1: cars.CarsListResponse.cars:type_name -> cars.Car
	2, // 2: cars.CarsService.GetCar:input_type -> cars.CarRequest
	4, // 3: cars.CarsService.GetCarsList:input_type -> cars.CarListRequest
	5, // 4: cars.CarsService.GetServerVersion:input_type -> cars.ServerVersionRequest
	1, // 5: cars.CarsService.GetCar:output_type -> cars.CarResponse
	3, // 6: cars.CarsService.GetCarsList:output_type -> cars.CarsListResponse
	6, // 7: cars.CarsService.GetServerVersion:output_type -> cars.ServerVersionResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_cars_proto_init() }
func file_proto_cars_proto_init() {
	if File_proto_cars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_proto_cars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarResponse); i {
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
		file_proto_cars_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarRequest); i {
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
		file_proto_cars_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarsListResponse); i {
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
		file_proto_cars_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarListRequest); i {
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
		file_proto_cars_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerVersionRequest); i {
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
		file_proto_cars_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerVersionResponse); i {
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
			RawDescriptor: file_proto_cars_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cars_proto_goTypes,
		DependencyIndexes: file_proto_cars_proto_depIdxs,
		MessageInfos:      file_proto_cars_proto_msgTypes,
	}.Build()
	File_proto_cars_proto = out.File
	file_proto_cars_proto_rawDesc = nil
	file_proto_cars_proto_goTypes = nil
	file_proto_cars_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CarsServiceClient is the client API for CarsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CarsServiceClient interface {
	GetCar(ctx context.Context, in *CarRequest, opts ...grpc.CallOption) (*CarResponse, error)
	GetCarsList(ctx context.Context, in *CarListRequest, opts ...grpc.CallOption) (CarsService_GetCarsListClient, error)
	GetServerVersion(ctx context.Context, in *ServerVersionRequest, opts ...grpc.CallOption) (*ServerVersionResponse, error)
}

type carsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarsServiceClient(cc grpc.ClientConnInterface) CarsServiceClient {
	return &carsServiceClient{cc}
}

func (c *carsServiceClient) GetCar(ctx context.Context, in *CarRequest, opts ...grpc.CallOption) (*CarResponse, error) {
	out := new(CarResponse)
	err := c.cc.Invoke(ctx, "/cars.CarsService/GetCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carsServiceClient) GetCarsList(ctx context.Context, in *CarListRequest, opts ...grpc.CallOption) (CarsService_GetCarsListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CarsService_serviceDesc.Streams[0], "/cars.CarsService/GetCarsList", opts...)
	if err != nil {
		return nil, err
	}
	x := &carsServiceGetCarsListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CarsService_GetCarsListClient interface {
	Recv() (*CarsListResponse, error)
	grpc.ClientStream
}

type carsServiceGetCarsListClient struct {
	grpc.ClientStream
}

func (x *carsServiceGetCarsListClient) Recv() (*CarsListResponse, error) {
	m := new(CarsListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *carsServiceClient) GetServerVersion(ctx context.Context, in *ServerVersionRequest, opts ...grpc.CallOption) (*ServerVersionResponse, error) {
	out := new(ServerVersionResponse)
	err := c.cc.Invoke(ctx, "/cars.CarsService/GetServerVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarsServiceServer is the server API for CarsService service.
type CarsServiceServer interface {
	GetCar(context.Context, *CarRequest) (*CarResponse, error)
	GetCarsList(*CarListRequest, CarsService_GetCarsListServer) error
	GetServerVersion(context.Context, *ServerVersionRequest) (*ServerVersionResponse, error)
}

// UnimplementedCarsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCarsServiceServer struct {
}

func (*UnimplementedCarsServiceServer) GetCar(context.Context, *CarRequest) (*CarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCar not implemented")
}
func (*UnimplementedCarsServiceServer) GetCarsList(*CarListRequest, CarsService_GetCarsListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCarsList not implemented")
}
func (*UnimplementedCarsServiceServer) GetServerVersion(context.Context, *ServerVersionRequest) (*ServerVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerVersion not implemented")
}

func RegisterCarsServiceServer(s *grpc.Server, srv CarsServiceServer) {
	s.RegisterService(&_CarsService_serviceDesc, srv)
}

func _CarsService_GetCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarsServiceServer).GetCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cars.CarsService/GetCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarsServiceServer).GetCar(ctx, req.(*CarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarsService_GetCarsList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CarListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CarsServiceServer).GetCarsList(m, &carsServiceGetCarsListServer{stream})
}

type CarsService_GetCarsListServer interface {
	Send(*CarsListResponse) error
	grpc.ServerStream
}

type carsServiceGetCarsListServer struct {
	grpc.ServerStream
}

func (x *carsServiceGetCarsListServer) Send(m *CarsListResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CarsService_GetServerVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarsServiceServer).GetServerVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cars.CarsService/GetServerVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarsServiceServer).GetServerVersion(ctx, req.(*ServerVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CarsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cars.CarsService",
	HandlerType: (*CarsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCar",
			Handler:    _CarsService_GetCar_Handler,
		},
		{
			MethodName: "GetServerVersion",
			Handler:    _CarsService_GetServerVersion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCarsList",
			Handler:       _CarsService_GetCarsList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/cars.proto",
}
