package main

import "fmt"

type Value struct {
	X int
}

func (v Value) Inc() {
	v.X++
}

func (v *Value) Inc2() {
	v.X++
}

func main() {
	a := Value{42}

	a.Inc()
	fmt.Println(a)

	a.Inc2()
	fmt.Println(a)

	(&a).Inc2()
	fmt.Println(a)
}
