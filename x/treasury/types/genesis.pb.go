// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/treasury/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// GenesisState defines the oracle module's genesis state.
type GenesisState struct {
	Params               Params                                   `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	TaxRate              github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,2,opt,name=tax_rate,json=taxRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tax_rate"`
	RewardWeight         github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,3,opt,name=reward_weight,json=rewardWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reward_weight"`
	TaxCaps              []TaxCap                                 `protobuf:"bytes,4,rep,name=tax_caps,json=taxCaps,proto3" json:"tax_caps"`
	TaxProceeds          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=tax_proceeds,json=taxProceeds,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"tax_proceeds"`
	EpochInitialIssuance github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=epoch_initial_issuance,json=epochInitialIssuance,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"epoch_initial_issuance"`
	EpochState           []EpochState                             `protobuf:"bytes,7,rep,name=epoch_state,json=epochState,proto3" json:"epoch_state"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c440a3f50aabab34, []int{0}
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

func (m *GenesisState) GetTaxCaps() []TaxCap {
	if m != nil {
		return m.TaxCaps
	}
	return nil
}

func (m *GenesisState) GetTaxProceeds() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TaxProceeds
	}
	return nil
}

func (m *GenesisState) GetEpochInitialIssuance() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.EpochInitialIssuance
	}
	return nil
}

func (m *GenesisState) GetEpochState() []EpochState {
	if m != nil {
		return m.EpochState
	}
	return nil
}

// TaxCap is the max tax amount can be charged for the given denom
type TaxCap struct {
	Denom  string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	TaxCap github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=tax_cap,json=taxCap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"tax_cap"`
}

func (m *TaxCap) Reset()         { *m = TaxCap{} }
func (m *TaxCap) String() string { return proto.CompactTextString(m) }
func (*TaxCap) ProtoMessage()    {}
func (*TaxCap) Descriptor() ([]byte, []int) {
	return fileDescriptor_c440a3f50aabab34, []int{1}
}
func (m *TaxCap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaxCap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaxCap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaxCap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaxCap.Merge(m, src)
}
func (m *TaxCap) XXX_Size() int {
	return m.Size()
}
func (m *TaxCap) XXX_DiscardUnknown() {
	xxx_messageInfo_TaxCap.DiscardUnknown(m)
}

var xxx_messageInfo_TaxCap proto.InternalMessageInfo

func (m *TaxCap) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

// EpochState is the record for each epoch state
type EpochState struct {
	Epoch             uint64                                 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	TaxReward         github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=tax_reward,json=taxReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tax_reward"`
	SeigniorageReward github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=seigniorage_reward,json=seigniorageReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"seigniorage_reward"`
	TotalStakedLuna   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=total_staked_luna,json=totalStakedLuna,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_staked_luna"`
}

func (m *EpochState) Reset()         { *m = EpochState{} }
func (m *EpochState) String() string { return proto.CompactTextString(m) }
func (*EpochState) ProtoMessage()    {}
func (*EpochState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c440a3f50aabab34, []int{2}
}
func (m *EpochState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochState.Merge(m, src)
}
func (m *EpochState) XXX_Size() int {
	return m.Size()
}
func (m *EpochState) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochState.DiscardUnknown(m)
}

var xxx_messageInfo_EpochState proto.InternalMessageInfo

func (m *EpochState) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "terra.treasury.v1beta1.GenesisState")
	proto.RegisterType((*TaxCap)(nil), "terra.treasury.v1beta1.TaxCap")
	proto.RegisterType((*EpochState)(nil), "terra.treasury.v1beta1.EpochState")
}

func init() {
	proto.RegisterFile("terra/treasury/v1beta1/genesis.proto", fileDescriptor_c440a3f50aabab34)
}

