/*
 * @Author: uyplayer
 * @Date: 2022/7/22 22:21
 * @Email: uyplayer@qq.com
 * @File: partition-list.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/linked list
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package linked_list

//
//  partition
//  @Description:
//  @param head
//  @param x
//  @return *ListNode
//
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2 := &ListNode{}, &ListNode{}
	p := head
	dummmy1 := p1
	dummmy2 := p2
	for p != nil {
		if p.Val >= x {
			// p2.Next 指向 p
			p2.Next = p
			// 哨兵往前走一步，准备下一次遍历使用
			p2 = p
		} else {
			p1.Next = p
			p1 = p
		}
		// 断开 先保存下一个阶段
		tmp := p.Next
		//断开
		p.Next = nil
		p = tmp
	}

	p1.Next = dummmy2.Next
	return dummmy1.Next

}
