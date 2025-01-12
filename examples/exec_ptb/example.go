package exec_ptb

import (
	"encoding/json"
	"fmt"

	"github.com/xchgn/suigo/client"
)

func Run() {
	cl := client.NewClient(client.MAINNET_URL)
	cl.InitAccountFromFile("seed_phrase.txt")
	tb := client.NewTransactionBuilder(cl)
	for i := 0; i < 1; i++ {
		cmd := client.NewTransactionBuilderMoveCall()
		cmd.PackageId = client.TEST_PACKAGE_ID
		cmd.ModuleName = "example"
		cmd.FunctionName = "ex1"
		cmd.Arguments = []interface{}{}
		tb.AddCommand(cmd)
	}
	res, err := cl.DryExecPTB(tb)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Status:", res.Effects.Status.Status)
	bs, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(bs))
	//fmt.Println("Result Digest:", res.Digest)
}
