[![GoDoc](https://godoc.org/github.com/lengpucheng/shear-plate-transmit?status.svg)](https://pkg.go.dev/github.com/lengpucheng/shear-plate-transmit)
[![Go Report Card](https://goreportcard.com/badge/github.com/lengpucheng/shear-plate-transmit)](https://goreportcard.com/report/github.com/lengpucheng/shear-plate-transmit)
[![License](https://img.shields.io/github/license/lengpucheng/shear-plate-transmit)](https://github.com/lengpucheng/shear-plate-transmit/blob/master/LICENSE)
[![Goproxy.cn](https://goproxy.cn/stats/github.com/lengpucheng/shear-plate-transmit/badges/download-count.svg)](https://goproxy.cn)
<center><h1>SHEAR-PLATE-TRANSMIT（SPT）</h1></center>

**SPT可以使用剪切板作为中介来传输数据，可以应用在某些跨平台设备只能传输文本的场景下实现任意数据的传输**

+ 😘仅需要使用接口下的send()和receive()发放即可完成数据承载收发
+ 😃提供串行和并行两种方式传输的选择
+ 😊自带文件读取-编码-传输-接收-解码-还原的实现应用
+ 😜既可以作为数据承载二次开发也可以作为文件传输根据编译后直接使用

[TOC]

# 一、安装

使用go get 拉取并编译安装，编译后的可执行文件在`GOPATH/bin`目录下

```shell
go get -u github.com/lengpucheng/shear-plate-transmit
```

若要直接使用，需要将`GOPATH/bin`设置到环境变量或者将编译后的文件移动到系统目录

+ Linux

```shell
cp ./shear-plate-transmit /usr/bin
```

+ Windows

```shell
PATH=%PATH%;%GOPATH%/bin
```

通常可以将编译后的文件重命名为`SPT`，会更加方便

# 二、使用

SPT已经完成了剪切板数据传输的基础功能，可以使用其作为数据承载进行开发，同时添加了文件传输的实现方法，具体使用如下

## 1. 传输文件

编译后为可执行文件，开启两个软件作分别作为传输和接收端，指定需要传输的文件（或目录）以及接收的文件目录，并设置参数即可

+ **使用**

```shell
shear-plate-transmit <-t> -p <path> -time <time> -max <max>
```  

+ 参数说明：
    + `-p`:路径，传输时为需传输的文件（或目录），接收时为文件保存目录
    + `-t`:**可选**，是否是传输操作，若没有使用则为接收
    + `-time`:**可选**，指定分片等待时间（毫秒），设置较大时可以避免剪切板死锁或由于网络问题导致的传输不稳定，一般不小于300
    + `-max`:**可选**，指定分片的最大大小（kb），当超过这一大小将对数据进行分片传输，一般不大于1024

使用时同时开启两个应用，一个用于发送，一个用于接收，传输途中会在控制台输出日志和进度信息，中途请勿操作剪切板

**若传输中日志提示错误，请等待，将自动重新建立链接并传输丢失的数据包分片**

## 2. 作为数据承载

目前该项目实现了通过剪切板的数据传输接口和底层，在项目中映入当前包可以使用其下接口`coreplate.PlateTransmit`完成通过剪切板的数据传输，方法如下：

+ `Send([]byte)`:使用剪切板发送数据，若数据过大会自动切片
+ `Receive()[]byte`:使用剪切板接收数据，若数据被切片会在所有分片接收完毕后，自动组装后返回

目前有两个实现类：

+ `PlateTransmitSingle`: **串行读写分离的传输**，将单协程传输或接收文件，并且一次只能读或者写
+ `PlateTransmitMulti`:**并行同时读写的传输**，将异步读写剪切板，可以同时传输和读写（未经过充分测试）

**若遇到传输错误或剪切板被外部修改时，会自动识别并重传发送异常的数据分片**，***以上请务必保证一次全局仅实例化一个对象，否则会导致剪切板读写异常***（由于剪切板是单例）

参数设置:

+ `coreplate.SetMaxsize(int64)`:可以设置全局数据单片的最大大小
+ `coreplate.SetDelayTime(int64)`:可以设置全局单片延时大小

# 三、传输原理

剪切板作为系统中的一个临时数据交换场景，主要用于存放临时数据，**使用虚拟机或并可在宿主机上跨虚拟机共享**，因此实现数据传输的关键就是实现剪切板读写

## 1. 数据包定义

本项目中使用[clipboard](github.com/atotto/clipboard)实现剪切板读写，通过定义数据结构`DataPack`完成对数据进行数据包封装

```go
type DataPack struct {
Id      int64  `json:"Id"`
Tid     int64  `json:"Tid"`
Size    int64  `json:"Size"`
Total   int64  `json:"Total"`
Data    []byte `json:"Data"`
IsEOF  bool   `json:"is_eof"`
ReBack bool   `json:"re_back"`
}
```

以上数据包详情如下：

+ `ID`---数据包唯一标识，由于剪切板的单例特性使用时间戳作为ID
+ `Size`---当前数据已经传输的大小（包括当前数据包）
+ `Total`---当前待传输数据的总大小
+ `Data`---实际传输的数据
+ `IsEOF`---是否传输完毕
+ `ReBack`---是否是响应回执

## 2. 数据切片封装

+ 数据切片 当调用`Send([]byte)`传入数据时，数据会和限制大小进行对比，若小于限制大小会直接传输，否则按照数据大小进行切片，并向数据包中写入总大小，然后依次发送，
  当最出现小于限制大小时，会将全部写入dada，并设置size=total以及isEof=true
+ 数据包序列化
    + 数据包在传输前会被序列化未`[]byte`并转换为`string`写入剪切板
    + 数据包在接收后会被反序列化为`datapack`并判断是否到末尾，否则会持续接收并拼接data

## 3. 传输和接收

+ 传输时每发送一个数据包，会等待接收端的回执相应，并判断回执ID是否和上一次相同，避免粘包的情况发生，为避免丢包只有获取到接收端回执才会进行下一次数据包传输，**若尝试一段时间`100x延时``
  依然为获取到回执，会尝试重新传输当前数据包**
+ 接收时每接收一个数据包，会判断是否和上一次接收的ID相同，如果相同会忽略，避免数据重复，只有获取到新ID的数据后，才会发送回执信息
+ 当读取剪切板发生错误或其他外部干扰数据时，会忽略错误并返回nil给底层的数据传输，当在延时等待后进行传输或接收的重试，用于传输容错
+ 并行传输时，接收的数据和发送的数据以及回执在被发送前或接收后会被写入对应的管道，调用`PlateTransmitMulti`中的send或receive方法会向管道中写入或读取数据
