package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	i "github.com/skilstak/go/input"
)
func getName() (name string) {
	name = "Matty"
	return
}

func printForever(message string) {
	for {
	fmt.Println(c.Clear + c.Multi("Hello World!"))
	}
}

func printRegular(message string) {
	fmt.Print("Hello World!")
}

func printMulti(message string) {
	fmt.Println(c.Clear + c.Multi("Hello World!"))
}

func printRandom(message string) {
	fmt.Println(c.Clear + c.Rc() + "Hello World!")
}

func name() string {
	
}
func main() {
	printMulti("Hello"  + getName())
	//a, hello := i.Prompt("--> ")
	//if hello != nil {
	//	panic(hello)
	}
}
