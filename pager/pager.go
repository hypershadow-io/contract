package pager

// MakePager creates a new Pager instance with the specified offset and count.
// 'from' defines the starting offset, and 'count' specifies the number of items to retrieve.
func MakePager(
	from int64,
	count int64,
) Pager {
	return page{
		from:  from,
		count: count,
	}
}

// page is an internal implementation of the Pager interface.
type page struct {
	from  int64
	count int64
}

func (a page) GetFrom() int64  { return a.from }
func (a page) GetCount() int64 { return a.count }
