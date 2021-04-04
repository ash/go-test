package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Привет"
	fmt.Println(len(s))
	
	r := []rune(s)
	fmt.Println(len(r))

	n := utf8.RuneCountInString(s)
	fmt.Println(n)
}
