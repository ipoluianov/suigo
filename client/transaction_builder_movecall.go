package client

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"strconv"

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
type ArgSharedObject string
type ArgImmObject string
type ArgBool bool
type ArgU8 uint8
type ArgU16 uint16
type ArgU32 uint32
type ArgU64 uint64
type ArgU128 big.Int
type ArgU256 big.Int

type ArgVecAddress []string
type ArgVecObject []string
type ArgVecBool []bool
type ArgVecU8 []uint8
type ArgVecU16 []uint16
type ArgVecU32 []uint32
type ArgVecU64 []uint64
type ArgVecU128 []big.Int
type ArgVecU256 []big.Int

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
		case ArgVecU16:
			pIndex := c.buildArgumentVecU16(builder.transactionData.V1.Kind.ProgrammableTransaction, v)
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
		case ArgVecU32:
			pIndex := c.buildArgumentVecU32(builder.transactionData.V1.Kind.ProgrammableTransaction, v)
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
		case ArgVecU64:
			pIndex := c.buildArgumentVecU64(builder.transactionData.V1.Kind.ProgrammableTransaction, v)
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
		case ArgVecU128:
			pIndex := c.buildArgumentVecU128(builder.transactionData.V1.Kind.ProgrammableTransaction, v)
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
		case ArgVecU256:
			pIndex := c.buildArgumentVecU256(builder.transactionData.V1.Kind.ProgrammableTransaction, v)
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
		case ArgVecAddress:
			pIndex := c.buildArgumentVecAddress(builder.transactionData.V1.Kind.ProgrammableTransaction, v)
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		case ArgSharedObject:
			var pIndex int
			if existingObjectRef, ok := builder.objectRefs[string(v)]; ok {
				pIndex = existingObjectRef
			} else {
				pIndex = c.buildArgumentSharedObject(builder, builder.transactionData.V1.Kind.ProgrammableTransaction, string(v))
				builder.objectRefs[string(v)] = pIndex
			}
			//pIndex = c.buildArgumentSharedObject(builder, builder.transactionData.V1.Kind.ProgrammableTransaction, string(v))
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		case ArgImmObject:
			pIndex := c.buildArgumentImmObject(builder, builder.transactionData.V1.Kind.ProgrammableTransaction, string(v))
			arg := txdata.Argument{}
			arg.ArgumentType = txdata.ArgumentTypeInput
			arg.ArgumentInput = txdata.ArgumentInput(pIndex)
			cmd.MoveCall.Arguments = append(cmd.MoveCall.Arguments, arg)
		case ArgVecObject:
			pIndex := c.buildArgumentVecObject(builder.transactionData.V1.Kind.ProgrammableTransaction, v)
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

