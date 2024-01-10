package entities

type Pagination struct {
	Offset      int
	Limit       int
	CurrentPage int
	Total       int64
}
