/*******************************************************
 *
 * Copyright (c) 2018 xxx, Inc. All Rights Reserved
 * 
 * @Author: Alihanniba
 *
 * @date: 2019/8/9 11:47
 *
 *******************************************************/

package client

import (
	"net"
	"reflect"
)

type Client struct {
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn}
}

func (c *Client) CallRpc(rn string, fPtr interface{}) {
	container := reflect.ValueOf(fPtr).Elem()

	container.Set(reflect.MakeFunc(container.Type(), f))
}