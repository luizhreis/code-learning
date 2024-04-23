package main

import "testing"

type testCase struct {
	n    int
	want []int
}

func TestCountBits(t *testing.T) {
	testCases := []testCase{}

	testCases = append(testCases, testCase{0, []int{0}})
	testCases = append(testCases, testCase{5, []int{0, 1, 1, 2, 1, 2}})
	testCases = append(testCases, testCase{8, []int{0, 1, 1, 2, 1, 2, 2, 3, 1}})

	for _, testCase := range testCases {
		if got := countBits(testCase.n); !slicesEqual(got, testCase.want) {
			t.Errorf("got =  %v want %v", got, testCase.want)
		}
	}
}

func slicesEqual(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
