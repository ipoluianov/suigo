package client

import (
	"encoding/json"
	"fmt"
)

type TransactionBlockResponse struct {
	Digest         string           `json:"digest"`
	TimestampMs    string           `json:"timestampMs"`
	CheckPoint     string           `json:"checkpoint"`
	Transaction    TransactionBlock `json:"transaction"`
	RawTransaction string           `json:"rawTransaction"`
}

func (c *Client) GetTransactionBlock(digest string, showParams TransactionBlockResponseOptions) (response TransactionBlockResponse, err error) {
	requestBody := RPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "sui_getTransactionBlock",
		Params:  []interface{}{digest, showParams},
	}

	res, err := c.rpcCall(requestBody)

	if err != nil {
		return
	}

	fmt.Println(string(res.Result))

	err = json.Unmarshal(res.Result, &response)
	return
}
