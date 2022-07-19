## LC 剑指 Offer 53 - II. 0～n-1中缺失的数字
[Leetcode连接](https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/)

- [索引对比方法](#solution1)
- [额外用一个字典](#solution1)


### <span id="solution1">索引对比方法</span>

一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。</br>

0开始到size(nums)- 1 ; 这不是数组的索引吗？ 
- key != value ；按照题义，索引等于元素值


### Go代码

``` go（此处换成你的语言，比如js，py 等）
func missingNumber(nums []int) int {


    for key,value := range nums {

        if key != value {
            return key 
        }

    }

    return len(nums)

}


```

**复杂度分析**
- 时间复杂度：O(N)，其中 N 为数组长度。每个元素遍历一遍
- 空间复杂度：O(N)，其中 N 为数组长度 ,key,value 每次存储变量，O(N)+O(N)=O(2N)；简化为O(N)

### <span id="solution2">额外用一个字典</span>
用额外字典来保存每一个元素,然后从i开始遍历，如果不再map 里面，即返回；

### Go代码


``` go（此处换成你的语言，比如js，py 等）

func missingNumber(nums []int) int {


   elems := make(map[int]bool)

    for _,value := range nums {
        elems[value]=true 
    }

    for i:=0;;i++ {

        if _,ok:=elems[i];ok == false {
            return i
        }
        
    }
}


```


**复杂度分析**
- 时间复杂度：时间复杂度：O(n)，其中 nn 是数组 nums 的长度加 1。遍历数组 nums 将元素加入哈希集合的时间复杂度是 O(n)，遍历从 0到 n - 1 的每个整数并判断是否在哈希集合中的时间复杂度也是 O(n)。
- 空间复杂度：O(n)，其中 n是数组nums 的长度加 1。哈希集合中需要存储 n - 1个整数。
