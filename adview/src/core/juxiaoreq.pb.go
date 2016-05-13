// Code generated by protoc-gen-go.
// source: juxiaoreq.proto
// DO NOT EDIT!

/*
Package juxiaoreq is a generated protocol buffer package.

It is generated from these files:
	juxiaoreq.proto

It has these top-level messages:
	BidRequest
*/
package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type BidRequest_Network_Type int32

const (
	BidRequest_NET_UNKNOWN BidRequest_Network_Type = 0
	BidRequest_NET_WIFI    BidRequest_Network_Type = 1
	BidRequest_NET_2G      BidRequest_Network_Type = 2
	BidRequest_NET_3G      BidRequest_Network_Type = 3
	BidRequest_NET_4G      BidRequest_Network_Type = 4
)

var BidRequest_Network_Type_name = map[int32]string{
	0: "NET_UNKNOWN",
	1: "NET_WIFI",
	2: "NET_2G",
	3: "NET_3G",
	4: "NET_4G",
}
var BidRequest_Network_Type_value = map[string]int32{
	"NET_UNKNOWN": 0,
	"NET_WIFI":    1,
	"NET_2G":      2,
	"NET_3G":      3,
	"NET_4G":      4,
}

func (x BidRequest_Network_Type) Enum() *BidRequest_Network_Type {
	p := new(BidRequest_Network_Type)
	*p = x
	return p
}
func (x BidRequest_Network_Type) String() string {
	return proto.EnumName(BidRequest_Network_Type_name, int32(x))
}
func (x *BidRequest_Network_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_Network_Type_value, data, "BidRequest_Network_Type")
	if err != nil {
		return err
	}
	*x = BidRequest_Network_Type(value)
	return nil
}
func (BidRequest_Network_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type BidRequest_DeviceId_Device_Id_Type int32

const (
	BidRequest_DeviceId_DEVICE_UNKNOWN BidRequest_DeviceId_Device_Id_Type = 0
	BidRequest_DeviceId_IMEI           BidRequest_DeviceId_Device_Id_Type = 1
	BidRequest_DeviceId_IDFA           BidRequest_DeviceId_Device_Id_Type = 2
	BidRequest_DeviceId_AAID           BidRequest_DeviceId_Device_Id_Type = 3
	BidRequest_DeviceId_MAC            BidRequest_DeviceId_Device_Id_Type = 4
	BidRequest_DeviceId_IDFV           BidRequest_DeviceId_Device_Id_Type = 5
	BidRequest_DeviceId_M2ID           BidRequest_DeviceId_Device_Id_Type = 6
	BidRequest_DeviceId_SERIALID       BidRequest_DeviceId_Device_Id_Type = 7
	BidRequest_DeviceId_IMSI           BidRequest_DeviceId_Device_Id_Type = 8
)

var BidRequest_DeviceId_Device_Id_Type_name = map[int32]string{
	0: "DEVICE_UNKNOWN",
	1: "IMEI",
	2: "IDFA",
	3: "AAID",
	4: "MAC",
	5: "IDFV",
	6: "M2ID",
	7: "SERIALID",
	8: "IMSI",
}
var BidRequest_DeviceId_Device_Id_Type_value = map[string]int32{
	"DEVICE_UNKNOWN": 0,
	"IMEI":           1,
	"IDFA":           2,
	"AAID":           3,
	"MAC":            4,
	"IDFV":           5,
	"M2ID":           6,
	"SERIALID":       7,
	"IMSI":           8,
}

func (x BidRequest_DeviceId_Device_Id_Type) Enum() *BidRequest_DeviceId_Device_Id_Type {
	p := new(BidRequest_DeviceId_Device_Id_Type)
	*p = x
	return p
}
func (x BidRequest_DeviceId_Device_Id_Type) String() string {
	return proto.EnumName(BidRequest_DeviceId_Device_Id_Type_name, int32(x))
}
func (x *BidRequest_DeviceId_Device_Id_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_DeviceId_Device_Id_Type_value, data, "BidRequest_DeviceId_Device_Id_Type")
	if err != nil {
		return err
	}
	*x = BidRequest_DeviceId_Device_Id_Type(value)
	return nil
}
func (BidRequest_DeviceId_Device_Id_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 1, 0}
}

type BidRequest_DeviceId_Hash_Type int32

const (
	BidRequest_DeviceId_NONE  BidRequest_DeviceId_Hash_Type = 0
	BidRequest_DeviceId_MD5   BidRequest_DeviceId_Hash_Type = 1
	BidRequest_DeviceId_SH1   BidRequest_DeviceId_Hash_Type = 2
	BidRequest_DeviceId_OTHER BidRequest_DeviceId_Hash_Type = 3
)

var BidRequest_DeviceId_Hash_Type_name = map[int32]string{
	0: "NONE",
	1: "MD5",
	2: "SH1",
	3: "OTHER",
}
var BidRequest_DeviceId_Hash_Type_value = map[string]int32{
	"NONE":  0,
	"MD5":   1,
	"SH1":   2,
	"OTHER": 3,
}

func (x BidRequest_DeviceId_Hash_Type) Enum() *BidRequest_DeviceId_Hash_Type {
	p := new(BidRequest_DeviceId_Hash_Type)
	*p = x
	return p
}
func (x BidRequest_DeviceId_Hash_Type) String() string {
	return proto.EnumName(BidRequest_DeviceId_Hash_Type_name, int32(x))
}
func (x *BidRequest_DeviceId_Hash_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_DeviceId_Hash_Type_value, data, "BidRequest_DeviceId_Hash_Type")
	if err != nil {
		return err
	}
	*x = BidRequest_DeviceId_Hash_Type(value)
	return nil
}
func (BidRequest_DeviceId_Hash_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 1, 1}
}

type BidRequest_Device_OS_Type int32

const (
	BidRequest_Device_OS_UNKNOWN BidRequest_Device_OS_Type = 0
	BidRequest_Device_OS_IOS     BidRequest_Device_OS_Type = 1
	BidRequest_Device_OS_ANDROID BidRequest_Device_OS_Type = 2
	BidRequest_Device_OS_WP      BidRequest_Device_OS_Type = 3
)

