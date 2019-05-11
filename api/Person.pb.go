// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Person.proto

package tutorial

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
}

func init() { proto.RegisterFile("Person.proto", fileDescriptor_841ab6396175eaf3) }

var fileDescriptor_841ab6396175eaf3 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xd1, 0x6a, 0xab, 0x40,
	0x10, 0x86, 0x8f, 0xc6, 0x48, 0x32, 0x86, 0xe0, 0x59, 0xc2, 0x41, 0xe4, 0x40, 0x25, 0xf4, 0x42,
	0x28, 0x6c, 0x20, 0x2d, 0xf4, 0xaa, 0x17, 0x15, 0x42, 0x5b, 0xda, 0x34, 0x22, 0x09, 0xbd, 0x2c,
	0xa6, 0x4e, 0x53, 0xa9, 0xba, 0x8b, 0xbb, 0x42, 0xf3, 0x4a, 0x7d, 0x85, 0xbe, 0x5c, 0xd1, 0xd5,
	0x10, 0x4a, 0xef, 0x66, 0x67, 0xbe, 0x19, 0x3f, 0x7f, 0x18, 0x85, 0x58, 0x0a, 0x56, 0x50, 0x5e,
	0x32, 0xc9, 0xc8, 0x40, 0x56, 0x92, 0x95, 0x69, 0x9c, 0xb9, 0x27, 0x3b, 0xc6, 0x76, 0x19, 0xce,
	0x9a, 0xfe, 0xb6, 0x7a, 0x9d, 0xc9, 0x34, 0x47, 0x21, 0xe3, 0x9c, 0x2b, 0x74, 0xfa, 0xa5, 0x83,
	0xa9, 0x76, 0x09, 0x01, 0xa3, 0x88, 0x73, 0x74, 0x34, 0x4f, 0xf3, 0x87, 0x51, 0x53, 0x93, 0x31,
	0xe8, 0x69, 0xe2, 0xe8, 0x9e, 0xe6, 0xf7, 0x23, 0x3d, 0x4d, 0xc8, 0x04, 0xfa, 0x98, 0xc7, 0x69,
	0xe6, 0xf4, 0x1a, 0x48, 0x3d, 0xc8, 0x05, 0x98, 0xfc, 0x8d, 0x15, 0x28, 0x1c, 0xc3, 0xeb, 0xf9,
	0xd6, 0xfc, 0x3f, 0xed, 0x04, 0x68, 0xeb, 0x15, 0xd6, 0xe3, 0xc7, 0x2a, 0xdf, 0x62, 0x19, 0xb5,
	0x2c, 0xb9, 0x82, 0x51, 0x16, 0x0b, 0xf9, 0x5c, 0xf1, 0x24, 0x96, 0x98, 0x38, 0x7d, 0x4f, 0xf3,
	0xad, 0xb9, 0x4b, 0x95, 0x32, 0xed, 0x94, 0xe9, 0xba, 0x53, 0x8e, 0xac, 0x9a, 0xdf, 0x28, 0xdc,
	0xdd, 0x80, 0x75, 0x74, 0x95, 0xfc, 0x03, 0xb3, 0x68, 0xaa, 0xd6, 0xbf, 0x7d, 0x11, 0x0a, 0x86,
	0xdc, 0x73, 0x6c, 0xfe, 0x61, 0x3c, 0x77, 0x7f, 0x37, 0x5b, 0xef, 0x39, 0x46, 0x0d, 0x37, 0x3d,
	0x83, 0xe1, 0xa1, 0x45, 0x00, 0xcc, 0xe5, 0x2a, 0xb8, 0x7b, 0x58, 0xd8, 0x7f, 0xc8, 0x00, 0x8c,
	0xdb, 0xd5, 0x72, 0x61, 0x6b, 0x75, 0xf5, 0xb4, 0x8a, 0xee, 0x6d, 0x7d, 0x7a, 0x09, 0xd6, 0x75,
	0x92, 0x94, 0x28, 0x44, 0xc0, 0xd8, 0x3b, 0xf1, 0xc1, 0xe4, 0xc8, 0x78, 0x56, 0x67, 0x58, 0xe7,
	0x60, 0xff, 0xfc, 0x5a, 0xd4, 0xce, 0x83, 0x10, 0x26, 0x2f, 0x2c, 0xa7, 0xf8, 0x11, 0xe7, 0x3c,
	0xc3, 0x03, 0x16, 0xfc, 0x3d, 0x3a, 0x17, 0xd6, 0x01, 0x88, 0x4f, 0xfd, 0xf4, 0x46, 0x05, 0x12,
	0x76, 0x81, 0x2c, 0xd4, 0x96, 0xa0, 0x47, 0xf0, 0xd6, 0x6c, 0xf2, 0x3a, 0xff, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x3a, 0xf1, 0xc5, 0xb5, 0x0a, 0x02, 0x00, 0x00,
}