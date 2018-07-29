package pathcompression

type (
	QuickUnionWeighted interface {
		Union(a, b int)
		Connected(a, b int) bool
		Root(a int) int
	}

	quickUnionWeighted struct {
		data []int
		size []int
	}
)

func New(s int) QuickUnionWeighted {
	data := make([]int, s)
	size := make([]int, s)

	for i := range data {
		data[i] = i
		size[i] = 1
	}

	return &quickUnionWeighted{
		data: data,
		size: size,
	}
}

func (q *quickUnionWeighted) Root(a int) int {
	for ; q.data[a] != a; {
		q.data[a] = q.data[q.data[a]]
		a = q.data[a]
	}
	return a
}

func (q *quickUnionWeighted) Union(a, b int) {
	i := q.Root(a)
	j := q.Root(b)
	if i == j {
		return
	}

	if q.size[i] < q.size[j] {
		q.data[i] = j
		q.size[j] += q.size[i]
		return
	}

	q.data[j] = i
	q.size[i] += q.size[j]
}

func (q *quickUnionWeighted) Connected(a, b int) bool {
	return q.Root(a) == q.Root(b)
}
