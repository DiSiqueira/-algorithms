package quickunion

type (
	QuickUnion interface {
		Union(a, b int)
		Connected(a, b int) bool
		Root(a int) int
	}

	quickUnion struct {
		data []int
	}
)

func New(size int) QuickUnion {
	data := make([]int, size)

	for i := range data {
		data[i] = i
	}

	return &quickUnion{
		data: data,
	}
}

func (q *quickUnion) Root(a int) int {
	for ; q.data[a] != a; a = q.data[a] {

	}
	return a
}

func (q *quickUnion) Union(a, b int) {
	q.data[q.Root(a)] = q.Root(b)
}

func (q *quickUnion) Connected(a, b int) bool {
	return q.Root(a) == q.Root(b)
}
