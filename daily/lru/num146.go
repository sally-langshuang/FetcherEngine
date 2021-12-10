package lru

import (
	"fmt"
	"strings"
)

type Node struct {
	K, V int
	Pre, Next *Node
}

type LRUCache struct {
	Capacity int
	HeadSentinel, TailSentinel *Node
	FindMap map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next, tail.Pre = tail, head
	return LRUCache{
		Capacity: capacity,
		HeadSentinel: head,
		TailSentinel: tail,
		FindMap: map[int]*Node{},
	}
}

func (this *LRUCache) String() string {
	b := &strings.Builder{}
	for i:= this.TailSentinel.Pre; i != this.HeadSentinel; i = i.Pre {
		fmt.Fprintf(b, "(%d,%d)--", i.K, i.V)
	}
	return b.String()
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.FindMap[key]
	if !ok {
		return -1
	}
	this.update(key, node.V)
	return node.V
}

func (this *LRUCache) update(key int, value int) {
	if this.Capacity <= 0 {
		return
	}
	var delNode, addNode *Node = nil, &Node{K:key, V:value}
	node, isExist := this.FindMap[key]
	if isExist {
		delNode, addNode, node.V = node, node, value
	} else if len(this.FindMap) == this.Capacity{
		delNode, addNode = this.HeadSentinel.Next, &Node{K:key, V:value}
	}
	if delNode != nil{
		delNode.Pre.Next = delNode.Next
		delNode.Next.Pre = delNode.Pre
		delete(this.FindMap, delNode.K)
	}
	addNode.Pre = this.TailSentinel.Pre
	addNode.Pre.Next = addNode
	addNode.Next = this.TailSentinel
	this.TailSentinel.Pre = addNode
	this.TailSentinel.Pre = addNode
	this.FindMap[key] = addNode
}

func (this *LRUCache) Put(key int, value int)  {
	this.update(key, value)
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */