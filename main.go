package main

import (
	"fmt"
	"idn/router"
	"idn/utils/helper"
)

func main() {
	listenAddress := helper.GetEnv("SERVER_PORT", ":40001")
	fmt.Println("Starting listen address: ", listenAddress)
	router.Server(listenAddress)
}
