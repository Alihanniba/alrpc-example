/*******************************************************
 *
 * Copyright (c) 2018 xxx, Inc. All Rights Reserved
 * 
 * @Author: Alihanniba
 *
 * @date: 2019/8/9 12:01
 *
 *******************************************************/

package proto

import (
	"github.com/gogo/protobuf/proto"
)

func AMarshal(da RpcData) ([]byte, error) {
	buf, err := proto.Marshal(&da);
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func AUnmarshal(buf []byte) (RpcData, error) {
	var rpcData RpcData
	if err := proto.Unmarshal(buf, &rpcData); err != nil {
		return rpcData, err
	}
	return rpcData, nil
}