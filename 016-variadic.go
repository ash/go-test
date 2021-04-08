package main

import "fmt"

func sum(vals ...int) int {
	r := 0
	for i, v := range vals {
		fmt.Println(i, v)
		r += v
	}

	return r
}

func main() {
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3))
}
