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
	a, thingsandstuff := i.Prompt("--> ")
	if thingsandstuff != nil {
		panic(thingsandstuff)

	}
	fmt.Println(a)
}
