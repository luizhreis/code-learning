package main

import (
	"testing"
)

type testCase struct {
	num  string
	k    int
	want string
}

func TestRemoveKDigits(t *testing.T) {
	testCases := []testCase{}

	testCases = append(testCases, testCase{"1432219", 3, "1219"})
	testCases = append(testCases, testCase{"10200", 1, "200"})
	testCases = append(testCases, testCase{"10", 2, "0"})
	testCases = append(testCases, testCase{"10", 1, "0"})
	testCases = append(testCases, testCase{"26370", 3, "20"})
	testCases = append(testCases, testCase{"9", 1, "0"})
	testCases = append(testCases, testCase{"112", 1, "11"})
	testCases = append(testCases, testCase{"333", 1, "33"})
	testCases = append(testCases, testCase{"90001801", 2, "101"})

	for _, testCase := range testCases {
		got := removeKdigits(testCase.num, testCase.k)
		if got != testCase.want {
			t.Errorf("removeKdigits(%v,%v):\ngot %v,\nwant %v", testCase.num, testCase.k, got, testCase.want)
		}
	}
}
