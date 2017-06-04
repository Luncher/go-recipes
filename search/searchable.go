package searchable

type SearchInterface interface {
	Compare(i int) int
	Len() int
}