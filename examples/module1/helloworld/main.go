package main

import (
	"flag"
	"fmt"
)

func main() {
	// return is the stress of string var
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	//fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", name) //0xc000088280 注意go里面有地址的概念
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)
}
