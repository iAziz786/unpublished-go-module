package main

import (
	"fmt"

	ipchecker "github.com/iAziz786/wipmodule"
)

func main() {
	fmt.Println("valid IP = ", ipchecker.IsValidIP("127.0.0.1"))
	fmt.Println("valid IP = ", ipchecker.IsValidIP("invalid-ip"))
}
