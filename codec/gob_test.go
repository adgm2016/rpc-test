/*
 * @Date: 2022-10-25 16:35:40
 * @LastEditors: chanyt chanyt166@163.com
 * @LastEditTime: 2022-10-25 16:42:41
 * @FilePath: /rpc-test/codec/gob_test.go
 */
package codec_test

import (
	"fmt"
	"testing"

	"github.com/adgm2016/rpc-test/codec"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	F1 string
	F2 int
}

func TestGob(t *testing.T) {
	should := assert.New(t)
	gobBytes, err := codec.GobEncode(&TestStruct{
		F1: "test_f1",
		F2: 10,
	})
	if should.NoError(err) {
		fmt.Println(gobBytes)
	}

	obj := &TestStruct{}
	err = codec.GobDecode(gobBytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj)
	}
}
