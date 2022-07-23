## LC 前缀问题
[Leetcode连接](https://leetcode.cn/problems/unique-substrings-in-wraparound-string/solution/xi-fa-dai-ni-xue-suan-fa-yi-ci-gao-ding-qian-zhui-/)

- [前缀和](#solution1)
- [连续子数组总个数](#solution2)

### <span id="solution1">前缀和</span>

题目一求前缀和
array  = [1,2,3,4,5,6] -》 pre = [1,3,6,10,15,21]
- pre[i] = pre[i-1] + array[i]


### Golang 代码

``` go

//
//  countSubArray
//  @Description:连续子数组总个数
//  @param nums
//  @return int
//
func countSubArray(nums []int) int {
	ans := 0
	pre := 0
	for _ = range nums {
		pre += 1
		ans += pre
	}
	return ans
}



```

**复杂度分析**
- 时间复杂度：O(N)，其中 N 为数组长度。
- 空间复杂度：O(1)



### <span id="solution2">连续子数组总个数</span>

### 思路

array = [1,3,4]

连续子数组 ： [1], [3], [4], [1,3], [3,4] , [1,3,4]

- 4 结尾的连续子数组 [4] ， [3,4] , [1,3,4]  ； 4 的 索引是2
- 3 结尾的连续子数组 [3]、[1,3] ； 3的索引是1
- 1 结尾的连续子数组 [1]； 1的索引是1 是 0

我们可以观察到连续子 【索引+1】

### 代码


``` go（此处换成你的语言，比如js，py 等）
（此处撰写代码）


```

**复杂度分析**
- 时间复杂度：O(N)，其中 N 为数组长度。
- 空间复杂度：O(1)