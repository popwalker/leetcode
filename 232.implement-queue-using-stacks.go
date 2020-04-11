/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

/*
使用两个栈来实现队列的功能
一个input栈，一个output栈
入队操作: 将数据压入input栈
出队操作: 将input栈依次取出，压入output栈，从output栈栈顶取出一个元素，即为出队操作
peek操作: 与出队操作相似
*/

// @lc code=start
type MyQueue struct {
	is *Stack
	os *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		is: newStack(),
		os: newStack(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.is.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.os.isEmpty() {
		for this.is.Len() > 0 {
			this.os.push(this.is.pop())
		}
		return this.os.pop()
	}

	return this.os.pop()

}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.os.isEmpty() {
		for this.is.Len() > 0 {
			this.os.push(this.is.pop())
		}
		return this.os.peek()
	}

	return this.os.peek()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.is.isEmpty() && this.os.isEmpty()
}

type Stack struct {
	elements []int
}

func newStack() *Stack {
	return &Stack{
		elements: []int{},
	}
}

func (s *Stack) push(x int) {
	s.elements = append(s.elements, x)
}
func (s *Stack) pop() int {
	if len(s.elements) > 0 {
		e := s.elements[len(s.elements)-1]
		s.elements = s.elements[:len(s.elements)-1]
		return e
	}
	return 0
}

func (s *Stack) isEmpty() bool {
	return len(s.elements) == 0
}
func (s *Stack) Len() int {
	return len(s.elements)
}

func (s *Stack) peek() int {
	if len(s.elements) > 0 {
		return s.elements[len(s.elements)-1]
	}
	return 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

