package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
)

func printMulti() {
	fmt.Println(c.Clear + c.Multi("Hello World!"))
}

func printRandom() {
	fmt.Println(c.Clear + c.Rc() + "Hello World!")
}

func main() {
	//printMulti()
	//printRandom()
	a, hello := i.Prompt("--> ")
	if hello != nil {
		panic(hello)
	}
	fmt.Println(input)
	if strings.Contains(input,"hello")
}
