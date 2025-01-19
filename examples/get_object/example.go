package get_object

import (
	"encoding/json"
	"fmt"

	"github.com/xchgn/suigo/client"
)

func Run() {
	// Create a new client
	c := client.NewClient(client.MAINNET_URL)

	// Get the object
	info, err := c.GetObject("0xf66d0ba03abedf834a4f51f50c3fd457e320261815152b1e9ea7ead6c26bdce6", client.GetObjectShowOptions{
		ShowType:                true,
		ShowOwner:               true,
		ShowPreviousTransaction: true,
		ShowDisplay:             true,
		ShowContent:             true,
		ShowBcs:                 true,
		ShowStorageRebate:       true,
	})

	bsInfo, _ := json.MarshalIndent(info, "", "  ")

	fmt.Println("Result:", string(bsInfo))

	if err != nil {
		fmt.Println("Error:", err)
	}
}
