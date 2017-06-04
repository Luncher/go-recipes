package comparable

type SortList []int

type Swap func(SortList, int, int)

type Compare func(int, int) int