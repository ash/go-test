package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
