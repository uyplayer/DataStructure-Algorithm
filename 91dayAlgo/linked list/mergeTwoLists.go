/*
 * @Author: uyplayer
 * @Date: 2022/7/22 00:09
 * @Email: uyplayer@qq.com
 * @File: mergeTwoLists.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/linked list
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package linked_list

//
//  mergeTwoLists
//  @Description:
//  @param list1
//  @param list2
//  @return *ListNode
//
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list2 == nil && list2 != nil {
		return nil
	}
	// 哨兵来啦
	nodes := &ListNode{}
	result := nodes
	for list1 != nil && list2 != nil {

		// 如果用 nodes.Val = list1.Val 的话 每次我们要创建一个节点，nodes.Next = 新节点。这样浪费空间
		if list1.Val < list2.Val {
			nodes.Next = list1
			list1 = list1.Next
		} else {
			nodes.Next = list2
			list2 = list2.Next
		}
		// 更新
		nodes = nodes.Next

	}

	if list1 == nil {
		nodes.Next = list2
	}
	if list2 == nil {
		nodes.Next = list1
	}
	// 更新
	nodes = nodes.Next
	return result

}
