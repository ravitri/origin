// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/common/segments.proto

package common // import "google.golang.org/genproto/googleapis/ads/googleads/v0/common"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import enums "google.golang.org/genproto/googleapis/ads/googleads/v0/enums"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Segment only fields.
type Segments struct {
	// Ad network type.
	AdNetworkType enums.AdNetworkTypeEnum_AdNetworkType `protobuf:"varint,3,opt,name=ad_network_type,json=adNetworkType,proto3,enum=google.ads.googleads.v0.enums.AdNetworkTypeEnum_AdNetworkType" json:"ad_network_type,omitempty"`
	// Conversion attribution event type.
	ConversionAttributionEventType enums.ConversionAttributionEventTypeEnum_ConversionAttributionEventType `protobuf:"varint,2,opt,name=conversion_attribution_event_type,json=conversionAttributionEventType,proto3,enum=google.ads.googleads.v0.enums.ConversionAttributionEventTypeEnum_ConversionAttributionEventType" json:"conversion_attribution_event_type,omitempty"`
	// Date to which metrics apply.
	// yyyy-MM-dd format, e.g., 2018-04-17.
	Date *wrappers.StringValue `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	// Day of the week, e.g., MONDAY.
	DayOfWeek enums.DayOfWeekEnum_DayOfWeek `protobuf:"varint,5,opt,name=day_of_week,json=dayOfWeek,proto3,enum=google.ads.googleads.v0.enums.DayOfWeekEnum_DayOfWeek" json:"day_of_week,omitempty"`
	// Device to which metrics apply.
	Device enums.DeviceEnum_Device `protobuf:"varint,1,opt,name=device,proto3,enum=google.ads.googleads.v0.enums.DeviceEnum_Device" json:"device,omitempty"`
	// Hotel booking window in days.
	HotelBookingWindowDays *wrappers.Int64Value `protobuf:"bytes,6,opt,name=hotel_booking_window_days,json=hotelBookingWindowDays,proto3" json:"hotel_booking_window_days,omitempty"`
	// Hotel center ID.
	HotelCenterId *wrappers.Int64Value `protobuf:"bytes,7,opt,name=hotel_center_id,json=hotelCenterId,proto3" json:"hotel_center_id,omitempty"`
	// Hotel check-in date. Formatted as yyyy-MM-dd.
	HotelCheckInDate *wrappers.StringValue `protobuf:"bytes,8,opt,name=hotel_check_in_date,json=hotelCheckInDate,proto3" json:"hotel_check_in_date,omitempty"`
	// Hotel check-in day of week.
	HotelCheckInDayOfWeek enums.DayOfWeekEnum_DayOfWeek `protobuf:"varint,9,opt,name=hotel_check_in_day_of_week,json=hotelCheckInDayOfWeek,proto3,enum=google.ads.googleads.v0.enums.DayOfWeekEnum_DayOfWeek" json:"hotel_check_in_day_of_week,omitempty"`
	// Hotel city.
	HotelCity *wrappers.StringValue `protobuf:"bytes,10,opt,name=hotel_city,json=hotelCity,proto3" json:"hotel_city,omitempty"`
	// Hotel class.
	HotelClass *wrappers.Int32Value `protobuf:"bytes,11,opt,name=hotel_class,json=hotelClass,proto3" json:"hotel_class,omitempty"`
	// Hotel country.
	HotelCountry *wrappers.StringValue `protobuf:"bytes,12,opt,name=hotel_country,json=hotelCountry,proto3" json:"hotel_country,omitempty"`
	// Hotel date selection type.
	HotelDateSelectionType enums.HotelDateSelectionTypeEnum_HotelDateSelectionType `protobuf:"varint,13,opt,name=hotel_date_selection_type,json=hotelDateSelectionType,proto3,enum=google.ads.googleads.v0.enums.HotelDateSelectionTypeEnum_HotelDateSelectionType" json:"hotel_date_selection_type,omitempty"`
	// Hotel length of stay.
	HotelLengthOfStay *wrappers.Int32Value `protobuf:"bytes,14,opt,name=hotel_length_of_stay,json=hotelLengthOfStay,proto3" json:"hotel_length_of_stay,omitempty"`
	// Hotel state.
	HotelState *wrappers.StringValue `protobuf:"bytes,15,opt,name=hotel_state,json=hotelState,proto3" json:"hotel_state,omitempty"`
	// Hour of day as a number between 0 and 23, inclusive.
	Hour *wrappers.Int32Value `protobuf:"bytes,16,opt,name=hour,proto3" json:"hour,omitempty"`
	// Month as represented by the date of the first day of a month. Formatted as
	// yyyy-MM-dd.
	Month *wrappers.StringValue `protobuf:"bytes,17,opt,name=month,proto3" json:"month,omitempty"`
	// Month of the year, e.g., January.
	MonthOfYear enums.MonthOfYearEnum_MonthOfYear `protobuf:"varint,18,opt,name=month_of_year,json=monthOfYear,proto3,enum=google.ads.googleads.v0.enums.MonthOfYearEnum_MonthOfYear" json:"month_of_year,omitempty"`
	// Partner hotel ID.
	PartnerHotelId *wrappers.StringValue `protobuf:"bytes,19,opt,name=partner_hotel_id,json=partnerHotelId,proto3" json:"partner_hotel_id,omitempty"`
	// Placeholder type. This is only used with feed item metrics.
	PlaceholderType enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,20,opt,name=placeholder_type,json=placeholderType,proto3,enum=google.ads.googleads.v0.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_type,omitempty"`
	// Quarter as represented by the date of the first day of a quarter.
	// Uses the calendar year for quarters, e.g., the second quarter of 2018
	// starts on 2018-04-01. Formatted as yyyy-MM-dd.
	Quarter *wrappers.StringValue `protobuf:"bytes,21,opt,name=quarter,proto3" json:"quarter,omitempty"`
	// Match type of the keyword that triggered the ad, including variants.
	SearchTermMatchType enums.SearchTermMatchTypeEnum_SearchTermMatchType `protobuf:"varint,22,opt,name=search_term_match_type,json=searchTermMatchType,proto3,enum=google.ads.googleads.v0.enums.SearchTermMatchTypeEnum_SearchTermMatchType" json:"search_term_match_type,omitempty"`
	// Position of the ad.
	Slot enums.SlotEnum_Slot `protobuf:"varint,23,opt,name=slot,proto3,enum=google.ads.googleads.v0.enums.SlotEnum_Slot" json:"slot,omitempty"`
	// Week as defined as Monday through Sunday, and represented by the date of
	// Monday. Formatted as yyyy-MM-dd.
	Week *wrappers.StringValue `protobuf:"bytes,24,opt,name=week,proto3" json:"week,omitempty"`
	// Year, formatted as yyyy.
	Year                 *wrappers.Int32Value `protobuf:"bytes,25,opt,name=year,proto3" json:"year,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Segments) Reset()         { *m = Segments{} }
func (m *Segments) String() string { return proto.CompactTextString(m) }
func (*Segments) ProtoMessage()    {}
func (*Segments) Descriptor() ([]byte, []int) {
	return fileDescriptor_segments_6c768d677b74fad4, []int{0}
}
func (m *Segments) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Segments.Unmarshal(m, b)
}
func (m *Segments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Segments.Marshal(b, m, deterministic)
}
func (dst *Segments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Segments.Merge(dst, src)
}
func (m *Segments) XXX_Size() int {
	return xxx_messageInfo_Segments.Size(m)
}
func (m *Segments) XXX_DiscardUnknown() {
	xxx_messageInfo_Segments.DiscardUnknown(m)
}

var xxx_messageInfo_Segments proto.InternalMessageInfo

func (m *Segments) GetAdNetworkType() enums.AdNetworkTypeEnum_AdNetworkType {
	if m != nil {
		return m.AdNetworkType
	}
	return enums.AdNetworkTypeEnum_UNSPECIFIED
}

func (m *Segments) GetConversionAttributionEventType() enums.ConversionAttributionEventTypeEnum_ConversionAttributionEventType {
	if m != nil {
		return m.ConversionAttributionEventType
	}
	return enums.ConversionAttributionEventTypeEnum_UNSPECIFIED
}

func (m *Segments) GetDate() *wrappers.StringValue {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Segments) GetDayOfWeek() enums.DayOfWeekEnum_DayOfWeek {
	if m != nil {
		return m.DayOfWeek
	}
	return enums.DayOfWeekEnum_UNSPECIFIED
}

func (m *Segments) GetDevice() enums.DeviceEnum_Device {
	if m != nil {
		return m.Device
	}
	return enums.DeviceEnum_UNSPECIFIED
}

func (m *Segments) GetHotelBookingWindowDays() *wrappers.Int64Value {
	if m != nil {
		return m.HotelBookingWindowDays
	}
	return nil
}

func (m *Segments) GetHotelCenterId() *wrappers.Int64Value {
	if m != nil {
		return m.HotelCenterId
	}
	return nil
}

func (m *Segments) GetHotelCheckInDate() *wrappers.StringValue {
	if m != nil {
		return m.HotelCheckInDate
	}
	return nil
}

func (m *Segments) GetHotelCheckInDayOfWeek() enums.DayOfWeekEnum_DayOfWeek {
	if m != nil {
		return m.HotelCheckInDayOfWeek
	}
	return enums.DayOfWeekEnum_UNSPECIFIED
}

func (m *Segments) GetHotelCity() *wrappers.StringValue {
	if m != nil {
		return m.HotelCity
	}
	return nil
}

func (m *Segments) GetHotelClass() *wrappers.Int32Value {
	if m != nil {
		return m.HotelClass
	}
	return nil
}

func (m *Segments) GetHotelCountry() *wrappers.StringValue {
	if m != nil {
		return m.HotelCountry
	}
	return nil
}

func (m *Segments) GetHotelDateSelectionType() enums.HotelDateSelectionTypeEnum_HotelDateSelectionType {
	if m != nil {
		return m.HotelDateSelectionType
	}
	return enums.HotelDateSelectionTypeEnum_UNSPECIFIED
}

func (m *Segments) GetHotelLengthOfStay() *wrappers.Int32Value {
	if m != nil {
		return m.HotelLengthOfStay
	}
	return nil
}

func (m *Segments) GetHotelState() *wrappers.StringValue {
	if m != nil {
		return m.HotelState
	}
	return nil
}

func (m *Segments) GetHour() *wrappers.Int32Value {
	if m != nil {
		return m.Hour
	}
	return nil
}

func (m *Segments) GetMonth() *wrappers.StringValue {
	if m != nil {
		return m.Month
	}
	return nil
}

func (m *Segments) GetMonthOfYear() enums.MonthOfYearEnum_MonthOfYear {
	if m != nil {
		return m.MonthOfYear
	}
	return enums.MonthOfYearEnum_UNSPECIFIED
}

func (m *Segments) GetPartnerHotelId() *wrappers.StringValue {
	if m != nil {
		return m.PartnerHotelId
	}
	return nil
}

func (m *Segments) GetPlaceholderType() enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderType
	}
	return enums.PlaceholderTypeEnum_UNSPECIFIED
}

func (m *Segments) GetQuarter() *wrappers.StringValue {
	if m != nil {
		return m.Quarter
	}
	return nil
}

func (m *Segments) GetSearchTermMatchType() enums.SearchTermMatchTypeEnum_SearchTermMatchType {
	if m != nil {
		return m.SearchTermMatchType
	}
	return enums.SearchTermMatchTypeEnum_UNSPECIFIED
}

func (m *Segments) GetSlot() enums.SlotEnum_Slot {
	if m != nil {
		return m.Slot
	}
	return enums.SlotEnum_UNSPECIFIED
}

func (m *Segments) GetWeek() *wrappers.StringValue {
	if m != nil {
		return m.Week
	}
	return nil
}

func (m *Segments) GetYear() *wrappers.Int32Value {
	if m != nil {
		return m.Year
	}
	return nil
}

func init() {
	proto.RegisterType((*Segments)(nil), "google.ads.googleads.v0.common.Segments")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/common/segments.proto", fileDescriptor_segments_6c768d677b74fad4)
}

var fileDescriptor_segments_6c768d677b74fad4 = []byte{
	// 966 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xdf, 0x6e, 0xdb, 0x36,
	0x14, 0xc6, 0xe1, 0x34, 0x4d, 0x1b, 0xba, 0x4e, 0x5c, 0xa6, 0xcd, 0xd8, 0x6c, 0x08, 0xba, 0x5c,
	0x05, 0xc3, 0x26, 0x79, 0x49, 0x91, 0x0b, 0x75, 0x1d, 0xe6, 0x38, 0x59, 0x9b, 0xad, 0x5d, 0x02,
	0xbb, 0x70, 0xb1, 0x21, 0x98, 0xc6, 0x48, 0xb4, 0x2c, 0x58, 0x22, 0x1d, 0x92, 0xb6, 0xa1, 0xab,
	0xdd, 0xec, 0x6e, 0x6f, 0xb1, 0xcb, 0xed, 0x4d, 0xf6, 0x24, 0xc3, 0x9e, 0x62, 0x20, 0x29, 0xc9,
	0xce, 0x3f, 0x49, 0xc0, 0xae, 0x4c, 0xd2, 0xdf, 0xef, 0x3b, 0x87, 0x87, 0x47, 0xa2, 0xc0, 0x17,
	0x01, 0x63, 0x41, 0x44, 0x6c, 0xec, 0x0b, 0xdb, 0x0c, 0xd5, 0x68, 0xda, 0xb2, 0x3d, 0x16, 0xc7,
	0x8c, 0xda, 0x82, 0x04, 0x31, 0xa1, 0x52, 0x58, 0x63, 0xce, 0x24, 0x83, 0xdb, 0x46, 0x63, 0x61,
	0x5f, 0x58, 0xb9, 0xdc, 0x9a, 0xb6, 0x2c, 0x23, 0xdf, 0xda, 0xbf, 0xcb, 0x8e, 0xd0, 0x49, 0x2c,
	0x6c, 0xec, 0xbb, 0x94, 0xc8, 0x19, 0xe3, 0x23, 0x57, 0x26, 0x63, 0x62, 0x4c, 0xb7, 0x8e, 0x8b,
	0x21, 0x8f, 0xd1, 0x29, 0xe1, 0x22, 0x64, 0xd4, 0xc5, 0x52, 0xf2, 0xf0, 0x62, 0x22, 0xd5, 0x98,
	0x4c, 0x09, 0x95, 0x8b, 0x36, 0x76, 0xb1, 0x8d, 0x8f, 0x13, 0x97, 0x0d, 0xdc, 0x19, 0x21, 0xa3,
	0x14, 0xf8, 0xac, 0x04, 0x20, 0xd3, 0xd0, 0xcb, 0xcc, 0x5f, 0x15, 0x6b, 0x87, 0x4c, 0x92, 0xc8,
	0xf5, 0xb1, 0x24, 0xae, 0x20, 0x11, 0xf1, 0x74, 0x86, 0x0b, 0xb9, 0x7d, 0x59, 0x8c, 0xc7, 0x8c,
	0xca, 0xa1, 0xca, 0x2e, 0x21, 0x98, 0xa7, 0xc8, 0x8b, 0x62, 0x64, 0x1c, 0x61, 0x8f, 0x0c, 0x59,
	0xe4, 0x13, 0xbe, 0x18, 0xc8, 0x29, 0xa6, 0x04, 0xc1, 0xdc, 0x1b, 0xba, 0x92, 0xf0, 0xd8, 0x8d,
	0xb1, 0x54, 0xc3, 0x39, 0xbb, 0x5b, 0xc2, 0x46, 0x4c, 0xa6, 0xca, 0xb4, 0x0d, 0x6c, 0x3d, 0xbb,
	0x98, 0x0c, 0xec, 0x19, 0xc7, 0xe3, 0x31, 0xe1, 0x69, 0x9b, 0xec, 0xfc, 0xd6, 0x04, 0x0f, 0x7b,
	0x69, 0xe7, 0xc0, 0x01, 0x58, 0xbf, 0x76, 0xee, 0xe8, 0xde, 0xf3, 0xda, 0xee, 0xda, 0xde, 0xd7,
	0xd6, 0x5d, 0xdd, 0xa4, 0x03, 0x5a, 0x6d, 0xff, 0x07, 0x03, 0xbd, 0x4f, 0xc6, 0xe4, 0x98, 0x4e,
	0xe2, 0xab, 0x2b, 0xdd, 0x06, 0x5e, 0x9c, 0xc2, 0xbf, 0x6a, 0xe0, 0xd3, 0xd2, 0x5e, 0x41, 0x4b,
	0x3a, 0xf4, 0x2f, 0x25, 0xa1, 0x3b, 0xb9, 0x4f, 0x7b, 0x6e, 0x73, 0xac, 0x5c, 0xf2, 0x5c, 0x8a,
	0x25, 0xdd, 0x6d, 0xaf, 0xf0, 0x7f, 0xd8, 0x02, 0xcb, 0xaa, 0x5d, 0xd0, 0xf2, 0xf3, 0xda, 0x6e,
	0x7d, 0xef, 0x93, 0x2c, 0x9f, 0xac, 0xa2, 0x56, 0x4f, 0xf2, 0x90, 0x06, 0x7d, 0x1c, 0x4d, 0x48,
	0x57, 0x2b, 0x61, 0x1f, 0xd4, 0x17, 0x7a, 0x18, 0xdd, 0xd7, 0x1b, 0x39, 0x28, 0xd9, 0xc8, 0x11,
	0x4e, 0x4e, 0x07, 0x1f, 0x08, 0x19, 0xe9, 0x9c, 0xf3, 0x59, 0x77, 0xd5, 0xcf, 0x86, 0xf0, 0x0d,
	0x58, 0x31, 0xad, 0x8e, 0x6a, 0xda, 0xb2, 0x55, 0x66, 0xa9, 0xc5, 0xc6, 0x4f, 0x0f, 0xbb, 0x29,
	0x0f, 0xfb, 0xe0, 0x99, 0x79, 0x10, 0x2e, 0x18, 0x1b, 0x85, 0x34, 0x70, 0x67, 0x21, 0xf5, 0xd9,
	0xcc, 0xf5, 0x71, 0x22, 0xd0, 0x8a, 0xde, 0xe8, 0xc7, 0x37, 0x36, 0x7a, 0x42, 0xe5, 0xc1, 0x0b,
	0xb3, 0xcf, 0x4d, 0x4d, 0x1f, 0x1a, 0xf8, 0x83, 0x66, 0x8f, 0x70, 0x22, 0x60, 0x07, 0xac, 0x1b,
	0x5f, 0x8f, 0x50, 0x49, 0xb8, 0x1b, 0xfa, 0xe8, 0x41, 0xb9, 0x5b, 0x43, 0x33, 0x1d, 0x8d, 0x9c,
	0xf8, 0xf0, 0x7b, 0xb0, 0x91, 0x9a, 0x0c, 0x89, 0x37, 0x72, 0x43, 0xaa, 0x1f, 0x57, 0xf4, 0xb0,
	0x42, 0xfd, 0x9b, 0xc6, 0x49, 0x71, 0x27, 0xf4, 0x48, 0x9d, 0xc5, 0x25, 0xd8, 0xba, 0x61, 0x36,
	0x3f, 0x9a, 0xd5, 0xff, 0x75, 0x34, 0x4f, 0xaf, 0x46, 0xcb, 0x8e, 0xe9, 0x25, 0x00, 0x69, 0xc8,
	0x50, 0x26, 0x08, 0x54, 0x48, 0x7b, 0xd5, 0x18, 0x85, 0x32, 0x81, 0x5f, 0x81, 0x7a, 0x0a, 0x47,
	0x58, 0x08, 0x54, 0xbf, 0xbb, 0x7a, 0xfb, 0x7b, 0x06, 0x36, 0xc1, 0x3a, 0x4a, 0x0e, 0xdb, 0xa0,
	0x91, 0xd2, 0x6c, 0x42, 0x25, 0x4f, 0xd0, 0xa3, 0x0a, 0xd1, 0x1f, 0x19, 0x03, 0x43, 0xc0, 0xdf,
	0x6b, 0x59, 0x6f, 0xdc, 0xf2, 0x92, 0x44, 0x0d, 0x5d, 0xb0, 0xb3, 0x92, 0x82, 0xbd, 0x51, 0xbc,
	0x2a, 0x7f, 0x2f, 0xa3, 0xf3, 0x87, 0xf1, 0xf6, 0xbf, 0xd2, 0x86, 0xba, 0xb1, 0x0e, 0xdf, 0x82,
	0x27, 0x26, 0x99, 0x88, 0xd0, 0xc0, 0xbc, 0x79, 0x85, 0xc4, 0x09, 0x5a, 0x2b, 0xaf, 0xcb, 0x63,
	0x0d, 0xbe, 0xd5, 0xdc, 0xe9, 0xa0, 0x27, 0x71, 0x02, 0x5f, 0x65, 0xc5, 0x15, 0x52, 0x75, 0xd4,
	0x7a, 0x85, 0xe2, 0x98, 0xea, 0xf6, 0x94, 0x1e, 0xda, 0x60, 0x79, 0xc8, 0x26, 0x1c, 0x35, 0xcb,
	0x83, 0x6b, 0x21, 0xdc, 0x03, 0xf7, 0xf5, 0x85, 0x81, 0x1e, 0x57, 0x88, 0x64, 0xa4, 0xf0, 0x67,
	0xd0, 0xb8, 0x72, 0xc9, 0x20, 0xa8, 0x4b, 0xee, 0x94, 0x94, 0xfc, 0x9d, 0x62, 0x4e, 0x07, 0x3f,
	0x12, 0xcc, 0x75, 0x9d, 0x17, 0xe6, 0xdd, 0x7a, 0x3c, 0x9f, 0xc0, 0x6f, 0x41, 0x73, 0x8c, 0xb9,
	0xa4, 0x84, 0xbb, 0xa6, 0x16, 0xa1, 0x8f, 0x36, 0x2a, 0xa4, 0xb7, 0x96, 0x52, 0xfa, 0x00, 0x4f,
	0x7c, 0x18, 0x83, 0xe6, 0xf5, 0x9b, 0x0d, 0x3d, 0xd1, 0xa9, 0x1e, 0x96, 0xa4, 0x7a, 0x36, 0xc7,
	0xf2, 0xb6, 0xb8, 0xb6, 0xd6, 0x5d, 0x1f, 0x5f, 0x5d, 0x80, 0x07, 0xe0, 0xc1, 0xe5, 0x04, 0x73,
	0x49, 0x38, 0x7a, 0x5a, 0x21, 0xdb, 0x4c, 0x0c, 0x7f, 0x05, 0x9b, 0xb7, 0x5f, 0xa5, 0x68, 0x53,
	0x27, 0xfb, 0x5d, 0x49, 0xb2, 0x3d, 0x0d, 0xbf, 0x27, 0x3c, 0x7e, 0xa7, 0xd0, 0x3c, 0xe1, 0x5b,
	0xd6, 0xbb, 0x1b, 0xe2, 0xe6, 0x22, 0xfc, 0x06, 0x2c, 0xab, 0xfb, 0x18, 0x7d, 0xa4, 0xc3, 0x7d,
	0x5e, 0x16, 0x2e, 0x62, 0xd2, 0xf8, 0x47, 0x4c, 0x76, 0x35, 0xa9, 0x2e, 0x20, 0xfd, 0xb2, 0x42,
	0x55, 0x2e, 0x20, 0xa5, 0x54, 0x8d, 0xaa, 0x5b, 0xe7, 0x59, 0x85, 0x46, 0x55, 0xc2, 0xc3, 0x7f,
	0x6a, 0x60, 0xc7, 0x63, 0xb1, 0x55, 0xfc, 0xd1, 0x78, 0xd8, 0xc8, 0x3e, 0x15, 0xce, 0x94, 0xd3,
	0x59, 0xed, 0xa7, 0xa3, 0x14, 0x08, 0x58, 0x84, 0x69, 0x60, 0x31, 0x1e, 0xd8, 0x01, 0xa1, 0x3a,
	0x4e, 0xf6, 0x61, 0x32, 0x0e, 0xc5, 0x5d, 0xdf, 0xac, 0x2f, 0xcd, 0xcf, 0x1f, 0x4b, 0xf7, 0x5e,
	0xb7, 0xdb, 0x7f, 0x2e, 0x6d, 0xbf, 0x36, 0x66, 0x6d, 0x5f, 0x58, 0x66, 0xa8, 0x46, 0xfd, 0x96,
	0xd5, 0xd1, 0xb2, 0xbf, 0x33, 0xc1, 0x79, 0xdb, 0x17, 0xe7, 0xb9, 0xe0, 0xbc, 0xdf, 0x3a, 0x37,
	0x82, 0x7f, 0x97, 0x76, 0xcc, 0xaa, 0xe3, 0xb4, 0x7d, 0xe1, 0x38, 0xb9, 0xc4, 0x71, 0xfa, 0x2d,
	0xc7, 0x31, 0xa2, 0x8b, 0x15, 0x9d, 0xdd, 0xfe, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x67,
	0x13, 0xd9, 0x50, 0x0b, 0x00, 0x00,
}
