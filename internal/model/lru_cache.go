package model

import "container/list"

type CacheItem struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int                   // Максимальная емкость кэша
	cache    map[int]*list.Element // Хеш-таблица для быстрого доступа
	order    *list.List            // Связанный список для отслеживания порядка использования
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		order:    list.New(),
	}
}

func (cache *LRUCache) Get(key int) int {
	if node, found := cache.cache[key]; found {
		cache.order.MoveToFront(node)
		return node.Value.(*CacheItem).value
	}
	// Если элемента нет, возвращаем -1
	return -1
}

func (cache *LRUCache) Put(key, value int) {
	if node, found := cache.cache[key]; found {
		node.Value.(*CacheItem).value = value
		cache.order.MoveToFront(node)
		return
	}

	// Если элемент не существует, добавляем его
	if cache.order.Len() == cache.capacity {
		// Удаляем наименее недавно использованный элемент (последний элемент в списке)
		leastUsed := cache.order.Back()
		cache.order.Remove(leastUsed)
		delete(cache.cache, leastUsed.Value.(*CacheItem).key)
	}

	// Добавляем новый элемент в начало списка
	newItem := &CacheItem{key: key, value: value}
	node := cache.order.PushFront(newItem)
	cache.cache[key] = node
}
