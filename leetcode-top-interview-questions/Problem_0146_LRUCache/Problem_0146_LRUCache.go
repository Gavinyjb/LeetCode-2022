package Problem_0146_LRUCache

import "container/list"

type LRUCache struct {
	maxCap  int
	usedCap int
	ll      *list.List
	cache   map[int]*list.Element
}
type entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		maxCap:  capacity,
		usedCap: 0,
		ll:      list.New(),
		cache:   make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.cache[key]; ok {
		this.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.cache[key]; ok {
		this.ll.MoveToFront(ele)
		kv := this.cache[key].Value.(*entry)
		kv.value = value
	} else {
		ele := this.ll.PushFront(&entry{
			key:   key,
			value: value,
		})
		this.cache[key] = ele
		this.usedCap += 1
	}
	for this.maxCap < this.usedCap {
		this.Remove()
	}
}
func (this *LRUCache) Remove() {
	ele := this.ll.Back()
	if ele != nil {
		this.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(this.cache, kv.key)
		this.usedCap -= 1
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
