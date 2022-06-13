package addTwoNumbers

import (
	"testing"

	"github.com/mailtokun/ALG/entity"
)

func Test_Unit(t *testing.T) {
	l1 := entity.ListNode{
		Val: 1,
		Next: &entity.ListNode{
			Val: 2,
			Next: &entity.ListNode{
				Val: 3,
			},
		},
	}
	l2 := entity.ListNode{
		Val: 1,
		Next: &entity.ListNode{
			Val: 2,
			Next: &entity.ListNode{
				Val: 3,
			},
		},
	}
	target := addTwoNumbers(&l1, &l2)

	if target.Val != 2 || target.Next.Val != 4 || target.Next.Next.Val != 6 {
		t.Fatal()
	}
}
func Test_Unit2(t *testing.T) {
	l1 := entity.ListNode{
		Val: 9,
		Next: &entity.ListNode{
			Val: 8,
			Next: &entity.ListNode{
				Val: 7,
			},
		},
	}
	l2 := entity.ListNode{
		Val: 1,
		Next: &entity.ListNode{
			Val: 2,
			Next: &entity.ListNode{
				Val: 3,
			},
		},
	}
	target := addTwoNumbers(&l1, &l2)

	if target.Val != 0 || target.Next.Val != 1 || target.Next.Next.Val != 1 || target.Next.Next.Next.Val != 1 {
		t.Fatal()
	}
}
func Test_Unit3(t *testing.T) {
	l1 := entity.ListNode{
		Val: 9,
		Next: &entity.ListNode{
			Val: 9,
			Next: &entity.ListNode{
				Val: 9,
				Next: &entity.ListNode{
					Val: 9,
					Next: &entity.ListNode{
						Val: 9,
						Next: &entity.ListNode{
							Val: 9,
							Next: &entity.ListNode{
								Val: 9,
							},
						},
					},
				},
			},
		},
	}
	l2 := entity.ListNode{
		Val: 9,
		Next: &entity.ListNode{
			Val: 9,
			Next: &entity.ListNode{
				Val: 9,
				Next: &entity.ListNode{
					Val: 9,
				},
			},
		},
	}
	target := addTwoNumbers(&l1, &l2)

	if target.Val != 8 || target.Next.Val != 9 || target.Next.Next.Val != 9 || target.Next.Next.Next.Val != 9 || target.Next.Next.Next.Next.Val != 0 || target.Next.Next.Next.Next.Next.Val != 0 || target.Next.Next.Next.Next.Next.Next.Val != 0 || target.Next.Next.Next.Next.Next.Next.Next.Val != 1 {
		t.Fatal()
	}
}
