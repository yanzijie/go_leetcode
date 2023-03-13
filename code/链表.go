package code

// MergeTwoListsMy /*****************************21. 合并两个有序链表*****************************/
/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
// 方法1, 定义新链表, 有头尾指针
func MergeTwoListsMy(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	head := &ListNode{}
	tail := head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			tail.Next = list2
			return head.Next
		}

		if list2 == nil {
			tail.Next = list1
			return head.Next
		}

		if list1.Val > list2.Val {
			tail.Next = list2
			list2 = list2.Next
			tail = tail.Next
			tail.Next = nil
		} else {
			tail.Next = list1
			list1 = list1.Next
			tail = tail.Next
			tail.Next = nil
		}
	}
	return head
}

// MergeTwoListTwoPoint 方法2：双指针法, 效率更高
func MergeTwoListTwoPoint(list1 *ListNode, list2 *ListNode) *ListNode {
	//初始化一个虚拟的头节点
	newList := &ListNode{}
	p := newList
	p1 := list1
	p2 := list2
	//遍历对比两个指针值的大小，有一个走完了就停止
	for p1 != nil && p2 != nil {
		//将值小的节点接到p指针后面
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		//p指针前进, 但是newList保留在头部,所以 newList 的下一个就是真正的链表头
		p = p.Next
	}
	//将未比较的剩余节点都放到p指针后面
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	//返回虚拟头节点的下一个节点就是真正的头节点
	return newList.Next
}

/*****************************21. 合并两个有序链表*****************************/
