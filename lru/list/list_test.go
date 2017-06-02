package list

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
	manager.Append(1)
	manager.Append(2)
	manager.Append(3)
	v1, _ := manager.Get(0)
	if v1 != 1 {
		t.Error("LRUCache Get index 0 must be equal 1, but Got:", v1)	
	}

	manager.Append(1)
	v1, _ = manager.Get(0)
	v2, _ := manager.Get(2)
	if v1 != 2 {
		t.Error("LRUCache Get index 0 must be equal 2, but Got:", v1)	
	}

	if v2 != 1 {
		t.Error("LRUCache Get index 2 must be equal 1, but Got:", v2)			
	}
}