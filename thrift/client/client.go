package client

import (
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"go.intra.alihanniba.com/alrpc/thrift/data/gen-go/data"
	"context"
)

var (
	bufferSize = 8192
	addr = "localhost:3400"
)

func Init() {
	transportFactory := thrift.NewTBufferedTransportFactory(bufferSize)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(addr)
	if err != nil {
		panic(err)
	}

	transportServer, err := transportFactory.GetTransport(transport)
	if err != nil {
		panic(err)
	}
	defer transportServer.Close()



	if err := transportServer.Open(); err  != nil {
		log.Fatalf("transport server open failed, err: %v", err)
	}

	iprot := protocolFactory.GetProtocol(transportServer)
	oprot := protocolFactory.GetProtocol(transportServer)

	client := data.NewRpcServiceClient(thrift.NewTStandardClient(iprot, oprot))

	req := data.Args_{
		Age:12,
		Name:"Alihanniba",
	}

	res, err := client.Hello(context.Background(), &req)
	if err != nil {
		log.Fatalf("client hello func called fail, err: %v", err)
	}

	log.Printf("call success, resut is \" %s \"", res.Title)

}