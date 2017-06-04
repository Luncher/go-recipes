package binarysearch

//Search export
func Search(n int, f func (int) int) int {
	lo := 0
	hi := n - 1
	
	for lo <= hi {
		mid := lo + (hi - lo) / 2
		ret := f(mid)
		if ret < 0 {
			lo = mid + 1
		} else if ret > 0 {
			hi = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

type IntSlice []int
type StringSlice []string

//SearchInt Export
func SearchInt(arr IntSlice, value int) int {
	return Search(len(arr), func (i int) int {
		return arr[i] - value
	})
}

//SearchString Export
func SearchString(arr StringSlice, value string) int {
	return Search(len(arr), func (i int) int {
		if arr[i] > value {
			return 1
		} else if arr[i] < value {
			return -1
		} else {
			return 0
		}
	})
}

//Search StringSlice
func (arr StringSlice)Search(v string) int {
	return SearchString(arr, v)
}

//Search IntSlice
func (arr IntSlice)Search(v int) int {
	return SearchInt(arr, v)
}