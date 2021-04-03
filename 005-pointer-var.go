package main

import "fmt"

func main() {
	var x int

	p := &x
	*p = 42

	fmt.Println(x)
}
