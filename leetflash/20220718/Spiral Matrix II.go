/*
 * @Author: uyplayer
 * @Date: 2022/7/18 09:51
 * @Email: uyplayer@qq.com
 * @File: Spiral Matrix II.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / leetflash/20220718
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package _0220718

//
//  generateMatrix
//  @Description: 分支
//  @param n
//  @return [][]int
//
func generateMatrix(n int) [][]int {
	topBound := 0
	lowerBound := n - 1
	leftBound := 0
	rightBound := n - 1

	// 二维
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	// 怎么遍历
	i := 1
	for i <= n*n {

		// 先从上面
		if topBound <= lowerBound {
			// 现在从左到右边界
			for j := leftBound; j <= rightBound; j++ {
				matrix[topBound][j] = i
				i++
			}
			topBound++
		}

		if leftBound <= rightBound {
			for j := topBound; j <= lowerBound; j++ {
				matrix[j][rightBound] = i
				i++
			}
			rightBound--

		}

		if topBound <= lowerBound {
			for j := rightBound; j <= leftBound; j-- {
				matrix[lowerBound][j] = i
				i++
			}
			lowerBound--
		}

		if leftBound <= rightBound {
			// 在左侧从下向上遍历
			for i := lowerBound; i >= rightBound; i-- {

				matrix[i][leftBound] = i
				i++
			}
			// 左边界右移
			leftBound++
		}

	}

	return matrix

}
