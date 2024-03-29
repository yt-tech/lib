// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: auth.proto

//包名，通过protoc生成时go文件时

package auth

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

//账号
type AuthInfo struct {
	Userid               string   `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Anchor               int32    `protobuf:"varint,3,opt,name=anchor,proto3" json:"anchor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthInfo) Reset()         { *m = AuthInfo{} }
func (m *AuthInfo) String() string { return proto.CompactTextString(m) }
func (*AuthInfo) ProtoMessage()    {}
func (*AuthInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}
func (m *AuthInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthInfo.Merge(m, src)
}
func (m *AuthInfo) XXX_Size() int {
	return m.Size()
}
func (m *AuthInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AuthInfo proto.InternalMessageInfo

func (m *AuthInfo) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *AuthInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AuthInfo) GetAnchor() int32 {
	if m != nil {
		return m.Anchor
	}
	return 0
}

type UpdateInfo struct {
	Table                string   `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Anchor               int32    `protobuf:"varint,3,opt,name=anchor,proto3" json:"anchor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateInfo) Reset()         { *m = UpdateInfo{} }
func (m *UpdateInfo) String() string { return proto.CompactTextString(m) }
func (*UpdateInfo) ProtoMessage()    {}
func (*UpdateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}
func (m *UpdateInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInfo.Merge(m, src)
}
func (m *UpdateInfo) XXX_Size() int {
	return m.Size()
}
func (m *UpdateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInfo proto.InternalMessageInfo

func (m *UpdateInfo) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *UpdateInfo) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *UpdateInfo) GetAnchor() int32 {
	if m != nil {
		return m.Anchor
	}
	return 0
}

func init() {
	proto.RegisterType((*AuthInfo)(nil), "auth.authInfo")
	proto.RegisterType((*UpdateInfo)(nil), "auth.updateInfo")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x82, 0xb8, 0x38, 0x40, 0xb4,
	0x67, 0x5e, 0x5a, 0xbe, 0x90, 0x18, 0x17, 0x5b, 0x69, 0x71, 0x6a, 0x51, 0x66, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x94, 0x07, 0x12, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x96, 0x60,
	0x52, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0xf2, 0x40, 0xe2, 0x89, 0x79, 0xc9, 0x19, 0xf9, 0x45, 0x12,
	0xcc, 0x10, 0x71, 0x08, 0x4f, 0xc9, 0x8f, 0x8b, 0xab, 0xb4, 0x20, 0x25, 0xb1, 0x24, 0x15, 0x6c,
	0xaa, 0x08, 0x17, 0x6b, 0x49, 0x62, 0x52, 0x4e, 0x2a, 0xd4, 0x50, 0x08, 0x47, 0x48, 0x88, 0x8b,
	0x25, 0x25, 0xb1, 0x24, 0x11, 0x6c, 0x22, 0x67, 0x10, 0x98, 0x8d, 0xcb, 0x3c, 0x27, 0x81, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc6, 0x63, 0x39, 0x86,
	0x24, 0x36, 0xb0, 0x17, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xfb, 0x3f, 0xd8, 0xd0,
	0x00, 0x00, 0x00,
}

func (m *AuthInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Userid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Userid)))
		i += copy(dAtA[i:], m.Userid)
	}
	if m.Status != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Status))
	}
	if m.Anchor != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Anchor))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *UpdateInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Table) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Table)))
		i += copy(dAtA[i:], m.Table)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAuth(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.Anchor != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAuth(dAtA, i, uint64(m.Anchor))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAuth(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *AuthInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Userid)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovAuth(uint64(m.Status))
	}
	if m.Anchor != 0 {
		n += 1 + sovAuth(uint64(m.Anchor))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpdateInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Table)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	if m.Anchor != 0 {
		n += 1 + sovAuth(uint64(m.Anchor))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAuth(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAuth(x uint64) (n int) {
	return sovAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuthInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: authInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: authInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Userid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Anchor", wireType)
			}
			m.Anchor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Anchor |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
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
func (m *UpdateInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: updateInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: updateInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Table", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Table = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Anchor", wireType)
			}
			m.Anchor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Anchor |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuth
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
func skipAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
				return 0, ErrInvalidLengthAuth
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAuth
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
				next, err := skipAuth(dAtA[start:])
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
	ErrInvalidLengthAuth = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuth   = fmt.Errorf("proto: integer overflow")
)
