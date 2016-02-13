package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
	//i "github.com/skilstak/go/input"
)

func printForever(message string) {
	for {
		fmt.Println(c.Clear + c.Multi(message))
	}
}

func printRegular(message string) {
	fmt.Print(message)
}

func printMulti(message string) {
	fmt.Println(c.Clear + c.Multi(message))
}

func printRandom(message string) {
	fmt.Println(c.Clear + c.Rc() + message)
}

//func name() string {

//}
