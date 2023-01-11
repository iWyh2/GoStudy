package lib2

import "fmt"

//lib2中的API定义
func Lib2TestAPI() {
	fmt.Println("lib2's func")
}


func init() {
	fmt.Println("lib2's initFunc is running...")
}