/*
--------------------------------------------------
@Create 2021-09-15 14:15
@Author lpc
@Program shear-plate-transmit
@Describe 并行通信
--------------------------------------------------
@Version 1.0 2021-09-15
@Memo create this file
*/

package coreplate

import (
	"log"
	"time"
)

var (
	// WriteChan 写数据包管道
	WriteChan chan DataPack
	// ReceiveChan 接收数据包管道
	ReceiveChan chan DataPack
	// 回执管道
	re chan bool
	// 自己传输的id
	myId int64
)

type PlateTransmitMulti struct{}

// NewTransmitMulti 实例化
func NewTransmitMulti() PlateTransmitMulti {
	log.Println("初始化chan")
	WriteChan = make(chan DataPack)
	re = make(chan bool)
	ReceiveChan = make(chan DataPack)
	log.Println("传输协程启动")
	// 接收消息的携程
	go receive()
	// 写消息的携程
	go trans()
	return PlateTransmitMulti{}
}

// Receive 接收数据
func (t *PlateTransmitMulti) Receive() []byte {
	pack := <-ReceiveChan
	data := *new([]byte)
	data = append(data, pack.Data...)
	for pack.Total != pack.Size && !pack.IsEOF {
		pack = <-ReceiveChan
		data = append(data, pack.Data...)
	}
	return data
}

// Send 发送数据
func (t *PlateTransmitMulti) Send(data []byte) {
	total := int64(len(data))
	if total > GetMaxSize() {
		for s, e := int64(0), GetMaxSize(); ; {
			pack := NewDataPack(e, total, data[s:e])
			WriteChan <- pack
			s = e
			if total-e >= 0 {
				// 下滑一个 max
				e = e + GetMaxSize()
			} else {
				// 否则把末尾发送后直接完毕
				pack = NewDataPack(total, total, data[s:], true)
				WriteChan <- pack
				break
			}
		}
	} else {
		// 否则直接发送
		pack := NewDataPack(total, total, data, true)
		WriteChan <- pack
	}
	log.Println("发送成功")
}

// 并行传输数据
func trans() {
	for {
		d := <-WriteChan
		myId = d.Id
		write(&d)
		<-re // 等待回执后写下一条 保证传输成功
	}
}

// 并行获取回执
func receive() {
	var myReId int64   // 自己的回执id
	var lastReid int64 // 上一条读id
	var lastId int64   // 上一条写id
	for {
		pack := read()
		if pack != nil {
			if pack.ReBack && lastReid != pack.Id && myReId != pack.Id {
				// 回执
				lastReid = pack.Id
				re <- true
			} else if !pack.ReBack && lastId != pack.Id && myId != pack.Id {
				// 消息
				lastId = pack.Id
				ReceiveChan <- *pack
				// 写回执
				black := NewDataPackReBlack()
				myReId = black.Id
				write(&black)
			}
		}
		time.Sleep(GetDelayTime())
	}
}
