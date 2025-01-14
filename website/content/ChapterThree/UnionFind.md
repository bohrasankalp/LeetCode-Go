---
title: 3.2 UnionFind
type: docs
weight: 2
---

# and lookup UnionFind

```go
package template

// UnionFind defind
//path compression + rank optimization
type UnionFind struct {
	parent, rank []int
	count        int
}

// Init define
func (uf *UnionFind) Init(n int) {
	uf.count = n
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i
	}
}

// Find define
func (uf *UnionFind) Find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	// compress path
	for p != uf.parent[p] {
		tmp := uf.parent[p]
		uf.parent[p] = root
		p = tmp
	}
	return root
}

// Union define
func (uf *UnionFind) Union(p, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if proot == qroot {
		return
	}
	if uf.rank[qroot] > uf.rank[proot] {
		uf.parent[proot] = qroot
	} else {
		uf.parent[qroot] = proot
		if uf.rank[proot] == uf.rank[qroot] {
			uf.rank[proot]++
		}
	}
	uf.count--
}

// TotalCount define
func (uf *UnionFind) TotalCount() int {
	return uf.count
}

// UnionFindCount define
//Calculate the number of elements in each collection + the maximum number of collection elements
type UnionFindCount struct {
	parent, count []int
	maxUnionCount int
}

// Init define
func (uf *UnionFindCount) Init(n int) {
	uf.parent = make([]int, n)
	uf.count = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i
		uf.count[i] = 1
	}
}

// Find define
func (uf *UnionFindCount) Find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	return root
}

//Without rank compression, the time complexity explodes, too high
// func (uf *UnionFindCount) union(p, q int) {
// 	proot := uf.find(p)
// 	qroot := uf.find(q)
// 	if proot == qroot {
// 		return
// 	}
// 	if proot != qroot {
// 		uf.parent[proot] = qroot
// 		uf.count[qroot] += uf.count[proot]
// 	}
// }

// Union define
func (uf *UnionFindCount) Union(p, q int) {
	proot := uf.Find(p)
	qroot := uf.Find(q)
	if proot == qroot {
		return
	}
	if proot == len(uf.parent)-1 {
		//proot is root
	} else if qroot == len(uf.parent)-1 {
		// qroot is root, always attach to root
		proot, qroot = qroot, proot
	} else if uf.count[qroot] > uf.count[proot] {
		proot, qroot = qroot, proot
	}

	//set relation[0] as parent
	uf.maxUnionCount = max(uf.maxUnionCount, (uf.count[proot] + uf.count[qroot]))
	uf.parent[qroot] = proot
	uf.count[proot] += uf.count[qroot]
}

// Count define
func (uf *UnionFindCount) Count() []int {
	return uf.count
}

// MaxUnionCount define
func (uf *UnionFindCount) MaxUnionCount() int {
	return uf.maxUnionCount
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

```



----------------------------------------------
<div style="display: flex;justify-content: space-between;align-items: center;">
<p><a href="https://books.halfrost.com/leetcode/ChapterThree/Segment_Tree/">⬅️previous page</a></p>
<p><a href="https://books.halfrost.com/leetcode/ChapterThree/LRUCache/">next page➡️</a></p>
</div>
