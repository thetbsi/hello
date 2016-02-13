package hello

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"string"
	"time"
	//"os"
	//i "github.com/skilstak/go/input"
)

func PrintForever(message string) {
	for {
		fmt.Println(message)
	}
}

func PrintRegular(message string) {
	fmt.Print(message)
}

func PrintMulti(message string) {
	for {
		fmt.Println(c.Clear + c.Multi(message) + " " + c.X)
		time.sleep(500 * time.Millisecond)
	}
}

func PrintColor(message string) {
	for {
		fmt.Println(c.Rc() + message + " " + c.X)
	}
}

//func name() string {

//}
