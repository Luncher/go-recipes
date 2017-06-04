package insertsort

import (
	"fmt"	
	"testing"
	"recipes/sort"
)

func TestInsertOps(t *testing.T) {
	list := comparable.SortList{3, 2, 1, 5, 0, 4}
	compare := func (a, b int) int {
		return a - b
	}
	swap := func (list comparable.SortList, i, j int) {
		list[i], list[j] = list[j], list[i]
	}

	InsertSort(list, compare, swap)
	fmt.Println("result:", list)
}