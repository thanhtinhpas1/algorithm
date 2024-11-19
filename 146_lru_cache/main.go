package main

import "fmt"

type Node struct {
	key, val   int
	prev, next *Node
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{-1, -1, nil, nil}
	tail := &Node{-1, -1, nil, nil}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) add(node *Node) {
	tmp := this.head.next
	tmp.prev = node
	node.next = tmp
	node.prev = this.head
	this.head.next = node
}

func (this *LRUCache) remove(node *Node) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) Get(key int) int {
	if node, exists := this.cache[key]; exists {
		this.remove(node)
		this.add(node)
		return node.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		this.remove(node)
		node.val = value
		this.add(node)
	} else {
		if len(this.cache) == this.capacity {
			lru := this.tail.prev
			this.remove(lru)
			delete(this.cache, lru.key)
		}

		newNode := &Node{key: key, val: value}
		this.add(newNode)
		this.cache[key] = newNode
	}
}

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // -1
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // -1
	fmt.Println(cache.Get(3)) // 3
	fmt.Println(cache.Get(4)) // 4
}
