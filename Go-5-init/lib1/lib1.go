package lib1 //一个包对应一个文件夹（除了main包）

import "fmt"

//lib1中的API定义
/*
	go的函数的定义
	函数名首字母大写为对外可访问 相当于 - public
	首字母小写为对外不可访问 只有当前包内可调用 - private
*/
func Lib1TestAPI() {
	fmt.Println("lib1's func")
}


func init() {
	fmt.Println("lib1's initFunc is running...")
}
