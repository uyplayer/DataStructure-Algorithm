/*
 * @Author: uyplayer
 * @Date: 2022/7/19 14:24
 * @Email: uyplayer@qq.com
 * @File: shortest-distance-to-a-character.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / leetflash/20220718
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package _0220718

import "math"

func shortestToChar(s string, c byte) []int {

	cLoc := make([]int, 0)
	for key := range s {
		if s[key] == c {
			cLoc = append(cLoc, key)
		}
	}
	result := make([]int, 0)
	for i := 0; i < len(s); i++ {
		min := 99999999
		for j := 0; j < len(cLoc); j++ {
			distance := int(math.Abs(float64(cLoc[j] - i)))
			if distance < min {
				min = distance
			}

		}
		result = append(result, min)
	}

	return result

}
