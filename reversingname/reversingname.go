package main

import "fmt"

func main() {
    input := "reshmask"
    reversed := reverseString(input)
    fmt.Println("Original:", input)
    fmt.Println("Reversed:", reversed)
}

func reverseString(input string) string {
    runes := []rune(input)
    length := len(runes)

    for i := 0; i < length/2; i++ {
        runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
    }

    return string(runes)
}
