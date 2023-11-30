package linkedlist

import "container/heap"

/**

23. 合并 K 个升序链表
困难
2.7K
相关企业
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。


*/

/**
思路： 分治 合并， 效率比顺序高
*/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)

}

func merge(lists []*ListNode, left int, right int) *ListNode {
	// 递归出口
	if left > right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)/2
	return mergeTwoList(merge(lists, left, mid), merge(lists, mid+1, right))
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoList(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoList(list1, list2.Next)
		return list2
	}
}

/**
使用小跟堆
*/

func mergeKLists2(lists []*ListNode) *ListNode {
	h := Hp{}
	for _, v := range lists {
		if v != nil {
			h = append(h, v)
		}
	}
	heap.Init(&h)
	dummy := &ListNode{}
	p := dummy
	for len(h) > 0 {
		v := heap.Pop(&h).(*ListNode)
		if v.Next != nil {
			heap.Push(&h, v.Next)
		}
		p.Next = v
		p = p.Next
	}
	return dummy.Next
}

type Hp []*ListNode

func (h Hp) Len() int {
	return len(h)
}

func (h Hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h Hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Hp) Push(v interface{}) {
	*h = append(*h, v.(*ListNode))
}

func (h *Hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
