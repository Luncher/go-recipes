package selectsort

import "recipes/sort"

func SelectSort(list comparable.SortList, compare comparable.Compare, swap comparable.Swap) {
	for i := 0; i < len(list); i++ {
		min := i
		for j := i + 1; j < len(list); j++ {
			if compare(list[min], list[j]) < 0 {
				min = j
			}
		}
		swap(list, i, min)
	}
}
