/*
 * @Author: uyplayer
 * @Date: 2022/7/21 21:18
 * @Email: uyplayer@qq.com
 * @File: rotateRight.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/linked list
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package linked_list

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//func rotateRight(head *ListNode, k int) *ListNode {
//
//	arr := make([]*ListNode, 0)
//
//	for head != nil {
//		arr = append(arr, head)
//		head = head.Next
//
//	}
//
//	count := 0
//	for count < k {
//		for i := 0; i < len(arr)-1; i++ {
//			if i == 0 {
//				arr[0] = arr[len(arr)-1]
//				continue
//			}
//			arr[i] = arr[i+1]
//		}
//	}
//
//	index := 0
//	for head != nil && index < len(arr) {
//
//		head.Next = arr[index]
//		head = head.Next
//
//	}
//
//	return head
//
//}

//
//  rotateRight1
//  @Description: 移动方法
//  @param head
//  @param k
//  @return *ListNode
//
func rotateRight1(head *ListNode, k int) *ListNode {

	if head == nil {
		return nil
	}
	//tail := &ListNode{
	//	Next: head,
	//}
	tail := head
	mid, cur := head, head
	num := 0
	// 找
	for cur != nil {
		num++
		tail = tail.Next
		cur = cur.Next
	}
	// 计算要移动步数
	k = k % num
	// 开始移动
	for i := 0; i < num-k-1; i++ {
		mid = mid.Next
	}
	// 换
	tail.Next = head
	// 连接头
	head = mid.Next
	// 断开
	mid.Next = nil

	return head

}

//
//  rotateRight
//  @Description: 快慢指针方法
//  @param head
//  @param k
//  @return *ListNode
//
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//找到tail
	curent := head
	listLength := 0
	// curent会指向当前的链表结尾节点
	for curent != nil {
		listLength++
		curent = curent.Next
	}
	// 计算移动次数
	k = k % listLength
	// 通过快慢指针找到头节点节点
	slow, fast := head, head
	for fast.Next != nil {
		fast = fast.Next
		// 移动 n-k-1 次 找到那个分解点
		k--
		if k < 0 {
			slow = slow.Next

		}
	}
	// 做环
	fast.Next = head
	res := slow.Next
	slow.Next = nil

	return res

}
