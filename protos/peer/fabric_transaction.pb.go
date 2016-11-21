// Code generated by protoc-gen-go.
// source: peer/fabric_transaction.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type InvalidTransaction_Cause int32

const (
	InvalidTransaction_TxIdAlreadyExists      InvalidTransaction_Cause = 0
	InvalidTransaction_RWConflictDuringCommit InvalidTransaction_Cause = 1
)

var InvalidTransaction_Cause_name = map[int32]string{
	0: "TxIdAlreadyExists",
	1: "RWConflictDuringCommit",
}
var InvalidTransaction_Cause_value = map[string]int32{
	"TxIdAlreadyExists":      0,
	"RWConflictDuringCommit": 1,
}

func (x InvalidTransaction_Cause) String() string {
	return proto.EnumName(InvalidTransaction_Cause_name, int32(x))
}
func (InvalidTransaction_Cause) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{1, 0} }

// This message is necessary to facilitate the verification of the signature
// (in the signature field) over the bytes of the transaction (in the
// transactionBytes field).
type SignedTransaction struct {
	// The bytes of the Transaction. NDD
	TransactionBytes []byte `protobuf:"bytes,1,opt,name=transactionBytes,proto3" json:"transactionBytes,omitempty"`
	// Signature of the transactionBytes The public key of the signature is in
	// the header field of TransactionAction There might be multiple
	// TransactionAction, so multiple headers, but there should be same
	// transactor identity (cert) in all headers
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignedTransaction) Reset()                    { *m = SignedTransaction{} }
func (m *SignedTransaction) String() string            { return proto.CompactTextString(m) }
func (*SignedTransaction) ProtoMessage()               {}
func (*SignedTransaction) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

