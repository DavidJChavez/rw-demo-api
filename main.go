package main

import (
	"DavidJChavez/rw-demo-api/cmd/server"
	"DavidJChavez/rw-demo-api/configs"
	"DavidJChavez/rw-demo-api/scripts"
	"os"
)

func main() {
	arg := os.Args[1]
	if arg == "generate" {
		scripts.Generate()
	}
	if arg == "dev" {
		configs.LoadEnv("./environments/dev.env")
		configs.ConnectDB()
		server.RunServer()
	}
}
