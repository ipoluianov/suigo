package txdata

import "errors"

type TransactionDataV2 struct {
}

func (c *TransactionDataV2) ToBytes() []byte {
	return []byte{}
}

func (c *TransactionDataV2) Parse(data []byte, offset int) (int, error) {
	return 0, errors.New("Not implemented")
}
