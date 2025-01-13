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
		cmd.FunctionName = "ex12"
		items := make([]string, 0)
		items = append(items, "0x1111")
		items = append(items, "0x2222")
		items = append(items, "0x3333")
		cmd.Arguments = []interface{}{
			client.ArgVecAddress(items),
		}
		tb.AddCommand(cmd)
	}
	res, err := cl.ExecPTB(tb)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	bs, _ := json.MarshalIndent(res, "", "  ")
	fmt.Println(string(bs))
	fmt.Println("Status:", res.Effects.Status.Status)
	//fmt.Println("Result Digest:", res.Digest)
}
