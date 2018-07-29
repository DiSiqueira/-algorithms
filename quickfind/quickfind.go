package quickfind

type (
	QuickFind interface {
		Union(a, b int)
		Connected(a, b int) bool
	}

	quickFind struct {
		data []int
	}
)

func New(size int) QuickFind {
	data := make([]int, size)

	for i := range data {
		data[i] = i
	}

	return &quickFind{
		data: data,
	}
}

func (q *quickFind) Union(a, b int) {
	for i,val := range q.data {
		if val == b {
			q.data[i] = q.data[a]
		}
	}
}

func (q *quickFind) Connected(a, b int) bool {
	return q.data[a] == q.data[b]
}
