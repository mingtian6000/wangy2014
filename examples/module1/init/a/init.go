package a

import (
	_ "example/init/b"
	"fmt"
)

func init() {
	fmt.Println("init from a")
}
