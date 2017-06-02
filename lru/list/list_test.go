package lrulist

import(
	"testing"
)
func TestLRUCacheOps (t *testing.T) {
	const MAX_SIZE = 3
	manager := NewLRUCacheManager(MAX_SIZE)
	dCap := manager.Cap()
	if dCap != 3 {
		t.Error("LRUCache Cap must be equal 3, but Got:", dCap)
	}
}