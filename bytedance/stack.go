package main

import (
	"fmt"
)

type Node struct {
	pre  *Node
	next *Node
	val  int
}

type NodeContainer struct {
	pre  *NodeContainer
	next *NodeContainer
	val  *Node
}

type MinStack struct {
	tail    *NodeContainer
	minTail *Node
}

/** initialize your data structure here. */
func Constructor() MinStack {

	s := MinStack{}
	s.tail = nil
	s.minTail = nil

	return s
}

func (this *MinStack) Push(x int) {

	n := &Node{val: x}
	nc := &NodeContainer{val: n}
	if this.tail == nil {
		this.tail = nc
		this.minTail = n
	} else {
		nc.pre = this.tail
		this.tail.next = nc
		this.tail = nc

		i := this.minTail
		for ; i != nil && i.pre != nil; i = i.pre {

			if i.val > x {
				n.pre = i
				n.next = i.next
				if i.next != nil {
					i.next.pre = n
				} else {
					this.minTail = n
				}
				i.next = n

				break
			}
		}

		// 1、空
		// 2、1 Node 最后一个 或者第一个

		if i == nil {
			// n.pre = i
			// n.next = i.next
			// i.next = n
			this.minTail = n
		} else if i.pre == nil {
			if i.val < x {
				n.next = i
				i.pre = n
			} else {
				if i.next != nil {
					n.pre = i
					n.next = i.next
					i.next.pre = n
					i.next = n
				} else {
					n.pre = i
					n.next = i.next
					i.next = n
					this.minTail = n
				}
			}
		}
	}
}

func (this *MinStack) Pop() {

	if this.tail == nil {
		return
	}

	if this.tail.pre == nil {
		this.tail = nil
		this.minTail = nil
	} else {
		n := this.tail.val

		this.tail = this.tail.pre
		this.tail.next = nil

		if n == this.minTail {
			this.minTail = n.pre
			if n.pre != nil {
				this.minTail.next = nil
			}
		} else {
			n.next.pre = n.pre
			if n.pre != nil {
				n.pre.next = n.next
			}
		}
	}
}

func (this *MinStack) Top() int {

	return this.tail.val.val
}

func (this *MinStack) GetMin() int {

	return this.minTail.val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

	obj := Constructor()
	obj.Push(1)
	obj.Push(3)
	obj.Push(0)
	obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.GetMin()

	fmt.Println(param_3)
	fmt.Println(param_4)
}
