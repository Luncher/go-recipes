package insertsort

import "recipes/sort"

func InsertSort(list comparable.SortList, compare comparable.Compare, swap comparable.Swap) {
	for i:= 1; i < len(list); i++ {
		for j := i; j > 0 && compare(list[j], list[j - 1]) > 0; j-- {
			swap(list, j, j - 1)
		}
	}
}