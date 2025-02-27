package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ipoluianov/gomisc/logger"
)

var mtx sync.Mutex
var counter int

func processRpc(url string) {
	//cl := client.NewClient(url)
	for {
		/*opts := client.GetObjectShowOptions{
			ShowType:    true,
			ShowContent: true,
		}
		obj, err := cl.GetObject("0xe01243f37f712ef87e556afb9b1d03d0fae13f96d324ec912daffc339dfdcbd2", opts)
		if err != nil {
			logger.Println("==================================== ERROR ====================================", url, err)
			return
		}

		mtx.Lock()
		counter++
		mtx.Unlock()

		_ = obj*/

		//fmt.Println(obj.Data.Content)
		//logger.Println(obj.Data.Version, url)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	fmt.Println("STARTED")

	/*go processRpc("https://fullnode.mainnet.sui.io/")
	go processRpc("https://sui-mainnet.public.blastapi.io/")
	go processRpc("https://cetus-mainnet-endpoint.blockvision.org/")
	go processRpc("https://fullnode.mainnet.sui.io/")
	go processRpc("https://rpc-mainnet.suiscan.xyz/")
	go processRpc("https://sui-rpc.publicnode.com/")
	go processRpc("https://mainnet.sui.rpcpool.com/")
	go processRpc("https://sui-mainnet.nodeinfra.com/")
	go processRpc("https://sui-mainnet-us-2.cosmostation.io/")
	go processRpc("https://sui-mainnet-us-1.cosmostation.io/")
	go processRpc("https://sui-mainnet-ca-2.cosmostation.io/")
	go processRpc("https://sui-mainnet-ca-1.cosmostation.io/")*/
	go processRpc("https://sui-mainnet-rpc.nodereal.io")

	lastCounter := 0
	dtLast := time.Now()
	for {
		time.Sleep(1 * time.Second)
		mtx.Lock()
		currentCounter := counter
		dtNow := time.Now()
		dtDiff := dtNow.Sub(dtLast)
		dtLast = dtNow
		countsPerSecond := float64(currentCounter-lastCounter) / dtDiff.Seconds()
		lastCounter = currentCounter
		mtx.Unlock()
		logger.Println("Counter: ", currentCounter, " CPS: ", countsPerSecond)
	}
}
