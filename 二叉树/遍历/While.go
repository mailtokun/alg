//        1
//       / \
//      2  3
//     / \  \
//    4  5   6
//      / \
//     7   8
//前序遍历 深度优先搜索算法（Depth First Search）：1  2  4  5  7  8  3  6
//中序遍历：4  2  7  5  8  1  3  6
//后序遍历：4  7  8  5  2  6  3  1
//层次遍历 广度优先搜索算法（Breadth First Search）：1  2  3  4  5  6  7  8
package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var res = make([]int, 0)
	node7 := TreeNode{Val: 7}
	node8 := TreeNode{Val: 8}
	node4 := TreeNode{Val: 4}
	node5 := TreeNode{Val: 5, Left: &node7, Right: &node8}
	node6 := TreeNode{Val: 6}
	node2 := TreeNode{Val: 2, Left: &node4, Right: &node5}
	node3 := TreeNode{Val: 3, Right: &node6}
	node1 := TreeNode{Val: 1, Left: &node2, Right: &node3}
	preorderTraversal(&node1, &res)
	fmt.Println(res)

}

// 前序遍历
// 输入：root = [1,null,2,3]
// 输出：[1,2,3]
func preorderTraversal(p *TreeNode, res *[]int) {
	fmt.Println("前序遍历")
	if p != nil {
		*res = append(*res, p.Val)
		preorderTraversal(p.Left, res)
		preorderTraversal(p.Right, res)
	}
}

// 后序遍历
func postorderTraversal(p *TreeNode, res *[]int) {
	fmt.Println("后序遍历")
	if p != nil {
		postorderTraversal(p.Left, res)
		postorderTraversal(p.Right, res)
		*res = append(*res, p.Val)
	}
}

// 层序遍历
func bfs(p *TreeNode) []int {
	fmt.Println("层序遍历")
	res := make([]int, 0)
	if p == nil {
		return res
	}
	queue := []*TreeNode{p}
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			length--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res = append(res, queue[0].Val)
			queue = queue[1:]
		}
	}
	return res
}
