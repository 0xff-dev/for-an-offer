package ImplementStackUsingQueues

type queue struct {
    data []int
}

func newQueue() *queue {
    return &queue{
        data: make([]int, 0),
    }
}
func (self *queue) push(val int) {
    self.data = append(self.data, val)
}

func (self *queue) front() int{
    return self.data[0]
}

func (self *queue) pop() int {
    tmpData := self.data[0]
    self.data = self.data[1:]
    return tmpData
}

func (self *queue) size() int {
    return len(self.data)
}

type MyStack struct {
    addQueue *queue
    aimQueue *queue
}

func Constructor() MyStack{
    return MyStack{
        addQueue: newQueue(),
        aimQueue: newQueue(),
    }
}

func (this *MyStack) Push(x int) {
    this.addQueue.push(x)
}

func (this *MyStack) aimFunc(isPop bool) int {
    if this.addQueue.size() > 0 {
        for this.addQueue.size() > 1 {
            this.aimQueue.push(this.addQueue.pop())
        }
        if !isPop {
            this.aimQueue.push(this.addQueue.front())
        }
        return this.addQueue.pop()
    } else {
        for this.aimQueue.size() > 1 {
            this.addQueue.push(this.aimQueue.pop())
        }
        if !isPop {
            this.addQueue.push(this.aimQueue.front())
        }
        return this.aimQueue.pop()
    }
}

func (this *MyStack) Pop() int {
    return this.aimFunc(true)
}

func (this *MyStack) Top() int {
    return this.aimFunc(false)
}

func (this *MyStack) Empty() bool {
    return this.addQueue.size() + this.aimQueue.size() == 0
}

