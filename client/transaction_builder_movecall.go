package client

import (
	"encoding/binary"
	"errors"

	"github.com/xchgn/suigo/txdata"
)

type TransactionBuilderMoveCall struct {
	PackageId    string
	ModuleName   string
	FunctionName string
	Arguments    []interface{}
}

type ArgObject string
type ArgU64 uint64

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
		case ArgU64:
			value := uint64(v)
			pIndex := c.buildArgumentU64(builder.transactionData.V1.Kind.ProgrammableTransaction, value)
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

func (c *TransactionBuilderMoveCall) buildArgumentU64(tx *txdata.ProgrammableTransaction, value uint64) int {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, value)
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
