package get_coin_info

import (
	"fmt"

	"github.com/ipoluianov/suigo/client"
)

func Run() {
	coinType := client.COIN_SUI
	cl := client.NewClient(client.MAINNET_URL)
	coinInfo, err := cl.GetCoinMetadata(coinType)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Coin Info:", coinType)
	fmt.Println("Name:", coinInfo.Name)
	fmt.Println("Symbol:", coinInfo.Symbol)
	fmt.Println("Decimals", coinInfo.Decimals)

	fmt.Println("Description:", coinInfo.Description)
	fmt.Println("IconURL:", coinInfo.IconUrl)
	fmt.Println("ObjectID:", coinInfo.Id)
}
