package test

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	slow, high := &ListNode{}, &ListNode{}
	p1, p2 := slow, high
	i := 0
	for head != nil {
		i++
		fmt.Println("循环...", i)
		fmt.Println("head = ", head, head.Next)
		if head.Val < x {
			p1.Next = head
			//fmt.Println("p1 = ", p1, p1.Next)
			p1 = p1.Next
			//fmt.Println("p1 = ", p1, p1.Next)
		} else {
			p2.Next = head
			p2 = p2.Next
		}
		fmt.Println("p1 = ", p1, "p1.Next", p1.Next)
		fmt.Println("p2 = ", p2, p2.Next)
		fmt.Println("slow = ", slow, slow.Next, slow.Next.Next, slow.Next.Next.Next)
		fmt.Println("high = ", high, high.Next)
		head = head.Next
	}
	p2.Next = nil
	p1.Next = high.Next
	return slow.Next
}

func Test86(t *testing.T) {
	l3 := &ListNode{
		Val: 2,
	}
	l2 := &ListNode{
		Val:  3,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  4,
		Next: l2,
	}
	l0 := &ListNode{
		Val:  1,
		Next: l1,
	}
	t.Log(partition(l0, 3))
}
