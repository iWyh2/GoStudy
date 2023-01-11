package main

import "fmt"

/*
	Go的函数 参数类型放在变量名之后
	返回值放在 ）与 { 之间
*/

//单个返回值
func func1(a string, b int) int {
	fmt.Println("a = ", a, "b = ", b)
	c := 20
	return c
}

/*
	多返回值 - 用括号包裹起来
	且还可以给返回值取名 未取名为匿名返回

	取名的多返回值可以看成是函数的形参，在定义时就有了初始化值，比如int为0等
	（Go的变量在定义时都有初始化值，防止野指针的出现）
*/
func func2(a int, b bool) (string, string) {
	fmt.Println("a = ", a, "b = ", b)
	return "iWyh2","20"
}

func func3(a int, b bool) (name string, age string) {
	fmt.Println("a = ", a, "b = ", b)

	//给有名称的返回值赋值
	name = "iWyh2"
	age = "20"

	//直接返回就无须再手动返回值，由带名称的返回值直接返回
	return
}

//多返回值类型相同时，类型可以放一起写
func func4(a int, b bool) (name , age string) {
	fmt.Println("a = ", a, "b = ", b)

	//给有名称的返回值赋值
	name = "iWyh2"
	age = "20"

	//直接返回就无须再手动返回值，由带名称的返回值直接返回
	return
}

func main() {
	a := func1("wyh", 20)
	fmt.Println("main's a = ", a)

	name, age := func2(20,true)
	fmt.Println("name = ", name, "age = ", age)

	name, age = func3(20,true)
	fmt.Println("name = ", name, "age = ", age)

	name, age = func4(20,true)
	fmt.Println("name = ", name, "age = ", age)
}