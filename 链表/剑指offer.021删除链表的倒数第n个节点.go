package 链表 //leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	p3 := &ListNode{0, head}
	p1, p2 := head, p3
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next
	return p3.Next
}
