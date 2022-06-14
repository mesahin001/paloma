// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/common.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ArbitrarySmartContractCall struct {
	Method     string `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Payload    []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	HexAddress string `protobuf:"bytes,3,opt,name=hexAddress,proto3" json:"hexAddress,omitempty"`
	Abi        []byte `protobuf:"bytes,4,opt,name=abi,proto3" json:"abi,omitempty"`
	// TODO: make a new type for turnstone messages
	// TODO: add TurnstoneID
	ViaTurnstone bool   `protobuf:"varint,5,opt,name=viaTurnstone,proto3" json:"viaTurnstone,omitempty"`
	Deadline     uint64 `protobuf:"varint,6,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (m *ArbitrarySmartContractCall) Reset()         { *m = ArbitrarySmartContractCall{} }
func (m *ArbitrarySmartContractCall) String() string { return proto.CompactTextString(m) }
func (*ArbitrarySmartContractCall) ProtoMessage()    {}
func (*ArbitrarySmartContractCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ef836bbf7879f0e, []int{0}
}
func (m *ArbitrarySmartContractCall) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ArbitrarySmartContractCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ArbitrarySmartContractCall.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ArbitrarySmartContractCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArbitrarySmartContractCall.Merge(m, src)
}
func (m *ArbitrarySmartContractCall) XXX_Size() int {
	return m.Size()
}
func (m *ArbitrarySmartContractCall) XXX_DiscardUnknown() {
	xxx_messageInfo_ArbitrarySmartContractCall.DiscardUnknown(m)
}

var xxx_messageInfo_ArbitrarySmartContractCall proto.InternalMessageInfo

func (m *ArbitrarySmartContractCall) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ArbitrarySmartContractCall) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ArbitrarySmartContractCall) GetHexAddress() string {
	if m != nil {
		return m.HexAddress
	}
	return ""
}

func (m *ArbitrarySmartContractCall) GetAbi() []byte {
	if m != nil {
		return m.Abi
	}
	return nil
}

func (m *ArbitrarySmartContractCall) GetViaTurnstone() bool {
	if m != nil {
		return m.ViaTurnstone
	}
	return false
}

func (m *ArbitrarySmartContractCall) GetDeadline() uint64 {
	if m != nil {
		return m.Deadline
	}
	return 0
}

func init() {
	proto.RegisterType((*ArbitrarySmartContractCall)(nil), "palomachain.paloma.evm.ArbitrarySmartContractCall")
}

func init() { proto.RegisterFile("evm/common.proto", fileDescriptor_2ef836bbf7879f0e) }

var fileDescriptor_2ef836bbf7879f0e = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0x6b, 0x5a, 0x4a, 0xb1, 0x3a, 0x54, 0x16, 0xaa, 0xac, 0x0e, 0x56, 0xd4, 0x29, 0x2c,
	0xc9, 0xc0, 0x09, 0x4a, 0x6f, 0x10, 0x98, 0xd8, 0xfe, 0xc4, 0x56, 0x62, 0x29, 0xf6, 0x1f, 0x39,
	0x6e, 0xd4, 0xdc, 0x82, 0x23, 0x31, 0x32, 0x76, 0x64, 0x44, 0xc9, 0x45, 0x50, 0xd2, 0x82, 0xca,
	0xf6, 0xbd, 0xa7, 0xf7, 0x96, 0x8f, 0xae, 0x54, 0x63, 0xe2, 0x0c, 0x8d, 0x41, 0x1b, 0x55, 0x0e,
	0x3d, 0xb2, 0x75, 0x05, 0x25, 0x1a, 0xc8, 0x0a, 0xd0, 0x36, 0x3a, 0x73, 0xa4, 0x1a, 0xb3, 0x79,
	0xc8, 0x31, 0xc7, 0x71, 0x12, 0x0f, 0x74, 0x5e, 0x6f, 0x3f, 0x08, 0xdd, 0xec, 0x5c, 0xaa, 0xbd,
	0x03, 0xd7, 0xbe, 0x18, 0x70, 0x7e, 0x8f, 0xd6, 0x3b, 0xc8, 0xfc, 0x1e, 0xca, 0x92, 0xad, 0xe9,
	0xdc, 0x28, 0x5f, 0xa0, 0xe4, 0x24, 0x20, 0xe1, 0x7d, 0x72, 0x49, 0x8c, 0xd3, 0xbb, 0x0a, 0xda,
	0x12, 0x41, 0xf2, 0x9b, 0x80, 0x84, 0xcb, 0xe4, 0x37, 0x32, 0x41, 0x69, 0xa1, 0x8e, 0x3b, 0x29,
	0x9d, 0xaa, 0x6b, 0x3e, 0x1d, 0x5f, 0x57, 0x0d, 0x5b, 0xd1, 0x29, 0xa4, 0x9a, 0xcf, 0xc6, 0xd7,
	0x80, 0x6c, 0x4b, 0x97, 0x8d, 0x86, 0xd7, 0x83, 0xb3, 0xb5, 0x47, 0xab, 0xf8, 0x6d, 0x40, 0xc2,
	0x45, 0xf2, 0xaf, 0x63, 0x1b, 0xba, 0x90, 0x0a, 0x64, 0xa9, 0xad, 0xe2, 0xf3, 0x80, 0x84, 0xb3,
	0xe4, 0x2f, 0x3f, 0xef, 0x3f, 0x3b, 0x41, 0x4e, 0x9d, 0x20, 0xdf, 0x9d, 0x20, 0xef, 0xbd, 0x98,
	0x9c, 0x7a, 0x31, 0xf9, 0xea, 0xc5, 0xe4, 0xed, 0x31, 0xd7, 0xbe, 0x38, 0xa4, 0x51, 0x86, 0x26,
	0xbe, 0xb2, 0x72, 0xe1, 0xf8, 0x18, 0x0f, 0xf2, 0x7c, 0x5b, 0xa9, 0x3a, 0x9d, 0x8f, 0x3a, 0x9e,
	0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x95, 0x01, 0x2c, 0x10, 0x50, 0x01, 0x00, 0x00,
}

func (m *ArbitrarySmartContractCall) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArbitrarySmartContractCall) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ArbitrarySmartContractCall) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Deadline != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Deadline))
		i--
		dAtA[i] = 0x30
	}
	if m.ViaTurnstone {
		i--
		if m.ViaTurnstone {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.Abi) > 0 {
		i -= len(m.Abi)
		copy(dAtA[i:], m.Abi)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Abi)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.HexAddress) > 0 {
		i -= len(m.HexAddress)
		copy(dAtA[i:], m.HexAddress)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.HexAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Method) > 0 {
		i -= len(m.Method)
		copy(dAtA[i:], m.Method)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Method)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ArbitrarySmartContractCall) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.HexAddress)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Abi)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.ViaTurnstone {
		n += 2
	}
	if m.Deadline != 0 {
		n += 1 + sovCommon(uint64(m.Deadline))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ArbitrarySmartContractCall) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: ArbitrarySmartContractCall: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArbitrarySmartContractCall: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HexAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HexAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abi", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Abi = append(m.Abi[:0], dAtA[iNdEx:postIndex]...)
			if m.Abi == nil {
				m.Abi = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ViaTurnstone", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ViaTurnstone = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			m.Deadline = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Deadline |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
