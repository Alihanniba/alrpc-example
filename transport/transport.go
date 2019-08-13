/*******************************************************
 *
 * Copyright (c) 2018 xxx, Inc. All Rights Reserved
 * 
 * @Author: Alihanniba
 *
 * @date: 2019/8/9 14:17
 *
 *******************************************************/

package transport

import (
	"net"
	"encoding/binary"
	"io"
)

type Transport struct {
	conn net.Conn
}

func NewTransport (conn net.Conn) *Transport {
	return &Transport{conn}
}

func (t *Transport) Send(data []byte) error {
	buf := make([]byte, 4+len(data))
	binary.BigEndian.PutUint32(buf[:4],uint32(len(data)))
	copy(buf[4:],data)
	_, err := t.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}


func (t *Transport) Read() ([]byte, error) {
	header := make([]byte, 4)
	_, err := io.ReadFull(t.conn, header)
	if err != nil {
		return nil, err
	}

	//在使用proto.Unmarshal(buf, message)对消息进行反序列化时，缓冲区buf的长度应当等于消息的实际长度。否则会报告如下错误消息
	//proto: protocol.Message: illegal tag 0 (wire type 0)
	dlen := binary.BigEndian.Uint32(header)
	data := make([]byte, dlen)
	_, err = io.ReadFull(t.conn,data)
	if err != nil {
		return nil, err
	}
	return data, nil
}


func (t *Transport) Write(data []byte) error {
	return nil
}


func (t *Transport) Close(data []byte) error {
	return nil
}
