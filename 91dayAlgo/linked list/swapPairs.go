/*
 * @Author: uyplayer
 * @Date: 2022/7/22 10:41
 * @Email: uyplayer@qq.com
 * @File: swapPairs.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/linked list
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package linked_list

//
//  swapPairs
//  @Description:给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//  @param head
//  @return *ListNode
//
func swapPairs(head *ListNode) *ListNode {

	nodes := &ListNode{
		Next: head,
	}

	//result := nodes
	//previous := nodes
	//for head != nil && head.Next != nil {
	//	// 当前的节点和下一个节点
	//	now := head
	//	next := head.Next
	//	// now 哨兵指向next的下一个节点
	//	now.Next = next.Next
	//	next.Next = previous.Next
	//	previous.Next = next
	//
	//	previous = now
	//	head = now.Next
	//
	//}

	//result := nodes
	//previous := nodes
	//for previous.Next != nil && previous.Next.Next != nil {
	//	now := previous.Next
	//	next := previous.Next.Next
	//
	//	previous.Next = next
	//	now.Next = next.Next
	//	next.Next = now
	//
	//	previous = now
	//
	//}

	previous := nodes
	current := previous.Next // 相当于head
	for current != nil && current.Next != nil {
		next := current.Next
		current.Next = next.Next
		next.Next = current
		previous.Next = next

		previous = current
		current = current.Next
	}

	return nodes.Next
}

//
//  swapPairs
//  @Description: 递归
//  @param head
//  @return *ListNode
//
func swapPairs01(head *ListNode) *ListNode {

	// 递归结束条件
	if head == nil && head.Next == nil {
		return head
	}

	now := head
	next := head.Next
	nextNext := swapPairs(next.Next)
	now.Next = nextNext
	next.Next = now

	return next

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs02(head *ListNode) *ListNode {

	nodes := &ListNode{
		Next: head,
	}

	result := nodes
	pre := nodes
	for head != nil && head.Next != nil {

		current := head
		next := head.Next
		current.Next = next.Next
		next.Next = pre.Next
		pre.Next = next

		pre = current
		head = current.Next

	}

	return result.Next

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs03(head *ListNode) *ListNode {

	nodes := &ListNode{
		Next: head,
	}

	result := nodes
	pre := nodes
	for nodes.Next != nil && nodes.Next.Next != nil {
		current := nodes.Next
		next := nodes.Next.Next
		// 当前节点 A
		current.Next = next.Next
		// 节点B  注意这里 如果写成  next.Next = current 形成一个环
		next.Next = pre.Next
		// preANext
		pre.Next = next

		// pre   修改指针位置，进行下一轮逆转
		pre = current
		nodes = nodes.Next

	}

	return result.Next

}
