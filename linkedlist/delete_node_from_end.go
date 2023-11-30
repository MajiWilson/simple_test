package linkedlist

/**
19. 删除链表的倒数第 N 个结点
提示
中等
2.7K
相关企业
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。


*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 删除头节点
	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 删除节点
	slow.Next = slow.Next.Next
	return head
}
