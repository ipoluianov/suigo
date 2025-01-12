package client

func (c *Client) ExecPTB(builder *TransactionBuilder) (*TransactionExecutionResult, error) {
	txBytes, err := builder.Build()
	if err != nil {
		return nil, err
	}

	txSigned, err := c.account.Signature(txBytes)
	if err != nil {
		return nil, err
	}

	result, err := c.ExecuteTransactionBlock(txSigned.TxBytes, txSigned.Signature)
	return result, err
}
