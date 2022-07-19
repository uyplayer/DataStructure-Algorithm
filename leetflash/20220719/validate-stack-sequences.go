/*
 * @Author: uyplayer
 * @Date: 2022/7/19 14:52
 * @Email: uyplayer@qq.com
 * @File: validate-stack-sequences.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / leetflash/20220719
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package _0220719

func validateStackSequences(pushed []int, popped []int) bool {

	length := len(pushed)
	result := make([]int, 0)
	index := 0

	for i := 0; i < length; i++ {
		result = append(result, pushed[i])

		for len(result) > 0 && index < len(popped) && popped[index] == result[len(result)-1] {
			result = result[:len(result)-1]
			index++
		}

	}

	return len(result) == 0

}
