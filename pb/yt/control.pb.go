// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: control.proto

package yt

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//协议类型
type ProtocolType int32

const (
	ProtocolType_CONNECT       ProtocolType = 0
	ProtocolType_CONNECTACK    ProtocolType = 1
	ProtocolType_SUB           ProtocolType = 2
	ProtocolType_SUBACK        ProtocolType = 3
	ProtocolType_UNSUB         ProtocolType = 4
	ProtocolType_UNSUBACK      ProtocolType = 5
	ProtocolType_HOLDMIC       ProtocolType = 6
	ProtocolType_HOLDMICACK    ProtocolType = 7
	ProtocolType_RELEASEMIC    ProtocolType = 8
	ProtocolType_RELEASEMICACK ProtocolType = 9
	ProtocolType_DISCONNECT    ProtocolType = 10
	ProtocolType_DISCONNECTACK ProtocolType = 11
	ProtocolType_REGISTER      ProtocolType = 12
	ProtocolType_REGISTERACK   ProtocolType = 13
)

var ProtocolType_name = map[int32]string{
	0:  "CONNECT",
	1:  "CONNECTACK",
	2:  "SUB",
	3:  "SUBACK",
	4:  "UNSUB",
	5:  "UNSUBACK",
	6:  "HOLDMIC",
	7:  "HOLDMICACK",
	8:  "RELEASEMIC",
	9:  "RELEASEMICACK",
	10: "DISCONNECT",
	11: "DISCONNECTACK",
	12: "REGISTER",
	13: "REGISTERACK",
}

var ProtocolType_value = map[string]int32{
	"CONNECT":       0,
	"CONNECTACK":    1,
	"SUB":           2,
	"SUBACK":        3,
	"UNSUB":         4,
	"UNSUBACK":      5,
	"HOLDMIC":       6,
	"HOLDMICACK":    7,
	"RELEASEMIC":    8,
	"RELEASEMICACK": 9,
	"DISCONNECT":    10,
	"DISCONNECTACK": 11,
	"REGISTER":      12,
	"REGISTERACK":   13,
}

func (x ProtocolType) String() string {
	return proto.EnumName(ProtocolType_name, int32(x))
}

func (ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{0}
}

type Control struct {
	ProtocolType         ProtocolType `protobuf:"varint,1,opt,name=protocolType,proto3,enum=yt.ProtocolType" json:"protocolType,omitempty"`
	UserID               uint32       `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	TopicID              uint32       `protobuf:"varint,3,opt,name=topicID,proto3" json:"topicID,omitempty"`
	AckCode              uint32       `protobuf:"varint,5,opt,name=ackCode,proto3" json:"ackCode,omitempty"`
	Subgroup             string       `protobuf:"bytes,7,opt,name=subgroup,proto3" json:"subgroup,omitempty"`
	Priority             uint32       `protobuf:"varint,8,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Control) Reset()         { *m = Control{} }
func (m *Control) String() string { return proto.CompactTextString(m) }
func (*Control) ProtoMessage()    {}
func (*Control) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c5120591600887d, []int{0}
}
func (m *Control) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Control) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Control.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Control) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Control.Merge(m, src)
}
func (m *Control) XXX_Size() int {
	return m.Size()
}
func (m *Control) XXX_DiscardUnknown() {
	xxx_messageInfo_Control.DiscardUnknown(m)
}

var xxx_messageInfo_Control proto.InternalMessageInfo

func (m *Control) GetProtocolType() ProtocolType {
	if m != nil {
		return m.ProtocolType
	}
	return ProtocolType_CONNECT
}

func (m *Control) GetUserID() uint32 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *Control) GetTopicID() uint32 {
	if m != nil {
		return m.TopicID
	}
	return 0
}

func (m *Control) GetAckCode() uint32 {
	if m != nil {
		return m.AckCode
	}
	return 0
}

func (m *Control) GetSubgroup() string {
	if m != nil {
		return m.Subgroup
	}
	return ""
}

func (m *Control) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterEnum("yt.ProtocolType", ProtocolType_name, ProtocolType_value)
	proto.RegisterType((*Control)(nil), "yt.Control")
}

func init() { proto.RegisterFile("control.proto", fileDescriptor_0c5120591600887d) }

