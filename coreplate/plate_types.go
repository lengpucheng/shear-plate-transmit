/*
--------------------------------------------------
@Create 2021-09-14 23:04
@Author lpc
@Program shear-plate-transmit
@Describe 剪切板数据传输的基本数据包定义
--------------------------------------------------
@Version 1.0 2021-09-14
@Memo create this file
*/

package coreplate

import "time"

// DataPack 数据包
type DataPack struct {
	Id     int64  `json:"Id"`
	Tid    int64  `json:"Tid"`
	Size   int64  `json:"Size"`
	Total  int64  `json:"Total"`
	Data   []byte `json:"Data"`
	IsEOF  bool   `json:"is_eof"`
	ReBack bool   `json:"re_back"`
}

// NewDataPack 初始化数据包
func NewDataPack(size int64, total int64, data []byte, end ...bool) DataPack {
	id := time.Now().UnixNano()
	return DataPack{
		Id:     id,
		Tid:    0,
		Size:   size,
		Total:  total,
		Data:   data,
		IsEOF:  end != nil && len(end) >= 1 && end[0],
		ReBack: false,
	}

}

// NewDataPackReBlack 接收回执
func NewDataPackReBlack() DataPack {
	pack := NewDataPack(0, 0, nil, true)
	pack.ReBack = true
	return pack
}
