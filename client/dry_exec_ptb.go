package client

func (c *Client) DryExecPTB(builder *TransactionBuilder) (*DryRunTransactionBlockResponse, error) {
	txBytes, err := builder.Build()
	if err != nil {
		return nil, err
	}

	result, err := c.DryRunTransactionBlock(txBytes)
	return result, err
}
