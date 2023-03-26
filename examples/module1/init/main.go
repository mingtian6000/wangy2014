package main

import (
	_ "example/init/a"
	_ "example/init/b"
	"fmt"
)

func init() {
	fmt.Println("main init")
}
func main() {

}
