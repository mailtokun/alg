package lengthOfLongestSubString

import (
	"testing"
)

func Test_Unit(t *testing.T) {
	s := "abcabcbb"
	if lengthOfLongestSubstring(s) != 3 {
		t.Fatal()
	}
	s = "bbbbb"
	if lengthOfLongestSubstring(s) != 1 {
		t.Fatal()
	}
	s = "pwwkew"
	if lengthOfLongestSubstring(s) != 3 {
		t.Fatal()
	}
}