// This is used to wrap an invalid Transaction with the cause
type InvalidTransaction struct {
	Transaction *Transaction2            `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
	Cause       InvalidTransaction_Cause `protobuf:"varint,2,opt,name=cause,enum=protos.InvalidTransaction_Cause" json:"cause,omitempty"`
}

func (m *InvalidTransaction) Reset()                    { *m = InvalidTransaction{} }
func (m *InvalidTransaction) String() string            { return proto.CompactTextString(m) }
func (*InvalidTransaction) ProtoMessage()               {}
func (*InvalidTransaction) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *InvalidTransaction) GetTransaction() *Transaction2 {
	if m != nil {
		return m.Transaction
	}
	return nil
}

// The transaction to be sent to the ordering service. A transaction contains
// one or more TransactionAction. Each TransactionAction binds a proposal to
// potentially multiple actions. The transaction is atomic meaning that either
// all actions in the transaction will be committed or none will.  Note that
// while a Transaction might include more than one Header, the Header.creator
// field must be the same in each.
// A single client is free to issue a number of independent Proposal, each with
// their header (Header) and request payload (ChaincodeProposalPayload).  Each
// proposal is independently endorsed generating an action
// (ProposalResponsePayload) with one signature per Endorser. Any number of
// independent proposals (and their action) might be included in a transaction
// to ensure that they are treated atomically.
type Transaction2 struct {
	// Version indicates message protocol version.
	Version int32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	// Timestamp is the local time that the
	// message was created by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// The payload is an array of TransactionAction. An array is necessary to
	// accommodate multiple actions per transaction
	Actions []*TransactionAction `protobuf:"bytes,3,rep,name=actions" json:"actions,omitempty"`
}

func (m *Transaction2) Reset()                    { *m = Transaction2{} }
func (m *Transaction2) String() string            { return proto.CompactTextString(m) }
func (*Transaction2) ProtoMessage()               {}
func (*Transaction2) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *Transaction2) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Transaction2) GetActions() []*TransactionAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

// TransactionAction binds a proposal to its action.  The type field in the
// header dictates the type of action to be applied to the ledger.
type TransactionAction struct {
	// The header of the proposal action, which is the proposal header
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The payload of the action as defined by the type in the header For
	// chaincode, it's the bytes of ChaincodeActionPayload
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *TransactionAction) Reset()                    { *m = TransactionAction{} }
func (m *TransactionAction) String() string            { return proto.CompactTextString(m) }
func (*TransactionAction) ProtoMessage()               {}
func (*TransactionAction) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

func init() {
	proto.RegisterType((*SignedTransaction)(nil), "protos.SignedTransaction")
	proto.RegisterType((*InvalidTransaction)(nil), "protos.InvalidTransaction")
	proto.RegisterType((*Transaction2)(nil), "protos.Transaction2")
	proto.RegisterType((*TransactionAction)(nil), "protos.TransactionAction")
	proto.RegisterEnum("protos.InvalidTransaction_Cause", InvalidTransaction_Cause_name, InvalidTransaction_Cause_value)
}

func init() { proto.RegisterFile("peer/fabric_transaction.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x6a, 0xe2, 0x40,
	0x18, 0x85, 0x37, 0x2b, 0x2a, 0xfe, 0x91, 0x45, 0x87, 0x5d, 0xc9, 0x4a, 0x4b, 0x25, 0x57, 0xb6,
	0x85, 0x04, 0x22, 0x48, 0xe9, 0x9d, 0x5a, 0x2f, 0xbc, 0x4d, 0x85, 0x42, 0xa1, 0x94, 0x49, 0x32,
	0xc6, 0x81, 0x24, 0x13, 0x66, 0x26, 0x62, 0x9e, 0xa4, 0xaf, 0xd3, 0x47, 0x2b, 0x66, 0x1c, 0x4d,
	0xb1, 0x37, 0x09, 0x67, 0xf2, 0xe5, 0xfc, 0x87, 0xf9, 0x0f, 0x5c, 0xe7, 0x84, 0x70, 0x77, 0x83,
	0x03, 0x4e, 0xc3, 0x77, 0xc9, 0x71, 0x26, 0x70, 0x28, 0x29, 0xcb, 0x9c, 0x9c, 0x33, 0xc9, 0x50,
	0xab, 0x7a, 0x89, 0xe1, 0x4d, 0xcc, 0x58, 0x9c, 0x10, 0xb7, 0x92, 0x41, 0xb1, 0x71, 0x25, 0x4d,
	0x89, 0x90, 0x38, 0xcd, 0x15, 0x68, 0xbf, 0x41, 0xff, 0x99, 0xc6, 0x19, 0x89, 0xd6, 0x67, 0x0f,
	0x74, 0x07, 0xbd, 0x9a, 0xe5, 0xbc, 0x94, 0x44, 0x58, 0xc6, 0xc8, 0x18, 0x77, 0xfd, 0x8b, 0x73,
	0x74, 0x05, 0x1d, 0x41, 0xe3, 0x0c, 0xcb, 0x82, 0x13, 0xeb, 0x77, 0x05, 0x9d, 0x0f, 0xec, 0x4f,
	0x03, 0xd0, 0x2a, 0xdb, 0xe1, 0x84, 0x7e, 0x1b, 0x30, 0x05, 0xb3, 0x66, 0x54, 0x79, 0x9b, 0xde,
	0x5f, 0x15, 0x49, 0x38, 0x35, 0xd2, 0xf3, 0xeb, 0x20, 0x9a, 0x42, 0x33, 0xc4, 0x85, 0x50, 0x83,
	0xfe, 0x78, 0x23, 0xfd, 0xc7, 0xe5, 0x08, 0x67, 0x71, 0xe0, 0x7c, 0x85, 0xdb, 0x8f, 0xd0, 0xac,
	0x34, 0xfa, 0x07, 0xfd, 0xf5, 0x7e, 0x15, 0xcd, 0x12, 0x4e, 0x70, 0x54, 0x2e, 0xf7, 0x54, 0x48,
	0xd1, 0xfb, 0x85, 0x86, 0x30, 0xf0, 0x5f, 0x16, 0x2c, 0xdb, 0x24, 0x34, 0x94, 0x4f, 0x05, 0xa7,
	0x59, 0xbc, 0x60, 0x69, 0x4a, 0x65, 0xcf, 0xb0, 0x3f, 0x0c, 0xe8, 0xd6, 0x13, 0x21, 0x0b, 0xda,
	0x3b, 0xc2, 0x85, 0x0e, 0xde, 0xf4, 0xb5, 0x44, 0x0f, 0xd0, 0x39, 0xdd, 0x6f, 0x15, 0xd1, 0xf4,
	0x86, 0x8e, 0xda, 0x80, 0xa3, 0x37, 0xe0, 0xac, 0x35, 0xe1, 0x9f, 0x61, 0x34, 0x81, 0xb6, 0xb2,
	0x17, 0x56, 0x63, 0xd4, 0x18, 0x9b, 0xde, 0xff, 0x1f, 0x2e, 0x63, 0x56, 0x3d, 0x7d, 0x4d, 0xda,
	0x4b, 0xe8, 0x5f, 0x7c, 0x45, 0x03, 0x68, 0x6d, 0x09, 0x8e, 0x08, 0x3f, 0x6e, 0xec, 0xa8, 0x0e,
	0xa9, 0x73, 0x5c, 0x26, 0x0c, 0x47, 0xc7, 0x2d, 0x69, 0x39, 0xbf, 0x7f, 0xbd, 0x8d, 0xa9, 0xdc,
	0x16, 0x81, 0x13, 0xb2, 0xd4, 0xdd, 0x96, 0x39, 0xe1, 0x09, 0x89, 0xe2, 0x53, 0xbd, 0x54, 0x79,
	0x84, 0x7b, 0x68, 0x5c, 0xa0, 0x8a, 0x35, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x67, 0xfe,
	0x54, 0x80, 0x02, 0x00, 0x00,
}
