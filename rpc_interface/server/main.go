/*
 * @Date: 2022-10-25 15:10:30
 * @LastEditors: chanyt chanyt166@163.com
 * @LastEditTime: 2022-10-25 15:52:03
 * @FilePath: /rpc-test/rpc_interface/server/main.go
 */
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/adgm2016/rpc-test/rpc_interface/service"
)

var _ service.HelloService = (*HelloService)(nil)

type HelloService struct {
}

func (s *HelloService) Hello(request string, response *string) error {
	*response = fmt.Sprintf("hello %s", request)
	return nil
}
func main() {

	//把rpc注册
	rpc.RegisterName(service.SERVICE_NAME, &HelloService{})
	//1. 先监听,
	listener, err := net.Listen("tcp", ":9527")
	if err != nil {
		log.Fatal("ListenTCP error : ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go rpc.ServeConn(conn)
	}

}
