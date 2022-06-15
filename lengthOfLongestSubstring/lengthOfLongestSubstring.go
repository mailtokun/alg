package lengthOfLongestSubString

import (
	"math"
)

func lengthOfLongestSubstring(s string) int {
	fast, low := 0, 0
	var maxSub float64 = 0
	checker := make(map[string]int)
	for i := 0; i < len(s); i++ {
		current := s[fast : fast+1]
		fast++
		checker[current]++
		for checker[current] > 1 {
			out := s[low : low+1]
			low++
			checker[out]--

		}
		maxSub = math.Max(float64(fast-low), float64(maxSub))
	}
	return int(maxSub)
}
