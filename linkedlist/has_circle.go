package linkedlist

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fastNode := head.Next.Next
	slowNode := head
	for fastNode != nil && fastNode.Next != nil {
		if fastNode == slowNode {
			return true
		}

		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next
	}
	return false
}

/**
判断是否有换并且找到入口
*/

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fastNode := head
	slowNode := head
	for fastNode != nil && fastNode.Next != nil {
		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next
		if fastNode == slowNode {
			break
		}
	}
	// 判断有无环
	if fastNode == nil || fastNode.Next == nil {
		return nil
	}

	slowNode = head
	for slowNode != fastNode {
		fastNode = fastNode.Next
		slowNode = slowNode.Next
	}

	return fastNode
}
