package 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome2(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	left := head
	p := left
	right := reverseL(slow)
	q := right
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	p.Next = reverseL(q)
	return true
}

func reverseL(head *ListNode) *ListNode {
	var pre, nxt *ListNode
	cur := head
	for cur != nil {
		nxt = cur.Next
		cur.Next = pre

		pre = cur
		cur = nxt
	}
	return pre
}
