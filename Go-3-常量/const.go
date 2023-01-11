package main

import "fmt"

/*
	const可以定义枚举类型 (枚举必须赋值) : 也就是多个常量定义
	在定义枚举时 可以使用关键字: iota
	第一行的iota值默认为0 iota在每一行都会累加1
	iota只能在const()中使用
*/
const(
	Spring = iota//此时 Spring = 0
	Summer//此时 Summer = 1
	Autumn//此时 Autumn = 2
	Winter//此时 Winter = 3
)

/*
	多个常量可以放在同一行
	枚举第一行的iota值默认为0
	iota还可以用在表达式中
*/
const(
	dog, cat = iota * 10, iota * 20//此时 dog = 0  cat = 0
	pig, kit//此时 pig = 10 kit = 20
	tgr, lon = iota * 10 + 5, iota * 20 + 10//此时 tgr = 25  lon = 50
)

func main() {
	//常量定义
	const length int = 10//只读
	//length = 20 is not allowed :-> 常量是不允许再被修改的

	//打印常量
	fmt.Println("Spring = ", Spring)
	fmt.Println("Summer = ", Summer)
	fmt.Println("Autumn = ", Autumn)
	fmt.Println("Winter = ", Winter)
	//打印常量
	fmt.Println("dog = ", dog, "cat = ", cat)
	fmt.Println("pig = ", pig, "kit = ", kit)
	fmt.Println("tgr = ", tgr, "lon = ", lon)
}