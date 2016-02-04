package hello

import (
	"fmt"
	c "github.com/skilstak/go/colors"
)

func PrintMulti(message string) {
	fmt.Println(c.Clear + c.Multi(message) + c.X)
}
func PrintForever(message string) {
	for { 
		fmt.Println(c.Clear + (message) + c.X)
}
func PrintPlain(message string) { 
		fmt.Println("Hello" + (message) + c.X)
}
func PrintRandom(message string) {
	fmt.Println(c.Clear + c.Rc() + (message) + c.X)
}
