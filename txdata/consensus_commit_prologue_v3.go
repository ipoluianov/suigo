package txdata

import "errors"

type ConsensusCommitPrologueV3 struct {
}

func (c *ConsensusCommitPrologueV3) ToBytes() []byte {
	return []byte{}
}

func (c *ConsensusCommitPrologueV3) Parse(data []byte, offset int) (int, error) {
	return 0, errors.New("not implemented")
}
