package main

func main() {

}

type Node struct {
	key, val   int
	prev, next *Node
}

type DoubleLinkedList struct {
	head, tail *Node
	size       int
}

type LRUCache struct {
	capacity int
	knMap    map[int]*Node
	cache    *DoubleLinkedList
}

func initNode(key, val int) *Node {
	return &Node{
		key: key,
		val: val,
	}
}

func initDoubleLinkedList() *DoubleLinkedList {
	dll := &DoubleLinkedList{
		head: initNode(0, 0),
		tail: initNode(0, 0),
		size: 0,
	}
	dll.head.next = dll.tail
	dll.tail.prev = dll.head
	return dll
}

// 在链表尾部添加 x 节点，时间 O(1)
func (this *DoubleLinkedList) addLast(x *Node) {
	x.prev = this.tail.prev
	x.next = this.tail

	this.tail.prev.next = x
	this.tail.prev = x

	this.size++
}

// 删除链表中的 x 节点，时间 O(1)
func (this *DoubleLinkedList) remove(x *Node) {
	x.next.prev = x.prev
	x.prev.next = x.next
	x.prev = nil
	x.next = nil
	this.size--
}

// 删除链表中的第一个节点，并返回该节点，时间 O(1)
func (this *DoubleLinkedList) removeFirst() *Node {
	if this.head.next == this.tail {
		return nil
	}
	first := this.head.next
	this.remove(first)
	return first
}

func Constructor(capacity int) LRUCache {
	lrucache := LRUCache{
		capacity: capacity,
		knMap:    map[int]*Node{},
		cache:    initDoubleLinkedList(),
	}
	return lrucache
}

// 将某个 key 调整为最近使用的元素
func (this *LRUCache) makeRecently(key int) {
	node := this.knMap[key]
	this.cache.remove(node)
	this.cache.addLast(node)
}

// 添加最近使用的元素
func (this *LRUCache) addRecently(key int, val int) {
	node := initNode(key, val)
	this.cache.addLast(node)
	this.knMap[key] = node
}

// 删除某一个 key 及对应元素
func (this *LRUCache) deleteKey(key int) {
	node := this.knMap[key]
	this.cache.remove(node)
	delete(this.knMap, key)
}

// 删除最近最少使用的元素
func (this *LRUCache) deleteLRU() {
	deletenode := this.cache.removeFirst()
	delete(this.knMap, deletenode.key)
}

func (this *LRUCache) Get(key int) int {
	node, exit := this.knMap[key]
	if !exit {
		return -1
	}
	this.makeRecently(key)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	_, exit := this.knMap[key]
	if exit {
		this.deleteKey(key)
		this.addRecently(key, value)
	} else {
		if this.capacity == this.cache.size {
			this.deleteLRU()
		}
		this.addRecently(key, value)
	}
}
