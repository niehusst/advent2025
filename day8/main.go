package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"math"
	"sort"
	"strconv"
	"container/heap"
)

type JBox struct {
	key string
	x int
	y int
	z int
}

type Edge struct {
	weight float64
	j1 *JBox
	i1 int // og index in coords. used for uf ops
	j2 *JBox
	i2 int // og index in coords. used for uf ops
}

type UnionFind struct {
	parent []int
	size []int
}

func (uf *UnionFind) init(length int) {
	uf.parent = make([]int, length)
	uf.size = make([]int, length)

	for i := 0; i < length; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
}

func (uf *UnionFind) find(n int) int {
	curr := n
	for uf.parent[curr] != curr {
		uf.parent[curr] = uf.parent[uf.parent[curr]]
		curr = uf.parent[curr]
	}
	return curr
}

func (uf *UnionFind) union(n1 int, n2 int) bool {
	p1 := uf.find(n1)
	p2 := uf.find(n2)
	if p1 == p2 {
		return true
	}

	// make bigger size the parrent
	if uf.size[p1] > uf.size[p2] {
		uf.parent[p2] = p1
		uf.size[p1] += uf.size[p2]
	} else {
		uf.parent[p1] = p2
		uf.size[p2] += uf.size[p1]
	}

	return false
}

//type Pair [2]int
type PairMaxHeap [][2]int

func (h PairMaxHeap) Len() int           { return len(h) }
func (h PairMaxHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h PairMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairMaxHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([2]int))
}

func (h *PairMaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


/*
part1

build MST from points, weight is distance between points
n^2 
actually union find???

limit 1000???
*/

func part1(coords []*JBox) {
	// find all edge weights
	edges := make([]Edge, 0)
	for i := 0; i < len(coords); i++ {
		for j := i+1; j < len(coords); j++ {
			c1 := coords[i]
			c2 := coords[j]

			w := math.Sqrt(math.Pow(float64(c1.x - c2.x), 2)+math.Pow(float64(c1.y - c2.y), 2)+math.Pow(float64(c1.z - c2.z), 2))
			edges = append(edges, Edge{weight:w, j1:c1, i1:i, j2:c2, i2:j})
		}
	}

	// sort edges
	sort.Slice(edges, func (i int, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	// for each edge take if adds new node to graph
	// limit 1000 edges???
	uf := UnionFind{}
	uf.init(len(coords))
	for i, edge := range edges {
		if i >= 1000 {
			fmt.Println("whew too many")
			break
		}
		//fmt.Printf("joining %s and %s\n", edge.j1.key, edge.j2.key)
		uf.union(edge.i1, edge.i2)
	}

	// multiply 3 largest circuit size together
	h := &PairMaxHeap{}
	heap.Init(h)
	for n, size := range uf.size {
		heap.Push(h, [2]int{size, n})
	}

	seen := make(map[int]bool)
	circuitProd := 1
	prodCount := 0
	for h.Len() > 0 && prodCount < 3 {
		p := heap.Pop(h).([2]int)
		node := uf.find(p[1])
		if !seen[node] {
			seen[node] = true
			circuitProd *= p[0]
			prodCount++
		}
	}
	//fmt.Println(uf)
	fmt.Println("Top 3 circuit sizes prod", circuitProd)
}

/*
part2

build full mst. mult x coords of last 2 nodes connected
*/

func part2(coords []*JBox) {
	// find all edge weights
	edges := make([]Edge, 0)
	for i := 0; i < len(coords); i++ {
		for j := i+1; j < len(coords); j++ {
			c1 := coords[i]
			c2 := coords[j]

			w := math.Sqrt(math.Pow(float64(c1.x - c2.x), 2)+math.Pow(float64(c1.y - c2.y), 2)+math.Pow(float64(c1.z - c2.z), 2))
			edges = append(edges, Edge{weight:w, j1:c1, i1:i, j2:c2, i2:j})
		}
	}

	// sort edges
	sort.Slice(edges, func (i int, j int) bool {
		return edges[i].weight < edges[j].weight
	})

	// for each edge take if adds new node to graph
	uf := UnionFind{}
	uf.init(len(coords))
	seen := make(map[string]bool)
	var lastEdge *Edge
	for _, edge := range edges {
		//fmt.Printf("joining %s and %s\n", edge.j1.key, edge.j2.key)
		if !uf.union(edge.i1, edge.i2) {
			// new union created
			seen[edge.j1.key] = true
			seen[edge.j2.key] = true
		}
		if len(seen) == len(coords) {
			// created last connection!
			lastEdge = &edge
			break
		}
	}

	if lastEdge == nil {
		log.Fatal("fuck no last edge??")
	}
	fmt.Println("jbox x coord prod", lastEdge.j1.x * lastEdge.j2.x)
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	coords := make([]*JBox, len(data))
	for i := 0; i < len(data); i++ {
		spoints := strings.Split(data[i], ",")
		points := make([]int, 3)
		for j := 0; j < 3; j++ {
			v, err := strconv.Atoi(spoints[j])
			if err != nil {
				log.Fatal(err)
			}
			points[j] = v
		}
		coords[i] = &JBox{
			key: data[i],
			x: points[0],
			y: points[1],
			z: points[2],
		}
	}
	part2(coords)
}
