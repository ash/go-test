package main

import "fmt"

func main() {
	n := 10
	f := 1

	for i := 1; i <= n; i++ {
		f *= i
	}

	fmt.Printf("%d! = %d\n", n, f)
}
