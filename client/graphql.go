package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func GetInitialSharedVersion(apiAddr string, sharedObjectAddress string) (uint64, error) {

	url := "https://sui-testnet.mystenlabs.com/graphql"
	query := `
{
  object(address: "<object_address>") {
    version
    owner {
      __typename
      ... on Shared {
        initialSharedVersion
      }
    }
  }
}
	`

	query = strings.ReplaceAll(query, "<object_address>", sharedObjectAddress)

	variables := map[string]interface{}{
		"id": "1",
	}

	requestBody := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")

	type InitialSharedVersionResponse struct {
		Data struct {
			Object struct {
				Owner struct {
					Typename             string `json:"__typename"`
					InitialSharedVersion uint64 `json:"initialSharedVersion"`
				} `json:"owner"`
				Version int `json:"version"`
			} `json:"object"`
		} `json:"data"`
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return 0, err
	}

	bsResp, _ := json.MarshalIndent(responseData, "", "  ")

	var res InitialSharedVersionResponse
	err = json.Unmarshal(bsResp, &res)
	if err != nil {
		log.Fatal(err)
	}

	return res.Data.Object.Owner.InitialSharedVersion, nil
}
