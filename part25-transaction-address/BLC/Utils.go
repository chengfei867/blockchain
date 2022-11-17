package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

// JsonToArray 字符串转换成出租
func JsonToArray(str string) []string {
	var sArr []string
	if err := json.Unmarshal([]byte(str), &sArr); err != nil {
		log.Panicln(err)
	}
	return sArr
}

// ReverseBytes 字节数组反转
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
