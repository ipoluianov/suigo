package client

import (
	"encoding/binary"
	"errors"
	"math/big"

	"github.com/xchgn/suigo/txdata"
	"github.com/xchgn/suigo/utils"
)

type TransactionBuilderMoveCall struct {
	PackageId    string
	ModuleName   string
	FunctionName string
	Arguments    []interface{}
}

type ArgAddress string
type ArgObject string
type ArgBool bool
type ArgU8 uint8
type ArgU16 uint16
type ArgU32 uint32
type ArgU64 uint64
type ArgU128 big.Int
type ArgU256 big.Int

type ArgVecU8 []uint8

func NewTransactionBuilderMoveCall() *TransactionBuilderMoveCall {
	var c TransactionBuilderMoveCall
	return &c
}

func (c *TransactionBuilderMoveCall) Build(builder *TransactionBuilder) error {
	var cmd txdata.Command
	cmd.Type = txdata.CommandTypeMoveCall
	cmd.MoveCall = &txdata.ProgrammableMoveCall{}
	cmd.MoveCall.Package.SetHex(c.PackageId)
	cmd.MoveCall.Module = c.ModuleName
	cmd.MoveCall.Function = c.FunctionName
	for _, arg := range c.Arguments {
		switch v := arg.(type) {
		case ArgBool:
			value := uint8(0)
			if v {
				value = 1
			}
			pIndex := c.buildArgumentU8(builder.transactionData.V1.Kind.ProgrammableTransaction, value)
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		case ArgU8:
			value := uint8(v)
			pIndex := c.buildArgumentU8(builder.transactionData.V1.Kind.ProgrammableTransaction, value)
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		case ArgVecU8:
			value := []uint8(v)
			pIndex := c.buildArgumentVecU8(builder.transactionData.V1.Kind.ProgrammableTransaction, value)
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		case ArgU16:
			value := uint16(v)
			pIndex := c.buildArgumentU16(builder.transactionData.V1.Kind.ProgrammableTransaction, value)
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		case ArgU32:
			value := uint32(v)
			pIndex := c.buildArgumentU32(builder.transactionData.V1.Kind.ProgrammableTransaction, value)
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		case ArgU64:
			value := uint64(v)
			pIndex := c.buildArgumentU64(builder.transactionData.V1.Kind.ProgrammableTransaction, value)
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		case ArgU128:
			value := big.Int(v)
			pIndex := c.buildArgumentU128(builder.transactionData.V1.Kind.ProgrammableTransaction, value)
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		case ArgU256:
			value := big.Int(v)
			pIndex := c.buildArgumentU256(builder.transactionData.V1.Kind.ProgrammableTransaction, value)
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		case ArgAddress:
			pIndex := c.buildArgumentAddress(builder.transactionData.V1.Kind.ProgrammableTransaction, string(v))
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		case ArgObject:
			pIndex := c.buildArgumentObject(builder.transactionData.V1.Kind.ProgrammableTransaction, string(v))
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		default:
			return errors.New("unsupported argument type")
		}
	}
	builder.transactionData.V1.Kind.ProgrammableTransaction.Commands = append(builder.transactionData.V1.Kind.ProgrammableTransaction.Commands, &cmd)

	return nil
}

func (c *TransactionBuilderMoveCall) buildArgumentU8(tx *txdata.ProgrammableTransaction, value uint8) int {
	bs := make([]byte, 1)
	bs[0] = value
	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: bs,
	})
	return len(tx.Inputs) - 1
}

func (c *TransactionBuilderMoveCall) buildArgumentVecU8(tx *txdata.ProgrammableTransaction, value []uint8) int {
	bs := make([]byte, 0)
	bsSize := txdata.SerializeULEB128(len(value))
	bs = append(bs, bsSize...)
	bs = append(bs, value...)
	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: bs,
	})
	return len(tx.Inputs) - 1
}

func (c *TransactionBuilderMoveCall) buildArgumentU16(tx *txdata.ProgrammableTransaction, value uint16) int {
	bs := make([]byte, 2)
	binary.LittleEndian.PutUint16(bs, value)
	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: bs,
	})
	return len(tx.Inputs) - 1
}

func (c *TransactionBuilderMoveCall) buildArgumentU32(tx *txdata.ProgrammableTransaction, value uint32) int {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, value)
	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: bs,
	})
	return len(tx.Inputs) - 1
}

func (c *TransactionBuilderMoveCall) buildArgumentU64(tx *txdata.ProgrammableTransaction, value uint64) int {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, value)
	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: bs,
	})
	return len(tx.Inputs) - 1
}

func (c *TransactionBuilderMoveCall) buildArgumentU128(tx *txdata.ProgrammableTransaction, value big.Int) int {
	bs := make([]byte, 16)
	bigEndBS := value.Bytes()
	for i := 0; i < len(bigEndBS) && i < 16; i++ {
		bs[i] = bigEndBS[len(bigEndBS)-1-i]
	}
	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: bs,
	})
	return len(tx.Inputs) - 1
}

func (c *TransactionBuilderMoveCall) buildArgumentU256(tx *txdata.ProgrammableTransaction, value big.Int) int {
	bs := make([]byte, 32)
	bigEndBS := value.Bytes()
	for i := 0; i < len(bigEndBS) && i < 32; i++ {
		bs[i] = bigEndBS[len(bigEndBS)-1-i]
	}
	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: bs,
	})
	return len(tx.Inputs) - 1
}

func (c *TransactionBuilderMoveCall) buildArgumentObject(tx *txdata.ProgrammableTransaction, objectId string) int {
	var arg txdata.CallArg
	arg.Type = txdata.CallArgTypeObject
	var objectArg txdata.ObjectArg
	objectArg.Type = txdata.ObjectArgTypeSharedObject
	var sharedObj txdata.SharedObject
	sharedObj.Id.SetHex(objectId)
	objectArg.SharedObject = &sharedObj
	arg.Object = &objectArg
	tx.Inputs = append(tx.Inputs, &arg)
	return len(tx.Inputs) - 1
}

func (c *TransactionBuilderMoveCall) buildArgumentAddress(tx *txdata.ProgrammableTransaction, value string) int {
	bs := make([]byte, 32)
	bsData := utils.ParseHex(value)
	if len(bsData) > 32 {
		bsData = bsData[len(bsData)-32:] // TODO: error
	}
	for len(bsData) < 32 {
		bsData = append([]byte{0}, bsData...)
	}
	copy(bs, bsData)

	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: bs,
	})
	return len(tx.Inputs) - 1
}
