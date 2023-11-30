package linkedlist

/**
21. 合并两个有序链表
简单
3.4K
相关企业
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。


*/

/**
递归
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 递归出口
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		listHead := mergeTwoLists(list1.Next, list2)
		list1.Next = listHead
		return list1
	} else {
		listHead := mergeTwoLists(list1, list2.Next)
		list2.Next = listHead
		return list2
	}
}
