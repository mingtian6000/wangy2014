package main

import (
	"fmt"
)

func main() {
	name := "testing"
	var i int
	i, _ = fmt.Printf("%d\n", name)
	fmt.Sprintf("%d\n", i)
	//fmt.Printf("%d\n", name)
	fmt.Printf("%s\n", name, name)
}
