package selectsort

import (
	"fmt"
	"testing"
	"recipes/sort"
)

func TestSelectOpt(t *testing.T) {
	list := comparable.SortList{3, 1, 5, 2, 6}
	compare := func(a, b int) int {
		return a - b
	}

	swap := func(list comparable.SortList, a, b int) {
		list[a], list[b] = list[b], list[a]
	}

	SelectSort(list, compare, swap)

	fmt.Println("result", list)
}
