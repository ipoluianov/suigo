package client

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/ipoluianov/suigo/txdata"
)

type CoinObject struct {
	DataType string `json:"dataType"`
	Fields   struct {
		Balance string `json:"balance"`
		ID      struct {
			ID string `json:"id"`
		} `json:"id"`
	} `json:"fields"`
	HasPublicTransfer bool   `json:"hasPublicTransfer"`
	Type              string `json:"type"`
}

func (c *CoinObject) GetBalanceUint64() uint64 {
	var err error
	var balance uint64
	balance, err = strconv.ParseUint(c.Fields.Balance, 10, 64)
	if err != nil {
		balance = 0
	}
	return balance
}

func parseCoinObject(obj interface{}) *CoinObject {
	var coin CoinObject
	jsonData, _ := json.Marshal(obj)
	json.Unmarshal(jsonData, &coin)
	return &coin
}

/*func (c *Client) GetGasCoinObjId(amount uint64) string {
	var query ObjectResponseQuery
	query.Options.ShowType = true
	query.Options.ShowOwner = true
	query.Options.ShowContent = true
	query.AddMatchStructType("0x2::coin::Coin<0x2::sui::SUI>")

	res, err := c.GetOwnedObjects(c.account.Address, "", 0, query)
	if err != nil {
		return ""
	}
	for _, obj := range res.Data {
		fmt.Println("OBJ:", obj.Data.ObjectId, obj.Data.Type)
		if obj.Data.Type == "0x2::coin::Coin<0x2::sui::SUI>" {
			coinObj := parseCoinObject(obj.Data.Content)
			if coinObj.GetBalanceUint64() >= amount {
				return coinObj.Fields.ID.ID
			}
		}
	}
	return ""
}*/

type GasDataInfo struct {
	ObjectId string
	Balance  string
	Digest   string
	SeqNum   int64
}

func (c *Client) GetGasCoinObj(amount uint64) (*GasDataInfo, error) {
	var query ObjectResponseQuery
	query.Options.ShowType = true
	query.Options.ShowOwner = true
	query.Options.ShowContent = true
	query.AddMatchStructType("0x2::coin::Coin<0x2::sui::SUI>")

	res, err := c.GetOwnedObjects(c.account.Address, "", 0, query)
	if err != nil {
		return nil, err
	}
	for _, obj := range res.Data {
		if obj.Data.Type == "0x2::coin::Coin<0x2::sui::SUI>" {
			coinObj := parseCoinObject(obj.Data.Content)
			if coinObj.GetBalanceUint64() >= amount {
				var version int64
				version, err = strconv.ParseInt(obj.Data.Version, 10, 64)
				if err != nil {
					return nil, err
				}
				return &GasDataInfo{
					ObjectId: obj.Data.ObjectId,
					Balance:  coinObj.Fields.Balance,
					Digest:   obj.Data.Digest,
					SeqNum:   version,
				}, nil
			}
		}
	}
	return nil, errors.New("No gas coin found")
}

type MoveCallParameters struct {
	PackageId    string        `json:"package_id"`
	ModuleName   string        `json:"module_name"`
	FunctionName string        `json:"function_name"`
	Arguments    []interface{} `json:"arguments"`
}

func (c *Client) ExecMoveCall(params MoveCallParameters, gasPrice uint64) (*TransactionExecutionResult, error) {
	// Prepare gas coin
	//gasBudget := uint64(100000000)
	//gasCoinObjId := c.GetGasCoinObjId(gasBudget)

	tb := NewTransactionBuilder(c)
	cmd := NewTransactionBuilderMoveCall()
	cmd.PackageId = params.PackageId
	cmd.ModuleName = params.ModuleName
	cmd.FunctionName = params.FunctionName
	cmd.Arguments = params.Arguments
	tb.AddCommand(cmd)
	txBytes, err := tb.Build(gasPrice)
	if err != nil {
		fmt.Println("BUILD ERROR:", err)
		return nil, err
	}

	// Prepare TxBytes
	/*	txBytes2, err := c.UnsafeMoveCall(gasCoinObjId, fmt.Sprint(gasBudget), cmd.PackageId, cmd.ModuleName, cmd.FunctionName, cmd.Arguments)
		if err != nil {
			return nil, err
		}

		if len(txBytes) != len(txBytes2.TxBytes) {
			fmt.Println("TXBYTES LENGTH MISMATCH")
			return nil, errors.New("TXBYTES LENGTH MISMATCH")
		}

		mismatch := false
		for i := 0; i < len(txBytes); i++ {
			if txBytes[i] != txBytes2.TxBytes[i] {
				mismatch = true
				break
			}
		}
		if mismatch {
			fmt.Println("TXBYTES MISMATCH")
			return nil, errors.New("TXBYTES MISMATCH")
		}

		fmt.Println("TXBYTES MATCH")
	*/
	//txBytesBS, _ := base64.StdEncoding.DecodeString(txBytes)
	txBytesBS1, _ := base64.StdEncoding.DecodeString(txBytes)
	//txBytesBS2, _ := base64.StdEncoding.DecodeString(txBytes2.TxBytes)

	fmt.Println("TXBYTES1:", hex.EncodeToString(txBytesBS1))
	//fmt.Println("TXBYTES2:", hex.EncodeToString(txBytesBS2))

	trDataParsed := txdata.NewTransactionData()
	_, err = trDataParsed.Parse(txBytesBS1, 0)
	if err != nil {
		fmt.Println("PARSE ERROR:", err)
		return nil, err
	}
	//bsParsed, _ := json.MarshalIndent(trDataParsed, "", "  ")
	//fmt.Println("PARSING SUCCESS:", string(bsParsed))

	// Signature
	txSigned, err := c.account.Signature(txBytes)
	//txSigned, err := c.account.Signature(txBytes.TxBytes)
	if err != nil {
		return nil, err
	}

	// Execute
	result, err := c.ExecuteTransactionBlock(txSigned.TxBytes, txSigned.Signature)
	return result, err
}
