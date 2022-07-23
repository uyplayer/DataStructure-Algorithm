/*
 * @Author: uyplayer
 * @Date: 2022/7/22 23:19
 * @Email: uyplayer@qq.com
 * @File: reverse-linked-list-ii.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/linked list
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package linked_list

//
//  reverseBetween
//  @Description: 题意反转两个区间内的元素
//  @param head
//  @param left
//  @param right
//  @return *ListNode
//
func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &ListNode{
		Next: head,
	}
	pre := dummyNode
	// 先处理left
	for i := 0; i < left-1; i++ {
		pre = pre.Next

	}
	rightNode := pre
	for i := 0; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	leftNode := pre.Next
	current := rightNode.Next
	pre.Next = nil
	rightNode.Next = nil
	// 反转
	reverseLinkedList(leftNode)

	// 连接
	pre.Next = rightNode
	leftNode.Next = current
	return dummyNode.Next
}

//
//  reverseLinkedList
//  @Description:
//  @param head
//
func reverseLinkedList(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
}

//
//  reverseBetween
//  @Description: 另一种迭代
//  @param head
//  @param left
//  @param right
//  @return *ListNode
//
func reverseBetween1(head *ListNode, left int, right int) *ListNode {

	dummyNode := &ListNode{Next: head}
	pre := dummyNode

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	current := pre.Next
	for i := 0; i < right-left; i++ {
		next := current.Next
		current.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummyNode.Next

}
