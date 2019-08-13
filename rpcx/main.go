/*******************************************************
 *
 * Copyright (c) 2018 xxx, Inc. All Rights Reserved
 * 
 * @Author: Alihanniba
 *
 * @date: 2019/8/12 11:55
 *
 *******************************************************/

package main

import (
	"go.intra.alihanniba.com/alrpc/rpcx/server"
	"go.intra.alihanniba.com/alrpc/rpcx/client"
)


func main() {

	go func() {
		server.Init()
	}()

	client.Init()
}
