package pkg

type Rack struct {
	RackID   string
	Capacity int
	Books    map[string]int // {bookId: copies}
}

func (r *Rack) BookCount() int {
	cnt := 0

	for _, bookCnt := range r.Books {
		cnt += bookCnt
	}

	return cnt
}
