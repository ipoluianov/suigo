package txdata

import "errors"

type ConsensusCommitPrologueV2 struct {
}

func (c *ConsensusCommitPrologueV2) ToBytes() []byte {
	return []byte{}
}

func (c *ConsensusCommitPrologueV2) Parse(data []byte, offset int) (int, error) {
	return 0, errors.New("not implemented")
}
