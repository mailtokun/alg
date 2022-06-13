package twosum

import (
	"fmt"
	"testing"
)

func Test_Unit(t *testing.T) {
	var nums []int = []int{1, 2, 3, 4, 5, 6}
	var target int = 11
	indexs := twoSum(nums, target)
	if nums[indexs[0]]+nums[indexs[1]] != target {
		t.Fatal()
	}
	fmt.Print(indexs)
}
