/*
--------------------------------------------------
@Create 2021-09-15 9:13
@Author lpc
@Program shear-plate-transmit
@Describe 传输文件
--------------------------------------------------
@Version 1.0 2021-09-15
@Memo create this file
*/

package transmit

import (
	"github.com/lengpucheng/shear-plate-transmit/coreplate"
	"github.com/lengpucheng/shear-plate-transmit/transcode"
)

type TransFile struct {
	transmit coreplate.PlateTransmit
}

func NewTransmitFile(plateTransmit coreplate.PlateTransmit) TransFile {
	return TransFile{transmit: plateTransmit}
}

// Upload 上传传输
func (f *TransFile) Upload(path string) {
	bytes := transcode.EncodeBytes(path)
	f.transmit.Send(bytes)
}

// Download 下载到本地
func (f *TransFile) Download(path string) {
	bytes := f.transmit.Receive()
	transcode.LoadForBytes(bytes, path)
}
