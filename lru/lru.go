package main

import (
	"fmt"
	"lru/list"
)

func main () {
	const maxSize = 3
	manager := list.NewLRUCacheManager(maxSize)
	manager.Append(3)
	manager.Append(2)
	manager.Append(1)
	values := manager.List()
	fmt.Println("values:", values)
}