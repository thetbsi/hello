package hello

import (
	"fmt"
	c "github.com/skilstak/go/colors"
	"os"
	//i "github.com/skilstak/go/input"
)

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
