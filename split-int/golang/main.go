package main

import "fmt"

func main() {
	n := 10932
	fmt.Printf("n = %d\nslc = %v", n, splitInt(n))
}

func splitInt(n int) []int {
	slc := []int{}

	for n > 0 {
		slc = append([]int{n % 10}, slc...)
		n = n / 10
	}

	return slc
}
