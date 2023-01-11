package main

import (
	_ "GoStudy/Go-5-init/lib1"   //导入包并起别名 此处为匿名 - 也就是导入包使用包的init方法 但不使用当前包内的方法函数
	wyh "GoStudy/Go-5-init/lib2" //导入包并起别名 此处为起别名 - 也就是正常导入并使用 只是用别名来调用了
	//. "GoStudy/Go-5-init/lib2" //使导入的包可以直接调用API 而不是使用包名调用（但一般情况下不建议这样使用，防止两个包有同名函数）
)



func main() {
	//lib1.Lib1TestAPI() 被匿名导入 无法调用包内的函数
	wyh.Lib2TestAPI()
}