package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 1000

	nums := make(map[int]bool)

	for x := 1; x <= n; x++ {
		s := x * x
		for y := 1; y <= n; y++ {
			s += y * y
			for z := 1; z <= n; z++ {
				if s + z * z == 3 * x * y * z {
					nums[x] = true
					nums[y] = true
					nums[z] = true
				}
			}
		}
	}

	res := make([]int, 0)
	for a, _ := range(nums) {
		res = append(res, a)
	}

	sort.Ints(res)
	fmt.Println(res)
}
