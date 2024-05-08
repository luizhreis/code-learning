package main

import "testing"

type testCase struct {
	s    string
	want int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := buildTestCases()

	for _, testCase := range testCases {
		got := lengthOfLongestSubstring(testCase.s)
		if got != testCase.want {
			t.Errorf("lengthOfLongestSubstring(%q): want %d, got %d", testCase.s, testCase.want, got)
		}
	}
}

func buildTestCases() []testCase {
	testCases := []testCase{}

	testCases = append(testCases, testCase{
		s:    "abcabcbb",
		want: 3,
	})
	testCases = append(testCases, testCase{
		s:    "bbbbb",
		want: 1,
	})
	testCases = append(testCases, testCase{
		s:    "pwwkew",
		want: 3,
	})
	testCases = append(testCases, testCase{
		s:    "dvdf",
		want: 3,
	})
	testCases = append(testCases, testCase{
		s:    "abcabdefg",
		want: 7,
	})
	testCases = append(testCases, testCase{
		s:    "abcadefgba",
		want: 7,
	})

	return testCases
}
