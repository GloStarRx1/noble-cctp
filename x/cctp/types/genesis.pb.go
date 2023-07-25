// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: noble/cctp/v1/genesis.proto

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

// GenesisState defines the cctp module's genesis state.
type GenesisState struct {
	Params                            Params                             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Authority                         *Authority                         `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
	AttesterList                      []Attester                         `protobuf:"bytes,3,rep,name=attester_list,json=attesterList,proto3" json:"attester_list"`
	PerMessageBurnLimitList           []PerMessageBurnLimit              `protobuf:"bytes,4,rep,name=per_message_burn_limit_list,json=perMessageBurnLimitList,proto3" json:"per_message_burn_limit_list"`
	BurningAndMintingPaused           *BurningAndMintingPaused           `protobuf:"bytes,5,opt,name=burning_and_minting_paused,json=burningAndMintingPaused,proto3" json:"burning_and_minting_paused,omitempty"`
	SendingAndReceivingMessagesPaused *SendingAndReceivingMessagesPaused `protobuf:"bytes,6,opt,name=sending_and_receiving_messages_paused,json=sendingAndReceivingMessagesPaused,proto3" json:"sending_and_receiving_messages_paused,omitempty"`
	MaxMessageBodySize                *MaxMessageBodySize                `protobuf:"bytes,7,opt,name=max_message_body_size,json=maxMessageBodySize,proto3" json:"max_message_body_size,omitempty"`
	NextAvailableNonce                *Nonce                             `protobuf:"bytes,8,opt,name=next_available_nonce,json=nextAvailableNonce,proto3" json:"next_available_nonce,omitempty"`
	SignatureThreshold                *SignatureThreshold                `protobuf:"bytes,9,opt,name=signature_threshold,json=signatureThreshold,proto3" json:"signature_threshold,omitempty"`
	TokenPairList                     []TokenPair                        `protobuf:"bytes,10,rep,name=token_pair_list,json=tokenPairList,proto3" json:"token_pair_list"`
	UsedNoncesList                    []Nonce                            `protobuf:"bytes,11,rep,name=used_nonces_list,json=usedNoncesList,proto3" json:"used_nonces_list"`
	TokenMessengerList                []TokenMessenger                   `protobuf:"bytes,12,rep,name=token_messenger_list,json=tokenMessengerList,proto3" json:"token_messenger_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_45b63e7866d42c8d, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetAuthority() *Authority {
	if m != nil {
		return m.Authority
	}
	return nil
}

func (m *GenesisState) GetAttesterList() []Attester {
	if m != nil {
		return m.AttesterList
	}
	return nil
}

func (m *GenesisState) GetPerMessageBurnLimitList() []PerMessageBurnLimit {
	if m != nil {
		return m.PerMessageBurnLimitList
	}
	return nil
}

func (m *GenesisState) GetBurningAndMintingPaused() *BurningAndMintingPaused {
	if m != nil {
		return m.BurningAndMintingPaused
	}
	return nil
}

func (m *GenesisState) GetSendingAndReceivingMessagesPaused() *SendingAndReceivingMessagesPaused {
	if m != nil {
		return m.SendingAndReceivingMessagesPaused
	}
	return nil
}

func (m *GenesisState) GetMaxMessageBodySize() *MaxMessageBodySize {
	if m != nil {
		return m.MaxMessageBodySize
	}
	return nil
}

func (m *GenesisState) GetNextAvailableNonce() *Nonce {
	if m != nil {
		return m.NextAvailableNonce
	}
	return nil
}

func (m *GenesisState) GetSignatureThreshold() *SignatureThreshold {
	if m != nil {
		return m.SignatureThreshold
	}
	return nil
}

func (m *GenesisState) GetTokenPairList() []TokenPair {
	if m != nil {
		return m.TokenPairList
	}
	return nil
}

func (m *GenesisState) GetUsedNoncesList() []Nonce {
	if m != nil {
		return m.UsedNoncesList
	}
	return nil
}

func (m *GenesisState) GetTokenMessengerList() []TokenMessenger {
	if m != nil {
		return m.TokenMessengerList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "noble.cctp.v1.GenesisState")
}

func init() { proto.RegisterFile("noble/cctp/v1/genesis.proto", fileDescriptor_45b63e7866d42c8d) }