var BidRequest_Device_OS_Type_name = map[int32]string{
	0: "OS_UNKNOWN",
	1: "OS_IOS",
	2: "OS_ANDROID",
	3: "OS_WP",
}
var BidRequest_Device_OS_Type_value = map[string]int32{
	"OS_UNKNOWN": 0,
	"OS_IOS":     1,
	"OS_ANDROID": 2,
	"OS_WP":      3,
}

func (x BidRequest_Device_OS_Type) Enum() *BidRequest_Device_OS_Type {
	p := new(BidRequest_Device_OS_Type)
	*p = x
	return p
}
func (x BidRequest_Device_OS_Type) String() string {
	return proto.EnumName(BidRequest_Device_OS_Type_name, int32(x))
}
func (x *BidRequest_Device_OS_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_Device_OS_Type_value, data, "BidRequest_Device_OS_Type")
	if err != nil {
		return err
	}
	*x = BidRequest_Device_OS_Type(value)
	return nil
}
func (BidRequest_Device_OS_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 2, 0}
}

type BidRequest_Device_Device_Type int32

const (
	BidRequest_Device_UNKNOWN BidRequest_Device_Device_Type = 0
	BidRequest_Device_TABLET  BidRequest_Device_Device_Type = 1
	BidRequest_Device_PHONE   BidRequest_Device_Device_Type = 2
)

var BidRequest_Device_Device_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "TABLET",
	2: "PHONE",
}
var BidRequest_Device_Device_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"TABLET":  1,
	"PHONE":   2,
}

func (x BidRequest_Device_Device_Type) Enum() *BidRequest_Device_Device_Type {
	p := new(BidRequest_Device_Device_Type)
	*p = x
	return p
}
func (x BidRequest_Device_Device_Type) String() string {
	return proto.EnumName(BidRequest_Device_Device_Type_name, int32(x))
}
func (x *BidRequest_Device_Device_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_Device_Device_Type_value, data, "BidRequest_Device_Device_Type")
	if err != nil {
		return err
	}
	*x = BidRequest_Device_Device_Type(value)
	return nil
}
func (BidRequest_Device_Device_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 2, 1}
}

type BidRequest_Device_Screen_Orientation int32

const (
	BidRequest_Device_SCREEN_ORIENTATION_UNKNOWN   BidRequest_Device_Screen_Orientation = 0
	BidRequest_Device_SCREEN_ORIENTATION_PORTRAIT  BidRequest_Device_Screen_Orientation = 1
	BidRequest_Device_SCREEN_ORIENTATION_LANDSCAPE BidRequest_Device_Screen_Orientation = 2
)

var BidRequest_Device_Screen_Orientation_name = map[int32]string{
	0: "SCREEN_ORIENTATION_UNKNOWN",
	1: "SCREEN_ORIENTATION_PORTRAIT",
	2: "SCREEN_ORIENTATION_LANDSCAPE",
}
var BidRequest_Device_Screen_Orientation_value = map[string]int32{
	"SCREEN_ORIENTATION_UNKNOWN":   0,
	"SCREEN_ORIENTATION_PORTRAIT":  1,
	"SCREEN_ORIENTATION_LANDSCAPE": 2,
}

func (x BidRequest_Device_Screen_Orientation) Enum() *BidRequest_Device_Screen_Orientation {
	p := new(BidRequest_Device_Screen_Orientation)
	*p = x
	return p
}
func (x BidRequest_Device_Screen_Orientation) String() string {
	return proto.EnumName(BidRequest_Device_Screen_Orientation_name, int32(x))
}
func (x *BidRequest_Device_Screen_Orientation) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_Device_Screen_Orientation_value, data, "BidRequest_Device_Screen_Orientation")
	if err != nil {
		return err
	}
	*x = BidRequest_Device_Screen_Orientation(value)
	return nil
}
func (BidRequest_Device_Screen_Orientation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 2, 2}
}

type BidRequest_Device_Carrier_Id int32

const (
	BidRequest_Device_CARRIER_UNKNOWN BidRequest_Device_Carrier_Id = 0
	BidRequest_Device_CHINA_MOBILE    BidRequest_Device_Carrier_Id = 70120
	BidRequest_Device_CHINA_TELECOM   BidRequest_Device_Carrier_Id = 70121
	BidRequest_Device_UNICOM          BidRequest_Device_Carrier_Id = 70123
)

var BidRequest_Device_Carrier_Id_name = map[int32]string{
	0:     "CARRIER_UNKNOWN",
	70120: "CHINA_MOBILE",
	70121: "CHINA_TELECOM",
	70123: "UNICOM",
}
var BidRequest_Device_Carrier_Id_value = map[string]int32{
	"CARRIER_UNKNOWN": 0,
	"CHINA_MOBILE":    70120,
	"CHINA_TELECOM":   70121,
	"UNICOM":          70123,
}

func (x BidRequest_Device_Carrier_Id) Enum() *BidRequest_Device_Carrier_Id {
	p := new(BidRequest_Device_Carrier_Id)
	*p = x
	return p
}
func (x BidRequest_Device_Carrier_Id) String() string {
	return proto.EnumName(BidRequest_Device_Carrier_Id_name, int32(x))
}
func (x *BidRequest_Device_Carrier_Id) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_Device_Carrier_Id_value, data, "BidRequest_Device_Carrier_Id")
	if err != nil {
		return err
	}
	*x = BidRequest_Device_Carrier_Id(value)
	return nil
}
func (BidRequest_Device_Carrier_Id) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 2, 3}
}

type BidRequest_Adspaces_Adspace_Type int32

const (
	BidRequest_Adspaces_BANNER       BidRequest_Adspaces_Adspace_Type = 1
	BidRequest_Adspaces_INTERSTITIAL BidRequest_Adspaces_Adspace_Type = 2
	BidRequest_Adspaces_VIDEO        BidRequest_Adspaces_Adspace_Type = 3
	BidRequest_Adspaces_NATIVE       BidRequest_Adspaces_Adspace_Type = 4
	BidRequest_Adspaces_EMBEDDED     BidRequest_Adspaces_Adspace_Type = 5
	BidRequest_Adspaces_OPENING      BidRequest_Adspaces_Adspace_Type = 6
	BidRequest_Adspaces_FOCUS        BidRequest_Adspaces_Adspace_Type = 7
)

