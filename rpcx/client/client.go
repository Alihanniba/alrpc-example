/*******************************************************
 *
 * Copyright (c) 2018 xxx, Inc. All Rights Reserved
 * 
 * @Author: Alihanniba
 *
 * @date: 2019/8/12 11:55
 *
 *******************************************************/

package client

import (
	"github.com/smallnest/rpcx/client"
	"flag"

	"go.intra.alihanniba.com/alrpc/rpcx/data"
	"context"
	"log"
)

var (
	addr = flag.String("addr", "localhost:3200", "server addr")
)

func Init() {
	flag.Parse()


	d := client.NewPeer2PeerDiscovery("tcp@" + *addr, "")

	xclient := client.NewXClient("Ali", client.Failtry, client.RandomSelect,d,client.DefaultOption)
	defer xclient.Close()

	args := data.Args{
		Age:18,
		Name:"Alihanniba",
	}
	reply := data.Reply{

	}

	err := xclient.Call(context.Background(), "GetTag", &args, &reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}



	log.Printf("call success, result is %s", reply.Title)

}