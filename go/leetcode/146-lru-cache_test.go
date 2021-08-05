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

	capacity = 2

	obj := Constructor(capacity)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3)
	val = obj.Get(1)
	fmt.Println(val)
}

type LRUCache struct {
	indexMap       map[int]*LruNode
	head           *LruNode
	tail           *LruNode
	capacity, size int
}

type LruNode struct {
	K, V int
	Prev *LruNode
	Next *LruNode
}

// 创造一个tail->head 的双向链表(head 为最近使用)
func Constructor(capacity int) LRUCache {
	head := &LruNode{}
	tail := &LruNode{}
	head.Prev = tail
	tail.Next = head

	return LRUCache{
		head:     head,
		tail:     tail,
		capacity: capacity,
		indexMap: make(map[int]*LruNode),
	}
}

func (this *LRUCache) Get(key int) int {
	var (
		r    int = -1
		node *LruNode
		ok   bool
	)

	// 不存在
	if node, ok = this.indexMap[key]; !ok {
		return -1
	}

	// 存在
	// 1. 找到位置(hash)
	// 2. 移动到链表头部节点
	r = node.V

	this.PutToHead(node)
	return r
}

func (this *LRUCache) Put(key int, value int) {
	// 判断是否存在
	var (
		node *LruNode
		ok   bool
	)

	if node, ok = this.indexMap[key]; ok {
		// 存在更新value
		node.V = value
	} else {
		// 不存在 新建一个节点
		node = &LruNode{
			K:    key,
			V:    value,
			Prev: nil,
			Next: nil,
		}
		this.indexMap[key] = node
		this.size++
	}

	this.PutToHead(node)

	// 容量是否超出容量
	if this.size > this.capacity {
		// 删除链表最后元素
		last := this.tail.Next
		// tail cur next
		last.Next.Prev = this.tail
		this.tail.Next = last.Next
		delete(this.indexMap, last.K)
		last = nil
		this.size--
	}
}

func (this *LRUCache) PutToHead(node *LruNode) {
	// 把当前节点从原有链表位置剔除
	if node.Next != nil && node.Prev != nil {
		next := node.Next
		prev := node.Prev
		prev.Next = next
		next.Prev = prev
	}

	// 节点放到链表头部(最近使用)
	//tail a b c head
	head := this.head.Prev
	head.Next = node
	node.Prev = head
	node.Next = this.head
	this.head = node
}
