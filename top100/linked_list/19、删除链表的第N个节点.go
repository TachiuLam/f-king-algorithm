package linked_list

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]


提示：

链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
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
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next = p2.Next.Next // 删除倒数第N个节点
	return p3.Next
}
