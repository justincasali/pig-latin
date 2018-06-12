// Justin Casali
package main

import "pig-latin/pig"
import "fmt"
import "os"

func main() {

    var text string
    
    switch len(os.Args) {
    case 1:
        text = "The quick brown fox jumps over the lazy dog"
    case 2:
        text = os.Args[1]
    default:
        panic("Too many arguments")
    }

    fmt.Println(text)
    pigLatin := pig.Latin(text)
    fmt.Println(pigLatin)
}
