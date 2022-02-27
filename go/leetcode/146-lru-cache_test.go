package leetcode

import (
	"fmt"
	"testing"
)

func Test146(*testing.T) {
	/**
	 * Your LRUCache object will be instantiated and called as such:
	 */
	var (
		capacity int
		val      int
	)

	capacity = 3

	obj := Constructor(capacity)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(1, 3)
	obj.Put(4, 4)
	//obj.Put(5, 5)
	val = obj.Get(1)
	fmt.Println(val)
}

/**
 * 实现方式
 *
 * hashmap + Doubly linked list
 * (dummy head, tail)
 *
 */

type LRUNode struct {
	prev     *LRUNode
	next     *LRUNode
	key, val int
}

type LRUCache struct {
	hash           map[int]*LRUNode
	head, tail     *LRUNode
	capacity, used int
}

func Constructor(capacity int) LRUCache {
	tail := &LRUNode{key: -2, val: -2}
	head := &LRUNode{key: -3, val: -3}
	tail.next = head
	head.prev = tail

	return LRUCache{
		hash:     make(map[int]*LRUNode, capacity),
		tail:     tail,
		head:     head,
		capacity: capacity,
		used:     0,
	}
}

/**
 * 1. map key to *node
 * 2. move *node to head
 * 3. relink old link
 */
func (this *LRUCache) Get(key int) int {
	var (
		node *LRUNode
		ok   bool
	)

	// not exists return
	if node, ok = this.hash[key]; !ok {
		return -1
	}

	this.move2head(node)
	return node.val
}

/**
 * 1. new *node
 * 2. node 2 head
 * 3. if len > cap
 * 3.1 tail to tail prev
 * 3.2 release old node
 */
func (this *LRUCache) Put(key int, value int) {
	var (
		node *LRUNode
		ok   bool
	)
	if node, ok = this.hash[key]; !ok {
		node = &LRUNode{key: key}
	}
	node.val = value
	this.hash[key] = node

	if !ok {
		this.used++
	}

	this.move2head(node)
	if this.used <= this.capacity {
		return
	}

	// 删除链表最后节点
	tail := this.tail.next

	last2 := this.tail.next.next
	this.tail.next = last2
	last2.prev = this.tail

	delete(this.hash, tail.key)
	this.used--
}

/**
 * 把一个节点从原来的位置移动到链表头指针
 */
func (this *LRUCache) move2head(node *LRUNode) {
	// 新创建的节点不需要从末端移动
	if node.prev != nil && node.next != nil {
		prev := node.prev
		next := node.next
		prev.next = next
		next.prev = prev
	}

	head := this.head.prev
	head.next = node
	node.prev = head
	this.head.prev = node
	node.next = this.head
}
