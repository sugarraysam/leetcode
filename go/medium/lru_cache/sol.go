package lru_cache

type LRUCache struct {
	dll *DLL
	hm  map[int]*Node
	cap int
	len int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		dll: &DLL{},
		hm:  make(map[int]*Node),
		cap: capacity,
		len: 0,
	}
}

func (this *LRUCache) Get(key int) int {
	// value exists
	if n, ok := this.hm[key]; ok {
		this.dll.appendExisting(n)
		return n.Val
	}
	// value does not exist
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	// value exists
	if n, ok := this.hm[key]; ok {
		n.Val = value
		this.dll.appendExisting(n)
	} else {
		// new value
		n := &Node{Val: value}
		this.dll.appendNew(n)
		this.hm[key] = n
		this.len++

		// need to evict LRU
		if this.len > this.cap {
			delKey := this.dll.removeHead()
			delete(this.hm, delKey)
			this.len--
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type DLL struct {
	Head *Node // LRU
	Tail *Node
}

// evict LRU - expect Head to exist
func (l *DLL) removeHead() int {
	res := l.Head.Val

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
	} else {
		l.Head = l.Head.Next
		l.Head.Prev = nil
	}
	return res
}

// put operation on a new value
func (l *DLL) appendNew(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		// at least 1 node
		l.Tail.Next = n
		n.Prev = l.Tail
		l.Tail = n
	}
}

// put operation on existing value
// get operator on existing value
func (l *DLL) appendExisting(n *Node) {
	// at least a node in list
	if l.Head == l.Tail {
		return
	}

	// node is not tail, list >= 2
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}

	// 2 - add node at the end
	l.appendNew(n)
}
