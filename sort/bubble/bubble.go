package bubblesort

import "recipes/sort"

//BubbleSort export
func BubbleSort(list comparable.SortList, compare comparable.Compare, swap comparable.Swap) {
	for i := 1; i < len(list); i++ {
		for j := 0; j < len(list) - i; j++ {
			if compare(list[j], list[j + 1]) > 0 {
				swap(list, j, j+1)
			}
		}
	}
}