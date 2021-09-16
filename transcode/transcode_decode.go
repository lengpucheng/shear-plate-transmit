/*
--------------------------------------------------
@Create 2021-09-15 8:56
@Author lpc
@Program shear-plate-transmit
@Describe 对编码后的bytes进行解码操作
--------------------------------------------------
@Version 1.0 2021-09-15
@Memo create this file
*/

package transcode

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Load 从文件中加载并解码到指定目录
func Load(file string, path ...string) {
	load := ""
	if path != nil && len(path) >= 1 && path[0] != "" {
		load = path[0]
	}
	decode(DecodePath(file), load)
}

// LoadForBytes 从bytes中加载并解码到指定目录
func LoadForBytes(bytes []byte, path ...string) {
	load := ""
	if path != nil && len(path) >= 1 && path[0] != "" {
		load = path[0]
	}
	decode(DecodeBytes(bytes), load)
}

// DecodePath 从指定目录加载并解码
func DecodePath(path string) TFData {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln(err)
	}
	return DecodeBytes(data)
}

// DecodeBytes 对bytes进行解码
func DecodeBytes(bytes []byte) TFData {
	var tf TFData
	if err := json.Unmarshal(bytes, &tf); err != nil {
		log.Panicln(err)
	}
	return tf
}

// decode 对TF解码
func decode(tf TFData, path string) {
	sep := string(os.PathSeparator)
	log.Printf("[Load]name : %s,isDir %v", path+tf.Name, tf.IsDir)
	if tf.IsDir {
		// 如果是文件夹
		pathDir := path + sep + tf.Name
		// 不存在就创建
		if _, err := os.Stat(pathDir); err != nil {
			// 不存在就创建
			if os.IsNotExist(err) {
				err = os.Mkdir(pathDir, os.ModePerm)
				if err != nil {
					log.Panicln(err)
				}
			} else {
				log.Panicln(err)
			}
		}
		// 递归
		for _, t := range tf.DirVal {
			decode(t, pathDir)
		}
	} else {
		var name string
		// 创建文件
		if path != "" {
			name = path + sep + tf.Name
		} else {
			name = tf.Name
		}
		_, err := os.Create(name)
		if err != nil {
			log.Panicln(err)
		}
		err = ioutil.WriteFile(name, tf.Val, os.ModePerm)
		if err != nil {
			log.Panicln(err)
		}
	}

}