var BidRequest_Adspaces_Adspace_Type_name = map[int32]string{
	1: "BANNER",
	2: "INTERSTITIAL",
	3: "VIDEO",
	4: "NATIVE",
	5: "EMBEDDED",
	6: "OPENING",
	7: "FOCUS",
}
var BidRequest_Adspaces_Adspace_Type_value = map[string]int32{
	"BANNER":       1,
	"INTERSTITIAL": 2,
	"VIDEO":        3,
	"NATIVE":       4,
	"EMBEDDED":     5,
	"OPENING":      6,
	"FOCUS":        7,
}

func (x BidRequest_Adspaces_Adspace_Type) Enum() *BidRequest_Adspaces_Adspace_Type {
	p := new(BidRequest_Adspaces_Adspace_Type)
	*p = x
	return p
}
func (x BidRequest_Adspaces_Adspace_Type) String() string {
	return proto.EnumName(BidRequest_Adspaces_Adspace_Type_name, int32(x))
}
func (x *BidRequest_Adspaces_Adspace_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_Adspaces_Adspace_Type_value, data, "BidRequest_Adspaces_Adspace_Type")
	if err != nil {
		return err
	}
	*x = BidRequest_Adspaces_Adspace_Type(value)
	return nil
}
func (BidRequest_Adspaces_Adspace_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 3, 0}
}

type BidRequest_Adspaces_Adspace_Position int32

const (
	BidRequest_Adspaces_NONE      BidRequest_Adspaces_Adspace_Position = 0
	BidRequest_Adspaces_FIRST_POS BidRequest_Adspaces_Adspace_Position = 1
	BidRequest_Adspaces_OTHERS    BidRequest_Adspaces_Adspace_Position = 2
)

var BidRequest_Adspaces_Adspace_Position_name = map[int32]string{
	0: "NONE",
	1: "FIRST_POS",
	2: "OTHERS",
}
var BidRequest_Adspaces_Adspace_Position_value = map[string]int32{
	"NONE":      0,
	"FIRST_POS": 1,
	"OTHERS":    2,
}

func (x BidRequest_Adspaces_Adspace_Position) Enum() *BidRequest_Adspaces_Adspace_Position {
	p := new(BidRequest_Adspaces_Adspace_Position)
	*p = x
	return p
}
func (x BidRequest_Adspaces_Adspace_Position) String() string {
	return proto.EnumName(BidRequest_Adspaces_Adspace_Position_name, int32(x))
}
func (x *BidRequest_Adspaces_Adspace_Position) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_Adspaces_Adspace_Position_value, data, "BidRequest_Adspaces_Adspace_Position")
	if err != nil {
		return err
	}
	*x = BidRequest_Adspaces_Adspace_Position(value)
	return nil
}
func (BidRequest_Adspaces_Adspace_Position) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 3, 1}
}

type BidRequest_Adspaces_Open_Type int32

const (
	BidRequest_Adspaces_ALL   BidRequest_Adspaces_Open_Type = 0
	BidRequest_Adspaces_INNER BidRequest_Adspaces_Open_Type = 1
	BidRequest_Adspaces_OUTER BidRequest_Adspaces_Open_Type = 2
)

var BidRequest_Adspaces_Open_Type_name = map[int32]string{
	0: "ALL",
	1: "INNER",
	2: "OUTER",
}
var BidRequest_Adspaces_Open_Type_value = map[string]int32{
	"ALL":   0,
	"INNER": 1,
	"OUTER": 2,
}

func (x BidRequest_Adspaces_Open_Type) Enum() *BidRequest_Adspaces_Open_Type {
	p := new(BidRequest_Adspaces_Open_Type)
	*p = x
	return p
}
func (x BidRequest_Adspaces_Open_Type) String() string {
	return proto.EnumName(BidRequest_Adspaces_Open_Type_name, int32(x))
}
func (x *BidRequest_Adspaces_Open_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_Adspaces_Open_Type_value, data, "BidRequest_Adspaces_Open_Type")
	if err != nil {
		return err
	}
	*x = BidRequest_Adspaces_Open_Type(value)
	return nil
}
func (BidRequest_Adspaces_Open_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 3, 2}
}

type BidRequest_Adspaces_Interaction_Type int32

const (
	BidRequest_Adspaces_ANY            BidRequest_Adspaces_Interaction_Type = 0
	BidRequest_Adspaces_NO_INTERACTION BidRequest_Adspaces_Interaction_Type = 1
	BidRequest_Adspaces_BROWSE         BidRequest_Adspaces_Interaction_Type = 2
	BidRequest_Adspaces_DOWNLOAD       BidRequest_Adspaces_Interaction_Type = 3
	BidRequest_Adspaces_DIALING        BidRequest_Adspaces_Interaction_Type = 4
	BidRequest_Adspaces_MESSAGE        BidRequest_Adspaces_Interaction_Type = 5
	BidRequest_Adspaces_MAIL           BidRequest_Adspaces_Interaction_Type = 6
)

var BidRequest_Adspaces_Interaction_Type_name = map[int32]string{
	0: "ANY",
	1: "NO_INTERACTION",
	2: "BROWSE",
	3: "DOWNLOAD",
	4: "DIALING",
	5: "MESSAGE",
	6: "MAIL",
}
var BidRequest_Adspaces_Interaction_Type_value = map[string]int32{
	"ANY":            0,
	"NO_INTERACTION": 1,
	"BROWSE":         2,
	"DOWNLOAD":       3,
	"DIALING":        4,
	"MESSAGE":        5,
	"MAIL":           6,
}

func (x BidRequest_Adspaces_Interaction_Type) Enum() *BidRequest_Adspaces_Interaction_Type {
	p := new(BidRequest_Adspaces_Interaction_Type)
	*p = x
	return p
}
func (x BidRequest_Adspaces_Interaction_Type) String() string {
	return proto.EnumName(BidRequest_Adspaces_Interaction_Type_name, int32(x))
}
func (x *BidRequest_Adspaces_Interaction_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(BidRequest_Adspaces_Interaction_Type_value, data, "BidRequest_Adspaces_Interaction_Type")
	if err != nil {
		return err
	}
	*x = BidRequest_Adspaces_Interaction_Type(value)
	return nil
}
func (BidRequest_Adspaces_Interaction_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor3, []int{0, 3, 3}
}

