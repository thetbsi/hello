package hello

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
	//i "github.com/skilstak/go/input"
)

func printMulti(message string) {
	fmt.Println(c.Clear + c.Multi(message) + c.X)
}
