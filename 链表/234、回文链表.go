package 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法一：将链表反转成另一链表，再比较，空间复杂度O(n)
// 解法二：反转链表中点到末尾的链表，和前半部分进行比较，空间复杂度O(n/2)
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// fast != nil 的话，说明链表长度是偶数，slow还需要前进一步，到达回文起点
	if fast != nil {
		slow = slow.Next
	}

	left := head
	//q : = right
	right := reverseListNodes(slow)
	// 判断是否相等
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		//p := left
		left = left.Next
		right = right.Next
	}
	// 如果需要恢复被反转的后半部分链表
	//p.Next = reverseL(q)
	return true
}

func reverseListNodes(head *ListNode) *ListNode {
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
