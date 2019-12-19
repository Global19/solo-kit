// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test/mocks/api/v2alpha1/mock_resources.proto

package v2alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//The best mock resource you ever done seen
type MockResource struct {
	Status   core.Status   `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	Metadata core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	// Types that are valid to be assigned to WeStuckItInAOneof:
	//	*MockResource_SomeDumbField
	//	*MockResource_Data
	WeStuckItInAOneof isMockResource_WeStuckItInAOneof `protobuf_oneof:"we_stuck_it_in_a_oneof"`
	// Types that are valid to be assigned to TestOneofFields:
	//	*MockResource_OneofOne
	//	*MockResource_OneofTwo
	TestOneofFields      isMockResource_TestOneofFields `protobuf_oneof:"test_oneof_fields"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MockResource) Reset()         { *m = MockResource{} }
func (m *MockResource) String() string { return proto.CompactTextString(m) }
func (*MockResource) ProtoMessage()    {}
func (*MockResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e9c669ea46dcbbe, []int{0}
}
func (m *MockResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MockResource.Unmarshal(m, b)
}
func (m *MockResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MockResource.Marshal(b, m, deterministic)
}
func (m *MockResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MockResource.Merge(m, src)
}
func (m *MockResource) XXX_Size() int {
	return xxx_messageInfo_MockResource.Size(m)
}
func (m *MockResource) XXX_DiscardUnknown() {
	xxx_messageInfo_MockResource.DiscardUnknown(m)
}

var xxx_messageInfo_MockResource proto.InternalMessageInfo

type isMockResource_WeStuckItInAOneof interface {
	isMockResource_WeStuckItInAOneof()
	Equal(interface{}) bool
}
type isMockResource_TestOneofFields interface {
	isMockResource_TestOneofFields()
	Equal(interface{}) bool
}

type MockResource_SomeDumbField struct {
	SomeDumbField string `protobuf:"bytes,100,opt,name=some_dumb_field,json=someDumbField,proto3,oneof" json:"some_dumb_field,omitempty"`
}
type MockResource_Data struct {
	Data string `protobuf:"bytes,1,opt,name=data,json=data.json,proto3,oneof" json:"data.json"`
}
type MockResource_OneofOne struct {
	OneofOne string `protobuf:"bytes,3,opt,name=oneof_one,json=oneofOne,proto3,oneof" json:"oneof_one,omitempty"`
}
type MockResource_OneofTwo struct {
	OneofTwo bool `protobuf:"varint,2,opt,name=oneof_two,json=oneofTwo,proto3,oneof" json:"oneof_two,omitempty"`
}

func (*MockResource_SomeDumbField) isMockResource_WeStuckItInAOneof() {}
func (*MockResource_Data) isMockResource_WeStuckItInAOneof()          {}
func (*MockResource_OneofOne) isMockResource_TestOneofFields()        {}
func (*MockResource_OneofTwo) isMockResource_TestOneofFields()        {}

func (m *MockResource) GetWeStuckItInAOneof() isMockResource_WeStuckItInAOneof {
	if m != nil {
		return m.WeStuckItInAOneof
	}
	return nil
}
func (m *MockResource) GetTestOneofFields() isMockResource_TestOneofFields {
	if m != nil {
		return m.TestOneofFields
	}
	return nil
}

func (m *MockResource) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *MockResource) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *MockResource) GetSomeDumbField() string {
	if x, ok := m.GetWeStuckItInAOneof().(*MockResource_SomeDumbField); ok {
		return x.SomeDumbField
	}
	return ""
}

func (m *MockResource) GetData() string {
	if x, ok := m.GetWeStuckItInAOneof().(*MockResource_Data); ok {
		return x.Data
	}
	return ""
}

func (m *MockResource) GetOneofOne() string {
	if x, ok := m.GetTestOneofFields().(*MockResource_OneofOne); ok {
		return x.OneofOne
	}
	return ""
}

func (m *MockResource) GetOneofTwo() bool {
	if x, ok := m.GetTestOneofFields().(*MockResource_OneofTwo); ok {
		return x.OneofTwo
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MockResource) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MockResource_SomeDumbField)(nil),
		(*MockResource_Data)(nil),
		(*MockResource_OneofOne)(nil),
		(*MockResource_OneofTwo)(nil),
	}
}

