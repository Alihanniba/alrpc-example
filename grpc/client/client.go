package client

import (
	"google.golang.org/grpc"
	"log"
	"go.intra.alihanniba.com/alrpc/grpc/data"
	"context"
	"time"
)

const (
	addr = "localhost:3300"

)

func Init() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("not connect: %v", err)
	}
	defer conn.Close()

	c := data.NewRpcServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := data.Args{
		Age:20,
		Name:"Alihanniba",
	}

	res, err := c.Hello(ctx, &req)
	if err != nil {
		log.Fatalf("curl failed: %v", err)
	}
	log.Printf("success, result is:\" %s \"", res.Title)


}