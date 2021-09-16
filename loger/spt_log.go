/*
--------------------------------------------------
@Create 2021-09-14 23:23
@Author lpc
@Program shear-plate-transmit
@Describe
--------------------------------------------------
@Version 1.0 2021-09-14
@Memo create this file
*/

package loger

import (
	"log"
	"os"
)

// Logger 结构体
type Logger struct {
}

// Log 日志实例
var Log *Logger
var logger *log.Logger

func init() {
	Log = &Logger{}
	logger = log.New(os.Stdout, "[SPT]", log.LstdFlags|log.Lshortfile)
}

// Info 普通日志
func (l *Logger) Info(info ...interface{}) {
	logger.Println(info)
}

// Debug 调试日志
func (l *Logger) Debug(info ...interface{}) {
	logger.Println(info)
}

// Log 输出日志
func (l *Logger) Log(format string, v ...interface{}) {
	logger.Printf(format+"\n", v)
}
