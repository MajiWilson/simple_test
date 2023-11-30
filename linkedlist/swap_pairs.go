package linkedlist

/**
24. 两两交换链表中的节点
中等
2.1K
相关企业
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。


*/

/**
非递归算法：需要借助一个Pre
*/
func swapPairs(head *ListNode) *ListNode {
	curNode := head
	tempHead := &ListNode{}
	preNode := tempHead
	preNode.Next = head

	for curNode != nil && curNode.Next != nil {
		nextNode := curNode.Next
		curNode.Next = curNode.Next.Next
		nextNode.Next = curNode
		preNode.Next = nextNode
		preNode = curNode
		curNode = preNode.Next
	}
	return tempHead.Next
}

/**
递归算法： 交互第一和第二个节点
*/

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next

	head.Next = swapPairs(head.Next.Next)
	newHead.Next = head
	return newHead
}
