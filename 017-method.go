package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) Display() {
	fmt.Printf("(%d, %d)\n", p.X, p.Y)
}

func main() {
	a := Point{3, 4}
	fmt.Println(a)
	a.Display()
}
