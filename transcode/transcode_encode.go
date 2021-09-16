/*
--------------------------------------------------
@Create 2021-09-15 8:56
@Author lpc
@Program shear-plate-transmit
@Describe 将文件进行编码
--------------------------------------------------
@Version 1.0 2021-09-15
@Memo create this file
*/

package transcode

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Save 将路径文件（夹）编码并保存到指定文件
func Save(path string, save ...string) {
	file := "save.json"
	if save != nil && len(save) >= 1 && save[0] != "" {
		file = save[0]
	}
	err := ioutil.WriteFile(file, EncodeBytes(path), 0666)
	if err != nil {
		log.Panicln(err)
	}
}

// EncodeBytes 将路径编码成[]byte json
func EncodeBytes(path string) []byte {
	tf := encode(path)
	jsons, err := json.Marshal(tf)
	if err != nil {
		log.Panicln(err)
	}
	return jsons
}

// encode 将路径下的文件或文件夹编码
func encode(path string) TFData {
	log.Printf("[SAVE]Path: %s", path)
	tf := NewTfJson()
	stat, err := os.Stat(path)
	if err != nil {
		log.Panicln(err)
	}
	sep := string(os.PathSeparator)
	tf.Name = stat.Name()
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)

	defer func() {
		errf := file.Close()
		if err != nil {
			log.Panicln(errf)
		}
	}()

	if stat.IsDir() {
		// dir
		tf.IsDir = true
		dir, err := file.ReadDir(-1)
		if err != nil {
			log.Panicln(err)
		}
		for _, info := range dir {
			infoPath := path + sep + info.Name()
			tf.DirVal = append(tf.DirVal, encode(infoPath))
		}
	} else {
		if err != nil {
			log.Panicln(err)
		}
		all, _ := ioutil.ReadAll(file)
		if err == io.EOF {
			log.Panicln(err)
		}
		tf.Val = all
	}
	return tf
}
