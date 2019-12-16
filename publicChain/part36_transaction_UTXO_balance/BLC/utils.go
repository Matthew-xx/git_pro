package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.BigEndian,num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

//标准的json字符串转数组
func JsonToArray(jsonString string) []string {
	//从json到string
	var sArray []string
	if err := json.Unmarshal([]byte(jsonString),&sArray);err != nil {
		log.Panic(err)
	}
	return sArray
}

