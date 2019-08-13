/*******************************************************
 *
 * Copyright (c) 2020 God Coming, Inc. All Rights Reserved
 * 
 * @Author: Alihanniba
 *
 * @date: 2019/8/12 16:16
 *
 *******************************************************/

package server

import (
	"context"
	"go.intra.alihanniba.com/alrpc/grpc/data"
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
)

var port = ":3300"

type Server struct {

}

func (s *Server) Hello(ctx context.Context, req *data.Args) (*data.Reply, error) {
	var res data.Reply
	res.Title = fmt.Sprintf("%s 今年 %d 了。", req.GetName(), req.GetAge())
	return &res, nil
}

func Init () {
	ne, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ser:= grpc.NewServer()
	data.RegisterRpcServiceServer(ser, &Server{})
	if err := ser.Serve(ne); err != nil {
		log.Fatalf("failed to server %v", err)
	}

}