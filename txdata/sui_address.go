package txdata

import (
	"encoding/hex"
	"fmt"
)

type SuiAddress [32]byte

func (c *SuiAddress) ToBytes() []byte {
	return c[:]
}

func (c *SuiAddress) Parse(data []byte, offset int) (int, error) {
	if len(data) < offset+32 {
		return 0, ErrNotEnoughData
	}

	copy(c[:], data[offset:offset+32])
	offset += 32
	fmt.Println("SuiAddress.Parse() called", hex.EncodeToString(c[:]))
	return offset, nil
}

func (c *SuiAddress) String() string {
	return "0x" + hex.EncodeToString(c[:])
}
