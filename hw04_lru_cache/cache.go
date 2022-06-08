package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	item, isKeyInMap := lc.items[key]

	if isKeyInMap {
		item.Value = cacheItem{key, value}
		lc.queue.MoveToFront(item)
		lc.items[key] = item

		return true
	}

	if len(lc.items) == lc.capacity {
		lc.Clear()
	}

	newCacheItem := cacheItem{key, value}
	newListItem := lc.queue.PushFront(newCacheItem)
	lc.items[key] = newListItem

	return false
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	if item, ok := lc.items[key]; ok {
		lc.queue.MoveToFront(item)

		result := item.Value.(cacheItem)

		return result.value, true
	}

	return nil, false
}

func (lc *lruCache) Clear() {
	if item := lc.queue.Back(); item != nil {
		lc.queue.Remove(item)
		delete(lc.items, item.Value.(cacheItem).key)
	}
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
