package client

import (
	"encoding/json"
	"fmt"
)

func (c *Client) UnsafeMoveCall(gasObj string, gasBudget string, packageId string, moduleName string, functionName string, arguments []interface{}) (*TransactionBlockBytes, error) {
	requestBody := RPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "unsafe_moveCall",
		Params: []interface{}{
			c.account.Address,
			packageId,
			moduleName,
			functionName,
			[]interface{}{},
			arguments,
			&gasObj,
			gasBudget,
		},
	}

	res, err := c.rpcCall(requestBody)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if res.Error != nil {
		fmt.Println("ERROR")
		fmt.Println(res.Error.Code, res.Error.Message)
		return nil, fmt.Errorf("ERROR: %d %s", res.Error.Code, res.Error.Message)
	}

	var txBytes TransactionBlockBytes

	json.Unmarshal(res.Result, &txBytes)

	return &txBytes, nil
}