type BidRequest struct {
	Bid              *string                  `protobuf:"bytes,1,req,name=bid" json:"bid,omitempty"`
	App              *BidRequest_App          `protobuf:"bytes,2,req,name=app" json:"app,omitempty"`
	Device           *BidRequest_Device       `protobuf:"bytes,3,req,name=device" json:"device,omitempty"`
	Adspaces         []*BidRequest_Adspaces   `protobuf:"bytes,4,rep,name=adspaces" json:"adspaces,omitempty"`
	Uid              *string                  `protobuf:"bytes,5,req,name=uid" json:"uid,omitempty"`
	Ip               *string                  `protobuf:"bytes,6,req,name=ip" json:"ip,omitempty"`
	UserAgent        *string                  `protobuf:"bytes,7,req,name=user_agent" json:"user_agent,omitempty"`
	NetworkType      *BidRequest_Network_Type `protobuf:"varint,8,req,name=network_type,enum=BidRequest_Network_Type" json:"network_type,omitempty"`
	Longitude        *float64                 `protobuf:"fixed64,9,opt,name=longitude" json:"longitude,omitempty"`
	Latitude         *float64                 `protobuf:"fixed64,10,opt,name=latitude" json:"latitude,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *BidRequest) Reset()                    { *m = BidRequest{} }
func (m *BidRequest) String() string            { return proto.CompactTextString(m) }
func (*BidRequest) ProtoMessage()               {}
func (*BidRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *BidRequest) GetBid() string {
	if m != nil && m.Bid != nil {
		return *m.Bid
	}
	return ""
}

func (m *BidRequest) GetApp() *BidRequest_App {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *BidRequest) GetDevice() *BidRequest_Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *BidRequest) GetAdspaces() []*BidRequest_Adspaces {
	if m != nil {
		return m.Adspaces
	}
	return nil
}

func (m *BidRequest) GetUid() string {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return ""
}

func (m *BidRequest) GetIp() string {
	if m != nil && m.Ip != nil {
		return *m.Ip
	}
	return ""
}

func (m *BidRequest) GetUserAgent() string {
	if m != nil && m.UserAgent != nil {
		return *m.UserAgent
	}
	return ""
}

func (m *BidRequest) GetNetworkType() BidRequest_Network_Type {
	if m != nil && m.NetworkType != nil {
		return *m.NetworkType
	}
	return BidRequest_NET_UNKNOWN
}

func (m *BidRequest) GetLongitude() float64 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *BidRequest) GetLatitude() float64 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

type BidRequest_App struct {
	AppName          *string `protobuf:"bytes,1,req,name=app_name" json:"app_name,omitempty"`
	PackageName      *string `protobuf:"bytes,2,req,name=package_name" json:"package_name,omitempty"`
	Category         *int32  `protobuf:"varint,3,req,name=category" json:"category,omitempty"`
	AppVersion       *string `protobuf:"bytes,4,req,name=app_version" json:"app_version,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BidRequest_App) Reset()                    { *m = BidRequest_App{} }
func (m *BidRequest_App) String() string            { return proto.CompactTextString(m) }
func (*BidRequest_App) ProtoMessage()               {}
func (*BidRequest_App) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func (m *BidRequest_App) GetAppName() string {
	if m != nil && m.AppName != nil {
		return *m.AppName
	}
	return ""
}

func (m *BidRequest_App) GetPackageName() string {
	if m != nil && m.PackageName != nil {
		return *m.PackageName
	}
	return ""
}

func (m *BidRequest_App) GetCategory() int32 {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return 0
}

func (m *BidRequest_App) GetAppVersion() string {
	if m != nil && m.AppVersion != nil {
		return *m.AppVersion
	}
	return ""
}

type BidRequest_DeviceId struct {
	DeviceId         *string                             `protobuf:"bytes,1,req,name=device_id" json:"device_id,omitempty"`
	DeviceIdType     *BidRequest_DeviceId_Device_Id_Type `protobuf:"varint,2,req,name=device_id_type,enum=BidRequest_DeviceId_Device_Id_Type" json:"device_id_type,omitempty"`
	HashType         *BidRequest_DeviceId_Hash_Type      `protobuf:"varint,3,req,name=hash_type,enum=BidRequest_DeviceId_Hash_Type" json:"hash_type,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *BidRequest_DeviceId) Reset()                    { *m = BidRequest_DeviceId{} }
func (m *BidRequest_DeviceId) String() string            { return proto.CompactTextString(m) }
func (*BidRequest_DeviceId) ProtoMessage()               {}
func (*BidRequest_DeviceId) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 1} }

func (m *BidRequest_DeviceId) GetDeviceId() string {
	if m != nil && m.DeviceId != nil {
		return *m.DeviceId
	}
	return ""
}

func (m *BidRequest_DeviceId) GetDeviceIdType() BidRequest_DeviceId_Device_Id_Type {
	if m != nil && m.DeviceIdType != nil {
		return *m.DeviceIdType
	}
	return BidRequest_DeviceId_DEVICE_UNKNOWN
}

func (m *BidRequest_DeviceId) GetHashType() BidRequest_DeviceId_Hash_Type {
	if m != nil && m.HashType != nil {
		return *m.HashType
	}
	return BidRequest_DeviceId_NONE
}

type BidRequest_Device struct {
	DeviceId          []*BidRequest_DeviceId                `protobuf:"bytes,1,rep,name=device_id" json:"device_id,omitempty"`
	OsType            *BidRequest_Device_OS_Type            `protobuf:"varint,2,req,name=os_type,enum=BidRequest_Device_OS_Type" json:"os_type,omitempty"`
	OsVersion         *string                               `protobuf:"bytes,3,req,name=os_version" json:"os_version,omitempty"`
	Brand             *string                               `protobuf:"bytes,4,req,name=brand" json:"brand,omitempty"`
	Model             *string                               `protobuf:"bytes,5,req,name=model" json:"model,omitempty"`
	DeviceType        *BidRequest_Device_Device_Type        `protobuf:"varint,6,opt,name=device_type,enum=BidRequest_Device_Device_Type" json:"device_type,omitempty"`
	ScreenWidth       *int32                                `protobuf:"varint,7,req,name=screen_width" json:"screen_width,omitempty"`
	ScreenHeight      *int32                                `protobuf:"varint,8,req,name=screen_height" json:"screen_height,omitempty"`
	ScreenDensity     *float64                              `protobuf:"fixed64,9,req,name=screen_density" json:"screen_density,omitempty"`
	ScreenOrientation *BidRequest_Device_Screen_Orientation `protobuf:"varint,10,opt,name=screen_orientation,enum=BidRequest_Device_Screen_Orientation" json:"screen_orientation,omitempty"`
	CarrierId         *BidRequest_Device_Carrier_Id         `protobuf:"varint,11,req,name=carrier_id,enum=BidRequest_Device_Carrier_Id" json:"carrier_id,omitempty"`
	XXX_unrecognized  []byte                                `json:"-"`
}

func (m *BidRequest_Device) Reset()                    { *m = BidRequest_Device{} }
func (m *BidRequest_Device) String() string            { return proto.CompactTextString(m) }
func (*BidRequest_Device) ProtoMessage()               {}
func (*BidRequest_Device) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 2} }

func (m *BidRequest_Device) GetDeviceId() []*BidRequest_DeviceId {
	if m != nil {
		return m.DeviceId
	}
	return nil
}

func (m *BidRequest_Device) GetOsType() BidRequest_Device_OS_Type {
	if m != nil && m.OsType != nil {
		return *m.OsType
	}
	return BidRequest_Device_OS_UNKNOWN
}

func (m *BidRequest_Device) GetOsVersion() string {
	if m != nil && m.OsVersion != nil {
		return *m.OsVersion
	}
	return ""
}

func (m *BidRequest_Device) GetBrand() string {
	if m != nil && m.Brand != nil {
		return *m.Brand
	}
	return ""
}

func (m *BidRequest_Device) GetModel() string {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return ""
}

func (m *BidRequest_Device) GetDeviceType() BidRequest_Device_Device_Type {
	if m != nil && m.DeviceType != nil {
		return *m.DeviceType
	}
	return BidRequest_Device_UNKNOWN
}

func (m *BidRequest_Device) GetScreenWidth() int32 {
	if m != nil && m.ScreenWidth != nil {
		return *m.ScreenWidth
	}
	return 0
}

func (m *BidRequest_Device) GetScreenHeight() int32 {
	if m != nil && m.ScreenHeight != nil {
		return *m.ScreenHeight
	}
	return 0
}

func (m *BidRequest_Device) GetScreenDensity() float64 {
	if m != nil && m.ScreenDensity != nil {
		return *m.ScreenDensity
	}
	return 0
}

func (m *BidRequest_Device) GetScreenOrientation() BidRequest_Device_Screen_Orientation {
	if m != nil && m.ScreenOrientation != nil {
		return *m.ScreenOrientation
	}
	return BidRequest_Device_SCREEN_ORIENTATION_UNKNOWN
}

func (m *BidRequest_Device) GetCarrierId() BidRequest_Device_Carrier_Id {
	if m != nil && m.CarrierId != nil {
		return *m.CarrierId
	}
	return BidRequest_Device_CARRIER_UNKNOWN
}

type BidRequest_Adspaces struct {
	AdspaceId        *string                                `protobuf:"bytes,1,req,name=adspace_id" json:"adspace_id,omitempty"`
	AdspaceType      *BidRequest_Adspaces_Adspace_Type      `protobuf:"varint,2,req,name=adspace_type,enum=BidRequest_Adspaces_Adspace_Type" json:"adspace_type,omitempty"`
	AdspacePosition  *BidRequest_Adspaces_Adspace_Position  `protobuf:"varint,3,opt,name=adspace_position,enum=BidRequest_Adspaces_Adspace_Position" json:"adspace_position,omitempty"`
	AllowedHtml      *bool                                  `protobuf:"varint,4,req,name=allowed_html" json:"allowed_html,omitempty"`
	Width            *int32                                 `protobuf:"varint,5,req,name=width" json:"width,omitempty"`
	Height           *int32                                 `protobuf:"varint,6,req,name=height" json:"height,omitempty"`
	ImpressionNum    *int32                                 `protobuf:"varint,7,req,name=impression_num" json:"impression_num,omitempty"`
	Keywords         *string                                `protobuf:"bytes,8,opt,name=keywords" json:"keywords,omitempty"`
	Channel          *string                                `protobuf:"bytes,9,opt,name=channel" json:"channel,omitempty"`
	OpenType         *BidRequest_Adspaces_Open_Type         `protobuf:"varint,10,req,name=open_type,enum=BidRequest_Adspaces_Open_Type" json:"open_type,omitempty"`
	InteractionType  []BidRequest_Adspaces_Interaction_Type `protobuf:"varint,11,rep,name=interaction_type,enum=BidRequest_Adspaces_Interaction_Type" json:"interaction_type,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *BidRequest_Adspaces) Reset()                    { *m = BidRequest_Adspaces{} }
func (m *BidRequest_Adspaces) String() string            { return proto.CompactTextString(m) }
func (*BidRequest_Adspaces) ProtoMessage()               {}
func (*BidRequest_Adspaces) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 3} }

