package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
	//i "github.com/skilstak/go/input"
)

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

//func name() string {

//}
