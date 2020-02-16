package main

import (
	"fmt"
)

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Get(1)
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))

}

// Map+双向链表 O(1)读 O(1)写
type LRUCache struct {
	NodeMap map[int]*DoubleLinkeNode
	// 虚头虚尾
	Head, Tail     *DoubleLinkeNode
	Capacity, Size int
}

type DoubleLinkeNode struct {
	Prev, Next *DoubleLinkeNode
	Key, Value int

	//API
	//1.添加结点到头
	//2.删除某一个结点
	//3.2+1 将存在结点重新放到头
	//4.2并返回该结点,用于map的删除
}

func Constructor(capacity int) LRUCache {
	dummyHead := new(DoubleLinkeNode)
	dummyTail := new(DoubleLinkeNode)
	dummyHead.Next = dummyTail
	dummyTail.Prev = dummyHead

	return LRUCache{Size: 0, Capacity: capacity, Head: dummyHead, Tail: dummyTail, NodeMap: make(map[int]*DoubleLinkeNode)}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.NodeMap[key]; ok {
		this.moveToHead(node)
		return node.Value
	} else {
		return -1
	}

}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.NodeMap[key]; ok {
		// 更新
		node.Value = value
		this.moveToHead(node)
	} else {
		node := new(DoubleLinkeNode)
		node.Key = key
		node.Value = value
		this.NodeMap[key] = node
		this.addNode(node)
		this.Size++
		if this.Size > this.Capacity {
			tail := this.popTail()
			delete(this.NodeMap, tail.Key)
			this.Size--
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// 更新头
func (this *LRUCache) addNode(node *DoubleLinkeNode) {
	// 结点指向头部的两个结点
	node.Prev = this.Head
	node.Next = this.Head.Next

	// 头部的两个结点断开,连到node
	this.Head.Next.Prev = node
	this.Head.Next = node
}

// 删除该结点
func (this *LRUCache) removeNode(node *DoubleLinkeNode) {
	// 获取结点的前后结点
	prev := node.Prev
	next := node.Next

	// 前后结点相连
	prev.Next = next
	next.Prev = prev
}

// 删结点,重放到头
func (this *LRUCache) moveToHead(node *DoubleLinkeNode) {
	this.removeNode(node)
	this.addNode(node)
}

// 抛出尾, 使map删除该结点
func (this *LRUCache) popTail() *DoubleLinkeNode {
	res := this.Tail.Prev
	this.removeNode(res)
	return res
}
