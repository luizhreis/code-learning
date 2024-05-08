package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var start uint16 = 0
	var end uint16 = 0
	var max uint16 = 0
	chars := make(map[byte]uint16)

	for end < uint16(len(s)) {
		lastPosition, exists := chars[s[end]]

		if exists && start <= lastPosition {
			start = lastPosition + 1
		}

		chars[s[end]] = end
		end++

		if end-start > max {
			max = end - start
		}
	}

	return int(max)
}

func lengthOfLongestSubstring1(s string) int {
	chars := make(map[byte]bool)
	length := 0
	start := 0
	max := length
	for i := 0; i < len(s); i++ {
		char := s[i]
		_, ok := chars[char]
		if ok {
			clear(chars)
			length = 0
			i = start
			start++
		} else {
			chars[char] = true
			length++

		}
		if max < length {
			max = length
		}
	}
	return max
}

func main() {
	s := "abcabcbb"
	fmt.Println("Length of Longest Substring: ", lengthOfLongestSubstring(s))
	fmt.Println("Length of Longest Substring: ", lengthOfLongestSubstring1(s))
}
