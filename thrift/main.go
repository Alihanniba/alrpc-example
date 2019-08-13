package main

import (
	"go.intra.alihanniba.com/alrpc/thrift/server"
	"go.intra.alihanniba.com/alrpc/thrift/client"
)

func main() {
	go func() {
		server.Init()
	}()
	client.Init()
}