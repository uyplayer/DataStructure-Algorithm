/*
 * @Author: uyplayer
 * @Date: 2022/7/21 01:24
 * @Email: uyplayer@qq.com
 * @File: findDuplicates.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/array
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package array

func findDuplicates(nums []int) []int {
	if len(nums) == 1 {
		return []int{}
	}

	array := make([]int, 0)
	dic := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dic[nums[i]]++
		if dic[nums[i]] == 2 {
			array = append(array, nums[i])
		}
	}

	return array

}
