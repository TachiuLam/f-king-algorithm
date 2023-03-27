package 双指针

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 双指针：
//1、 p1 = head; p2 = head - 1; p1 先走n步
//2、p1,p2 同时走，p1到达链表末尾时，p2到达倒数第n个节点前
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	p3 := &ListNode{0, head}
	p1, p2 := head, p3
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	p2.Next = p2.Next.Next
	return p3.Next
}
