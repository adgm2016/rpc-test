/*
 * @Date: 2022-10-26 10:07:00
 * @LastEditors: chanyt chanyt166@163.com
 * @LastEditTime: 2022-10-26 10:14:01
 * @FilePath: /rpc-test/protobuf/pb/hello_test.go
 */
package pb_test

import (
	"fmt"
	"testing"

	"github.com/adgm2016/rpc-test/protobuf/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestMarshal(t *testing.T) {
	should := assert.New(t)
	str := &pb.String{Value: "hello"}

	pbBytes, err := proto.Marshal(str)
	if should.NoError(err) {
		fmt.Println(pbBytes)
	}
	obj := pb.String{}
	err = proto.Unmarshal(pbBytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj.Value)
	}
}
