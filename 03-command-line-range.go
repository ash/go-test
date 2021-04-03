package main

import (
	"fmt"
	"os"
)

func main() {
	for n, value := range os.Args[1:] {
		fmt.Println(n, ":\t", value)
	}
}
