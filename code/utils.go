package code

import "fmt"

/*****************************创建两个链表*****************************/

type ListNode struct {
	data int
	Next *ListNode
}

// 尾插法插入节点
func (p *ListNode) Append(data int) {
	for p.Next != nil {
		p = p.Next
	}
	var newNode *ListNode = new(ListNode)
	newNode.data = data
	p.Next = newNode

	fmt.Printf("插入数据：%d\n", data)
}

// 遍历输出
func (p *ListNode) Traverse() {
	fmt.Printf("遍历结果：")
	for p.Next != nil {
		fmt.Printf("%d ", p.Next.data)
		p = p.Next
	}
	fmt.Printf("\n")
}

func DoSingleNode(head1 *ListNode, head2 *ListNode) {

	head1.Append(1)
	head1.Append(2)
	head1.Append(4)
	head1.Traverse()

	head2.Append(1)
	head2.Append(3)
	head2.Append(4)
	head2.Traverse()

}

/*****************************创建两个链表*****************************/
