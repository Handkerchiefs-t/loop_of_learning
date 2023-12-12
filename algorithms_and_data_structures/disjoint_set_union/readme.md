# 并查集

## 介绍
并查集是一种树型的数据结构，用于处理一些不相交集合（Disjoint Sets）的合并及查询问题。并查集的核心思想是维护一个集合的划分，每个集合通过一个代表来表示。

### 操作
并查集常见的操作有两个：查找（Find）和合并（Union）。
1. 查找: 就是找到某个元素所在的集合。
2. 合并: 就是把两个集合合成一个集合。

前面说过, 并查集的核心是通过代表来表示集合, 所以查找和合并的操作都是围绕着代表来进行的。

### 优化
并查集的优化主要有两个方面: 路径压缩和按秩合并.
1. 路径压缩: 在查找的过程中, 使路径上的每个节点都指向根节点, 从而减少树的高度.
2. 按秩合并: 在合并的过程中, 使秩小的树的根节点指向秩大的树的根节点, 从而减少树的高度.

## 实现
并查集通常使用数组来实现, 数组的下标表示元素, 数组的值表示该元素的父节点, 如果该元素是根节点, 则其父节点为自身.

### 代码
使用力扣的题目来展示并查集的实现:
[1631. 最小体力消耗路径](https://leetcode.cn/problems/path-with-minimum-effort)
在这个题目中,可以将每个点看做一个元素,每两个节点之间都存在一条边,边的权重是两个节点的高度差的绝对值.
题目要求找到一条路径,使得路径上的边的权重的最大值最小.这个问题可以转化为:找到两个节点之间的路径,使得路径上的边的权重的最大值最小.

这个问题的一个解决思路是,将所有的边都看做一个元素,加入到并查集中. 
然后将所有的边按照权重从小到大排序,并将这些边依次连通(即使用并查集连通不同节点),
当键入某个边后,起点和终点连通,这条边的权重就是路径上的边的权重的最大值,也就是题目要求的答案.
```go
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
```

## 应用
以最朴素的观点来看,并查集可以将多个节点划分为多个不相交的集合,并实现快速的查找和合并.
但是这种使用方式并不常见,因为对于查找和合并来说,map 可能是更好的选择.

在某些情况下,使用并查集来解决一些判断连通性的问题是比较好的选择,比如在图中,判断两个节点是否连通