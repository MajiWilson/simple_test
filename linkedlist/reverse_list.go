package linkedlist

/**
头插法（可以不用新建头节点，原位进行， 因为每个节点只会遍历一次
*/
func reverseList1(head *ListNode) *ListNode {
	var newList *ListNode
	curNode := head
	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = newList
		newList = curNode
		curNode = nextNode
	}
	return newList

}

/**
递归方法
*/
func reverseList2(head *ListNode) *ListNode {
	// 递归出口
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList2(head.Next)

	head.Next.Next = head
	head.Next = nil
	return newHead

}
