package client

import (
	"encoding/json"
	"fmt"
)

// Params:
// parent_object_id< ObjectID >
// name< DynamicFieldName >

type DynamicFieldName struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

func (c *Client) GetDynamicFieldObject(parentObjectId string, fieldType string, value interface{}) (objData SuiObjectResponse, err error) {
	var dynFieldName DynamicFieldName
	dynFieldName.Type = fieldType
	dynFieldName.Value = value
	requestBody := RPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "suix_getDynamicFieldObject",
		Params: []interface{}{
			parentObjectId,
			dynFieldName,
		},
	}

	requestBodyBytes, _ := json.MarshalIndent(requestBody, "", "  ")
	fmt.Println("Request:", string(requestBodyBytes))

	res, err := c.rpcCall(requestBody)
	if err != nil {
		return
	}

	fmt.Println(string(res.Result))
	fmt.Println("")
	fmt.Println("")

	err = json.Unmarshal(res.Result, &objData)
	if err != nil {
		return
	}

	return
}

/*
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "suix_getDynamicFieldObject",
  "params": [
    "0x5ea6f7a348f4a7bd1a9ab069eb7f63865de3075cc5a4e62432f634b50fd2bb2b",
    {
      "type": "0x0000000000000000000000000000000000000000000000000000000000000009::test::TestField",
      "value": "some_value"
    }
  ]
}

response:

{
  "jsonrpc": "2.0",
  "result": {
    "data": {
      "objectId": "0x5ea6f7a348f4a7bd1a9ab069eb7f63865de3075cc5a4e62432f634b50fd2bb2b",
      "version": "1",
      "digest": "FnxePMX8y7AqX5mRL4nCcK4xecSrpHrd85c3sJDmh5uG",
      "type": "0x0000000000000000000000000000000000000000000000000000000000000009::test::TestField",
      "owner": {
        "AddressOwner": "0x013d1eb156edcc1bedc3b1af1be1fe41671856fd3450dc5574abd53c793c9f22"
      },
      "previousTransaction": "Faiv4yqGR4HjAW8WhMN1NHHNStxXgP3u22dVPyvLad2z",
      "storageRebate": "100",
      "content": {
        "dataType": "moveObject",
        "type": "0x0000000000000000000000000000000000000000000000000000000000000009::test::TestField",
        "hasPublicTransfer": true,
        "fields": {}
      }
    }
  },
  "id": 1
}

*/
