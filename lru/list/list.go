package lrulist

import (
	"errors"
)

type LRUCache struct {
	value interface {}
}

type LRUCacheManager struct {
	MaxSize int
	items []LRUCache
}

func NewLRUCacheManager(size int)  *LRUCacheManager {
	return &LRUCacheManager{size, make([]LRUCache, 0, size)}
}

func (manager *LRUCacheManager) FindIndex(v interface {}) int {
	for index, value := range manager.items {
		if value == v {
			return index
		}
	}
	return -1
}

func (manager *LRUCacheManager) Push(v interface {}) {
	size := manager.MaxSize
	index := manager.FindIndex(v)

	if index != -1 {
		item := manager.items[index]
		manager.items = append(manager.items[0:index], manager.items[index:]...)
		manager.items = append(manager.items, item)
	} else {
		if size >= len(manager.items) {
			manager.items = manager.items[1:size]
		}
		item := LRUCache{v}
		manager.items = append(manager.items, item)
	}
}

func (manager *LRUCacheManager) Get(index int) (interface {}, error) {
	if index >= len(manager.items) || index < 0 {
		return nil, errors.New("Not Found Item by index")
	} else {
		return manager.items[index].value, nil
	}
}

func (manager *LRUCacheManager) Len() int {
	return len(manager.items)
}

func (manager *LRUCacheManager) Rest() int {
	length := manager.Len()
	return manager.MaxSize - length
}

func (manager *LRUCacheManager) Cap() int {
	return manager.MaxSize
}

func (manager *LRUCacheManager) List() []interface{} {
	var list = make([] interface {}, len(manager.items))

	for value, _ := range manager.items {
		list = append(list, value)
	}
	return list
}