package main

import "fmt"

func main() {
	var arr = [][]int{{0, 0, 0}, {1, 0, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(arr))
}
func uniquePathsWithObstacles(data [][]int) int {
	if len(data) == 0 {
		return 0
	}
	if data[0][0] == 1 {
		return 0
	}
	m, n := len(data), len(data[0])
	dp := data
	// 第一列
	for i := 1; i < m; i++ {
		if data[i][0] == 0 && data[i-1][0] == 1 {
			dp[i][0] = 1
		}
	}
	fmt.Println(dp)
	// 第一行
	for i := 1; i < n; i++ {
		if data[0][i] == 0 && data[0][i-1] == 1 {
			dp[0][i] = 1
		}
	}
	fmt.Println(dp)
	// 后续列
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if data[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}
