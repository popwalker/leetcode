/*
 * @lc app=leetcode id=225 lang=golang
 *
 * [225] Implement Stack using Queues
 */

// @lc code=start
type MyStack struct {
	enque []int
	deque []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		enque: []int{},
		deque: []int{},
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.enque = append(this.enque, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	length := len(this.enque)
	for i := 0; i < length-1; i++ {
		this.deque = append(this.deque, this.enque[0])
		this.enque = this.enque[1:]
	}

	topEle := this.enque[0]
	this.enque = this.deque
	this.deque = nil

	return topEle
}

/** Get the top element. */
func (this *MyStack) Top() int {
	topEle := this.Pop()
	this.Push(topEle)
	return topEle
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.enque) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

