/*
 * @Author: uyplayer
 * @Date: 2022/7/22 17:16
 * @Email: uyplayer@qq.com
 * @File: add-to-array-form-of-integer.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / leetflash/20220722
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package _0220722

import (
	"fmt"
	"strconv"
)

//
//  addToArrayForm
//  @Description: 989. 数组形式的整数加法 , 该方法如果位数太长了无效
//  @param num
//  @param k
//  @return []int
//
func addToArrayForm01(num []int, k int) []int {
	s := ""
	for i := 0; i < len(num); i++ {
		result := strconv.Itoa(num[i])
		s = s + result

	}
	result, _ := strconv.Atoi(s)
	sumInt := result + k
	sumString := strconv.Itoa(sumInt)
	fmt.Println(sumString)
	arr := make([]int, 0)
	for i := 0; i < len(sumString); i++ {
		n, _ := strconv.Atoi(string(sumString[i]))
		arr = append(arr, n)
	}

	return arr
}

//
//  addToArrayForm
//  @Description:
//  @param num
//  @param k
//  @return []int
//
func addToArrayForm(num []int, k int) []int {

	// 保存结果
	result := make([]int, 0)
	// 最后一位开始遍历
	for i := len(num) - 1; i >= 0; i-- {
		// 任何两位数加起来都是0-18之间
		// 先加起来，k%10取最后一位，这里搞k%10 获取最后一位
		sum := num[i] + k%10
		// 获取剩下的部分
		k = k / 10
		// sum 大于等于10的情况
		if sum >= 10 {
			// 大于10 的时候 ，前面的第一位的 1 加到 k 上
			k++
			// 上面 前面的第一位的 1 加到 k 上
			sum = sum - 10
			result = append(result, sum)

		} else {
			result = append(result, sum)
		}

	}

	// 好了没有  。当然没有，如果k开没加完或者k的位数大于数组的位数呢？那就继续加
	for ; k > 0; k /= 10 {
		result = append(result, k%10)
	}

	// 反转
	for i := 0; i < len(result)/2; i++ {

		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]

	}

	return result

}

//
//  addToArrayForm02
//  @Description:  上面的方法和这方法区别在于，这方法价值加到k里面取
//  @param num
//  @param k
//  @return []int
//
func addToArrayForm02(num []int, k int) []int {

	// 保存结果
	result := make([]int, 0)
	// 最后一位开始遍历
	for i := len(num) - 1; i >= 0 || k > 0; i-- {

		if i >= 0 {
			k = k + num[i]
		}

		result = append(result, k%10)
		k /= 10

	}

	// 反转
	for i := 0; i < len(result)/2; i++ {

		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]

	}

	return result

}
