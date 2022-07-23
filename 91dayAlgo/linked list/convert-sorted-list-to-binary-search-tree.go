/*
 * @Author: uyplayer
 * @Date: 2022/7/23 12:06
 * @Email: uyplayer@qq.com
 * @File: convert-sorted-list-to-binary-search-tree.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/linked list
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package linked_list

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//
//  sortedListToBST
//  @Description:
//  @param head
//  @return *TreeNode
//
func sortedListToBST(head *ListNode) *TreeNode {
	// 如果 head 空
	if head == nil {
		return nil
	}
	// 开始i遍历
	return dfs(head, nil)

}

func dfs(head, tail *ListNode) *TreeNode {
	// 递归推出题条件
	if head == tail {
		return nil
	}
	// 快慢指针
	fast, slow := head, head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 创建root节点
	root := &TreeNode{Val: slow.Val}
	root.Left = dfs(head, slow)
	root.Right = dfs(slow.Next, tail)
	return root

}

func sortedListToBST1(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	array := make([]int, 0)
	// 每个节点存到数组中
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}
	// 开始调用递归
	return dfs1(array, 0, len(array))
}

func dfs1(res []int, start, end int) *TreeNode {
	//  遍历完
	if start >= end {
		return nil
	}
	//  获取中间的节点
	mid := (end-start)/2 + start
	root := &TreeNode{Val: res[mid]}
	root.Left = dfs1(res, start, mid)
	root.Right = dfs1(res, mid+1, end)
	return root

}
