package ImplementQueueUsingStacks

type stack struct {
    data []int
}

func newStack() *stack{
    return &stack{
        data: make([]int, 0),
    }
}

func (self *stack) size() int {
    return len(self.data)
}

func (self *stack) push(val int) {
    self.data = append(self.data, val)
}

func (self *stack) pop() int {
    length := len(self.data)
    resData := self.data[length-1]
    self.data = self.data[:length-1]
    return resData
}

func (self *stack) top() int {
    length := len(self.data)
    return self.data[length-1]
}

func (self *stack) empty() bool {
    return len(self.data) == 0
}

type MyQueue struct {
    storeStack *stack
    aimStack   *stack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{
        storeStack: newStack(),
        aimStack: newStack(),
    }
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.storeStack.push(x)
}


func (self *MyQueue) aimQueue(isPop bool) int{
    if !self.aimStack.empty() {
        if isPop {
            return self.aimStack.pop()
        }
        return self.aimStack.top()
    }
    for self.storeStack.size() > 1 {
        self.aimStack.push(self.storeStack.pop())
    }
    topData := self.storeStack.pop()
    if !isPop {
        self.aimStack.push(topData)
    }
    return topData
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    return this.aimQueue(true)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.aimQueue(false)
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return this.storeStack.empty() && this.aimStack.empty()
}
