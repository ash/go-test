package main

import "fmt"

func main() {
	s := "Привет"

	r := []rune(s)
	fmt.Println(r)
	fmt.Println(len(r))

	fmt.Println(string(r[2:4]))
}