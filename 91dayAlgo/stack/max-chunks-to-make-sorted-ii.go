/*
 * @Author: uyplayer
 * @Date: 2022/7/20 22:29
 * @Email: uyplayer@qq.com
 * @File: max-chunks-to-make-sorted-ii.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/stack
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package stack

import "sort"

func maxChunksToSorted(arr []int) int {

	length := len(arr)
	if length <= 1 {
		return 1
	}
	array := make([]int, length)
	copy(array, arr)
	sort.Ints(array)
	sum, ret := 0, 0
	for i := 0; i < length; i++ {
		sum += (arr[i] - array[i])
		if sum == 0 {
			ret++
		}
	}
	return ret

}
