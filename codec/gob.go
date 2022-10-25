/*
 * @Date: 2022-10-25 16:27:44
 * @LastEditors: chanyt chanyt166@163.com
 * @LastEditTime: 2022-10-25 16:34:34
 * @FilePath: /rpc-test/codec/gob.go
 */
package codec

import (
	"bytes"
	"encoding/gob"
)

// object -gob -> []byte
func GobEncode(obj interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	//编码后的结果输出到buf中
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(obj); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func GobDecode(data []byte, obj interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	return decoder.Decode(obj)
}
