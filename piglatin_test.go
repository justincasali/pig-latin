package main

import (
	"fmt"
	"testing"

	"github.com/justincasali/piglatin/pig"
)

func TestFunc(t *testing.T) {
	text := "The quick brown fox jumps over the lazy dog"
	expected := "ethay uickqay ownbray oxfay umpsjay overway ethay azylay ogday"
	pigLatin := pig.Latin(text)
	if pigLatin != expected {
		t.Fail()
	}
}

func ExampleFunc() {
	pigLatin := pig.Latin("Hello")
	fmt.Println(pigLatin)
	// Output: ellohay
}
