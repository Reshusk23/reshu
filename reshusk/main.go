package main

import (
	"fmt"

	"github.com/Reshusk23/reshu/reshmask"
)

func main() {
	input := "mam"

	if reshmask.IsPalindrome(input) {
		fmt.Println(input, "is a palindrome.")
	} else {
		fmt.Println(input, "is not a palindrome.")
	}
}
