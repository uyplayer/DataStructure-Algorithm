/*
 * @Author: uyplayer
 * @Date: 2022/7/20 11:30
 * @Email: uyplayer@qq.com
 * @File: Substring.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/slide-window
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package slide_window

//
//  countSubArray
//  @Description: 前缀和
//  @param array
//  @return int
//
func sumSubArray(array []int) []int {

	pre := make([]int, len(array))
	pre[0] = 1
	for i := 1; i < len(array); i++ {
		pre[i] = pre[i-1] + array[i]
	}

	return pre

}

//
//  countSubArray
//  @Description:连续子数组总个数
//  @param nums
//  @return int
//
func countSubArray(nums []int) int {
	ans := 0
	pre := 0
	for _ = range nums {
		pre += 1
		ans += pre
	}
	return ans
}
