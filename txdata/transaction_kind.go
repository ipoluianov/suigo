package txdata

import "errors"

type TransactionKind struct {
	Type TransactionKindType

	ProgrammableTransaction   *ProgrammableTransaction
	ChangeEpoch               *ChangeEpoch
	Genesis                   *GenesisTransaction
	ConsensusCommitPrologue   *ConsensusCommitPrologue
	AuthenticatorStateUpdate  *AuthenticatorStateUpdate
	EndOfEpochTransaction     []*EndOfEpochTransactionKind
	RandomnessStateUpdate     *RandomnessStateUpdate
	ConsensusCommitPrologueV2 *ConsensusCommitPrologueV2
	ConsensusCommitPrologueV3 *ConsensusCommitPrologueV3
}

type TransactionKindType int

const (
	ProgrammableTransactionType   TransactionKindType = 0
	ChangeEpochType               TransactionKindType = 1
	GenesisType                   TransactionKindType = 2
	ConsensusCommitPrologueType   TransactionKindType = 3
	AuthenticatorStateUpdateType  TransactionKindType = 4
	EndOfEpochTransactionType     TransactionKindType = 5
	RandomnessStateUpdateType     TransactionKindType = 6
	ConsensusCommitPrologueV2Type TransactionKindType = 7
	ConsensusCommitPrologueV3Type TransactionKindType = 8
)

func (c *TransactionKind) ToBytes() []byte {
	var data []byte

	// Serialize the transaction kind
	data = append(data, SerializeULEB128(int(c.Type))...) // second 0

	// Serialize the transaction data
	switch c.Type {
	case ProgrammableTransactionType:
		data = append(data, c.ProgrammableTransaction.ToBytes()...)
	case ChangeEpochType:
		data = append(data, c.ChangeEpoch.ToBytes()...)
	case GenesisType:
		data = append(data, c.Genesis.ToBytes()...)
	case ConsensusCommitPrologueType:
		data = append(data, c.ConsensusCommitPrologue.ToBytes()...)
	case AuthenticatorStateUpdateType:
		data = append(data, c.AuthenticatorStateUpdate.ToBytes()...)
	case EndOfEpochTransactionType:
		for _, v := range c.EndOfEpochTransaction {
			data = append(data, v.ToBytes()...)
		}
	case RandomnessStateUpdateType:
		data = append(data, c.RandomnessStateUpdate.ToBytes()...)
	case ConsensusCommitPrologueV2Type:
		data = append(data, c.ConsensusCommitPrologueV2.ToBytes()...)
	case ConsensusCommitPrologueV3Type:
		data = append(data, c.ConsensusCommitPrologueV3.ToBytes()...)
	}

	return data
}

func (c *TransactionKind) Parse(data []byte, offset int) (int, error) {
	var kind int
	var err error

	// Parse the transaction kind
	kind, offset, err = ParseULEB128(data, offset)
	if err != nil {
		return 0, err
	}
	if kind < 0 || kind > 8 {
		return 0, ErrInvalidEnumValue
	}
	c.Type = TransactionKindType(kind)

	// Parse the transaction data
	switch c.Type {
	case ProgrammableTransactionType:
		c.ProgrammableTransaction = &ProgrammableTransaction{}
		offset, err = c.ProgrammableTransaction.Parse(data, offset)
		if err != nil {
			return 0, err
		}
	case ChangeEpochType:
		c.ChangeEpoch = &ChangeEpoch{}
		offset, err = c.ChangeEpoch.Parse(data, offset)
		if err != nil {
			return 0, err
		}
	case GenesisType:
		c.Genesis = &GenesisTransaction{}
		offset, err = c.Genesis.Parse(data, offset)
		if err != nil {
			return 0, err
		}
	case ConsensusCommitPrologueType:
		c.ConsensusCommitPrologue = &ConsensusCommitPrologue{}
		offset, err = c.ConsensusCommitPrologue.Parse(data, offset)
		if err != nil {
			return 0, err
		}
	case AuthenticatorStateUpdateType:
		c.AuthenticatorStateUpdate = &AuthenticatorStateUpdate{}
		offset, err = c.AuthenticatorStateUpdate.Parse(data, offset)
		if err != nil {
			return 0, err
		}
	case EndOfEpochTransactionType:
		return 0, errors.New("not implemented")
		/*c.EndOfEpochTransaction = &EndOfEpochTransactionKind{}
		offset, err = c.EndOfEpochTransaction.Parse(data, offset)
		if err != nil {
			return 0, err
		}*/
	case RandomnessStateUpdateType:
		c.RandomnessStateUpdate = &RandomnessStateUpdate{}
		offset, err = c.RandomnessStateUpdate.Parse(data, offset)
		if err != nil {
			return 0, err
		}
	case ConsensusCommitPrologueV2Type:
		c.ConsensusCommitPrologueV2 = &ConsensusCommitPrologueV2{}
		offset, err = c.ConsensusCommitPrologueV2.Parse(data, offset)
		if err != nil {
			return 0, err
		}
	case ConsensusCommitPrologueV3Type:
		c.ConsensusCommitPrologueV3 = &ConsensusCommitPrologueV3{}
		offset, err = c.ConsensusCommitPrologueV3.Parse(data, offset)
		if err != nil {
			return 0, err
		}
	}

	return offset, nil
}
