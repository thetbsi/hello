package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
)

func main() {
	fmt.Println(c.Clear + c.Multi("Hello World!"))
}
