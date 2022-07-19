/*
 * @Author: uyplayer
 * @Date: 2022/7/19 15:09
 * @Email: uyplayer@qq.com
 * @File: decode-string.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / leetflash/20220719
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package _0220719

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {

	stack := make([]string, 0)
	index := 0

	for index < len(s) {
		current := s[index]
		if current >= '0' && current <= '9' {
			digit := getDigits(s, &index)
			stack = append(stack, digit)
		} else if current >= 'a' && current <= 'z' || current >= 'A' && current <= 'Z' || current == '[' {
			stack = append(stack, string(current))
			index++
		} else {
			index++
			sub := []string{}
			for stack[len(stack)-1] != "[" {
				// 出栈
				sub = append(sub, stack[len(stack)-1])
				// 更新str
				stack = stack[:len(stack)-1]
			}
			// 反转
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			//   stk[len(stk)-1] != "[" 已经碰到右括号']'，去掉它
			stack = stack[:len(stack)-1]
			// 出栈 的时候数字
			repTime, _ := strconv.Atoi(stack[len(stack)-1])
			// 更新栈
			stack = stack[:len(stack)-1]
			// 重复
			t := strings.Repeat(getString(sub), repTime)
			// 添加到数组中
			stack = append(stack, t)

		}

	}

	return getString(stack)

}

//
//  getDigits
//  @Description:
//  @param s
//  @param ptr
//  @return string
//
func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

//
//  getString
//  @Description:
//  @param v
//  @return string
//
func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}
