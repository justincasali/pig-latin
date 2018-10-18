// Package pig provides functions for converting text to pig latin
package pig

import "strings"

const vowels = "aeiou"
const vSuffix = "way"
const cSuffix = "ay"

func convert(input string) string {
	v := strings.IndexAny(input, vowels)

	if v == 0 { // starts with vowel
		return input + vSuffix
	}

	if v > 0 { // starts with consontant
		return input[v:] + input[:v] + cSuffix
	}

	return input
}

// Latin converts text to pig latin
func Latin(input string) string {
	// convert string to lower case and split up words
	lower := strings.ToLower(input)
	words := strings.Fields(lower)

	// loop over words
	for i, w := range words {
		words[i] = convert(w)
	}

	// return joined words
	return strings.Join(words, " ")
}
