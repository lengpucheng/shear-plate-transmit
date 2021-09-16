/*
--------------------------------------------------
@Create 2021-09-15 13:53
@Author lpc
@Program shear-plate-transmit
@Describe 读写分离的传输 串行通信
--------------------------------------------------
@Version 1.0 2021-09-15
@Memo create this file
*/

package coreplate

import (
	"log"
	"time"
)

var lastID int64

// 回执获取尝试次数
const trySum = 100

// PlateTransmitSingle 串行传输
type PlateTransmitSingle struct{}

// NewTransmitSingle 实例化
func NewTransmitSingle() PlateTransmitSingle {
	return PlateTransmitSingle{}
}

// Receive 接收数据
func (t *PlateTransmitSingle) Receive() []byte {
	log.Println("开始接收数据")
	data := *new([]byte)
	for {
		pack := read()
		log.Printf("读取ing.....--->%d", pack.Id)
		if pack != nil && pack.Id != 0 && !pack.ReBack && lastID != pack.Id {
			showProcess(pack)
			lastID = pack.Id
			data = append(data, pack.Data...)
			callBack()
			if pack.IsEOF || pack.Size == pack.Total {
				log.Println("数据接收完成")
				return data
			}
		}
		time.Sleep(GetDelayTime())
	}
}

// Send 发送数据
func (t *PlateTransmitSingle) Send(data []byte) {
	log.Println("开始发送数据")
	if data == nil || len(data) == 0 {
		log.Println("数据为空，传输失败")
		return
	}
	total := int64(len(data))
	var s int64 = 0
	var e int64
	for {
		if total-s <= GetMaxSize() {
			pack := NewDataPack(total, total, data[s:], true)
			showProcess(&pack)
			writeTry(&pack)
			log.Println("数据发送完成")
			break
		} else {
			e = s + GetMaxSize()
			pack := NewDataPack(e, total, data[s:e], false)
			showProcess(&pack)
			writeTry(&pack)
			s = e
		}
	}
}

// 尝试写
func writeTry(pack *DataPack) {
	for {
		write(pack)
		if getCallBack() {
			// 获得了回执就返回 否则一直尝试传输数据
			return
		}
	}
}

// 获取回执
func getCallBack() bool {
	var sum = 0
	for sum <= trySum {
		pack := read()
		if pack != nil && pack.Id != 0 && pack.ReBack && lastID != pack.Id {
			lastID = pack.Id
			log.Printf("读取回执ing.....--->%d", pack.Id)
			return true
		}
		sum++
		time.Sleep(GetDelayTime() / 2)
	}
	// 尝试后未获取则重试
	return false
}

// 写回执
func callBack() {
	pack := NewDataPackReBlack()
	log.Printf("写回执ing.....--->%d", pack.Id)
	write(&pack)
}

// 显示进度
func showProcess(pack *DataPack) {
	f := (float64(pack.Size) / float64(pack.Total)) * 100
	log.Printf("正在收发数据包.....ID--->%d", pack.Id)
	log.Printf("当前进度--[%3.2f %%]---> %d/%d", f, pack.Size, pack.Total)
}
