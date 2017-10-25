// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package nebletpb is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Config
	P2PConfig
	RPCConfig
	PowConfig
	AccountConfig
*/
package nebletpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Neblet global configurations.
type Config struct {
	// P2P network config.
	P2P *P2PConfig `protobuf:"bytes,1,opt,name=p2p" json:"p2p,omitempty"`
	// RPC server config.
	Rpc *RPCConfig `protobuf:"bytes,2,opt,name=rpc" json:"rpc,omitempty"`
	// Pow config
	Pow *PowConfig `protobuf:"bytes,3,opt,name=pow" json:"pow,omitempty"`
	// Account manager config
	Account *AccountConfig `protobuf:"bytes,4,opt,name=account" json:"account,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetP2P() *P2PConfig {
	if m != nil {
		return m.P2P
	}
	return nil
}

func (m *Config) GetRpc() *RPCConfig {
	if m != nil {
		return m.Rpc
	}
	return nil
}

func (m *Config) GetPow() *PowConfig {
	if m != nil {
		return m.Pow
	}
	return nil
}

func (m *Config) GetAccount() *AccountConfig {
	if m != nil {
		return m.Account
	}
	return nil
}

type P2PConfig struct {
	// P2P seed node addresses.
	Seed string `protobuf:"bytes,1,opt,name=seed,proto3" json:"seed,omitempty"`
	// P2P node port number.
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// P2P node chainID.
	ChainId uint32 `protobuf:"varint,3,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// P2P node version.
	Version uint32 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *P2PConfig) Reset()                    { *m = P2PConfig{} }
func (m *P2PConfig) String() string            { return proto.CompactTextString(m) }
func (*P2PConfig) ProtoMessage()               {}
func (*P2PConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *P2PConfig) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *P2PConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *P2PConfig) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *P2PConfig) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

