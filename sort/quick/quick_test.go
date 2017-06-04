package quicksort

import (
	"fmt"
	"testing"
	"recipes/sort"
)

func TestQuickSortOps(t *testing.T) {
	list := comparable.SortList{3, 2, 1, 4, 5}

	compare := func (a, b int) int {
		return a - b
	}
	swap := func (list comparable.SortList, i, j int) {
		list[i], list[j] = list[j], list[i]
	}

	QuickSort(list, compare, swap)

	fmt.Println("result:", list)
}