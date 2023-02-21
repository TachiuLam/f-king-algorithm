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

func merge(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	p := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return res.Next
}

func TestDe(t *testing.T) {
	t.Log((7 + 1) >> 1)
	s := make([]int, 2, 5)
	t.Log(s[0], s[1])
	{
		s := []int{1, 2, 3, 4, 5, 6}

		fmt.Println(s, len(s), cap(s))

		sc := make([]int, 6, 6)

		copy(sc, s)
		sc[0] = 100
		fmt.Println(s, len(s), cap(s))
		fmt.Println(sc, len(sc), cap(sc))
	}
}
