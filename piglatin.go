// Justin Casali
package main

import (
	"fmt"
	"os"

	"github.com/justincasali/piglatin/pig"
)

func main() {

	var text string

	switch len(os.Args) {
	case 1:
		text = "The quick brown fox jumps over the lazy dog"
	case 2:
		text = os.Args[1]
	default:
		fmt.Println("too many arguments")
		return
	}

	pigLatin := pig.Latin(text)
	fmt.Println(pigLatin)
}
