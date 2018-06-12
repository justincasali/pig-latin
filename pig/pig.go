// Justin Casali
package pig

import s "strings"

const vowels = "aeiou"
const vSuffix = "ay"
const cSuffix = "way"

func Latin(input string) string {
    // convert string to lower case and split up words
    lower := s.ToLower(input)
    words := s.Fields(lower)

    // loop over words
    for i, w := range words {
        v := s.IndexAny(w, vowels)
        
        switch {
        case v == 0:    // starts with vowel
            words[i] = w + cSuffix

        case v > 0:     // starts with consontant
            words[i] = w[v:] + w[:v] + vSuffix

        default:        // word has no vowels
            words[i] = w
        }
        
    }

    // return joined words
    return s.Join(words, " ")
}
