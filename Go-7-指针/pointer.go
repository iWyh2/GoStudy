package main

import (
	"fmt"
)

/*
	快速掌握指针的小例子 - 交换两个数据的值
*/

//按值传递
// func exchange(a int, b int) {
// 	temp := a
// 	a = b
// 	b = temp
// }

//按址传递
func swap(pa *int, pb *int) {
	temp := *pa
	*pa = *pb
	*pb = temp
}

func main() {
	value1 := 10
	value2 := 100

	//exchange(value1,value2)
	//fmt.Println("After exchange: value1 = ",value1,"value2 = ",value2)

	swap(&value1, &value2)
	fmt.Println("After swap: value1 = ",value1,"value2 = ",value2)

	var p *int = &value1
	fmt.Println("p = ",p,"*p = ",*p)
	pp := &p//二级指针
	fmt.Println("pp = ",pp,"*pp = ",*pp,"**pp = ",*(*pp))
}