package lru_cache

type LRUCache struct {
	root, tail *node
	cache      map[int]*node
	size, cap  int
}

type node struct {
	prev, next *node
	key, val   int
}

func Constructor(capacity int) LRUCache {
	r, t := &node{}, &node{}
	r.next, t.prev = t, r
	return LRUCache{
		root:  r,
		tail:  t,
		cache: make(map[int]*node),
		cap:   capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.del(key)
	this.add(key, n.val)
	return n.val
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.cache[key]
	if ok {
		this.cache[key].val = value
		this.del(key)
	} else {
		if this.size == this.cap {
			this.del(this.tail.prev.key)
		}
	}
	this.add(key, value)
}

func (this *LRUCache) del(key int) {
	n, ok := this.cache[key]
	if !ok {
		return
	}
	delete(this.cache, key)
	pr, ne := n.prev, n.next
	pr.next, ne.prev = ne, pr
	this.size--
}

func (this *LRUCache) add(key, val int) {
	n := &node{
		prev: this.root,
		next: this.root.next,
		key:  key,
		val:  val,
	}
	this.cache[key] = n
	this.root.next.prev = n
	this.root.next = n
	this.size++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
