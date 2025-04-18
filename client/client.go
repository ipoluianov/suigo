package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Client struct {
	rpcURL                string
	account               *Account
	initialSharedVersions map[string]uint64
}

func NewClient(rpcURL string) *Client {
	var c Client
	c.rpcURL = rpcURL
	c.initialSharedVersions = make(map[string]uint64)
	// c.account, _ = NewAccountFromMnemonic("")
	return &c
}

type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *RPCError       `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *Client) InitAccountFromFile(filePath string) error {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	c.account, err = NewAccountFromMnemonic(string(bs))
	return err
}

func (c *Client) Account() *Account {
	return c.account
}

func (c *Client) GetInitialSharedVersion(sharedObjectAddress string) (uint64, error) {
	if ver, ok := c.initialSharedVersions[sharedObjectAddress]; ok {
		return ver, nil
	}

	ver, err := GetInitialSharedVersion(TESTNET_GRAPHQL_URL, sharedObjectAddress)
	if err != nil {
		return 0, err
	}

	return ver, nil
}

func (c *Client) rpcCall(request RPCRequest) (response *RPCResponse, err error) {
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return
	}

	resp, err := http.Post(c.rpcURL, "application/json", bytes.NewBuffer(requestJSON))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return
	}

	return
}
