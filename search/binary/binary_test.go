package binarysearch

import (
	"testing"
	"fmt"
)

func TestBinarySearchOps(t *testing.T) {
	arr := IntSlice{1, 2, 3, 4, 5}
	index := arr.Search(5)
	fmt.Println("result:", index)
}