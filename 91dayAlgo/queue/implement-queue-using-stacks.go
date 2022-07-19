/*
 * @Author: uyplayer
 * @Date: 2022/7/19 22:11
 * @Email: uyplayer@qq.com
 * @File: implement-queue-using-stacks.go
 * @Software: GoLand
 * @Dir: DataStructure-Algorithm / 91dayAlgo/queue
 * @Project_Name: DataStructure-Algorithm
 * @Description:
 */

package queue

//
//  MyQueue
//  @Description: 队列结构体
//
type MyQueue struct {
	inStack  []int
	outStack []int
}

//
//  Constructor
//  @Description: 创建
//  @return MyQueue
//
func Constructor() MyQueue {

	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}

}

//
//  Push
//  @Description: 入队列
//  @receiver this
//  @param x
//
func (this *MyQueue) Push(x int) {

	this.inStack = append(this.inStack, x)

}

//
//  Pop
//  @Description: 出队列
//  @receiver this
//  @return int
//
func (this *MyQueue) Pop() int {

	// outStack 长度等于0
	if len(this.outStack) == 0 {
		for len(this.inStack) > 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}

	}
	// 获取最后一个元素
	elem := this.outStack[len(this.outStack)-1]
	// 更新
	this.outStack = this.outStack[:len(this.outStack)-1]
	return elem

}

//
//  Peek
//  @Description: 最后一个元素
//  @receiver this
//  @return int
//
func (this *MyQueue) Peek() int {

	if len(this.outStack) == 0 {
		for len(this.inStack) > 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}
	return this.outStack[len(this.outStack)-1]

}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
