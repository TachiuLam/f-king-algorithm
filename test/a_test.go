package test

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestDumpPoint(t *testing.T) {
	l3 := &ListNode{
		Val: 4,
	}
	l1 := new(ListNode)
	l2 := l1
	l1.Val = 2
	l2.Val = 3
	l2.Next = l3
	l2 = l2.Next
	t.Log(l1, l2)
}
