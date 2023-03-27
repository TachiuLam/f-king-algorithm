package 双指针

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快慢指针解法：快指针到达末尾时，慢指针刚好到达中点
// 时间复杂度O(N)，空间复杂度O(1)
func middleNode(head *ListNode) *ListNode {
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		head = head.Next
	}
	return head
}
