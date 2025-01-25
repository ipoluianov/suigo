package table_item

import (
	"encoding/json"
	"fmt"

	"github.com/xchgn/suigo/client"
)

func Run() {
	// Create a new client
	c := client.NewClient(client.TESTNET_URL)

	// Get the object
	info, err := c.GetDynamicFieldObject(
		"0x280a7e3d8cad32cc6e7506d31eef826fffd6c5997c651aa7d1eb0486a4bff760", "u32", 1)

	bsInfo, _ := json.MarshalIndent(info, "", "  ")

	fmt.Println("Result:", string(bsInfo))

	if err != nil {
		fmt.Println("Error:", err)
	}
}
