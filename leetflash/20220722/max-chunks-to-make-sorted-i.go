/*
 * @Author: uyplayer
 * @Date: 2022/7/22 19:03
 * @Email: uyplayer@qq.com
 * @File: max-chunks-to-make-sorted-i.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / leetflash/20220722
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package _0220722

import "sort"

//简单理解一下
//
//第一种方法
//
//排序前数组前n个如果和排序后数组前n个元素一样，那前n个就是可以被分块的，此时n要找到最小的
//
//重复这个过程
//
//第二种方法
//
//每个块内：排序后第k个，一定大于等于前k个数
//
//找到这个最小k，然后重复

//
//  maxChunksToSorted
//  @Description: 如果两个数组的前i项和相等，即arr的前i项和%20==%20sortArr的前i项和，则增加计数，因为它们可以组成一个块，如果不相等，继续遍历到相等的位置j，表示i到j这些元素也可以组成一个块，增加计数
//  @param arr
//  @return int
//

func maxChunksToSorted(arr []int) int {

	n := len(arr)
	array := make([]int, n)

	copy(array, arr)
	sort.Ints(array)

	//排序前数组前n个如果和排序后数组前n个元素一样，那前n个就是可以被分块的，此时n要找到最小的
	sum, ret := 0, 0
	for i := 0; i < n; i++ {
		sum += (arr[i] - array[i])
		if sum == 0 {

			ret++

		}

	}

	return ret

}
