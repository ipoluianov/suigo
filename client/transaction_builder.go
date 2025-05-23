package client

import (
	"encoding/base64"
	"errors"

	"github.com/ipoluianov/suigo/txdata"
)

type TransactionBuilder struct {
	client    *Client
	gasBudget uint64

	commands []*TransactionBuilderMoveCall

	transactionData *txdata.TransactionData
	objectRefs      map[string]int
}

func NewTransactionBuilder(client *Client) *TransactionBuilder {
	var c TransactionBuilder
	c.client = client
	c.gasBudget = uint64(10000000)
	c.objectRefs = make(map[string]int)
	return &c
}

func (c *TransactionBuilder) AddCommand(cmd *TransactionBuilderMoveCall) {
	c.commands = append(c.commands, cmd)
}

func (c *TransactionBuilder) Build(gasPrice uint64) (string, error) {
	if c.client.account == nil {
		return "", errors.New("account is not set")
	}

	c.transactionData = txdata.NewTransactionData()
	c.transactionData.V1 = &txdata.TransactionDataV1{}
	senderAddrBS := ParseAddress(c.client.account.Address)
	c.transactionData.V1.Sender = senderAddrBS
	c.transactionData.V1.Expiration = txdata.NewTransactionExpiration()

	var gasData txdata.GasData
	gasData.Owner = senderAddrBS
	gasData.Price = gasPrice
	gasData.Budget = c.gasBudget
	// Get GAS coin information
	gasCoinObj, err := c.client.GetGasCoinObj(c.gasBudget)
	if err != nil {
		return "", err
	}
	var payment txdata.ObjectRef
	payment.ObjectID.SetHex(gasCoinObj.ObjectId)
	payment.ObjectDigest.SetBase58(gasCoinObj.Digest)
	payment.SequenceNumber = txdata.SequenceNumber(gasCoinObj.SeqNum)
	gasData.Payment = append(gasData.Payment, payment)
	c.transactionData.V1.GasData = &gasData

	c.transactionData.V1.Kind = &txdata.TransactionKind{}
	c.transactionData.V1.Kind.Type = txdata.ProgrammableTransactionType
	c.transactionData.V1.Kind.ProgrammableTransaction = &txdata.ProgrammableTransaction{}

	for _, cmd := range c.commands {
		cmd.Build(c)
	}

	bs := c.transactionData.ToBytes()

	return base64.StdEncoding.EncodeToString(bs), nil
}
