package 双指针

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 递归解法
func mergeKLists2(lists []*ListNode) *ListNode {
	return merge(0, len(lists)-1, lists)
}

func merge(left, right int, lists []*ListNode) *ListNode {
	if left > right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	mid := (left + right) >> 1
	return mergeTow(merge(left, mid, lists), merge(mid+1, right, lists))
}

func mergeTow(list1, list2 *ListNode) *ListNode {
	p := &ListNode{}
	head := p
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
	return head.Next
}
