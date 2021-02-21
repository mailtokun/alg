// 逆序对
// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
package main

import (
	"fmt"
)

func main() {

	ages := []int{7, 5, 6, 4, 2}

	fmt.Println("逆序对:", reversePairs(ages))
}

// 依次遍历的笨办法
func reversePairs(arr []int) int {
	var count int
	for i, value := range arr {
		if i+1 > len(arr) {
			fmt.Println("结束")
			break
		}
		subArr := arr[i+1:]
		for _, subValue := range subArr {
			if subValue < value {
				count++
			}
		}
	}
	return count
}
