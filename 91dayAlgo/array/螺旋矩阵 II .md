## LC 59. 螺旋矩阵 II

[Leetcode连接](https://leetcode.cn/problems/spiral-matrix-ii/)

- [X-Y方向方法](#solution1)
- [分支方法方向方法](#solution2)

### <span id="solution1">X-Y方向方法</span>

（此处撰写思路）

### Golang 代码

``` go
func generateMatrix(n int) [][]int {

    type pair struct{ x, y int }
    var dirs = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上



    // 首先这里它们都是二位数组，所以第一步我们得构造二维数组
    matrix := make([][]int, n)
    for i:=0;i<n;i++{
        matrix[i] = make([]int, n)
        
    }
    // 二维数组里面每个元素有方向
    row, col, dirIdx := 0, 0, 0
    // 遍历元素
     for i := 1; i <= n*n; i++ {
         matrix[row][col] = i
         dir := dirs[dirIdx]
         if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
            dirIdx = (dirIdx + 1) % 4 // 顺时针旋转至下一个方向
            dir = dirs[dirIdx]
        }
        row += dir.x
        col += dir.y
     }
     
    return matrix
}


```

**复杂度分析**
- 时间复杂度：O(N^2)，N*N遍历。
- 空间复杂度：O(1) ，除了返回的矩阵以外，空间复杂度是常数




### <span id="solution2">分支方法方向方法</span>


加了边界条件
// 设定4边界
- upper_bound := 0
- left_bound := 0
- lower_bound := n-1
- right_bound := n-1
- 然后优化for num <= n*n ， num =1

### 代码


``` go（此处换成你的语言，比如js，py 等）
func generateMatrix(n int) [][]int {
     // 首先这里它们都是二位数组，所以第一步我们得构造二维数组
    matrix := make([][]int, n)
    for i:=0;i<n;i++{
        matrix[i] = make([]int, n)
        
    }

    // 设定4边界
    upper_bound := 0
    left_bound := 0
    lower_bound := n-1
    right_bound := n-1

    num := 1
    for num <= n*n {


        if upper_bound <= lower_bound {
            // 在顶部从左向右遍历
            for  j := left_bound; j <= right_bound; j++{
                 matrix[upper_bound][j] = num
                 num++
            }
            // 上边界下移
            upper_bound++
        }

        if left_bound <= right_bound {
            // 在右侧从上向下遍历
            for i := upper_bound; i <= lower_bound; i++{
                
                matrix[i][right_bound] = num
                num++
            }
            // 右边界左移
            right_bound--
        }

        if upper_bound <= lower_bound{
            // 在底部从右向左遍历
            for  j := right_bound; j >= left_bound; j-- {
              
                matrix[lower_bound][j] = num
                  num++
            }
            // 下边界上移
            lower_bound--
        }
        
        if left_bound <= right_bound{
            // 在左侧从下向上遍历
            for i := lower_bound; i >= upper_bound; i--{
                
                matrix[i][left_bound] = num
                num++
            }
            // 左边界右移
            left_bound++
        }

    }



    return matrix

    
}


```

**复杂度分析**
- 时间复杂度：O(N)，其中 N 为数组长度。
- 空间复杂度：O(1)