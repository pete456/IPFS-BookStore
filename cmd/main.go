package main

import (
	"bs/internal/app"
	"bs/internal/com"
	"bs/internal/ipfs"
	"bs/internal/logger"

	"fmt"
)

func main() {
	Service := new(app.Service)
	Service.Log = logger.NewLog()
	ipfs, err := ipfs.New()
	if err != nil {
		Service.Log.Fatal(fmt.Sprintf("Failed to initialize ipfs: %s",err))
	}
	Service.IPFS = ipfs

	err = com.StartServer(Service,":8080")
	if err != nil {
		Service.Log.Fatal(fmt.Sprintf("Failed to initialize ipfs: %s",err))
	}
}
