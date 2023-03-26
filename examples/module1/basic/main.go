package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	c, python, java bool
)

func swap(x, y string) (string, string) {
	return y, x
}

func add(x int, y int) int {
	return x + y
}

func sum(a []int) int {
	b := 0
	for i := 0; i < len(a); i++ {
		b += a[i]
	}
	return b
}
func main() {
	fmt.Println("Welcome to the Go World!")

	fmt.Println("The basic is", time.Now())
	for i := 0; i < 3; i++ {
		fmt.Println("My favorite number is", rand.Intn(10))
	}
	//截断？？
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(c, python, java)
	list1 := []int{1, 2, 3, 4, 5}

	fmt.Println(sum(list1))

	fmt.Println(swap("abcd", "xyz"))
}
