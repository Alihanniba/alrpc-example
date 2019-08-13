/*******************************************************
 *
 * Copyright (c) 2018 xxx, Inc. All Rights Reserved
 * 
 * @Author: Alihanniba
 *
 * @date: 2019/8/9 11:55
 *
 *******************************************************/

package server

import (
	"reflect"
	"log"
	"fmt"
	"net"
	"io"
	"go.intra.alihanniba.com/alrpc/proto"
	"go.intra.alihanniba.com/alrpc/transport"
)

type RpcServer struct {
	addr string
	funcs map[string]reflect.Value
}

func NewServer(addr string) *RpcServer {
	return &RpcServer{addr:addr,funcs:make(map[string]reflect.Value)}
}

func (rs *RpcServer) Register(fn string, ff interface{}) {
	if _, ok := rs.funcs[fn]; !ok {
		rs.funcs[fn] = reflect.ValueOf(ff)
	}
}

func (rs *RpcServer) Run() {
	//发起tcp请求
	l, err := net.Listen("tcp", rs.addr)
	if err != nil {
		log.Println(fmt.Sprintf("listen o %s err: %v \n", err))
		return
	}

	//监听
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(fmt.Sprintf("accept err: %v \n", err))
			continue
		}
		//握手之后
		go func() {
			connTransport := transport.NewTransport(conn)
			for {
				res, err := connTransport.Read()
				if err != nil && err != io.EOF {
					log.Println(fmt.Sprintf("read trans err: %v\n", err))
					return
				}

				umRes, err := proto.AUnmarshal(res)
				if err != nil {
					log.Println(fmt.Sprintf("unmarshal res err: %v\n", err))
					return
				}


				err = connTransport.Send(b)
				if err != nil {
					log.Println(fmt.Sprintf("transport write err: %v\n", err))
				}
			}
		}()
	}
}

