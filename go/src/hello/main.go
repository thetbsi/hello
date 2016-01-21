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
	name := "You"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	printMulti("Hello" + name)
}
