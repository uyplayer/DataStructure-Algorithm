/*
 * @Author: uyplayer
 * @Date: 2022/7/22 00:39
 * @Email: uyplayer@qq.com
 * @File: deleteDuplicates.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/linked list
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package linked_list

//
//  deleteDuplicates
//  @Description: 去掉重复的 节点，只留下唯一的,d但是这个方法没法去掉所有的，想去掉的话利用字典
// 额外的空间来保存重复的元素，再一次循环删除
//  @param head
//  @return *ListNode
//
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next

		} else {
			current = current.Next
		}

	}

	return head

}

//
//  deleteDuplicates01
//  @Description: 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
// 看起来双指针比较好
//  @param head
//  @return *ListNode
//
func deleteDuplicates01(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	nodes := &ListNode{
		Next: head,
	}
	result := nodes
	for nodes.Next != nil {

		if nodes.Next.Val == nodes.Next.Next.Val {
			nodes.Next = nodes.Next.Next

		} else {
			nodes = nodes.Next
		}

	}

	return result.Next

}
