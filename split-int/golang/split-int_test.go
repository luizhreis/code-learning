package main

import "testing"

type testCase struct {
	n    int
	want []int
}

func TestSplitInt(t *testing.T) {
	testCases := []testCase{}

	testCases = append(testCases, testCase{10932, []int{1, 0, 9, 3, 2}})
	testCases = append(testCases, testCase{567, []int{5, 6, 7}})
	testCases = append(testCases, testCase{0, []int{}})

	for _, testCase := range testCases {
		if got := splitInt(testCase.n); !slicesEqual(got, testCase.want) {
			t.Errorf("got = %v\nwant %v", got, testCase.want)
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
