package main

import "sort"

type unionFind struct {
	parent []int
	size   []int // the height of the tree rooted at i
}

// newUnionFind returns a new union-find data structure with n elements.
func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent, size}
}

// find returns the root of the set that x belongs to.
// path compression is used to set the parents of all nodes on the path from x to the root node as root nodes
func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

// union merges the sets that x and y belong to.
func (uf *unionFind) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	// union by size, always merge the smaller set into the larger set
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
}

// inSameSet returns true if x and y belong to the same set.
func (uf *unionFind) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

// edge represents an edge between two vertices.
type edge struct {
	v    int
	w    int
	diff int
}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	edges := []edge{}
	// add all edges to the list
	for i, row := range heights {
		for j, h := range row {
			id := i*m + j
			if i > 0 {
				edges = append(edges, edge{id - m, id, abs(h - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(h - heights[i][j-1])})
			}
		}
	}

	// sort edges by diff in ascending order
	sort.Slice(edges, func(i, j int) bool { return edges[i].diff < edges[j].diff })

	// find the minimum diff that connects the first and last vertices
	uf := newUnionFind(n * m)
	for _, e := range edges {
		uf.union(e.v, e.w)
		if uf.inSameSet(0, n*m-1) {
			return e.diff
		}
	}
	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
