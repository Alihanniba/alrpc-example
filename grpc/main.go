package main

import (
	"go.intra.alihanniba.com/alrpc/grpc/server"
	"go.intra.alihanniba.com/alrpc/grpc/client"
)

func main() {
	go func() {
		server.Init()
	}()

	client.Init()
}