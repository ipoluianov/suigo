package txdata

type Upgrade struct {
}

func (c *Upgrade) ToBytes() []byte {
	return []byte{}
}

func (c *Upgrade) Parse(data []byte, offset int) (int, error) {
	return 0, ErrNotImplemented
}
