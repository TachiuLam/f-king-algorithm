package 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// base case
	if left == 1 {
		return reverseN(head, right)
	}
	// 前进到反转的起点触发 base case
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var successor *ListNode

// 反转链表前N个节点
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}
