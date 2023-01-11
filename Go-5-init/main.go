package main

import (
	"GoStudy/Go-5-init/lib1" //需要写完整gopath的相对路径 不然会找不到包 之后会用mod解决 注意还要关闭go-module
	"GoStudy/Go-5-init/lib2" //导入的包必须需要使用 不然会报错
	"GoStudy/Go-5-init/swap"
	"fmt"
)

func main() {
	lib1.Lib1TestAPI()
	lib2.Lib2TestAPI()
	fmt.Println("-----------")
	fmt.Println("|swap func|")
	fmt.Println("-----------")
	value1 := 100
	value2 := 50
	fmt.Println("before swap: value1 = ",value1,"value2 = ",value2)
	fmt.Println(swap.Swap(&value1, &value2))
	fmt.Println("before swap: value1 = ",value1,"value2 = ",value2)
}