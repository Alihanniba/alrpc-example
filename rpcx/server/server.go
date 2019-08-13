/*******************************************************
 *
 * Copyright (c) 2018 xxx, Inc. All Rights Reserved
 * 
 * @Author: Alihanniba
 *
 * @date: 2019/8/12 11:55
 *
 *******************************************************/

package server

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/server"

	"flag"
)

type Args struct {
	Age int
	Name string
}

type Reply struct {
	Title string
}

type Ali int


func (ali *Ali) GetTag(ctx context.Context, args *Args, reply *Reply) error {
	reply.Title = fmt.Sprintf("%s 今年 %d 岁", args.Name, args.Age)
	return nil
}

func Init() {
	flag.Parse()

	s := server.NewServer()
	s.Register(new(Ali),"")
	err := s.Serve("tcp", ":3200")
	if err != nil {
		panic(err)
	}

}