package 链表

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针 + 虚拟头节点( 时间复杂度O(N))
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := new(ListNode)
	p := res // res和p是同一个对象，指向同一地址
	for list1 != nil && list2 != nil {
		// 第一次循环，res.Nest 指针就是 p.Next 指向的地址
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		// p 指针前进
		p = p.Next // 一次循环后，p 被赋予下一个对象的地址空间，不再与res一致，res.Nest也不再变化
	}
	// list1不为空，放在p后面
	if list1 != nil {
		p.Next = list1
	}
	// list2不为空，放在p后面
	if list2 != nil {
		p.Next = list2
	}
	return res.Next
}

//leetcode submit region end(Prohibit modification and deletion)
