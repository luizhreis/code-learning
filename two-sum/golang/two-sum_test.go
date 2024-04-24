package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums   []int
	target int
	want   []int
}

func TestTwoSum(t *testing.T) {
	testCases := buildTestCases()

	for _, testCase := range testCases {
		got := twoSum(testCase.nums, testCase.target)
		assert.Equal(t, testCase.want, got, "Failed for nums=%v and target=%d", testCase.nums, testCase.target)
	}
}

func buildTestCases() []testCase {
	testCases := []testCase{}

	testCases = append(
		testCases,
		testCase{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
	)
	testCases = append(
		testCases,
		testCase{
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
	)
	testCases = append(
		testCases,
		testCase{
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	)

	return testCases
}