type FrequentlyChangingAnnotationsResource struct {
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	Blah                 string        `protobuf:"bytes,1,opt,name=blah,proto3" json:"blah,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FrequentlyChangingAnnotationsResource) Reset()         { *m = FrequentlyChangingAnnotationsResource{} }
func (m *FrequentlyChangingAnnotationsResource) String() string { return proto.CompactTextString(m) }
func (*FrequentlyChangingAnnotationsResource) ProtoMessage()    {}
func (*FrequentlyChangingAnnotationsResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e9c669ea46dcbbe, []int{1}
}
func (m *FrequentlyChangingAnnotationsResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequentlyChangingAnnotationsResource.Unmarshal(m, b)
}
func (m *FrequentlyChangingAnnotationsResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequentlyChangingAnnotationsResource.Marshal(b, m, deterministic)
}
func (m *FrequentlyChangingAnnotationsResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequentlyChangingAnnotationsResource.Merge(m, src)
}
func (m *FrequentlyChangingAnnotationsResource) XXX_Size() int {
	return xxx_messageInfo_FrequentlyChangingAnnotationsResource.Size(m)
}
func (m *FrequentlyChangingAnnotationsResource) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequentlyChangingAnnotationsResource.DiscardUnknown(m)
}

var xxx_messageInfo_FrequentlyChangingAnnotationsResource proto.InternalMessageInfo

func (m *FrequentlyChangingAnnotationsResource) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *FrequentlyChangingAnnotationsResource) GetBlah() string {
	if m != nil {
		return m.Blah
	}
	return ""
}

func init() {
	proto.RegisterType((*MockResource)(nil), "testing.solo.io.MockResource")
	proto.RegisterType((*FrequentlyChangingAnnotationsResource)(nil), "testing.solo.io.FrequentlyChangingAnnotationsResource")
}

func init() {
	proto.RegisterFile("test/mocks/api/v2alpha1/mock_resources.proto", fileDescriptor_5e9c669ea46dcbbe)
}

var fileDescriptor_5e9c669ea46dcbbe = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0xad, 0xa7, 0x61, 0x1e, 0x86, 0xaa, 0xc2, 0xad, 0xaa, 0xa8, 0x40, 0x3b, 0xaa, 0x84, 0x94,
	0x05, 0x4d, 0xd4, 0x22, 0xa4, 0xaa, 0x3b, 0x0c, 0xaa, 0x66, 0x53, 0x21, 0x05, 0x56, 0x6c, 0x2c,
	0x27, 0xe3, 0xc9, 0x98, 0x24, 0xbe, 0x43, 0xec, 0x30, 0x65, 0x3b, 0x9f, 0xc0, 0x57, 0xf0, 0x29,
	0x7c, 0x03, 0x8b, 0x2e, 0x58, 0xb2, 0x9b, 0xdd, 0x2c, 0x51, 0xec, 0xcc, 0x94, 0xc7, 0x8e, 0xdd,
	0xb9, 0xf7, 0x9c, 0x63, 0x5f, 0x5f, 0x1f, 0xfc, 0xcc, 0x08, 0x6d, 0xa2, 0x12, 0xd2, 0x5c, 0x47,
	0x7c, 0x26, 0xa3, 0x4f, 0xe7, 0xbc, 0x98, 0x4d, 0xf9, 0x99, 0x6d, 0xb1, 0x4a, 0x68, 0xa8, 0xab,
	0x54, 0xe8, 0x70, 0x56, 0x81, 0x01, 0xb2, 0xdb, 0xa8, 0xa5, 0xca, 0x42, 0x0d, 0x05, 0x84, 0x12,
	0x0e, 0x8f, 0x1a, 0x70, 0x9a, 0x4b, 0xe3, 0xcc, 0x67, 0x51, 0x29, 0x0c, 0x1f, 0x73, 0xc3, 0x9d,
	0xe1, 0xf0, 0xf1, 0xdf, 0xbc, 0x36, 0xdc, 0xd4, 0xed, 0x71, 0xff, 0xba, 0xd7, 0x75, 0xcb, 0xef,
	0x67, 0x90, 0x81, 0x85, 0x51, 0x83, 0xda, 0x2e, 0x11, 0x37, 0xc6, 0x35, 0xc5, 0x4d, 0xab, 0x3c,
	0xf9, 0xde, 0xc1, 0x0f, 0xae, 0x21, 0xcd, 0xe3, 0x76, 0x60, 0xf2, 0x02, 0x77, 0xdd, 0x55, 0x7e,
	0x77, 0x88, 0x82, 0xfb, 0xe7, 0xfb, 0x61, 0x0a, 0x95, 0x58, 0xcf, 0x1d, 0xbe, 0xb5, 0x1c, 0xed,
	0xad, 0x28, 0xfa, 0x76, 0x7b, 0xbc, 0x15, 0xb7, 0x62, 0x72, 0x81, 0xfb, 0xeb, 0x17, 0xf8, 0x3d,
	0x6b, 0x3c, 0xf8, 0xd3, 0x78, 0xdd, 0xb2, 0xd4, 0xb3, 0xbe, 0x8d, 0x9a, 0x9c, 0xe2, 0x5d, 0x0d,
	0xa5, 0x60, 0xe3, 0xba, 0x4c, 0xd8, 0x44, 0x8a, 0x62, 0xec, 0x8f, 0x87, 0x28, 0x18, 0xd0, 0xed,
	0x15, 0x45, 0xa3, 0xad, 0x78, 0xa7, 0x61, 0x5f, 0xd7, 0x65, 0x72, 0xd5, 0x70, 0x24, 0xc0, 0x9e,
	0xbd, 0x04, 0x59, 0xcd, 0xce, 0xcf, 0xdb, 0xe3, 0x81, 0x5d, 0xdb, 0x07, 0x0d, 0x6a, 0xb4, 0x15,
	0xdf, 0x15, 0xe4, 0x09, 0x1e, 0x80, 0x12, 0x30, 0x61, 0xa0, 0x84, 0xbf, 0xdd, 0xc8, 0x47, 0x28,
	0xee, 0xdb, 0xd6, 0x1b, 0x25, 0xee, 0x68, 0x33, 0x07, 0xbf, 0x33, 0x44, 0x41, 0x7f, 0x43, 0xbf,
	0x9b, 0xc3, 0xe5, 0xde, 0x62, 0xe9, 0x79, 0xb8, 0x53, 0xe6, 0x8b, 0xa5, 0xd7, 0x23, 0xf7, 0xec,
	0x47, 0x53, 0x1f, 0x1f, 0xcc, 0x05, 0xd3, 0xa6, 0x4e, 0x73, 0x26, 0x0d, 0x93, 0x8a, 0x71, 0x66,
	0x1d, 0x74, 0x0f, 0x3f, 0x6c, 0xbe, 0xd8, 0x55, 0xee, 0x19, 0xfa, 0xe4, 0x0b, 0xc2, 0x4f, 0xaf,
	0x2a, 0xf1, 0xb1, 0x16, 0xca, 0x14, 0x9f, 0x5f, 0x4d, 0xb9, 0xca, 0xa4, 0xca, 0x5e, 0x2a, 0x05,
	0x86, 0x1b, 0x09, 0x4a, 0x6f, 0xb6, 0xfe, 0xff, 0xeb, 0x23, 0xd8, 0x4b, 0x0a, 0x3e, 0x75, 0xfb,
	0x88, 0x2d, 0xbe, 0x7c, 0xb4, 0x58, 0x7a, 0x5d, 0xec, 0x4d, 0x52, 0x5e, 0xb9, 0xe9, 0x1b, 0xa4,
	0x17, 0x4b, 0xaf, 0x13, 0x20, 0x7a, 0xb1, 0xa2, 0xe8, 0xeb, 0x8f, 0x23, 0xf4, 0x3e, 0xca, 0xa4,
	0x99, 0xd6, 0x49, 0x98, 0x42, 0xe9, 0xf2, 0x23, 0x61, 0x93, 0xa3, 0xe8, 0xb7, 0x74, 0xaf, 0x93,
	0x9d, 0x74, 0x6d, 0x64, 0x9e, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xea, 0x3a, 0x52, 0xfb,
	0x02, 0x00, 0x00,
}

func (this *MockResource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource)
	if !ok {
		that2, ok := that.(MockResource)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if that1.WeStuckItInAOneof == nil {
		if this.WeStuckItInAOneof != nil {
			return false
		}
	} else if this.WeStuckItInAOneof == nil {
		return false
	} else if !this.WeStuckItInAOneof.Equal(that1.WeStuckItInAOneof) {
		return false
	}
	if that1.TestOneofFields == nil {
		if this.TestOneofFields != nil {
			return false
		}
	} else if this.TestOneofFields == nil {
		return false
	} else if !this.TestOneofFields.Equal(that1.TestOneofFields) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MockResource_SomeDumbField) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource_SomeDumbField)
	if !ok {
		that2, ok := that.(MockResource_SomeDumbField)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SomeDumbField != that1.SomeDumbField {
		return false
	}
	return true
}
func (this *MockResource_Data) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource_Data)
	if !ok {
		that2, ok := that.(MockResource_Data)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Data != that1.Data {
		return false
	}
	return true
}
func (this *MockResource_OneofOne) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource_OneofOne)
	if !ok {
		that2, ok := that.(MockResource_OneofOne)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.OneofOne != that1.OneofOne {
		return false
	}
	return true
}
func (this *MockResource_OneofTwo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource_OneofTwo)
	if !ok {
		that2, ok := that.(MockResource_OneofTwo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.OneofTwo != that1.OneofTwo {
		return false
	}
	return true
}
func (this *FrequentlyChangingAnnotationsResource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FrequentlyChangingAnnotationsResource)
	if !ok {
		that2, ok := that.(FrequentlyChangingAnnotationsResource)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if this.Blah != that1.Blah {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
