package 双指针

//leetcode submit region begin(Prohibit modification and deletion)
/**
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
// 分治思想：
// 时间复杂度：渐进时间复杂度为 O(kn×logk)
// 递归空间复杂度：O(logn)
/*
将k 个链表配对并将同一对中的链表合并；
第一轮合并为 k/2个链表，依次合并
*/
func mergeKLists(lists []*ListNode) *ListNode {
	var pre, cur *ListNode
	n := len(lists)
	if n <= 0 {
		return nil
	}
	pre = lists[0]
	for i := 1; i < n; i++ {
		cur = lists[i]
		pre = mergeTowList(pre, cur)
	}
	return pre
}

func mergeTowList(list1, list2 *ListNode) *ListNode {
	res := &ListNode{}
	p := res
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return res.Next
}
