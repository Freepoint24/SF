package main

import (
	"fmt"

	//	"fmt"
	"strings"
	"unicode"
)

func main() {
	value := "АаааИИИИ"
	word := IsIsogram(value)
	fmt.Println(value)
	fmt.Println(word)

}

// IsIsogram determines if a word or phrase is an isogram.

func IsIsogram(word string) bool {
	var bitset uint32 = 0
	for _, c := range strings.ToLower(word) {
		if unicode.IsLetter(c) {
			if index := c - 'a'; bitset&(1<<index) == 0 {
				bitset |= 1 << index
			} else {
				return false
			}
		}
	}
	return true
}
