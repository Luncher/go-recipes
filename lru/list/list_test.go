package lrulist

import(
	"testing"
)
func TestLRUCacheOps (t *testing.T) {
	const maxSize = 3
	manager := NewLRUCacheManager(maxSize)
	dCap := manager.Cap()
	if dCap != 3 {
		t.Error("LRUCache Cap must be equal 3, but Got:", dCap)
	}
}