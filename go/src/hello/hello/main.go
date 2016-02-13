package main

import (
	h "hello"
	//"mt"
	//c "github.com/skilstak/go/colors"
	"os"
)

func main() {
	name := " You"
	if len(os.Args) > 1 {
		name = os.Args[1]

	switch option { 
	case "-m":
		h.PrintMulti(message)
	case "-c":
		h.PrintColor(message)
	case "-f":
		h.PrintForever(message)
	default:
		h.PrintPlain(message)
	}
}
