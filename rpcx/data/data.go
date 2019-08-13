/*******************************************************
 *
 * Copyright (c) 2018 xxx, Inc. All Rights Reserved
 * 
 * @Author: Alihanniba
 *
 * @date: 2019/8/12 11:54
 *
 *******************************************************/

package data

type Args struct {
	Age int `msg:"age"`
	Name string `msg:"name"`
}

type Reply struct {
	Title string `msg:"title"`
}

