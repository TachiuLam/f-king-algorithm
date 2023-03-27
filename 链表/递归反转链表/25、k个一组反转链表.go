package 递归反转链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	a, b := head, head
	for i := 0; i < k; i++ {
		// base case ，不足k个点的直接返回初始指针
		if b == nil {
			return head
		}
		b = b.Next
	}
	// 反转前k个节点
	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 反转[a, b)之间的节点
func reverse(a *ListNode, b *ListNode) *ListNode {
	pre := &ListNode{}
	cur, nxt := a, a
	for cur != b {
		// 反转下一个节点
		nxt = cur.Next
		cur.Next = pre
		// 更新指针位置
		pre = cur
		cur = nxt
	}
	return pre
}
