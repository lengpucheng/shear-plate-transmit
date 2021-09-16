/*
--------------------------------------------------
@Create 2021-09-15 9:50
@Author lpc
@Program shear-plate-transmit
@Describe 文件传输测试
--------------------------------------------------
@Version 1.0 2021-09-15
@Memo create this file
*/

package test

import (
	"github.com/lengpucheng/shear-plate-transmit/coreplate"
	"github.com/lengpucheng/shear-plate-transmit/transcode"
)

// FileUploadTest 文件上传测试
func FileUploadTest(path string) {
	bytes := transcode.EncodeBytes(path)
	pack := coreplate.NewDataPack(10, 10, bytes, true)
	coreplate.TestWrite(&pack)
}

// FileDownloadTest 文件下载测试
func FileDownloadTest(path string) {
	pack := coreplate.TestRead()
	transcode.LoadForBytes(pack.Data, path)
}
