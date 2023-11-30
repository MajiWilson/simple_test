package linkedlist

/**
234. 回文链表
简单
1.8K
相关企业
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
*/

/**
(1) 遍历后存储在数组中判断
（2）利用递归方法
(3) 翻转后半部分， 然后判断即可
*/
func isPalindrome(head *ListNode) bool {
	preNode := head
	var check func(curNode *ListNode) bool
	check = func(curNode *ListNode) bool {
		if curNode == nil {
			return true
		}
		if !check(curNode.Next) {
			return false
		}

		if curNode.Val != preNode.Val {
			return false
		}

		preNode = preNode.Next

		return true
	}

	return check(head)
}

func isPalindrome2(head *ListNode) bool {
	// 先找到中点
	if head == nil || head.Next == nil {
		return true
	}

	fastNode := head.Next.Next
	slowNode := head
	for fastNode != nil && fastNode.Next != nil {
		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next
	}
	// 逆转后半部分
	fastNode = reverse(slowNode)
	slowNode = head

	// 比较(同步）
	for fastNode != nil && slowNode != nil {
		if fastNode.Val != slowNode.Val {
			return false
		}
		fastNode = fastNode.Next
		slowNode = slowNode.Next
	}
	return true

}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
