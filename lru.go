package lru

import "container/list"

type Cache interface {
	Get(key int) int    // look up key's value, if not found return KeyNotFound const
	Put(key, value int) // adds value to the cache
}

const KeyNotFound = -1

func New(cap uint) Cache {

	return &lru{
		cap:       cap,
		evictList: list.New(),
		items:     make(map[int]*list.Element),
	}
}

type entry struct {
	key, value int
}

type lru struct {
	cap       uint
	evictList *list.List
	items     map[int]*list.Element
}

func (c *lru) Get(key int) int {
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		return ent.Value.(*entry).value
	}

	return KeyNotFound
}

func (c *lru) Put(key, value int) {
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value.(*entry).value = value
		return
	}

	c.items[key] = c.evictList.PushFront(&entry{key, value})

	if c.evictList.Len() > int(c.cap) {
		if ent := c.evictList.Back(); ent != nil {

			c.evictList.Remove(ent)
			delete(c.items, ent.Value.(*entry).key)
		}
	}
}
