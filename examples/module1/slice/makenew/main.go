package main

import (
	"fmt"
)

func main() {

	//slice 的类型仅由它所包含的元素的类型决定（与元素个数无关）。
	//要创建一个长度不为 0 的空 slice，需要使用内建函数 make。
	mySlice1 := new([]int)
	mySlice2 := make([]int, 0)
	mySlice3 := make([]int, 10)
	mySlice4 := make([]int, 10, 20)

	//slice 和数组是不同的类型，但它们通过 fmt.Println 打印的输出结果是类似的。
	fmt.Printf("mySlice1 %+v\n", mySlice1)
	fmt.Printf("mySlice2 %+v\n", mySlice2)
	fmt.Printf("mySlice3 %+v\n", mySlice3)
	fmt.Printf("mySlice4 %+v\n", mySlice4)
}
