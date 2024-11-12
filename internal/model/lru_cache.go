package model

import "container/list"

// LRUCache представляет LRU-кэш
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

// Constructor создает новый LRUCache
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

// Get возвращает значение по ключу и перемещает элемент в конец списка
func (c *LRUCache) Get(key int) int {
	if element, found := c.cache[key]; found {
		c.list.MoveToFront(element) // Перемещаем элемент в начало
		return element.Value.(*LRUNode).Value
	}
	return -1 // Элемент не найден
}

// Put добавляет или обновляет элемент в кэше
func (c *LRUCache) Put(key int, value int) {
	if element, found := c.cache[key]; found {
		// Обновляем существующий элемент
		element.Value.(*LRUNode).Value = value
		c.list.MoveToFront(element) // Перемещаем его в начало
	} else {
		// Добавляем новый элемент
		newNode := &LRUNode{Key: key, Value: value}
		newElement := c.list.PushFront(newNode)
		c.cache[key] = newElement

		// Если кэш переполнен, удаляем последний элемент
		if c.list.Len() > c.capacity {
			back := c.list.Back()
			if back != nil {
				c.list.Remove(back)
				delete(c.cache, back.Value.(*LRUNode).Value)
			}
		}
	}
}
