/*
 * @Date: 2022-10-25 15:10:44
 * @LastEditors: chanyt chanyt166@163.com
 * @LastEditTime: 2022-10-25 15:36:06
 * @FilePath: /rpc-test/rpc_simple/client/main.go
 */
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:9527")
	if err != nil {
		log.Fatal(err)
	}
	var resp string
	err = client.Call("HelloService.Hello", "bob", &resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}