func (c *TransactionBuilderMoveCall) buildArgumentVecU16(tx *txdata.ProgrammableTransaction, value []uint16) int {
	result := make([]byte, 0)
	bsSize := txdata.SerializeULEB128(len(value))
	result = append(result, bsSize...)

	bs := make([]byte, 2)
	for _, v := range value {
		binary.LittleEndian.PutUint16(bs, v)
		result = append(result, bs...)
	}

	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: result,
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

func (c *TransactionBuilderMoveCall) buildArgumentVecU32(tx *txdata.ProgrammableTransaction, value []uint32) int {
	result := make([]byte, 0)
	bsSize := txdata.SerializeULEB128(len(value))
	result = append(result, bsSize...)

	bs := make([]byte, 4)
	for _, v := range value {
		binary.LittleEndian.PutUint32(bs, v)
		result = append(result, bs...)
	}

	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: result,
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

func (c *TransactionBuilderMoveCall) buildArgumentVecU64(tx *txdata.ProgrammableTransaction, value []uint64) int {
	result := make([]byte, 0)
	bsSize := txdata.SerializeULEB128(len(value))
	result = append(result, bsSize...)

	bs := make([]byte, 8)
	for _, v := range value {
		binary.LittleEndian.PutUint64(bs, v)
		result = append(result, bs...)
	}

	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: result,
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

func (c *TransactionBuilderMoveCall) buildArgumentVecU128(tx *txdata.ProgrammableTransaction, value []big.Int) int {
	result := make([]byte, 0)
	bsSize := txdata.SerializeULEB128(len(value))
	result = append(result, bsSize...)

	bs := make([]byte, 16)
	for _, v := range value {
		bigEndBS := v.Bytes()
		for i := 0; i < len(bigEndBS) && i < 16; i++ {
			bs[i] = bigEndBS[len(bigEndBS)-1-i]
		}
		result = append(result, bs...)
	}

	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: result,
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

func (c *TransactionBuilderMoveCall) buildArgumentVecU256(tx *txdata.ProgrammableTransaction, value []big.Int) int {
	result := make([]byte, 0)
	bsSize := txdata.SerializeULEB128(len(value))
	result = append(result, bsSize...)

	bs := make([]byte, 32)
	for _, v := range value {
		bigEndBS := v.Bytes()
		for i := 0; i < len(bigEndBS) && i < 32; i++ {
			bs[i] = bigEndBS[len(bigEndBS)-1-i]
		}
		result = append(result, bs...)
	}

	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: result,
	})
	return len(tx.Inputs) - 1
}

func (c *TransactionBuilderMoveCall) buildArgumentSharedObject(builder *TransactionBuilder, tx *txdata.ProgrammableTransaction, objectId string) int {
	//var digest string
	var version uint64

	/*objData, err := builder.client.GetObject(objectId, GetObjectShowOptions{})
	if err == nil {
		//digest = objData.Data.Digest
		version, _ = strconv.ParseUint(objData.Data.Version, 10, 64)
	}*/

	var err error

	version, err = builder.client.GetInitialSharedVersion(objectId)
	if err != nil {
		fmt.Println("GetInitialSharedVersion Error:", err)
	}

	mutable := true

	if objectId == CLOCK_OBJECT_ID {
		mutable = false
	}

	// fmt.Println("OBJ VERSION FOR SHARED:", version)
	// fmt.Println("OBJ DIGEST FOR SHARED:", digest)

	var arg txdata.CallArg
	arg.Type = txdata.CallArgTypeObject
	var objectArg txdata.ObjectArg
	objectArg.Type = txdata.ObjectArgTypeSharedObject
	var sharedObj txdata.SharedObject
	sharedObj.Id.SetHex(objectId)
	sharedObj.InitialSharedVersion = txdata.SequenceNumber(version)
	sharedObj.Mutable = mutable
	objectArg.SharedObject = &sharedObj
	arg.Object = &objectArg
	tx.Inputs = append(tx.Inputs, &arg)
	return len(tx.Inputs) - 1
}

func (c *TransactionBuilderMoveCall) buildArgumentImmObject(builder *TransactionBuilder, tx *txdata.ProgrammableTransaction, objectId string) int {
	var digest string
	var version uint64
	// Get Object Version
	objData, err := builder.client.GetObject(objectId, GetObjectShowOptions{})
	if err == nil {
		fmt.Println("OBJ DIGEST:", objData.Data.Digest)
		digest = objData.Data.Digest
		version, _ = strconv.ParseUint(objData.Data.Version, 10, 64)
	}

	// Get Object Digest

	var arg txdata.CallArg
	arg.Type = txdata.CallArgTypeObject
	var objectArg txdata.ObjectArg
	objectArg.Type = txdata.ObjectArgTypeImmOrOwnedObject
	var immObj txdata.ObjectRef
	immObj.ObjectID.SetHex(objectId)
	immObj.SequenceNumber = txdata.SequenceNumber(version)
	immObj.ObjectDigest.SetBase58(digest)
	//immObj.ObjectDigest.
	//sharedObj.Id.SetHex(objectId)
	objectArg.ImmOrOwnedObject = &immObj
	arg.Object = &objectArg
	tx.Inputs = append(tx.Inputs, &arg)
	return len(tx.Inputs) - 1
}

func (c *TransactionBuilderMoveCall) buildArgumentVecObject(tx *txdata.ProgrammableTransaction, objectId []string) int {
	panic("not implemented")
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

func (c *TransactionBuilderMoveCall) buildArgumentVecAddress(tx *txdata.ProgrammableTransaction, value []string) int {
	result := make([]byte, 0)
	bsSize := txdata.SerializeULEB128(len(value))
	result = append(result, bsSize...)

	bs := make([]byte, 32)
	for i := 0; i < len(value); i++ {
		bsData := utils.ParseHex(value[i])
		if len(bsData) > 32 {
			bsData = bsData[len(bsData)-32:] // TODO: error
		}
		for len(bsData) < 32 {
			bsData = append([]byte{0}, bsData...)
		}
		copy(bs, bsData)
		result = append(result, bs...)
	}

	tx.Inputs = append(tx.Inputs, &txdata.CallArg{
		Type: txdata.CallArgTypePure,
		Pure: result,
	})
	return len(tx.Inputs) - 1
}
