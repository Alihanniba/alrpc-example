package server

import (
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"fmt"
	"go.intra.alihanniba.com/alrpc/thrift/data/gen-go/data"
	"context"
)

var (
	port = ":3400"
	bufferSize = 8192
)

type HelloServer struct {

}

func (h *HelloServer) Hello (ctx context.Context,req *data.Args_) (*data.Reply, error) {
	res := data.Reply{
		Title:fmt.Sprintf("%s 明年 %d 岁", req.Name, req.Age),
	}
	return &res, nil
}

func Init() {
	transportFactory := thrift.NewTBufferedTransportFactory(bufferSize)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(port)
	if err != nil {
		//panic
		log.Fatalf("create server socket err: %v", err)
	}
	handler := HelloServer{}
	processor := data.NewRpcServiceProcessor(&handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory,protocolFactory)

	log.Printf("service con")

	if err = server.Serve(); err != nil {
		log.Fatalf("service start err: %v", err)
	}
}