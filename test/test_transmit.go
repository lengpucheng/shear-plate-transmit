/*
--------------------------------------------------
@Create 2021-09-14 23:34
@Author lpc
@Program shear-plate-transmit
@Describe 剪切板通信的测试
--------------------------------------------------
@Version 1.0 2021-09-14
@Memo create this file
*/

package test

import (
	"flag"
	"github.com/lengpucheng/shear-plate-transmit/coreplate"
	"io/ioutil"
	"log"
)

var name string
var cmd bool

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.StringVar(&name, "f", "", "传输的文件")
	flag.BoolVar(&cmd, "s", false, "是否为显示")
}

// TransmitTest 传输测试
func TransmitTest(shows bool) {
	flag.Parse()
	log.Println("TFS已启动.....")
	if cmd || shows {
		show(name)
	} else {
		send(name)
	}
}

// 显示操作
func show(name string) {
	for {
		log.Println("显示模式.....")
		transmit := coreplate.NewTransmitMulti()
		bytes := transmit.Receive()
		log.Printf(string(bytes))
		log.Printf("接收完毕")
	}
}

// 发送操作
func send(name string) {
	log.Printf("发送模式，路径文件为 %s", name)
	file, err := ioutil.ReadFile(name)
	if err != nil {
		log.Panicln(err)
	}
	transmit := coreplate.NewTransmitMulti()
	transmit.Send(file)
	for {
	}
}

// WriteSendTest 写发送测试
func WriteSendTest(path string) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		log.Panicln(err)
	}
	pack := coreplate.NewDataPack(723, 723, file, true)
	coreplate.TestWrite(&pack)
}
