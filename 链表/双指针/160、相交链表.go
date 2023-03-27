package 双指针

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 双指针，每个指针走两个链表 (A+B) （B+A） 如果链表相交，则一定会在相交点的入口相遇
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 一个链表为空则一定不存在相交点
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	// 不存在交点时，p1 p2 都同时走完headA+headB，指向nil，返回nil
	for p1 != p2 {
		// p1 走到headA尽头，则指向headB
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		// p2 走到headB尽头，则指向headA
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
