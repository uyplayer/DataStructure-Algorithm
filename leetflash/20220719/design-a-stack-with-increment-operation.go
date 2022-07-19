/*
 * @Author: uyplayer
 * @Date: 2022/7/19 14:58
 * @Email: uyplayer@qq.com
 * @File: design-a-stack-with-increment-operation.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / leetflash/20220719
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package _0220719

type CustomStack struct {
	maxSize int
	stack   []int
}

func Constructor(maxSize int) CustomStack {

	return CustomStack{
		maxSize: maxSize,
		stack:   make([]int, 0),
	}

}

func (this *CustomStack) Push(x int) {

	if len(this.stack) < this.maxSize {
		this.stack = append(this.stack, x)
	}

}

func (this *CustomStack) Pop() int {

	if len(this.stack) == 0 {
		return -1
	}
	tmp := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return tmp

}

func (this *CustomStack) Increment(k int, val int) {

	if len(this.stack) <= k {
		for i := 0; i < len(this.stack); i++ {
			this.stack[i] += val
		}
	} else {
		for i := 0; i < k; i++ {
			this.stack[i] += val
		}
	}

}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
