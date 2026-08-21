package main

import (
	"Coal_company/factory"
	"Coal_company/factory/items"
	"Coal_company/factory/miners"
	"Coal_company/server"
	"fmt"
)

func main() {
	miners := miners.NewMiners()
	items := items.NewItems()
	factory := factory.NewFactory(*items, *miners)
	httpHandlers := server.NewHTTPHandlers(factory)
	server := server.NewSerber(httpHandlers)

	if err := server.ServerStart(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
