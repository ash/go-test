package main

import "fmt"

func main() {
	var a []int

	for i := 0; i < 20; i++ {
		fmt.Printf("len=%d\tcap=%d\n", len(a), cap(a))
		a = append(a, i)
	}
}
