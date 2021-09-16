/*
--------------------------------------------------
@Create 2021-09-14 23:09
@Author lpc
@Program shear-plate-transmit
@Describe 剪切板读写操作
--------------------------------------------------
@Version 1.0 2021-09-14
@Memo create this file
*/

package coreplate

import (
	"encoding/json"
	"github.com/atotto/clipboard"
	"log"
)

// 异常字符串
var errStr string

// write 写入数据到剪切板
func write(data *DataPack) {
	jsons, err := json.Marshal(*data)
	if err != nil {
		log.Panicln(err)
	}
	str := string(jsons)
	err = clipboard.WriteAll(str)
	if err != nil {
		log.Println(err)
	}
}

// read 从剪切板读取
func read() (pack *DataPack) {
	pack = &DataPack{}
	str, err := clipboard.ReadAll()
	if err != nil {
		log.Printf("[ERROR]读取剪切板错误,str=%s,错误为:\n%v\n", str, err)
		return
	}
	// 相同的错误就忽略
	if str != errStr {
		err = json.Unmarshal([]byte(str), pack)
		if err != nil {
			errStr = str
			log.Printf("[ERROR]读取剪切板转换,pake=%v \n错误:%v\n", pack, err)
			return
		}
	}
	return
}

// TestWrite 写测试
func TestWrite(data *DataPack) {
	write(data)
}

// TestRead 读测试
func TestRead() *DataPack {
	return read()
}