func (m *BidRequest_Adspaces) GetAdspaceId() string {
	if m != nil && m.AdspaceId != nil {
		return *m.AdspaceId
	}
	return ""
}

func (m *BidRequest_Adspaces) GetAdspaceType() BidRequest_Adspaces_Adspace_Type {
	if m != nil && m.AdspaceType != nil {
		return *m.AdspaceType
	}
	return BidRequest_Adspaces_BANNER
}

func (m *BidRequest_Adspaces) GetAdspacePosition() BidRequest_Adspaces_Adspace_Position {
	if m != nil && m.AdspacePosition != nil {
		return *m.AdspacePosition
	}
	return BidRequest_Adspaces_NONE
}

func (m *BidRequest_Adspaces) GetAllowedHtml() bool {
	if m != nil && m.AllowedHtml != nil {
		return *m.AllowedHtml
	}
	return false
}

func (m *BidRequest_Adspaces) GetWidth() int32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *BidRequest_Adspaces) GetHeight() int32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

func (m *BidRequest_Adspaces) GetImpressionNum() int32 {
	if m != nil && m.ImpressionNum != nil {
		return *m.ImpressionNum
	}
	return 0
}

func (m *BidRequest_Adspaces) GetKeywords() string {
	if m != nil && m.Keywords != nil {
		return *m.Keywords
	}
	return ""
}

