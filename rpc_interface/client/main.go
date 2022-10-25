/*
 * @Date: 2022-10-25 15:10:44
 * @LastEditors: chanyt chanyt166@163.com
 * @LastEditTime: 2022-10-25 15:58:26
 * @FilePath: /rpc-test/rpc_interface/client/main.go
 */
package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/adgm2016/rpc-test/rpc_interface/service"
)

var _ service.HelloService = (*HelloServiceClient)(nil)

func NewHelloServiceClient(method, ip string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(method, ip)
	if err != nil {
		log.Fatal(err)
	}
	return &HelloServiceClient{
		cliect: client,
	}, nil
}

type HelloServiceClient struct {
	cliect *rpc.Client
}

func (c *HelloServiceClient) Hello(request string, response *string) error {
	return c.cliect.Call(fmt.Sprintf("%s.Hello", service.SERVICE_NAME), request, response)

}

func main() {

	c, err := NewHelloServiceClient("tcp", ":9527")
	if err != nil {
		panic(err)
	}

	var resp string
	if err := c.Hello("bob", &resp); err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

}
