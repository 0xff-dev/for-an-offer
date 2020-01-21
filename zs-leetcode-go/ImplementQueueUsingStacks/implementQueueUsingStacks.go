package ImplementQueueUsingStacks

type stack struct {
    data []int
}

func newStack() *stack {
    return &stack{
        data: make([]int, 0),
    }
}

func (self *stack) push(val int) {
    self.data = append(self.data, val)
}

func (self *stack) pop() int {
    length := len(self.data)
    val := self.data[length-1]
    self.data = self.data[:length-1]
    return val
}

func (self *stack) top() int {
    return self.data[len(self.data)-1]
}

func (self *stack) empty() bool {
    return len(self.data) == 0
}

func (self *stack) size() int {
    return len(self.data)
}


type MyQueue struct {
    dataStack *stack
    aimStack *stack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{
        dataStack: newStack(),
        aimStack: newStack(),
    }
}


func (this *MyQueue) aimFunc(isPop bool) int {
    if !this.aimStack.empty() {
        if isPop {
            return this.aimStack.pop()
        }
        return this.aimStack.top()
    }
    for this.dataStack.size() > 1 {
        this.aimStack.push(this.dataStack.pop())
    }
    topData := this.dataStack.pop()
    if !isPop {
        this.aimStack.push(topData)
    }
    return topData
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
   this.dataStack.push(x) 
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    return this.aimFunc(true)
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.aimFunc(false)
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return this.dataStack.size() + this.aimStack.size() == 0
}


