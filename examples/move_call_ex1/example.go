package move_call_ex1

import (
	"fmt"

	"github.com/xchgn/suigo/client"
)

func Run() {
	cl := client.NewClient(client.MAINNET_URL)
	err := cl.InitAccountFromFile("seed_phrase.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	var p client.MoveCallParameters
	p.PackageId = client.TEST_PACKAGE_ID
	p.ModuleName = "example"
	p.FunctionName = "ex01"
	p.Arguments = []interface{}{}
	res, err := cl.ExecMoveCall(p)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("RESULT:", res)
}
