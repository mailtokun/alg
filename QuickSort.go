package main

import "fmt"

func main() {

	ages :=[] int {5,6,3,4,1,2,7,0 }
	qc(ages,0, len(ages)-1)
	//exchange(ages,1,2)
	fmt.Println(ages)
}

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
	toRight := true
	for i <= j  {
		fmt.Println(arr, "i=",i,"j=",j, "benchmarkIndex=", benchmarkIndex,"benchmarkValue=", benchmarkValue)
		if toRight {
			if benchmarkValue > arr[j] {
				exchange(arr, j, benchmarkIndex)
				benchmarkIndex = j
				i++
				toRight = false
			} else {
				j--
			}
		} else {
			if arr[i] > benchmarkValue {
				exchange(arr, i, benchmarkIndex)
				benchmarkIndex = i
				j--
				toRight = true
			}else {
				i++
			}
		}
	}
	// process left side
	fmt.Println("=== left side")
	if benchmarkIndex == left {
		qc(arr, left +1 , right)
	} else {
		qc(arr, left, benchmarkIndex)
	}
	// process right side
	fmt.Println("=== right side")
	start2 := benchmarkIndex+1
	end2 := right
	qc(arr,start2,end2)
}

func exchange(arr []int, x, y int){
	xValue := arr[x]
	arr[x] = arr[y]
	arr[y] = xValue
}