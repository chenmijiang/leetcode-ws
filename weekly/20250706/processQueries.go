package main

import (
	"fmt"
)

// UnionFind 数据结构
type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return
	}
	// 按秩合并
	if uf.rank[px] < uf.rank[py] {
		uf.parent[px] = py
	} else if uf.rank[px] > uf.rank[py] {
		uf.parent[py] = px
	} else {
		uf.parent[py] = px
		uf.rank[px]++
	}
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
	// 初始化并查集
	uf := NewUnionFind(c + 1)

	// 构建电网连接关系
	for _, conn := range connections {
		uf.Union(conn[0], conn[1])
	}

	// 记录电站状态
	online := make([]bool, c+1)
	for i := 1; i <= c; i++ {
		online[i] = true
	}

	result := make([]int, 0)
	// 处理查询
	for _, q := range queries {
		op, x := q[0], q[1]
		if op == 2 {
			// 电站离线
			online[x] = false
		} else {
			// 维护检查
			if online[x] {
				result = append(result, x)
			} else {
				// 找到同一电网中最小的在线电站
				root := uf.Find(x)
				minOnline := -1
				// 遍历所有电站，找到同一电网中最小的在线电站
				for i := 1; i <= c; i++ {
					if online[i] && uf.Find(i) == root {
						if minOnline == -1 || i < minOnline {
							minOnline = i
						}
					}
				}
				result = append(result, minOnline)
			}
		}
	}
	return result
}

func main() {
	c := 5
	connections := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	queries := [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}}
	fmt.Println(processQueries(c, connections, queries))

	c = 3
	connections = [][]int{}
	queries = [][]int{{1, 1}, {2, 1}, {1, 1}}
	fmt.Println(processQueries(c, connections, queries))

	c = 4
	// [[4,3],[3,1],[4,2],[3,2],[4,1]]
	connections = [][]int{{4, 3}, {3, 1}, {4, 2}, {3, 2}, {4, 1}}
	// [[2,3],[1,2],[2,4],[1,1],[2,2],[1,2],[1,2],[2,2],[1,3],[2,3],[2,4],[2,3],[2,4],[1,2],[1,1]]
	queries = [][]int{{2, 3}, {1, 2}, {2, 4}, {1, 1}, {2, 2}, {1, 2}, {1, 2}, {2, 2}, {1, 3}, {2, 3}, {2, 4}, {2, 3}, {2, 4}, {1, 2}, {1, 1}}
	//expect: [2,1,1,1,1,1,1]
	// print : [2 1 -1 -1 1 -1 1] , error
	fmt.Println(processQueries(c, connections, queries))
}