var fileDescriptor_0c5120591600887d = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0x19, 0xb8, 0xb4, 0xe5, 0xd0, 0x72, 0xc7, 0xb3, 0x30, 0x13, 0x17, 0x0d, 0x71, 0x45,
	0x5c, 0xb0, 0x50, 0x5f, 0x00, 0xda, 0x46, 0x1b, 0x11, 0xcc, 0x14, 0x1e, 0x40, 0x2a, 0x31, 0x44,
	0xe2, 0x34, 0xc3, 0xb0, 0xe8, 0x9b, 0xf8, 0x3e, 0x6e, 0x5c, 0xba, 0x77, 0x63, 0xf0, 0x45, 0xcc,
	0x29, 0x2d, 0xe0, 0x6e, 0xbe, 0xff, 0xfb, 0xe7, 0xcc, 0x49, 0x06, 0xbc, 0x54, 0xbd, 0x1a, 0xad,
	0x56, 0xfd, 0x4c, 0x2b, 0xa3, 0xb0, 0x9e, 0x9b, 0xf3, 0x77, 0x06, 0x76, 0xb0, 0x4b, 0xf1, 0x1a,
	0xdc, 0x42, 0xa4, 0x6a, 0x35, 0xcd, 0xb3, 0x85, 0x60, 0x5d, 0xd6, 0xeb, 0x5c, 0xf2, 0x7e, 0x6e,
	0xfa, 0x0f, 0x47, 0xb9, 0xfc, 0xd3, 0xc2, 0x53, 0xb0, 0x36, 0xeb, 0x85, 0x8e, 0x43, 0x51, 0xef,
	0xb2, 0x9e, 0x27, 0x4b, 0x42, 0x01, 0xb6, 0x51, 0xd9, 0x32, 0x8d, 0x43, 0xd1, 0x28, 0x44, 0x85,
	0x64, 0x1e, 0xd3, 0x97, 0x40, 0x3d, 0x2d, 0x44, 0x73, 0x67, 0x4a, 0xc4, 0x33, 0x70, 0xd6, 0x9b,
	0xf9, 0xb3, 0x56, 0x9b, 0x4c, 0xd8, 0x5d, 0xd6, 0x6b, 0xc9, 0x3d, 0x93, 0xcb, 0xf4, 0x52, 0xe9,
	0xa5, 0xc9, 0x85, 0x53, 0x5c, 0xdb, 0xf3, 0xc5, 0x17, 0x03, 0xf7, 0x78, 0x45, 0x6c, 0x83, 0x1d,
	0x4c, 0xc6, 0xe3, 0x28, 0x98, 0xf2, 0x1a, 0x76, 0x00, 0x4a, 0x18, 0x04, 0x77, 0x9c, 0xa1, 0x0d,
	0x8d, 0x64, 0x36, 0xe4, 0x75, 0x04, 0xb0, 0x92, 0xd9, 0x90, 0xc2, 0x06, 0xb6, 0xa0, 0x39, 0x1b,
	0x53, 0xfc, 0x0f, 0x5d, 0x70, 0x8a, 0x23, 0x89, 0x26, 0x8d, 0xba, 0x9d, 0x8c, 0xc2, 0xfb, 0x38,
	0xe0, 0x16, 0x8d, 0x2a, 0x81, 0xa4, 0x4d, 0x2c, 0xa3, 0x51, 0x34, 0x48, 0x22, 0xf2, 0x0e, 0x9e,
	0x80, 0x77, 0x60, 0xaa, 0xb4, 0xa8, 0x12, 0xc6, 0x49, 0xb5, 0x0d, 0x50, 0xe5, 0xc0, 0x54, 0x69,
	0xd3, 0x83, 0x32, 0xba, 0x89, 0x93, 0x69, 0x24, 0xb9, 0x8b, 0xff, 0xa1, 0x5d, 0x11, 0x69, 0x6f,
	0xc8, 0x3f, 0xb6, 0x3e, 0xfb, 0xdc, 0xfa, 0xec, 0x7b, 0xeb, 0xb3, 0xb7, 0x1f, 0xbf, 0x36, 0xb7,
	0x8a, 0x1f, 0xb8, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x82, 0x07, 0xab, 0xd1, 0x01, 0x00,
	0x00,
}

func (m *Control) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Control) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ProtocolType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintControl(dAtA, i, uint64(m.ProtocolType))
	}
	if m.UserID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintControl(dAtA, i, uint64(m.UserID))
	}
	if m.TopicID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintControl(dAtA, i, uint64(m.TopicID))
	}
	if m.AckCode != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintControl(dAtA, i, uint64(m.AckCode))
	}
	if len(m.Subgroup) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintControl(dAtA, i, uint64(len(m.Subgroup)))
		i += copy(dAtA[i:], m.Subgroup)
	}
	if m.Priority != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintControl(dAtA, i, uint64(m.Priority))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintControl(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Control) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProtocolType != 0 {
		n += 1 + sovControl(uint64(m.ProtocolType))
	}
	if m.UserID != 0 {
		n += 1 + sovControl(uint64(m.UserID))
	}
	if m.TopicID != 0 {
		n += 1 + sovControl(uint64(m.TopicID))
	}
	if m.AckCode != 0 {
		n += 1 + sovControl(uint64(m.AckCode))
	}
	l = len(m.Subgroup)
	if l > 0 {
		n += 1 + l + sovControl(uint64(l))
	}
	if m.Priority != 0 {
		n += 1 + sovControl(uint64(m.Priority))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovControl(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozControl(x uint64) (n int) {
	return sovControl(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Control) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowControl
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Control: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Control: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolType", wireType)
			}
			m.ProtocolType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProtocolType |= (ProtocolType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			m.UserID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TopicID", wireType)
			}
			m.TopicID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TopicID |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckCode", wireType)
			}
			m.AckCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AckCode |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subgroup", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthControl
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subgroup = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowControl
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipControl(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthControl
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipControl(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowControl
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
					return 0, ErrIntOverflowControl
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowControl
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthControl
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowControl
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipControl(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthControl = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowControl   = fmt.Errorf("proto: integer overflow")
)