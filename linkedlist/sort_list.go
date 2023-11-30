package linkedlist

/**
方法一： 归并排序
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	midNode := findMidNode(head)
	rightHead := midNode.Next
	midNode.Next = nil

	leftSortedHead := sortList(head)
	rightSortedHead := sortList(rightHead)

	return merge2List(leftSortedHead, rightSortedHead)
}

/**
	中点有两个， 左中点和右中点
	1     1
    2     1 2    如果fastNode = head， 那么永远都会循环在2的长度里面， 所以要注意这里！！！！
    3     2
    4     2 3
    5     3
    6     3 4
    7     4
*/
func findMidNode(head *ListNode) *ListNode {
	slowNode := head
	// 注意
	fastNode := head.Next
	for fastNode != nil && fastNode.Next != nil {
		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next
	}

	return slowNode
}

/**
这样也可以保证使用的是左中点
*/
func findMidNode2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slowNode := head
	fastNode := head
	for fastNode.Next != nil && fastNode.Next.Next != nil {
		fastNode = fastNode.Next.Next
		slowNode = slowNode.Next
	}

	return slowNode
}

func merge2List(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = merge2List(list1.Next, list2)
		return list1
	} else {
		list2.Next = merge2List(list1, list2.Next)
		return list2
	}
}

/**
方法2: 快排
*/
func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 划分
	pNode := partition(head)

	// 找到左边界
	if pNode == head {
		head.Next = sortList(head.Next)
		return head
	}
	preNode := head
	for preNode.Next != pNode {
		preNode = preNode.Next
	}
	preNode.Next = nil

	//分治
	leftList := sortList(head)
	rightList := sortList(pNode.Next)

	preNode = leftList
	for preNode.Next != nil {
		preNode = preNode.Next
	}
	preNode.Next = pNode
	pNode.Next = rightList

	return leftList
}

func partition(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pivot := head.Val
	curNode := head.Next
	preNode := head

	for curNode != nil {
		if curNode.Val < pivot {
			preNode = preNode.Next
			temp := curNode.Val
			curNode.Val = preNode.Val
			preNode.Val = temp
		}
		curNode = curNode.Next
	}

	head.Val = preNode.Val
	preNode.Val = pivot

	return preNode
}
