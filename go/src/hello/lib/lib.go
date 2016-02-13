package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	//"os"
	//i "github.com/skilstak/go/input"
)

func PrintForever(message string) {
	for {
		fmt.Println(c.Clear + c.Multi(message))
	}
}

func PrintRegular(message string) {
	fmt.Print(message)
}

func PrintMulti(message string) {
	fmt.Println(c.Clear + c.Multi(message))
}

func PrintRandom(message string) {
	fmt.Println(c.Clear + c.Rc() + message)
}

//func name() string {

//}
