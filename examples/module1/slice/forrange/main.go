package main

import (
	"fmt"
)

func main() {
	mySlice := []int{10, 20, 30, 40, 50}
	//这里的数组也类似与map的结构
	for _, value := range mySlice {
		value *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)
	for index, _ := range mySlice {
		mySlice[index] *= 2
	}
	fmt.Printf("mySlice %+v\n", mySlice)
}
