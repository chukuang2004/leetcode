package main

import (
	"container/list"
	"fmt"
	"sync"
)

type Node struct {
	key   int
	value int
}

type LRUCache struct {
	index map[int]*list.Element
	data  *list.List
	size  int
	lock  sync.Mutex
}

func Constructor(capacity int) LRUCache {

	c := LRUCache{size: capacity}
	c.index = make(map[int]*list.Element)
	c.data = list.New()

	return c
}

func (this *LRUCache) Get(key int) int {

	this.lock.Lock()
	defer this.lock.Unlock()

	e, ok := this.index[key]
	if !ok {
		return -1
	}

	this.data.Remove(e)
	n, _ := e.Value.(*Node)

	e = this.data.PushFront(n)
	this.index[key] = e

	return n.value
}

func (this *LRUCache) Put(key int, value int) {

	this.lock.Lock()
	e, ok := this.index[key]
	if ok {
		n, _ := e.Value.(*Node)
		n.value = value
		this.data.Remove(e)
		e = this.data.PushFront(n)
		this.index[key] = e
	} else {
		n := &Node{key: key, value: value}
		size := len(this.index)
		if size >= this.size {
			e := this.data.Back()
			this.data.Remove(e)
			n, _ := e.Value.(*Node)
			delete(this.index, n.key)
		}

		e := this.data.PushFront(n)
		this.index[key] = e
	}
	this.lock.Unlock()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {

	cache := Constructor(2 /* 缓存容量 */)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 返回  1
	cache.Put(3, 3)           // 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)           // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回  3
	fmt.Println(cache.Get(4)) // 返回  4
	cache.Put(5, 5)
	fmt.Println(cache.Get(4)) // 返回  4
	fmt.Println(cache.Get(3)) // 返回  -1
	fmt.Println(cache.Get(5)) // 返回  5
	cache.Put(6, 6)
	cache.Put(7, 7)
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(6, 8)
	fmt.Println(cache.Get(5)) // 返回  5
	fmt.Println(cache.Get(7)) // 返回  7
	fmt.Println(cache.Get(6)) // 返回  8
}
