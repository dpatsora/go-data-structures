package disjoint_set

type DisjointSet struct {
	data  []int
	ranks []int // approximate height of the tree
	size  int
}

func NewDisjointSet(size int) *DisjointSet {
	data := make([]int, size)
	ranks := make([]int, size)
	// make each node independent
	for i, _ := range data {
		data[i] = i
	}

	return &DisjointSet{
		data:  data,
		ranks: ranks,
		size:  size,
	}
}

// Find returns root node for the given element.
// This method uses path compression algorithm
func (ds *DisjointSet) Find(x int) int {
	// Incorrect value passed
	if x >= ds.size && x < 0 {
		return -1
	}

	// recursively find root node and apply path compression if given element is not a root one
	if ds.data[x] != x {
		ds.data[x] = ds.Find(ds.data[x])
		return ds.data[x]
	}

	// given element is root
	return x
}

// Union merges 2 nodes into 1 set. It uses rank union
func (ds *DisjointSet) Union(x, y int) {
	rootX := ds.Find(x)
	rootY := ds.Find(y)

	if rootX != rootY {
		if ds.ranks[rootX] > ds.ranks[rootY] {
			ds.data[rootY] = rootX
		} else if ds.ranks[rootX] < ds.ranks[rootY] {
			ds.data[rootX] = rootY
		} else {
			ds.data[rootY] = rootX
			ds.ranks[rootX] += 1
		}
	}
}

func (ds *DisjointSet) IsConnected(x, y int) bool {
	return ds.Find(x) == ds.Find(y)
}

func (ds *DisjointSet) IndependentNodesCount() int {
	var num int
	for i, v := range ds.data {
		if i == v {
			num += 1
		}
	}

	return num
}
