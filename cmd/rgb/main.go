package main

import "ReGio/internal/server"
import "ReGio/internal/conf"

func main() {
	server.Start(conf.NewConfig())
}