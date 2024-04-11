package main

import (
	"fmt"
	"strings"
)

func main() {
	num := "1432219"
	k := 3
	fmt.Println(removeKdigits(num, k))
}

func removeKdigits(num string, k int) string {
	stack := []string{}

	digits := strings.Split(num, "")

	for _, digit := range digits {
		for len(stack) > 0 && k > 0 && digit < stack[len(stack)-1] {
			k--
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, digit)

	}

	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	if len(stack) == 0 {
		return "0"
	}

	stack = removeLeadingZeroes(stack)

	return strings.Join(stack, "")
}

func removeLeadingZeroes(digits []string) []string {
	for len(digits) > 1 && digits[0] == "0" {
		digits = digits[1:]
	}

	return digits
}
