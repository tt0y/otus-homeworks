package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int               // максимальная ёмкость кэша
	queue    List              // очередь из элементов (на основе двусвязного списка)
	items    map[Key]*ListItem // словарь для быстрого доступа к элементам
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

// Set установка значения в кеш по ключу
func (c *lruCache) Set(key Key, value interface{}) bool {
	// Проверяем присуствует ли элемент в кеше
	if item, found := c.items[key]; found {
		// Обновляем значение и перемещаем элемент в начало списка
		item.Value.(*cacheItem).value = value
		c.queue.MoveToFront(item)
		return true
	}

	// Добавляем метод по ключу
	newItem := &cacheItem{key: key, value: value}
	listItem := c.queue.PushFront(newItem)
	c.items[key] = listItem

	// Проверяем, не превысили ли мы ёмкость кэша
	if c.queue.Len() > c.capacity {
		// Удаляем последний элемент из списка и словаря
		lastItem := c.queue.Back()
		c.queue.Remove(lastItem)
		delete(c.items, lastItem.Value.(*cacheItem).key)
	}

	return false
}

// Get получает значение из кэша по ключу
func (c *lruCache) Get(key Key) (interface{}, bool) {
	// Проверяем, присутствует ли элемент в кэше
	if item, found := c.items[key]; found {
		// Перемещаем элемент в начало списка
		c.queue.MoveToFront(item)
		return item.Value.(*cacheItem).value, true
	}

	// Если элемент не найден
	return nil, false
}

// Clear очищает кэш
func (c *lruCache) Clear() {
	// Очищаем очередь и словарь
	c.queue = NewList()
	c.items = make(map[Key]*ListItem)
}
