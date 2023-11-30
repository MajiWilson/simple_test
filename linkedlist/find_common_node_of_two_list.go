package linkedlist

/**
160. 相交链表
简单
2.3K
相关企业
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。
*/

/**
一定会在交汇处相遇或者 都走完，
列表A长度： a + c
列表B长度： b + c
每个指针都走过
	a + b + c
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA := headA
	nodeB := headB

	for nodeA != nodeB {
		if nodeA == nil {
			nodeA = headB
		} else {
			nodeA = nodeA.Next
		}

		if nodeB == nil {
			nodeB = headA
		} else {
			nodeB = nodeB.Next
		}
	}

	return nodeB
}
