package exec_ptb

import (
	"encoding/json"
	"fmt"

	"github.com/ipoluianov/suigo/client"
)

func Run() {
	cl := client.NewClient(client.TESTNET_URL)
	cl.InitAccountFromFile("seed_phrase.txt")
	tb := client.NewTransactionBuilder(cl)
	for i := 0; i < 1; i++ {
		cmd := client.NewTransactionBuilderMoveCall()
		cmd.PackageId = client.TEST_PACKAGE_ID
		cmd.ModuleName = "fund"
		cmd.FunctionName = "ex4"
		cmd.Arguments = []interface{}{
			client.ArgSharedObject(client.TEST_FUND_ID),
			client.ArgSharedObject(client.CLOCK_OBJECT_ID),
		}
		tb.AddCommand(cmd)
	}
	res, err := cl.ExecPTB(tb, client.DEFAULT_GAS_PRICE)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	bs, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(bs))
	fmt.Println("Status:", res.Effects.Status.Status)
	//fmt.Println("Result Digest:", res.Digest)
}

// sui client ptb --move-call 0xef1846d0644254415d1165fe973cceaf0f0a0a3da012dedc8b2fe6c4475e4889::fund::ex2 @0x328e39beb44c849a79a0030f90afdd1f6d91457057cda077c942e6693ea22fc3
