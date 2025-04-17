package movecallclock

import (
	"fmt"

	"github.com/ipoluianov/suigo/client"
)

const (
	PACKAGE_ID     = "0x7674735315216af4fc71ef20bb5775ad3b1e1f162c19150f0da9772f142528dd"
	FUND_OBJECT_ID = "0xfe53b68fbce0daa159c7abe893633584835bd398f9b6b8612a6d72fd72e9f1ff"
	XCHG_ADDR      = "0x9337d82d8b18a3fdf294d020319483e3f383716fbe89490955ff71b4cb518a76"
)

func Run() {
	cl := client.NewClient(client.TESTNET_URL)
	err := cl.InitAccountFromFile("seed_phrase.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	var p client.MoveCallParameters
	p.PackageId = PACKAGE_ID
	p.ModuleName = "fund"
	p.FunctionName = "get_cheques_ids"
	p.Arguments = []interface{}{
		client.ArgSharedObject(FUND_OBJECT_ID),
		client.ArgAddress(XCHG_ADDR),
		client.ArgU32(10),
		client.ArgSharedObject(client.CLOCK_OBJECT_ID),
	}
	res, err := cl.ExecMoveCall(p, 1000)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("RESULT:", res)
}
