package reshmask

import (
	"strings"
)

func IsPalindrome(s string) bool {
	// Convert the input string to lowercase and remove spaces.
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	// Compare the string with its reverse.
	return s == reverseString(s)
}

func reverseString(s string) string {
	runes := []rune(s)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	return string(runes)
}
