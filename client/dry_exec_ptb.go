package client

func (c *Client) DryExecPTB(builder *TransactionBuilder, gasPrice uint64) (*DryRunTransactionBlockResponse, error) {
	txBytes, err := builder.Build(gasPrice)
	if err != nil {
		return nil, err
	}

	result, err := c.DryRunTransactionBlock(txBytes)
	return result, err
}
