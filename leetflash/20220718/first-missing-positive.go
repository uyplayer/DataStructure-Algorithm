/*
 * @Author: uyplayer
 * @Date: 2022/7/18 09:15
 * @Email: uyplayer@qq.com
 * @File: first-missing-positive.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / leetflash/20220718
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package _0220718

//
//  firstMissingPositive
//  @Description:
//  @param nums
//  @return int
//
func firstMissingPositive(nums []int) int {

	length := len(nums)
	for i := 0; i < length; i++ {
		for nums[i] > 0 && nums[i] <= length && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < length; i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}

	return length + 1

}
