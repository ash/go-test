package main

import "fmt"

func main() {
	x := new(int)
	*x = 42
	fmt.Println(x)
	fmt.Println(*x)
}