var fileDescriptor_45b63e7866d42c8d = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x4f, 0xd4, 0x4e,
	0x18, 0xdf, 0xfd, 0xc3, 0x1f, 0x64, 0x00, 0x35, 0xe3, 0x12, 0xea, 0x22, 0x15, 0x30, 0x2a, 0x1a,
	0xed, 0x0a, 0x24, 0x26, 0x1e, 0x77, 0x63, 0xf0, 0xc2, 0x1a, 0xd2, 0xc5, 0x8b, 0x97, 0x66, 0xda,
	0x0e, 0xdd, 0x89, 0xed, 0xb4, 0x99, 0x99, 0x6e, 0x76, 0x39, 0xfa, 0x09, 0xfc, 0x58, 0x1c, 0x39,
	0x7a, 0x32, 0x06, 0x2e, 0x7e, 0x0c, 0xd3, 0x79, 0x81, 0x74, 0xa8, 0xd1, 0x5b, 0xfb, 0xfc, 0xde,
	0xe6, 0x99, 0x79, 0x66, 0xc0, 0x06, 0xcd, 0xc3, 0x14, 0xf7, 0xa2, 0x48, 0x14, 0xbd, 0xc9, 0x5e,
	0x2f, 0xc1, 0x14, 0x73, 0xc2, 0xbd, 0x82, 0xe5, 0x22, 0x87, 0xab, 0x12, 0xf4, 0x2a, 0xd0, 0x9b,
	0xec, 0x75, 0x37, 0xeb, 0x5c, 0x54, 0x8a, 0x71, 0xce, 0x88, 0x98, 0x29, 0x76, 0xd7, 0xab, 0xc3,
	0x61, 0xc9, 0x28, 0xa1, 0x49, 0x80, 0x68, 0x1c, 0x64, 0x84, 0x8a, 0xea, 0xbb, 0x40, 0x25, 0xc7,
	0xb1, 0xe6, 0xbf, 0xa8, 0xf3, 0x33, 0x34, 0x0d, 0x32, 0xcc, 0x39, 0x4a, 0x70, 0x10, 0xe6, 0xf1,
	0x2c, 0xe0, 0xe4, 0x0c, 0x6b, 0xea, 0xc3, 0x3a, 0x95, 0xe6, 0x34, 0x32, 0x50, 0xb7, 0x0e, 0x15,
	0x88, 0xa1, 0x4c, 0xaf, 0xbf, 0xfb, 0xd2, 0xc2, 0x30, 0xbb, 0x49, 0x28, 0x19, 0x0d, 0x52, 0x92,
	0x11, 0xa1, 0xb9, 0x8f, 0xac, 0xe6, 0x84, 0xc0, 0x5c, 0x60, 0xa6, 0xd1, 0x77, 0x75, 0x94, 0x63,
	0x1a, 0x9b, 0xde, 0x18, 0x8e, 0x30, 0x99, 0x54, 0x7f, 0xda, 0x9b, 0xd7, 0xdb, 0x7c, 0x6e, 0x49,
	0x49, 0x42, 0x91, 0x28, 0x19, 0x0e, 0xc4, 0x98, 0x61, 0x3e, 0xce, 0x53, 0x43, 0x74, 0xeb, 0x44,
	0x91, 0x7f, 0xc1, 0x34, 0x28, 0x10, 0x31, 0x6b, 0x78, 0xd2, 0x84, 0x57, 0x99, 0x98, 0x26, 0xd7,
	0x0b, 0xed, 0x24, 0x79, 0x92, 0xcb, 0xcf, 0x5e, 0xf5, 0xa5, 0xaa, 0x3b, 0xbf, 0x16, 0xc1, 0xca,
	0x07, 0x75, 0xb4, 0x23, 0x81, 0x04, 0x86, 0x07, 0x60, 0x41, 0xed, 0x94, 0xd3, 0xde, 0x6a, 0xef,
	0x2e, 0xef, 0xaf, 0x79, 0xb5, 0xa3, 0xf6, 0x8e, 0x25, 0x38, 0x98, 0x3f, 0xff, 0xf1, 0xb8, 0xe5,
	0x6b, 0x2a, 0x7c, 0x0b, 0x96, 0xae, 0xcf, 0xdc, 0xf9, 0x4f, 0xea, 0x1c, 0x4b, 0xd7, 0x37, 0xb8,
	0x7f, 0x43, 0x85, 0x03, 0xb0, 0x6a, 0xb6, 0x33, 0x48, 0x09, 0x17, 0xce, 0xdc, 0xd6, 0xdc, 0xee,
	0xf2, 0xfe, 0xba, 0xad, 0xd5, 0x1c, 0x9d, 0xba, 0x62, 0x34, 0x47, 0x84, 0x0b, 0x78, 0x0a, 0x36,
	0x9a, 0x8f, 0x4f, 0x39, 0xce, 0x4b, 0xc7, 0x1d, 0xbb, 0x0b, 0xcc, 0x86, 0x4a, 0x30, 0x28, 0x19,
	0x3d, 0xaa, 0xe8, 0xda, 0x7c, 0xbd, 0xb8, 0x0d, 0xc9, 0x9c, 0x08, 0x74, 0xff, 0x3c, 0xb8, 0xce,
	0xff, 0xb2, 0xe9, 0x67, 0x56, 0xcc, 0x40, 0x09, 0xfa, 0x34, 0x1e, 0x2a, 0xfa, 0xb1, 0x64, 0xfb,
	0xeb, 0x61, 0x33, 0x00, 0xbf, 0xb6, 0xc1, 0xd3, 0x7f, 0x1a, 0x21, 0x67, 0x41, 0x06, 0xbe, 0xb1,
	0x02, 0x47, 0x4a, 0xdb, 0xa7, 0xb1, 0x6f, 0x94, 0xba, 0x19, 0xae, 0xa3, 0xb7, 0xf9, 0xdf, 0x28,
	0xf0, 0x04, 0xac, 0x35, 0x5e, 0x39, 0x67, 0x51, 0x66, 0x6e, 0x5b, 0x99, 0x43, 0x34, 0x35, 0x1b,
	0x96, 0xc7, 0xb3, 0x11, 0x39, 0xc3, 0x3e, 0xcc, 0x6e, 0xd5, 0xe0, 0x21, 0xe8, 0x50, 0x3c, 0x15,
	0x01, 0x9a, 0x20, 0x92, 0xa2, 0x30, 0xc5, 0x81, 0xbc, 0xac, 0xce, 0x1d, 0x69, 0xda, 0xb1, 0x4c,
	0x3f, 0x56, 0x98, 0x0f, 0x2b, 0x45, 0xdf, 0x08, 0x64, 0x0d, 0xfa, 0xe0, 0x41, 0xc3, 0x4d, 0x71,
	0x96, 0x1a, 0xd7, 0x36, 0x32, 0xcc, 0x13, 0x43, 0xf4, 0x21, 0xbf, 0x55, 0x83, 0x87, 0xe0, 0xde,
	0xcd, 0xa5, 0x52, 0x73, 0x03, 0xe4, 0xdc, 0xd8, 0x53, 0x7c, 0x52, 0xb1, 0x8e, 0x11, 0x31, 0xa3,
	0xb8, 0x2a, 0x4c, 0x41, 0xce, 0xc8, 0x7b, 0x70, 0xbf, 0xda, 0x41, 0xd5, 0x19, 0x57, 0x46, 0xcb,
	0xd2, 0xa8, 0xb1, 0x3f, 0x6d, 0x72, 0xb7, 0xd2, 0xc8, 0x02, 0x97, 0x2e, 0x9f, 0x40, 0xc7, 0xba,
	0xc2, 0xca, 0x69, 0x45, 0x3a, 0x6d, 0x36, 0x2d, 0x69, 0x68, 0x98, 0xda, 0x12, 0x8a, 0x5a, 0xb5,
	0xb2, 0x1d, 0x1c, 0x9e, 0x5f, 0xba, 0xed, 0x8b, 0x4b, 0xb7, 0xfd, 0xf3, 0xd2, 0x6d, 0x7f, 0xbb,
	0x72, 0x5b, 0x17, 0x57, 0x6e, 0xeb, 0xfb, 0x95, 0xdb, 0xfa, 0xfc, 0x2a, 0x21, 0x62, 0x5c, 0x86,
	0x5e, 0x94, 0x67, 0xbd, 0x88, 0xb0, 0x28, 0xc5, 0xa7, 0x84, 0xf6, 0x64, 0xcc, 0x6b, 0xf9, 0xa8,
	0x4c, 0xd5, 0xdb, 0x22, 0x66, 0x05, 0xe6, 0xe1, 0x82, 0x7c, 0x39, 0x0e, 0x7e, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xf4, 0xf0, 0xdd, 0x4a, 0x21, 0x06, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenMessengerList) > 0 {
		for iNdEx := len(m.TokenMessengerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenMessengerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if len(m.UsedNoncesList) > 0 {
		for iNdEx := len(m.UsedNoncesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UsedNoncesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.TokenPairList) > 0 {
		for iNdEx := len(m.TokenPairList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenPairList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.SignatureThreshold != nil {
		{
			size, err := m.SignatureThreshold.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.NextAvailableNonce != nil {
		{
			size, err := m.NextAvailableNonce.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.MaxMessageBodySize != nil {
		{
			size, err := m.MaxMessageBodySize.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.SendingAndReceivingMessagesPaused != nil {
		{
			size, err := m.SendingAndReceivingMessagesPaused.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.BurningAndMintingPaused != nil {
		{
			size, err := m.BurningAndMintingPaused.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PerMessageBurnLimitList) > 0 {
		for iNdEx := len(m.PerMessageBurnLimitList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PerMessageBurnLimitList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.AttesterList) > 0 {
		for iNdEx := len(m.AttesterList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AttesterList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Authority != nil {
		{
			size, err := m.Authority.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.Authority != nil {
		l = m.Authority.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.AttesterList) > 0 {
		for _, e := range m.AttesterList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PerMessageBurnLimitList) > 0 {
		for _, e := range m.PerMessageBurnLimitList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.BurningAndMintingPaused != nil {
		l = m.BurningAndMintingPaused.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.SendingAndReceivingMessagesPaused != nil {
		l = m.SendingAndReceivingMessagesPaused.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MaxMessageBodySize != nil {
		l = m.MaxMessageBodySize.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.NextAvailableNonce != nil {
		l = m.NextAvailableNonce.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.SignatureThreshold != nil {
		l = m.SignatureThreshold.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.TokenPairList) > 0 {
		for _, e := range m.TokenPairList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UsedNoncesList) > 0 {
		for _, e := range m.UsedNoncesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TokenMessengerList) > 0 {
		for _, e := range m.TokenMessengerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authority == nil {
				m.Authority = &Authority{}
			}
			if err := m.Authority.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttesterList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttesterList = append(m.AttesterList, Attester{})
			if err := m.AttesterList[len(m.AttesterList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerMessageBurnLimitList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PerMessageBurnLimitList = append(m.PerMessageBurnLimitList, PerMessageBurnLimit{})
			if err := m.PerMessageBurnLimitList[len(m.PerMessageBurnLimitList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurningAndMintingPaused", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BurningAndMintingPaused == nil {
				m.BurningAndMintingPaused = &BurningAndMintingPaused{}
			}
			if err := m.BurningAndMintingPaused.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendingAndReceivingMessagesPaused", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SendingAndReceivingMessagesPaused == nil {
				m.SendingAndReceivingMessagesPaused = &SendingAndReceivingMessagesPaused{}
			}
			if err := m.SendingAndReceivingMessagesPaused.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMessageBodySize", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxMessageBodySize == nil {
				m.MaxMessageBodySize = &MaxMessageBodySize{}
			}
			if err := m.MaxMessageBodySize.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextAvailableNonce", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NextAvailableNonce == nil {
				m.NextAvailableNonce = &Nonce{}
			}
			if err := m.NextAvailableNonce.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignatureThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SignatureThreshold == nil {
				m.SignatureThreshold = &SignatureThreshold{}
			}
			if err := m.SignatureThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPairList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenPairList = append(m.TokenPairList, TokenPair{})
			if err := m.TokenPairList[len(m.TokenPairList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsedNoncesList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UsedNoncesList = append(m.UsedNoncesList, Nonce{})
			if err := m.UsedNoncesList[len(m.UsedNoncesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenMessengerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenMessengerList = append(m.TokenMessengerList, TokenMessenger{})
			if err := m.TokenMessengerList[len(m.TokenMessengerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)