func (m *BidRequest_Adspaces) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

func (m *BidRequest_Adspaces) GetOpenType() BidRequest_Adspaces_Open_Type {
	if m != nil && m.OpenType != nil {
		return *m.OpenType
	}
	return BidRequest_Adspaces_ALL
}

func (m *BidRequest_Adspaces) GetInteractionType() []BidRequest_Adspaces_Interaction_Type {
	if m != nil {
		return m.InteractionType
	}
	return nil
}

func init() {
	proto.RegisterType((*BidRequest)(nil), "BidRequest")
	proto.RegisterType((*BidRequest_App)(nil), "BidRequest.App")
	proto.RegisterType((*BidRequest_DeviceId)(nil), "BidRequest.DeviceId")
	proto.RegisterType((*BidRequest_Device)(nil), "BidRequest.Device")
	proto.RegisterType((*BidRequest_Adspaces)(nil), "BidRequest.Adspaces")
	proto.RegisterEnum("BidRequest_Network_Type", BidRequest_Network_Type_name, BidRequest_Network_Type_value)
	proto.RegisterEnum("BidRequest_DeviceId_Device_Id_Type", BidRequest_DeviceId_Device_Id_Type_name, BidRequest_DeviceId_Device_Id_Type_value)
	proto.RegisterEnum("BidRequest_DeviceId_Hash_Type", BidRequest_DeviceId_Hash_Type_name, BidRequest_DeviceId_Hash_Type_value)
	proto.RegisterEnum("BidRequest_Device_OS_Type", BidRequest_Device_OS_Type_name, BidRequest_Device_OS_Type_value)
	proto.RegisterEnum("BidRequest_Device_Device_Type", BidRequest_Device_Device_Type_name, BidRequest_Device_Device_Type_value)
	proto.RegisterEnum("BidRequest_Device_Screen_Orientation", BidRequest_Device_Screen_Orientation_name, BidRequest_Device_Screen_Orientation_value)
	proto.RegisterEnum("BidRequest_Device_Carrier_Id", BidRequest_Device_Carrier_Id_name, BidRequest_Device_Carrier_Id_value)
	proto.RegisterEnum("BidRequest_Adspaces_Adspace_Type", BidRequest_Adspaces_Adspace_Type_name, BidRequest_Adspaces_Adspace_Type_value)
	proto.RegisterEnum("BidRequest_Adspaces_Adspace_Position", BidRequest_Adspaces_Adspace_Position_name, BidRequest_Adspaces_Adspace_Position_value)
	proto.RegisterEnum("BidRequest_Adspaces_Open_Type", BidRequest_Adspaces_Open_Type_name, BidRequest_Adspaces_Open_Type_value)
	proto.RegisterEnum("BidRequest_Adspaces_Interaction_Type", BidRequest_Adspaces_Interaction_Type_name, BidRequest_Adspaces_Interaction_Type_value)
}

