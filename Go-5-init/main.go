package main

import (
	"GoStudy/Go-5-init/lib1" //需要写完整gopath的相对路径 不然会找不到包 之后会用mod解决 注意还要关闭go-module
	"GoStudy/Go-5-init/lib2" //导入的包必须需要使用 不然会报错
)

func main() {
	lib1.Lib1TestAPI()
	lib2.Lib2TestAPI()
}