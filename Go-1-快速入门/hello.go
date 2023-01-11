package main //表示该文件所在的包为main

// import "fmt" //引入了一个包为fmt，引入了之后才可以使用这个包里面的函数：Println
// import "time"

import ( //导入多个包
	"fmt" //包含一些基本的输入输出
	"time"
)

//func表示是一个函数；main是主函数，程序的入口
func main () {//函数的‘{’必须和函数名在同一行
	fmt.Println("Hello Go, Change The World!")//不加分号

	time.Sleep(1 * time.Second)
}