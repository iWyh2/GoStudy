/*
四种变量的声明方式
*/
package main

import "fmt"

/*
	声明全局变量时，方法一、二、三没有区别 都是可以的
	但是方法四的(:=)不支持声明全局变量
*/
var glob_a int = 50
var glob_b int
var glob_c string = "iWyh2"


func main() {
	//方法一: 声明一个变量 默认值为0
	var a int
	fmt.Println("a = ", a)

	//默认值为空
	var as string
	fmt.Println("as = ", as)

	//方法二: 声明一个变量 声明的同时进行初始化
	var b int = 50
	fmt.Println("b = ", b)

	//方法三: 声明一个变量并对其进行初始化时 省去变量的数据类型 通过初始化值自动匹配当前变量的数据类型
	var c = 100
	fmt.Printf("type of c is %T\n", c)//格式化输出 输出当前变量的数据类型用%T
	fmt.Println("c = ", c)

	//方法四: 省去var关键字 直接自动匹配进行声明并初始化赋值 - 常用这种方法 (但不支持声明全局变量 只适合在方法体内进行使用)
	d := 100
	fmt.Printf("type of d is %T\n", d)//格式化输出 输出当前变量的数据类型用%T
	fmt.Println("d = ", d)

	//输出全局变量
	fmt.Printf("全局变量a, b, c为: %d %d %s\n", glob_a, glob_b, glob_c)

	//多变量声明 - 方式一: 一行书写
	var x, y int = 10, 20//相同类型
	var z, w = 0, "wang"//不同类型 自动匹配
	fmt.Println("x = ", x, "y = ", y)//可以多字符串拼接式输出
	fmt.Println("z = ", z, "w = ", w)

	//多变量声明 - 方式二: 多行书写 (一般适用于多且 不同变量类型 时)
	var (
		xx int
		yy bool = false
		zz string = ""
	)
	fmt.Println("xx = ", xx,"yy = ", yy,"zz = ", zz)
}