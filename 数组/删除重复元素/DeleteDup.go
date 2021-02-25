package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{6, 1, 2, 6, 3, 3}
	ages = removeDuplicateElement(ages)
	fmt.Println(ages)
}
func removeDuplicateElement(A []int) []int {
	if len(A) <= 1 {
		return A
	}
	sort.Ints(A)
	newArray := make([]int, 0)
	newArray = append(newArray, A[0])
	for i := 1; i < len(A); i++ {
		if A[i] != newArray[len(newArray)-1] {
			newArray = append(newArray, A[i])
		}
	}
	return newArray
}
