/*
--------------------------------------------------
@Create 2021-09-14 23:14
@Author lpc
@Program shear-plate-transmit
@Describe 剪切板测试
--------------------------------------------------
@Version 1.0 2021-09-14
@Memo create this file
*/

package test

import (
	"github.com/atotto/clipboard"
	"io/ioutil"
)

// ReadTest 读取测试
func ReadTest() {
	old := ""
	for {
		str, err := clipboard.ReadAll()
		if err != nil {
			panic(err)
		}
		if old != str {
			println(str)
			old = str
		}
	}
}

// WriteFileTest 写文件测试
func WriteFileTest(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	WriteByteTest(file)
}

// WriteByteTest 写byte测试
func WriteByteTest(data []byte) {
	WriteTest(string(data))
}

// WriteTest 写测试
func WriteTest(str string) {
	err := clipboard.WriteAll(str)
	if err != nil {
		panic(err)
	}
}