var fileDescriptor_c440a3f50aabab34 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x18, 0x86, 0xb3, 0x6d, 0x9a, 0xda, 0x69, 0x44, 0x3a, 0x84, 0xb2, 0xf6, 0xb0, 0x09, 0x41, 0x25,
	0x97, 0xce, 0x5a, 0xbd, 0x0a, 0x42, 0xa2, 0xd4, 0x80, 0x42, 0xd9, 0x08, 0x42, 0x41, 0xc2, 0x97,
	0xcd, 0xc7, 0x66, 0x6c, 0x32, 0xb3, 0xcc, 0x4c, 0x6c, 0x7a, 0xf4, 0x1f, 0xf8, 0x3b, 0xc4, 0x1f,
	0xd2, 0x63, 0x8f, 0xe2, 0xa1, 0x4a, 0x72, 0xf5, 0x47, 0xc8, 0xcc, 0x6c, 0xd3, 0x1e, 0xac, 0x48,
	0xf0, 0x94, 0xec, 0xee, 0x3b, 0xcf, 0xfb, 0xee, 0xf7, 0xbd, 0x2c, 0x79, 0x60, 0x50, 0x29, 0x88,
	0x8d, 0x42, 0xd0, 0x53, 0x75, 0x16, 0x7f, 0x3c, 0x18, 0xa0, 0x81, 0x83, 0x38, 0x43, 0x81, 0x9a,
	0x6b, 0x96, 0x2b, 0x69, 0x24, 0xdd, 0x75, 0x2a, 0x76, 0xa5, 0x62, 0x85, 0x6a, 0xaf, 0x96, 0xc9,
	0x4c, 0x3a, 0x49, 0x6c, 0xff, 0x79, 0xf5, 0xde, 0xc3, 0x5b, 0x98, 0xcb, 0xe3, 0x5e, 0x16, 0xa5,
	0x52, 0x4f, 0xa4, 0x8e, 0x07, 0xa0, 0x71, 0xa9, 0x49, 0x25, 0x17, 0xfe, 0x79, 0xf3, 0x57, 0x99,
	0x54, 0x0f, 0x7d, 0x8c, 0x9e, 0x01, 0x83, 0xf4, 0x19, 0xa9, 0xe4, 0xa0, 0x60, 0xa2, 0xc3, 0xa0,
	0x11, 0xb4, 0xb6, 0x9f, 0x44, 0xec, 0xcf, 0xb1, 0xd8, 0x91, 0x53, 0xb5, 0xcb, 0xe7, 0x97, 0xf5,
	0x52, 0x52, 0x9c, 0xa1, 0x5d, 0x72, 0xc7, 0xc0, 0xac, 0xaf, 0xc0, 0x60, 0xb8, 0xd6, 0x08, 0x5a,
	0x5b, 0x6d, 0x66, 0x9f, 0x7f, 0xbf, 0xac, 0x3f, 0xca, 0xb8, 0x19, 0x4d, 0x07, 0x2c, 0x95, 0x93,
	0xb8, 0xc8, 0xe4, 0x7f, 0xf6, 0xf5, 0xf0, 0x24, 0x36, 0x67, 0x39, 0x6a, 0xf6, 0x02, 0xd3, 0x64,
	0xd3, 0xc0, 0x2c, 0xb1, 0x41, 0x7a, 0xe4, 0xae, 0xc2, 0x53, 0x50, 0xc3, 0xfe, 0x29, 0xf2, 0x6c,
	0x64, 0xc2, 0xf5, 0x95, 0x78, 0x55, 0x0f, 0x79, 0xe7, 0x18, 0xf4, 0xb9, 0xcf, 0x97, 0x42, 0xae,
	0xc3, 0x72, 0x63, 0xfd, 0x6f, 0xef, 0xf7, 0x16, 0x66, 0x1d, 0xc8, 0x8b, 0xf7, 0xb3, 0xa9, 0x3a,
	0x90, 0x6b, 0x2a, 0x48, 0xd5, 0x02, 0x72, 0x25, 0x53, 0xc4, 0xa1, 0x0e, 0x37, 0x1c, 0xe4, 0x3e,
	0xf3, 0xde, 0xcc, 0x8e, 0x79, 0x49, 0xe8, 0x48, 0x2e, 0xda, 0x8f, 0xed, 0xf9, 0x2f, 0x3f, 0xea,
	0xad, 0x7f, 0xc8, 0x6b, 0x0f, 0xe8, 0x64, 0xdb, 0xc0, 0xec, 0xa8, 0xe0, 0xd3, 0x4f, 0x01, 0xd9,
	0xc5, 0x5c, 0xa6, 0xa3, 0x3e, 0x17, 0xdc, 0x70, 0x18, 0xf7, 0xb9, 0xd6, 0x53, 0x10, 0x29, 0x86,
	0x95, 0xff, 0x6f, 0x5d, 0x73, 0x56, 0x5d, 0xef, 0xd4, 0x2d, 0x8c, 0x68, 0x97, 0x6c, 0xfb, 0x08,
	0xda, 0x36, 0x24, 0xdc, 0x74, 0xbe, 0xcd, 0xdb, 0xe6, 0xf6, 0xd2, 0x4a, 0x5d, 0x97, 0x8a, 0xd9,
	0x11, 0x5c, 0xde, 0x69, 0x66, 0xa4, 0xe2, 0xe7, 0x4a, 0x6b, 0x64, 0x63, 0x88, 0x42, 0x4e, 0x5c,
	0xcd, 0xb6, 0x12, 0x7f, 0x41, 0x0f, 0xc9, 0x66, 0xb1, 0x9f, 0x15, 0xea, 0xd3, 0x15, 0x26, 0xa9,
	0xf8, 0x45, 0x35, 0xbf, 0xae, 0x11, 0x72, 0x9d, 0xc4, 0xba, 0xb9, 0x14, 0xce, 0xad, 0x9c, 0xf8,
	0x0b, 0xfa, 0x86, 0x10, 0xd7, 0x56, 0xd7, 0x90, 0x15, 0xfb, 0xba, 0x65, 0xfb, 0xea, 0x00, 0xf4,
	0x3d, 0xa1, 0x1a, 0x79, 0x26, 0xb8, 0x54, 0x90, 0xe1, 0x15, 0x76, 0xb5, 0xda, 0xee, 0xdc, 0x20,
	0x15, 0xf8, 0x63, 0xb2, 0x63, 0xa4, 0x81, 0xb1, 0x5d, 0xc3, 0x09, 0x0e, 0xfb, 0xe3, 0xa9, 0x80,
	0xb0, 0xbc, 0xd2, 0x94, 0xee, 0x39, 0x50, 0xcf, 0x71, 0x5e, 0x4f, 0x05, 0xb4, 0x5f, 0x9d, 0xcf,
	0xa3, 0xe0, 0x62, 0x1e, 0x05, 0x3f, 0xe7, 0x51, 0xf0, 0x79, 0x11, 0x95, 0x2e, 0x16, 0x51, 0xe9,
	0xdb, 0x22, 0x2a, 0x1d, 0xb3, 0x1b, 0x48, 0xb7, 0xf1, 0xfd, 0x5c, 0xc9, 0x0f, 0x98, 0x9a, 0x38,
	0x95, 0x0a, 0xe3, 0xd9, 0xf5, 0x17, 0xc8, 0xe1, 0x07, 0x15, 0xf7, 0x5d, 0x79, 0xfa, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x5c, 0xcb, 0xcb, 0xa7, 0xf4, 0x04, 0x00, 0x00,
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
	if len(m.EpochState) > 0 {
		for iNdEx := len(m.EpochState) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EpochState[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.EpochInitialIssuance) > 0 {
		for iNdEx := len(m.EpochInitialIssuance) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EpochInitialIssuance[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.TaxProceeds) > 0 {
		for iNdEx := len(m.TaxProceeds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TaxProceeds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.TaxCaps) > 0 {
		for iNdEx := len(m.TaxCaps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TaxCaps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size := m.RewardWeight.Size()
		i -= size
		if _, err := m.RewardWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.TaxRate.Size()
		i -= size
		if _, err := m.TaxRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
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

func (m *TaxCap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaxCap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaxCap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TaxCap.Size()
		i -= size
		if _, err := m.TaxCap.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EpochState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalStakedLuna.Size()
		i -= size
		if _, err := m.TotalStakedLuna.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.SeigniorageReward.Size()
		i -= size
		if _, err := m.SeigniorageReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.TaxReward.Size()
		i -= size
		if _, err := m.TaxReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Epoch != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x8
	}
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
	l = m.TaxRate.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.RewardWeight.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.TaxCaps) > 0 {
		for _, e := range m.TaxCaps {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TaxProceeds) > 0 {
		for _, e := range m.TaxProceeds {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.EpochInitialIssuance) > 0 {
		for _, e := range m.EpochInitialIssuance {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.EpochState) > 0 {
		for _, e := range m.EpochState {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *TaxCap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.TaxCap.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *EpochState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovGenesis(uint64(m.Epoch))
	}
	l = m.TaxReward.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.SeigniorageReward.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.TotalStakedLuna.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field TaxRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TaxRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxCaps", wireType)
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
			m.TaxCaps = append(m.TaxCaps, TaxCap{})
			if err := m.TaxCaps[len(m.TaxCaps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxProceeds", wireType)
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
			m.TaxProceeds = append(m.TaxProceeds, types.Coin{})
			if err := m.TaxProceeds[len(m.TaxProceeds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochInitialIssuance", wireType)
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
			m.EpochInitialIssuance = append(m.EpochInitialIssuance, types.Coin{})
			if err := m.EpochInitialIssuance[len(m.EpochInitialIssuance)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochState", wireType)
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
			m.EpochState = append(m.EpochState, EpochState{})
			if err := m.EpochState[len(m.EpochState)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *TaxCap) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: TaxCap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaxCap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxCap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TaxCap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *EpochState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EpochState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TaxReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeigniorageReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SeigniorageReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalStakedLuna", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalStakedLuna.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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