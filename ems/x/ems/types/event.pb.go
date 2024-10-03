// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ems/ems/event.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Event struct {
	EventName        string `protobuf:"bytes,1,opt,name=eventName,proto3" json:"eventName,omitempty"`
	EventDescription string `protobuf:"bytes,2,opt,name=eventDescription,proto3" json:"eventDescription,omitempty"`
	StartDate        string `protobuf:"bytes,3,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate          string `protobuf:"bytes,4,opt,name=endDate,proto3" json:"endDate,omitempty"`
	Location         string `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	NumFtTickets     uint64 `protobuf:"varint,6,opt,name=numFtTickets,proto3" json:"numFtTickets,omitempty"`
	Organizer        string `protobuf:"bytes,7,opt,name=organizer,proto3" json:"organizer,omitempty"`
	Creator          string `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
	Id               uint64 `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	TicketPrice      uint64 `protobuf:"varint,10,opt,name=ticketPrice,proto3" json:"ticketPrice,omitempty"`
	TicketsLeft      uint64 `protobuf:"varint,11,opt,name=ticketsLeft,proto3" json:"ticketsLeft,omitempty"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_c885920ccb4cf9f9, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Event.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return m.Size()
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *Event) GetEventDescription() string {
	if m != nil {
		return m.EventDescription
	}
	return ""
}

func (m *Event) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *Event) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

func (m *Event) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Event) GetNumFtTickets() uint64 {
	if m != nil {
		return m.NumFtTickets
	}
	return 0
}

func (m *Event) GetOrganizer() string {
	if m != nil {
		return m.Organizer
	}
	return ""
}

func (m *Event) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Event) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetTicketPrice() uint64 {
	if m != nil {
		return m.TicketPrice
	}
	return 0
}

func (m *Event) GetTicketsLeft() uint64 {
	if m != nil {
		return m.TicketsLeft
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "ems.ems.Event")
}

func init() { proto.RegisterFile("ems/ems/event.proto", fileDescriptor_c885920ccb4cf9f9) }

var fileDescriptor_c885920ccb4cf9f9 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xeb, 0xd0, 0x36, 0x8d, 0x8b, 0x00, 0x99, 0xc5, 0x42, 0xc8, 0x8a, 0x3a, 0x15, 0x86,
	0x32, 0xf0, 0x0f, 0x50, 0x61, 0x42, 0x08, 0x55, 0x4c, 0x6c, 0x26, 0x39, 0x90, 0x05, 0x89, 0x23,
	0xfb, 0x40, 0xc0, 0xaf, 0xe0, 0x67, 0x31, 0x76, 0x41, 0x62, 0x44, 0xc9, 0x1f, 0x41, 0xbe, 0x40,
	0x5a, 0xc4, 0x70, 0xc3, 0xfb, 0xde, 0xd3, 0x7b, 0xc3, 0xf1, 0x5d, 0x28, 0xfc, 0x11, 0xdd, 0x13,
	0x94, 0x38, 0xab, 0x9c, 0x45, 0x2b, 0x62, 0x28, 0xfc, 0x0c, 0x0a, 0x3f, 0xf9, 0x88, 0xf8, 0xe0,
	0x34, 0x18, 0x62, 0x9f, 0x27, 0x94, 0xb8, 0xd0, 0x05, 0x48, 0x96, 0xb2, 0x69, 0xb2, 0x58, 0x01,
	0x71, 0xc8, 0x77, 0x48, 0xcc, 0xc1, 0x67, 0xce, 0x54, 0x68, 0x6c, 0x29, 0x23, 0x0a, 0xfd, 0xe3,
	0xa1, 0xc9, 0xa3, 0x76, 0x38, 0xd7, 0x08, 0x72, 0xa3, 0x6d, 0xea, 0x80, 0x90, 0x3c, 0x86, 0x32,
	0x27, 0xaf, 0x4f, 0xde, 0xaf, 0x14, 0x7b, 0x7c, 0xf4, 0x60, 0x33, 0x4d, 0xdd, 0x03, 0xb2, 0x3a,
	0x2d, 0x26, 0x7c, 0xb3, 0x7c, 0x2c, 0xce, 0xf0, 0xca, 0x64, 0xf7, 0x80, 0x5e, 0x0e, 0x53, 0x36,
	0xed, 0x2f, 0xfe, 0xb0, 0xb0, 0x6b, 0xdd, 0x9d, 0x2e, 0xcd, 0x2b, 0x38, 0x19, 0xb7, 0xbb, 0x1d,
	0x08, 0xbb, 0x99, 0x03, 0x8d, 0xd6, 0xc9, 0x51, 0xbb, 0xfb, 0x23, 0xc5, 0x16, 0x8f, 0x4c, 0x2e,
	0x13, 0x6a, 0x8c, 0x4c, 0x2e, 0x52, 0x3e, 0x46, 0xaa, 0xbc, 0x74, 0x26, 0x03, 0xc9, 0xc9, 0x58,
	0x47, 0xab, 0x84, 0x3f, 0x87, 0x5b, 0x94, 0xe3, 0xf5, 0x04, 0xa1, 0x93, 0x83, 0xf7, 0x5a, 0xb1,
	0x65, 0xad, 0xd8, 0x57, 0xad, 0xd8, 0x5b, 0xa3, 0x7a, 0xcb, 0x46, 0xf5, 0x3e, 0x1b, 0xd5, 0xbb,
	0xde, 0x0e, 0xbf, 0x78, 0xa6, 0x8f, 0xe0, 0x4b, 0x05, 0xfe, 0x66, 0x48, 0x2f, 0x39, 0xfe, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x42, 0x42, 0x2c, 0x3a, 0xa9, 0x01, 0x00, 0x00,
}

func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Event) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TicketsLeft != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.TicketsLeft))
		i--
		dAtA[i] = 0x58
	}
	if m.TicketPrice != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.TicketPrice))
		i--
		dAtA[i] = 0x50
	}
	if m.Id != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Organizer) > 0 {
		i -= len(m.Organizer)
		copy(dAtA[i:], m.Organizer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Organizer)))
		i--
		dAtA[i] = 0x3a
	}
	if m.NumFtTickets != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NumFtTickets))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Location) > 0 {
		i -= len(m.Location)
		copy(dAtA[i:], m.Location)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Location)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.EndDate) > 0 {
		i -= len(m.EndDate)
		copy(dAtA[i:], m.EndDate)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.EndDate)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.StartDate) > 0 {
		i -= len(m.StartDate)
		copy(dAtA[i:], m.StartDate)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.StartDate)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EventDescription) > 0 {
		i -= len(m.EventDescription)
		copy(dAtA[i:], m.EventDescription)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.EventDescription)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EventName) > 0 {
		i -= len(m.EventName)
		copy(dAtA[i:], m.EventName)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.EventName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Event) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EventName)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.EventDescription)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.StartDate)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.EndDate)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.NumFtTickets != 0 {
		n += 1 + sovEvent(uint64(m.NumFtTickets))
	}
	l = len(m.Organizer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovEvent(uint64(m.Id))
	}
	if m.TicketPrice != 0 {
		n += 1 + sovEvent(uint64(m.TicketPrice))
	}
	if m.TicketsLeft != 0 {
		n += 1 + sovEvent(uint64(m.TicketsLeft))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventDescription", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventDescription = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumFtTickets", wireType)
			}
			m.NumFtTickets = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumFtTickets |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Organizer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Organizer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TicketPrice", wireType)
			}
			m.TicketPrice = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TicketPrice |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TicketsLeft", wireType)
			}
			m.TicketsLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TicketsLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)