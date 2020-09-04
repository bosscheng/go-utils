package utils

import (
	"container/list"
	"sync"
)

// RWMutex是读写互斥锁。该锁可以被同时多个读取者持有或唯一个写入者持有。RWMutex可以创建为其他结构体的字段；零值为解锁状态。
// RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators
type Key interface {
}

type entry struct {
	key   Key
	value *ByteBuf
}

// Cache is an LRU cache. It is not safe for concurrent access.
type Cache struct {
	sync.RWMutex

	MaxBytes uint64

	current uint64

	OnEvicted func(key Key, value *ByteBuf)
	// 双向链表
	ll *list.List
	// Element类型代表是双向链表的一个元素。
	cache map[interface{}]*list.Element
}

// Pool a mem pool interface
type Pool interface {
	Alloc(int) []byte
	Free([]byte)
}


func NewLRUCache(maxBytes uint64, evictedFunc func(key Key, value *ByteBuf)) *Cache {
	return &Cache{
		MaxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[interface{}]*list.Element),
		OnEvicted: evictedFunc,
	}
}

func (c *Cache) Add(key Key, value *ByteBuf) {
	c.Lock()
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)

		entry := ee.Value.(*entry)
		c.current -= uint64(value.Readable())
		c.current += uint64(value.Readable())
		entry.value = value
		c.Unlock()
		return
	}

}

func (c *Cache) Get(key Key) (value *ByteBuf, ok bool) {
	c.RLock()

	if c.cache == nil {
		c.RUnlock()
		return
	}

	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		c.RUnlock()
		return ele.Value.(*entry).value, true
	}

	c.RUnlock()
	return
}

func (c *Cache) Remove(key Key) {
	c.Lock()

	if c.cache == nil {
		c.Unlock()
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
	}

	c.Unlock()
}

func (c *Cache) removeOldest() {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cache, kv.key)
	c.current -= uint64(kv.value.Readable())
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

func (c *Cache) Len() int {
	c.RLock()
	if c.cache == nil {
		c.RUnlock()
		return 0
	}
	value := c.ll.Len()
	c.RUnlock()
	return value
}

func (c *Cache) Clear() {
	c.Lock()
	if c.OnEvicted != nil {
		for _, e := range c.cache {
			kv := e.Value.(*entry)
			c.OnEvicted(kv.key, kv.value)
		}
	}
	c.ll = nil
	c.cache = nil
	c.current = 0
	c.Unlock()
}
