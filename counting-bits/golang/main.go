package main

import "fmt"

func countBits(n int) []int {
	ans := []int{}

	for i := 0; i <= n; i++ {
		ans = append(ans, countOnes(i))
	}

	return ans
}

func countOnes(n int) int {
	count := 0

	for n > 0 {
		n = n & (n - 1)
		count++
	}

	return count
}

func main() {
	fmt.Println(countBits(5))
}
