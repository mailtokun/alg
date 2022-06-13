package addTwoNumbers

import "github.com/mailtokun/ALG/entity"

/**
* 输入：l1 = [2,4,3], l2 = [5,6,4]
* 输出：[7,0,8]
* 解释：342 + 465 = 807.
**/
func addTwoNumbers(l1 *entity.ListNode, l2 *entity.ListNode) *entity.ListNode {
	var header entity.ListNode
	current := &header
	currentX := l1
	currentY := l2
	modulus := 0
	for currentX != nil || currentY != nil {
		x, y := 0, 0
		if currentX != nil {
			x = currentX.Val
			currentX = currentX.Next
		}
		if currentY != nil {
			y = currentY.Val
			currentY = currentY.Next
		}
		t := (x + y + modulus) % 10
		modulus = (x + y + modulus) / 10
		current.Next = &entity.ListNode{Val: t}
		current = current.Next
	}
	if modulus != 0 {
		current.Next = &entity.ListNode{Val: modulus}
	}
	return header.Next
}
