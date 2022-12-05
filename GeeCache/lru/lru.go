package lru

import "container/list"

// LRU cache 结构体
type Cache struct {
	maxBytes  int64                         // 允许使用的最大内存
	nbytes    int64                         // 当前已使用的内存
	ll        *list.List                    // cache链表
	cache     map[string]*list.Element      // 查找键值对的字典
	OnEvicted func(key string, value Value) // 某条记录被移除时的回调函数
}

// 双向链表节点的数据类型
// 主要目的是为了删除节点后能从字典中删除该键值对
type entry struct {
	key   string
	value Value
}

// 值的类型可以是任意的，定义一个空接口，实现Len()方法返回值的占用空间大小

// Len the number of cache entries
func (c *Cache) Len() int {
	return c.ll.Len()
}

type Value interface {
	Len() int // 包含一个方法返回值占用的内存大小
}

// 工厂模式，返回实例化的cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// 查找功能，在字典中进行查找，然后移动到队尾（Front）
func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// LRU删除策略：从队首（Back）拿到节点，然后将其删除
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len()) // 更新当前已经使用的内存
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// 新增节点/修改节点
func (c *Cache) Add(key string, value Value) {
	// 如果在链表中找到则将其更新，同时更新占用的空间大小等，并移动到队列尾端
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else { // 如果找不到则直接插入
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	// 如果占用空间超过了链表的最大空间，则删除掉队首的节点
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}
