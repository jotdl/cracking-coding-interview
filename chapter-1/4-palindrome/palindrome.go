package palindrome

import (
	"strings"
)

func IsPermutationOfPalindrome(src string) bool {
	length := 0
	runes := make(map[rune]int)

	for _, r := range strings.ToLower(src) {
		if r == ' ' {
			continue
		}

		length++
		runes[r]++
	}

	consumedMiddle := length%2 == 0
	for _, c := range runes {
		if c%2 == 0 {
			continue
		}

		if consumedMiddle {
			return false
		}
		consumedMiddle = true
	}

	return true
}
