package main

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/justincasali/piglatin/pig"
)

//go:embed data/test_input.txt
var testInput string

//go:embed data/test_output.txt
var testOutput string

func TestFunc(t *testing.T) {

	pigLatin := pig.Latin(testInput)
	if pigLatin != testOutput {
		t.Fail()
	}

}

func ExampleFunc() {

	pigLatin := pig.Latin("Hello")
	fmt.Println(pigLatin)
	// Output: ellohay

}
