package hello

import (
	"fmt"
	c "github.com/skilstak/go/colors"
)

func PrintMulti(message string) {
	fmt.Println(c.Clear + c.Multi(message) + c.X)
}
func PrintForever(message string) {
	
