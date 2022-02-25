package main

import (
	"flag"
	"fmt"
	"github.com/lengpucheng/shear-plate-transmit/coreplate"
	"github.com/lengpucheng/shear-plate-transmit/transmit"
)

var (
	t bool
	p string
)

func init() {
	flag.BoolVar(&t, "t", false, "是否为传输，不加默认为接收")
	flag.StringVar(&p, "p", "", "路径 传输则是需要传输的路径  否则为接收的数据加载路径")
}

func main() {
	flag.Parse()
	fmt.Println(Test())
	single := coreplate.NewTransmitSingle()
	trans := transmit.NewTransmitFile(&single)
	if t {
		trans.Upload(p)
	} else {
		trans.Download(p)
	}
}

func Test() string {
	return "ok"
}