type RPCConfig struct {
	// RPC server listening port number.
	Port uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (m *RPCConfig) Reset()                    { *m = RPCConfig{} }
func (m *RPCConfig) String() string            { return proto.CompactTextString(m) }
func (*RPCConfig) ProtoMessage()               {}
func (*RPCConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

func (m *RPCConfig) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type PowConfig struct {
	// pow miner's coinbase
	Coinbase string `protobuf:"bytes,1,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
}

func (m *PowConfig) Reset()                    { *m = PowConfig{} }
func (m *PowConfig) String() string            { return proto.CompactTextString(m) }
func (*PowConfig) ProtoMessage()               {}
func (*PowConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

func (m *PowConfig) GetCoinbase() string {
	if m != nil {
		return m.Coinbase
	}
	return ""
}

type AccountConfig struct {
	// Account transaction sign algorithm type
	Signature uint32 `protobuf:"varint,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Account addr encrypt algorithm type
	Encrypt uint32 `protobuf:"varint,2,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
	// Account key file directory
	KeyDir string `protobuf:"bytes,3,opt,name=key_dir,json=keyDir,proto3" json:"key_dir,omitempty"`
	// Account test keyfile's passphrase
	TestPassphrase string `protobuf:"bytes,4,opt,name=test_passphrase,json=testPassphrase,proto3" json:"test_passphrase,omitempty"`
}

func (m *AccountConfig) Reset()                    { *m = AccountConfig{} }
func (m *AccountConfig) String() string            { return proto.CompactTextString(m) }
func (*AccountConfig) ProtoMessage()               {}
func (*AccountConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{4} }

func (m *AccountConfig) GetSignature() uint32 {
	if m != nil {
		return m.Signature
	}
	return 0
}

func (m *AccountConfig) GetEncrypt() uint32 {
	if m != nil {
		return m.Encrypt
	}
	return 0
}

func (m *AccountConfig) GetKeyDir() string {
	if m != nil {
		return m.KeyDir
	}
	return ""
}

func (m *AccountConfig) GetTestPassphrase() string {
	if m != nil {
		return m.TestPassphrase
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "nebletpb.Config")
	proto.RegisterType((*P2PConfig)(nil), "nebletpb.P2PConfig")
	proto.RegisterType((*RPCConfig)(nil), "nebletpb.RPCConfig")
	proto.RegisterType((*PowConfig)(nil), "nebletpb.PowConfig")
	proto.RegisterType((*AccountConfig)(nil), "nebletpb.AccountConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xd1, 0x4a, 0xf3, 0x30,
	0x14, 0xc7, 0xe9, 0xb7, 0xb1, 0x2e, 0xe7, 0xb3, 0x0a, 0xf1, 0x62, 0x55, 0x04, 0xa5, 0x20, 0xf3,
	0x6a, 0xe0, 0x7c, 0x02, 0x99, 0x37, 0xde, 0x95, 0xbc, 0xc0, 0x68, 0xd3, 0xb8, 0x85, 0x49, 0x12,
	0x92, 0xcc, 0xb1, 0x47, 0xf0, 0x6d, 0x7c, 0x44, 0xc9, 0x59, 0x93, 0xa2, 0xde, 0xe5, 0x9c, 0xf3,
	0xe3, 0xcf, 0xef, 0x0f, 0x81, 0x33, 0xae, 0xd5, 0x9b, 0xdc, 0x2c, 0x8c, 0xd5, 0x5e, 0xd3, 0xa9,
	0x12, 0xed, 0xbb, 0xf0, 0xa6, 0xad, 0xbe, 0x32, 0x98, 0xac, 0xf0, 0x44, 0xef, 0x61, 0x64, 0x96,
	0xa6, 0xcc, 0xee, 0xb2, 0x87, 0xff, 0xcb, 0xcb, 0x45, 0x44, 0x16, 0xf5, 0xb2, 0x3e, 0x11, 0x2c,
	0xdc, 0x03, 0x66, 0x0d, 0x2f, 0xff, 0xfd, 0xc6, 0x58, 0xbd, 0x8a, 0x98, 0x35, 0x1c, 0xd3, 0xf4,
	0xa1, 0x1c, 0xfd, 0x49, 0xd3, 0x87, 0x94, 0xa6, 0x0f, 0xf4, 0x11, 0xf2, 0x86, 0x73, 0xbd, 0x57,
	0xbe, 0x1c, 0x23, 0x3a, 0x1b, 0xd0, 0xe7, 0xd3, 0xa1, 0xc7, 0x23, 0x57, 0x6d, 0x81, 0x24, 0x25,
	0x4a, 0x61, 0xec, 0x84, 0xe8, 0xd0, 0x9a, 0x30, 0x7c, 0x87, 0x9d, 0xd1, 0xd6, 0xa3, 0x62, 0xc1,
	0xf0, 0x4d, 0xaf, 0x60, 0xca, 0xb7, 0x8d, 0x54, 0x6b, 0xd9, 0xa1, 0x53, 0xc1, 0x72, 0x9c, 0x5f,
	0x3b, 0x5a, 0x42, 0xfe, 0x21, 0xac, 0x93, 0x5a, 0xa1, 0x42, 0xc1, 0xe2, 0x58, 0xdd, 0x02, 0x49,
	0xad, 0x52, 0x6a, 0x36, 0xa4, 0x56, 0x73, 0x20, 0xa9, 0x0f, 0xbd, 0x86, 0x29, 0xd7, 0x52, 0xb5,
	0x8d, 0x13, 0xbd, 0x4e, 0x9a, 0xab, 0xcf, 0x0c, 0x8a, 0x1f, 0x75, 0xe8, 0x0d, 0x10, 0x27, 0x37,
	0xaa, 0xf1, 0x7b, 0x2b, 0xfa, 0xcc, 0x61, 0x11, 0x9c, 0x84, 0xe2, 0xf6, 0x68, 0x62, 0x8b, 0x38,
	0xd2, 0x19, 0xe4, 0x3b, 0x71, 0x5c, 0x77, 0xd2, 0x62, 0x0f, 0xc2, 0x26, 0x3b, 0x71, 0x7c, 0x91,
	0x96, 0xce, 0xe1, 0xc2, 0x0b, 0xe7, 0xd7, 0xa6, 0x71, 0xce, 0x6c, 0x6d, 0xb0, 0x18, 0x23, 0x70,
	0x1e, 0xd6, 0x75, 0xda, 0xb6, 0x13, 0xfc, 0x03, 0x4f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x14,
	0x93, 0xbb, 0xc8, 0x13, 0x02, 0x00, 0x00,
}
