/*
 * @Date: 2022-10-25 15:40:06
 * @LastEditors: chanyt chanyt166@163.com
 * @LastEditTime: 2022-10-25 15:51:44
 * @FilePath: /rpc-test/rpc_interface/service/interface.go
 */
package service

const (
	SERVICE_NAME = "HelloService"
)

type HelloService interface {
	Hello(request string, response *string) error
}
