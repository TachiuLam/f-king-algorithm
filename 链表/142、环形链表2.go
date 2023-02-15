package 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快慢指针
/* 解题思路
1、慢指针入环后，第一圈一定被快指针追上：快指针先入环，且快指针走两圈=慢指针走一圈
2、相遇时：慢指针行走距离：a + b (a:起点到入环点；b：入环点-相遇点)；快指针行走距离：a + b + n(b +c) (n：圈数， b+c=环长度)
故：2(a + b) = a + b + n(b+c) => a = (n - 1)(b + c) + c
所以相遇后，新指针从链表起点出发，慢指针从相遇点触发，两个指针步长一致，必定在a点相遇
*/
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 相遇后，再计算离环入口距离
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
