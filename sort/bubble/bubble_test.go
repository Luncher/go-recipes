package bubblesort

import (
	"fmt"
	"testing"
	"recipes/sort"
)

func TestBubbleSortOps(t *testing.T) {
	list := comparable.SortList{2, 0, 1, 5, 4, 6}

	compare := func (a, b int) int {
		return a - b
	}

	swap := func (list comparable.SortList, i, j int) {
		list[i], list[j] = list[j], list[i]
	}

	BubbleSort(list, compare, swap)

	fmt.Println("result:", list)
}