package main

import (
	"fmt"
	"math"
)

// 在不同类型的项之间赋值时需要显式转换
func main() {
	var x, y int = 3, 4
	//var f float64 = math.Sqrt(x*x + y*y)
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)

	v := 123 // 修改这里！
	fmt.Printf("v is of type %T\n", v)
}
