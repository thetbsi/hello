package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
	//i "github.com/skilstak/go/input"
)

func getName(title string) string {
	var name string = "Matty"
	return title + " " + name
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

//func name() string {

//}
func main() {
	//message := "Hello World"
	sentence := getName("Hello")
	printMulti(sentence)
}
