package main

import (
	l "hello"
	//"mt"
	//c "github.com/skilstak/go/colors"
	"os"
)

func main() {
	name := "You"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	l.PrintMulti("Hello" + name)
}
