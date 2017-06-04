package quicksort

import "fmt"
import "recipes/sort"


func partition(list comparable.SortList, compare comparable.Compare, swap comparable.Swap, lo, hi int) int {
	i := lo
	j := hi + 1
	v := list[lo]

	for i < j {
		i = i + 1
		for compare(list[i], v) < 0 {
			i = i + 1			
			if i == hi {
				break
			}
		}
		j = j - 1		
		for compare(v, list[j]) < 0 {
			j = j - 1			
			if j == lo {
				break
			}
		}

		swap(list, i, j)
	}

	swap(list, lo, j)
	
	return j
}

func doSort(list comparable.SortList, compare comparable.Compare, swap comparable.Swap, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(list, compare, swap, lo, hi)
	fmt.Println("list:", list)
	fmt.Println("p:", p, "lo:", lo, "hi:", hi)
	doSort(list, compare, swap, lo, p - 1)
	doSort(list, compare, swap, p + 1, hi)
}

//QuickSort export
func QuickSort(list comparable.SortList, compare comparable.Compare, swap comparable.Swap) {
	doSort(list, compare, swap, 0, len(list) - 1)
}