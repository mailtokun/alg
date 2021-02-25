package main

import "fmt"

func main() {
	t := generateTrees(1, 3)
	for _, item := range t {
		pintT(item, "")
		fmt.Println("===============")
	}
}
func pintT(t *Tree, way string) {
	if t == nil {
		return
	}
	fmt.Println(t.node, way)
	if t.left != nil {
		pintT(t.left, "l")
	}
	if t.right != nil {
		pintT(t.right, "r")
	}
}

type Tree struct {
	node  int
	left  *Tree
	right *Tree
}

func generateTrees(start, end int) []*Tree {
	tree := []*Tree{}
	if start > end {
		tree = append(tree, nil)
		return tree
	}
	left, right := []*Tree{}, []*Tree{}
	for i := start; i <= end; i++ {
		left = generateTrees(start, i-1)
		right = generateTrees(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &Tree{node: i, left: l, right: r}
				tree = append(tree, root)
				fmt.Println("append", root)
			}
		}
	}
	return tree
}