var fileDescriptor3 = []byte{
	// 1122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x55, 0xdd, 0x72, 0xe3, 0x44,
	0x13, 0xfd, 0x62, 0xf9, 0xb7, 0xed, 0x38, 0xfa, 0xc4, 0x42, 0xb9, 0xcc, 0xb2, 0x04, 0x53, 0xc0,
	0x16, 0x54, 0xa5, 0x2a, 0x59, 0xa8, 0xbd, 0xe0, 0x02, 0x64, 0x69, 0x92, 0x4c, 0x61, 0x8f, 0x8c,
	0xa4, 0x24, 0x70, 0xa5, 0xd2, 0x5a, 0x53, 0xb1, 0x88, 0x2d, 0x69, 0x25, 0x79, 0x43, 0x9e, 0x67,
	0xdf, 0x02, 0x5e, 0x80, 0x47, 0x81, 0xe2, 0x96, 0x07, 0xa0, 0x67, 0x24, 0xff, 0x64, 0x6d, 0xee,
	0x46, 0x3d, 0xa7, 0xa7, 0xfb, 0xf4, 0x9c, 0x39, 0x82, 0xa3, 0x5f, 0x96, 0xbf, 0x86, 0x7e, 0x9c,
	0xf2, 0xd7, 0x27, 0x49, 0x1a, 0xe7, 0xf1, 0xe0, 0x1f, 0x15, 0x60, 0x18, 0x06, 0x36, 0x7f, 0xbd,
	0xe4, 0x59, 0xae, 0xb5, 0x41, 0x79, 0x15, 0x06, 0xbd, 0x83, 0xe3, 0xca, 0xf3, 0x96, 0xf6, 0x14,
	0x14, 0x3f, 0x49, 0x7a, 0x15, 0xfc, 0x68, 0x9f, 0x1d, 0x9d, 0x6c, 0x60, 0x27, 0x7a, 0x92, 0x68,
	0x03, 0xa8, 0x07, 0xfc, 0x4d, 0x38, 0xe5, 0x3d, 0x45, 0x02, 0xb4, 0x6d, 0x80, 0x29, 0x77, 0xb4,
	0xcf, 0xa1, 0xe9, 0x07, 0x59, 0xe2, 0x4f, 0x79, 0xd6, 0xab, 0x1e, 0x2b, 0x88, 0x7a, 0xf2, 0xe8,
	0x98, 0x72, 0x4f, 0x94, 0x5d, 0x62, 0xd9, 0x9a, 0x2c, 0x0b, 0x50, 0x09, 0x93, 0x5e, 0x5d, 0xae,
	0xf1, 0x63, 0x99, 0xf1, 0xd4, 0xf3, 0x6f, 0x79, 0x94, 0xf7, 0x1a, 0x32, 0x76, 0x02, 0x9d, 0x88,
	0xe7, 0xf7, 0x71, 0x7a, 0xe7, 0xe5, 0x0f, 0x09, 0xef, 0x35, 0x31, 0xda, 0x3d, 0xeb, 0x6d, 0x1f,
	0xcc, 0xca, 0x7d, 0x17, 0xf7, 0xb5, 0xff, 0x43, 0x6b, 0x1e, 0x47, 0xb7, 0x61, 0xbe, 0x0c, 0x78,
	0xaf, 0x75, 0x7c, 0xf0, 0xfc, 0x40, 0x53, 0xa1, 0x39, 0xf7, 0xf3, 0x22, 0x02, 0x22, 0xd2, 0x77,
	0x41, 0x11, 0xa4, 0x70, 0x03, 0x29, 0x7b, 0x91, 0xbf, 0xe0, 0xe5, 0x10, 0x9e, 0x40, 0x07, 0x7b,
	0xbc, 0xc3, 0x06, 0x8a, 0x68, 0x45, 0x46, 0x11, 0x37, 0xf5, 0x73, 0x7e, 0x1b, 0xa7, 0x0f, 0x92,
	0x7e, 0x4d, 0x7b, 0x0f, 0xda, 0x22, 0xf3, 0x0d, 0x4f, 0xb3, 0x30, 0x8e, 0x90, 0x2d, 0xc2, 0xfa,
	0xbf, 0x55, 0xa0, 0x59, 0x8c, 0x82, 0x06, 0xa2, 0x8f, 0x62, 0x60, 0xde, 0x7a, 0xc2, 0xdf, 0x42,
	0x77, 0x1d, 0x2a, 0xc8, 0x54, 0x24, 0x99, 0x4f, 0x77, 0x67, 0x49, 0x83, 0x72, 0xe1, 0xd1, 0xa0,
	0xe0, 0x75, 0x0a, 0xad, 0x99, 0x9f, 0xcd, 0x8a, 0x3c, 0x45, 0xe6, 0x3d, 0xdb, 0x9b, 0x77, 0x29,
	0x50, 0x22, 0x65, 0x70, 0x0f, 0xdd, 0x77, 0x0e, 0xd1, 0x30, 0x42, 0xae, 0xa9, 0x41, 0xbc, 0x2b,
	0xf6, 0x03, 0xb3, 0x6e, 0x98, 0xfa, 0x3f, 0xad, 0x09, 0x55, 0x3a, 0x26, 0x54, 0x3d, 0x90, 0x2b,
	0xf3, 0x5c, 0x57, 0x2b, 0x62, 0xa5, 0xeb, 0xd4, 0x54, 0x15, 0xad, 0x01, 0xca, 0x58, 0x37, 0xd4,
	0x6a, 0xb9, 0x79, 0xad, 0xd6, 0xc4, 0x6a, 0x7c, 0x86, 0x9b, 0x75, 0xad, 0x03, 0x4d, 0x87, 0xd8,
	0x54, 0x1f, 0xe1, 0x57, 0xa3, 0x38, 0xc8, 0xa1, 0x6a, 0x73, 0x70, 0x06, 0xad, 0x75, 0x17, 0x22,
	0xcc, 0x2c, 0x46, 0xb0, 0x92, 0x38, 0xcb, 0xfc, 0x06, 0x0b, 0xe1, 0xc2, 0xb9, 0x3c, 0xc5, 0x3a,
	0x2d, 0xa8, 0x59, 0xee, 0x25, 0xb1, 0x55, 0xa5, 0xff, 0x47, 0x0d, 0xea, 0xa5, 0x8e, 0xbe, 0x78,
	0x3c, 0xba, 0x1d, 0x21, 0xad, 0x67, 0xfc, 0x15, 0x34, 0xe2, 0x6c, 0x7b, 0x92, 0xfd, 0x5d, 0xd8,
	0x89, 0xe5, 0xac, 0xb8, 0x03, 0x82, 0x57, 0x37, 0xa6, 0xc8, 0x1b, 0x39, 0x84, 0xda, 0xab, 0xd4,
	0x8f, 0x82, 0xe2, 0x02, 0xc5, 0xe7, 0x22, 0x0e, 0xf8, 0xbc, 0x94, 0xe6, 0x0b, 0x68, 0x97, 0x7d,
	0xc8, 0x12, 0x75, 0x94, 0xce, 0xde, 0xa1, 0xaf, 0xae, 0x4a, 0x96, 0x41, 0x05, 0x65, 0xd3, 0x94,
	0xf3, 0xc8, 0xbb, 0x0f, 0x83, 0x7c, 0x26, 0x55, 0x5c, 0xd3, 0xde, 0x87, 0xc3, 0x32, 0x3a, 0xe3,
	0xe1, 0xed, 0x2c, 0x97, 0x32, 0xae, 0x69, 0x1f, 0x40, 0xb7, 0x0c, 0x07, 0x3c, 0xca, 0xc2, 0xfc,
	0x01, 0x15, 0x5b, 0x41, 0xc5, 0xea, 0xa0, 0x95, 0xf1, 0x38, 0x0d, 0xf1, 0x2d, 0xa0, 0x7a, 0xb1,
	0x67, 0x90, 0x0d, 0x7c, 0xb6, 0xa7, 0x01, 0xa7, 0x00, 0x5b, 0x1b, 0x30, 0xea, 0x05, 0xa6, 0x7e,
	0x8a, 0x81, 0x54, 0x4c, 0xb1, 0x2d, 0xc7, 0xf3, 0xd1, 0x9e, 0x54, 0xa3, 0x04, 0xd1, 0x60, 0xf0,
	0x3d, 0x34, 0x56, 0xc3, 0xea, 0x02, 0xe0, 0x72, 0x23, 0x12, 0x80, 0x3a, 0x7e, 0x53, 0xcb, 0xc1,
	0xdb, 0x2b, 0xf6, 0x74, 0x66, 0xda, 0x16, 0xde, 0x7b, 0x71, 0x89, 0x8e, 0x77, 0x33, 0x51, 0x95,
	0xc1, 0x29, 0xb4, 0xb7, 0x67, 0xd1, 0x86, 0xc6, 0xa3, 0x23, 0x5c, 0x7d, 0x38, 0x22, 0x2e, 0x1e,
	0x81, 0x29, 0x93, 0x4b, 0x21, 0x8a, 0x0a, 0x8a, 0x54, 0xdb, 0xd3, 0xfd, 0x33, 0xe8, 0x3b, 0x86,
	0x4d, 0x08, 0xf3, 0x2c, 0x9b, 0x12, 0xe6, 0xea, 0x2e, 0xb5, 0xd8, 0x56, 0x3f, 0x1f, 0xc3, 0x87,
	0x7b, 0xf6, 0x27, 0x96, 0xed, 0xda, 0x3a, 0x15, 0x15, 0x8e, 0xe1, 0xe9, 0x1e, 0xc0, 0x08, 0xbb,
	0x76, 0x0c, 0x7d, 0x22, 0x0a, 0xff, 0x04, 0xb0, 0xe1, 0x8e, 0x0f, 0xfa, 0xc8, 0xd0, 0x6d, 0x84,
	0xda, 0x5b, 0x55, 0x34, 0xe8, 0x18, 0x97, 0x94, 0xe9, 0xde, 0xd8, 0x1a, 0xd2, 0x11, 0x51, 0xff,
	0x7c, 0x5b, 0x45, 0xe0, 0x61, 0x11, 0x73, 0xc9, 0x88, 0x18, 0xd6, 0x58, 0xfd, 0x0b, 0x83, 0x1d,
	0xa8, 0x5f, 0x31, 0x2a, 0xbe, 0xfe, 0x7e, 0x5b, 0xed, 0xff, 0x5e, 0x83, 0xe6, 0xda, 0xec, 0x90,
	0x77, 0x69, 0x8a, 0x1b, 0x23, 0x78, 0x09, 0x9d, 0x55, 0x6c, 0x4b, 0xbc, 0x9f, 0xec, 0x33, 0xcb,
	0xd5, 0xa2, 0x18, 0xe8, 0x77, 0xa0, 0xae, 0x12, 0x93, 0x18, 0xf5, 0x52, 0x28, 0x79, 0x47, 0x15,
	0x3b, 0xc9, 0x93, 0x12, 0x2c, 0xd4, 0xe9, 0xcf, 0xe7, 0xf1, 0x3d, 0x0f, 0xbc, 0x59, 0xbe, 0x98,
	0x4b, 0xdd, 0x37, 0x85, 0xee, 0x0b, 0xb1, 0xd6, 0xa4, 0x2a, 0xbb, 0x50, 0x2f, 0x55, 0x5a, 0x5f,
	0xa9, 0x34, 0x5c, 0x24, 0x29, 0xcf, 0xc4, 0xcb, 0xf1, 0xa2, 0xe5, 0xa2, 0x14, 0x35, 0xda, 0xe2,
	0x1d, 0x7f, 0x40, 0xeb, 0x0d, 0x32, 0xd4, 0xf3, 0x01, 0x12, 0x3b, 0x82, 0xc6, 0x74, 0xe6, 0x47,
	0x11, 0x3e, 0xa1, 0x96, 0x0c, 0xa0, 0x6b, 0xc5, 0x09, 0xde, 0xad, 0xa4, 0x09, 0xbb, 0xae, 0xb5,
	0xee, 0xd4, 0x12, 0xa8, 0x15, 0xc7, 0x30, 0xca, 0x79, 0xea, 0x4f, 0x45, 0xc7, 0x45, 0x66, 0x1b,
	0x4d, 0xe0, 0xbf, 0x38, 0xd2, 0x2d, 0xb0, 0xb4, 0xbd, 0x10, 0x3a, 0x8f, 0x86, 0x86, 0xc2, 0x1b,
	0xea, 0x8c, 0xa1, 0xcb, 0x88, 0x5f, 0x41, 0x87, 0x32, 0x97, 0xd8, 0x8e, 0x4b, 0x5d, 0xf4, 0xad,
	0x42, 0xbd, 0xd7, 0xd4, 0x24, 0x16, 0x7a, 0x1d, 0x02, 0x19, 0xea, 0xe4, 0x9a, 0xa8, 0xe2, 0x46,
	0x9b, 0x64, 0x3c, 0x24, 0xa6, 0x49, 0x4c, 0xb4, 0x3c, 0x14, 0xb2, 0x35, 0x21, 0x8c, 0xb2, 0x0b,
	0x74, 0x3d, 0xcc, 0x38, 0xb7, 0x8c, 0x2b, 0x47, 0x6d, 0x0c, 0x5e, 0x82, 0xba, 0x33, 0xe2, 0x8d,
	0xdf, 0x1d, 0x42, 0xeb, 0x9c, 0x62, 0x29, 0xd4, 0xa5, 0x78, 0x37, 0xe2, 0x0d, 0x09, 0xb3, 0x73,
	0x50, 0x7c, 0x5f, 0x42, 0x6b, 0xc3, 0x18, 0xed, 0x50, 0x1f, 0x8d, 0x30, 0x01, 0x4f, 0xa6, 0x65,
	0xa3, 0xe2, 0x51, 0x5d, 0x61, 0xa3, 0x88, 0x5d, 0x80, 0xfa, 0x2e, 0x47, 0x99, 0xc2, 0x7e, 0x96,
	0x12, 0xed, 0x32, 0xcb, 0x93, 0x9c, 0x74, 0x43, 0x68, 0xbc, 0x28, 0x34, 0xb4, 0xad, 0x1b, 0x07,
	0x55, 0x2e, 0x78, 0x98, 0x28, 0xe6, 0x91, 0xa5, 0x0b, 0x37, 0x47, 0x1e, 0xa6, 0xb0, 0x6b, 0xe4,
	0x51, 0x15, 0x1f, 0x63, 0xe2, 0x38, 0xfa, 0x05, 0x29, 0x4d, 0x5d, 0xa7, 0x23, 0xb5, 0x3e, 0xf8,
	0x11, 0x3a, 0x8f, 0x7e, 0xa8, 0x47, 0xd0, 0x66, 0xc4, 0xdd, 0x7a, 0x15, 0x78, 0xa4, 0x08, 0xdc,
	0xd0, 0x73, 0x5a, 0x14, 0x13, 0x5f, 0x67, 0x17, 0x58, 0xac, 0x5c, 0xbf, 0xb8, 0x28, 0x87, 0x89,
	0xeb, 0xaf, 0xb1, 0xd2, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xdf, 0xd6, 0x84, 0x88, 0x08,
	0x00, 0x00,
}