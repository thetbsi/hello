package hello

import (
	"fmt"
	c "github.com/skilstak/go/colors"
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
		fmt.Println(c.Clear + c.Multi(message))
		time.sleep(500 * time.Millisecond)
	}
}

func PrintColor(message string) {
	fmt.Println(c.Clear + c.Rc() + message)
}

//func name() string {

//}
