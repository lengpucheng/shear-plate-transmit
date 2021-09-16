/*
--------------------------------------------------
@Create 2021-09-15 13:54
@Author lpc
@Program shear-plate-transmit
@Describe 核心接口和常量
--------------------------------------------------
@Version 1.0 2021-09-15
@Memo create this file
*/

package coreplate

import (
	"flag"
	"time"
)

var (
	// max 单读写的最大限制   测试最大  566kb<  s   <576KB
	max int64 = 512
	// delay 延时 最高速度为= MAXIMUM*delay/1000
	delay int64 = 500
)

func init() {
	flag.Int64Var(&max, "max", 512, "(KB)单片数据包最大上限，建议不大于640")
	flag.Int64Var(&delay, "time", 500, "(ms)剪切板通信延时时间，建议不小于128")
}

// PlateTransmit 剪切板传输核心接口
type PlateTransmit interface {
	Send([]byte)
	Receive() []byte
}

// GetMaxSize 获取单片最大大小
func GetMaxSize() int64 {
	if max <= 10 {
		max = 128
	}
	return 1024 * max
}

// SetMaxsize 设置单片限制大小
func SetMaxsize(limit int64) {
	max = limit
}

// GetDelayTime 获取延时时间
func GetDelayTime() time.Duration {
	if delay < 100 {
		delay = 100
	}
	return time.Millisecond * time.Duration(delay)
}

// SetDelayTime 设置延时时间
func SetDelayTime(time int64) {
	delay = time
}
