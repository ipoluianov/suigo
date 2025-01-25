package client

import (
	"encoding/json"
)

func (c *Client) GetObject(objectId string, options GetObjectShowOptions) (objData SuiObjectResponse, err error) {
	requestBody := RPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "sui_getObject",
		Params: []interface{}{
			objectId,
			options,
		},
	}

	res, err := c.rpcCall(requestBody)
	if err != nil {
		return
	}

	/*fmt.Println(string(res.Result))
	fmt.Println("")
	fmt.Println("")*/

	err = json.Unmarshal(res.Result, &objData)
	if err != nil {
		return
	}

	return
}
