package client

import (
	"encoding/json"
	"errors"
)

func (c *Client) GetCoinMetadata(coinType string) (coinMetadata SuiCoinMetadata, err error) {
	if coinType == "" {
		return SuiCoinMetadata{}, errors.New("coinType is required")
	}

	requestBody := RPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "suix_getCoinMetadata",
		Params:  []interface{}{coinType},
	}

	requestBody.Params = append(requestBody.Params, coinType)

	res, err := c.rpcCall(requestBody)
	if err != nil {
		return
	}

	err = json.Unmarshal(res.Result, &coinMetadata)
	return
}
