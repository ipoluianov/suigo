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
	p.PackageId = "0xbe66e3956632c8b8cb90211ecb329b9bb03afef9ba5d72472a7c240d3afe19fd"
	p.ModuleName = "example"
	p.FunctionName = "ex1"
	p.Arguments = []interface{}{}
	res, err := cl.ExecMoveCall(p)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("RESULT:", res)
}
