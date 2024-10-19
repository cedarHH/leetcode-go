package exercises

import "fmt"

type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func NewUnionFind(n, count int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		count:  count,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
		uf.count--
	}
}

func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) GetCount() int {
	return uf.count
}

func numIslands(grid [][]byte) int {
	var count int
	row, col := len(grid), len(grid[0])
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
			}
		}
	}
	unionFind := NewUnionFind(row*col, count)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				if i > 0 && grid[i-1][j] == '1' {
					unionFind.Union(i*col+j, (i-1)*col+j)
				}
				if j > 0 && grid[i][j-1] == '1' {
					unionFind.Union(i*col+j, i*col+j-1)
				}
			}
		}
	}
	return unionFind.GetCount()
}

func NumIslands() {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(grid))
}
