// 快速排序
//
//
package main

import "fmt"

func main() {

	ages := []int{6, 1, 2, 7, 9, 3}
	qc(ages, 0, len(ages)-1)
	//exchange(ages,1,2)
	fmt.Println(ages)
}

// 对一个数组或者子数组进行排序
// left:为开始坐标
// right: 为结束坐标
func qc(arr []int, left, right int) {
	if left >= right {
		return
	}
	fmt.Println("Quick Sort", left, right)
	fmt.Println(arr)
	i := left
	j := right
	// 基准值
	benchmarkValue := arr[left]
	// 基准下标
	benchmarkIndex := left
	toLeft := true
	for i <= j {
		var currentWay string = "向右"
		if toLeft {
			currentWay = "向左"
		}
		fmt.Println(arr, currentWay, "i=", i, "j=", j, "benchmarkIndex=", benchmarkIndex, "benchmarkValue=", benchmarkValue)
		if toLeft {
			if benchmarkValue > arr[j] {
				exchange(arr, j, benchmarkIndex)
				benchmarkIndex = j
				i++
				toLeft = false
			} else {
				j--
			}
		} else {
			if benchmarkValue < arr[i] {
				exchange(arr, i, benchmarkIndex)
				benchmarkIndex = i
				j--
				toLeft = true
			} else {
				i++
			}
		}
	}
	// process left side
	fmt.Println("=== left side")
	if benchmarkIndex == left {
		qc(arr, left+1, right)
	} else {
		qc(arr, left, benchmarkIndex)
	}
	// process right side
	fmt.Println("=== right side")
	start2 := benchmarkIndex + 1
	end2 := right
	qc(arr, start2, end2)
}

func exchange(arr []int, x, y int) {
	xValue := arr[x]
	arr[x] = arr[y]
	arr[y] = xValue
}
