/*
 * @Author: uyplayer
 * @Date: 2022/7/21 10:55
 * @Email: uyplayer@qq.com
 * @File: findSubstringInWraproundString.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/slide-window
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package slide_window

import "math"

//
//  findSubstringInWraproundString
//  @Description: 求没有重复的子串 ，可以用双嵌套循环，但是会超时呀 哈哈 ，解决方案滑动窗口
//  @param p
//  @return int
//
func findSubstringInWraproundString(p string) int {
	p = string('^') + p
	w := 1
	// 记录
	dic := make(map[string]int)

	// 遍历
	for i := 1; i < len(p); i++ {
		// 判断是否连续
		if (p[i]-p[i-1]) == 1 || -(p[i]-p[i-1]) == 25 {

			w++

		} else {
			w = 1
		}
		dic[string(p[i])] = int(math.Max(float64(dic[string(p[i])]), float64(w)))

	}
	sum := 0
	for _, value := range dic {
		sum = sum + value
	}
	return sum
}
