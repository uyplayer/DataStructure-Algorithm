/*
 * @Author: uyplayer
 * @Date: 2022/7/18 09:29
 * @Email: uyplayer@qq.com
 * @File: buddy-strings.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / leetflash/20220718
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package _0220718

func buddyStrings(s string, goal string) bool {

	sLength := len(s)
	gLength := len(goal)

	if sLength != gLength {
		return false
	}

	// 记录不一样的字符
	diffIndex := make([]int, 0)
	//记录值出现的次数
	byteRepeat := map[byte]int{}
	//记录是否存在重复
	duplicateByte := false
	for i := 0; i < sLength; i++ {
		if s[i] != goal[i] {
			diffIndex = append(diffIndex, i)
			// 超过2 说明不能
			if len(diffIndex) > 2 {
				return false
			}
		} else {

			byteRepeat[s[i]]++
			// 如果有重复的
			if byteRepeat[s[i]] > 1 {
				duplicateByte = true
			}
		}
	}

	if len(diffIndex) == 2 && s[diffIndex[0]] == goal[diffIndex[1]] && s[diffIndex[1]] == goal[diffIndex[0]] {
		return true
	}

	// 查看重复的
	if len(diffIndex) == 0 && duplicateByte {

		return true

	}

	return false

}
