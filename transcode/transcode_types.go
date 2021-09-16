/*
--------------------------------------------------
@Create 2021-09-15 8:54
@Author lpc
@Program shear-plate-transmit
@Describe 数据编码结构体
--------------------------------------------------
@Version 1.0 2021-09-15
@Memo create this file
*/

package transcode

// TFData 传输文件结构体
type TFData struct {
	Name string
	Val  []byte
	// 是否是文件夹
	IsDir bool
	// 如果是文件夹则内
	DirVal []TFData
}

// NewTfJson 实例化
func NewTfJson() TFData {
	t := TFData{}
	return t
}
