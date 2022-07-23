/*
 * @Author: uyplayer
 * @Date: 2022/7/21 12:54
 * @Email: uyplayer@qq.com
 * @File: numSubarrayBoundedMax.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/slide-window
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package slide_window

func numSubarrayBoundedMax(nums []int, left int, right int) int {

	return find(right, nums) - find(left-1, nums)
}

func find(n int, nums []int) int {

	result := 0
	pre := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] <= n {
			pre++
		} else {
			pre = 0
		}
		result = result + pre
	}

	return result
}
