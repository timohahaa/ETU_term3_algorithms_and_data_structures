package graph

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/timohahaa/ETU_term3_algorithms_and_data_structures/lds"
	"github.com/timohahaa/ETU_term3_algorithms_and_data_structures/timsort"
)

// Kruskals algo implementation
// Graph struct is for undirected weighted connected graph

type Vertex string

type Edge struct {
	V1     Vertex
	V2     Vertex
	Weight int
}

type Graph struct {
	Vertecies   *lds.DynamicArray[Vertex]
	Edges       *lds.DynamicArray[Edge]
	AdjMatrix   *lds.DynamicArray[*lds.DynamicArray[int]]
	MST         *lds.DynamicArray[Edge]
	VertexCount int
	EdgeCount   int
	// disjoint set is used to check if adding an edge forms a cycle
	disjointSet map[Vertex]Vertex
}

func read1Dsclice(r io.Reader, size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		var num int
		fmt.Fscan(r, &num)
		arr[i] = num
	}
	fmt.Fscanln(r)
	return arr
}

func read2Dslice(r io.Reader, rows, cols int) [][]int {
	arr := make([][]int, rows)
	for i := 0; i < rows; i++ {
		arr[i] = read1Dsclice(r, cols)
	}
	return arr
}

func NewGraph(filePath string) (*Graph, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	g := &Graph{}

	freader := bufio.NewReader(file)
	firstLine, err := freader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	vertecies := strings.Fields(firstLine)
	g.VertexCount = len(vertecies)
	adjMat := read2Dslice(freader, g.VertexCount, g.VertexCount)

	// init the adjacent matrix {
	g.AdjMatrix = lds.NewDArray[*lds.DynamicArray[int]]()
	for i := 0; i < g.VertexCount; i++ {
		row := lds.NewDArray[int]()
		g.AdjMatrix.PushBack(row)
	}
	// now fill it with read values
	for i := 0; i < len(adjMat); i++ {
		for j := 0; j < len(adjMat); j++ {
			g.AdjMatrix.GetNoError(i).PushBack(adjMat[i][j])
		}
	}

	// init vertecies
	g.Vertecies = lds.NewDArray[Vertex]()
	for _, v := range vertecies {
		g.Vertecies.PushBack(Vertex(v))
	}

	// now parse and form the edges
	// because graph is acyclic, only need to check values in the upper part
	g.Edges = lds.NewDArray[Edge]()
	for i := 0; i < g.VertexCount; i++ {
		for j := i + 1; j < g.VertexCount; j++ {
			w := g.AdjMatrix.GetNoError(i).GetNoError(j)
			// no edge exists if weight is zero
			if w == 0 {
				continue
			}
			e := Edge{
				V1:     g.Vertecies.GetNoError(i),
				V2:     g.Vertecies.GetNoError(j),
				Weight: w,
			}
			g.Edges.PushBack(e)
			g.EdgeCount++
		}
	}
	return g, nil
}

// finds a group in a disjoint disjoint set
func (g *Graph) findGroup(v Vertex) Vertex {
	for v != g.disjointSet[v] {
		v = g.disjointSet[v]
	}
	return v
}

// does a union in a set
// no need to check the rank or size, as it doesnt matter
// only thing that matters is joining, no matter in witch order
func (g *Graph) union(a, b Vertex) {
	parentA := g.findGroup(a)
	parentB := g.findGroup(b)
	if parentA == parentB {
		return
	}
	g.disjointSet[parentA] = parentB
}

func (e Edge) SortedEdgeName() string {
	if strings.Compare(string(e.V1), string(e.V2)) == -1 {
		return string(e.V1) + "<->" + string(e.V2)
	}
	return string(e.V2) + "<->" + string(e.V1)
}

// computes MST of a graph
func (g *Graph) ComputeMST() *lds.DynamicArray[Edge] {
	g.MST = lds.NewDArray[Edge]()
	g.disjointSet = make(map[Vertex]Vertex)

	// sort the edges by weight first
	timsort.TimSort[Edge](g.Edges, func(a, b Edge) int {
		if a.Weight > b.Weight {
			return 1
		}
		if a.Weight < b.Weight {
			return -1
		}
		return 0
	})

	// add vertecies to disjoint set
	for i := 0; i < g.Vertecies.Len(); i++ {
		v := g.Vertecies.GetNoError(i)
		g.disjointSet[v] = v
	}

	// MST has EXACTLY (numOfVertecies - 1) edges
	edgesTaken := 0
	iter := 0
	for edgesTaken < g.VertexCount-1 {
		// take the smallest available edge
		e := g.Edges.GetNoError(iter)
		iter++
		// now check if this edge will add a cycle or not
		// find groups of both verceties in a disjoint set
		v1Parent := g.findGroup(e.V1)
		v2Parent := g.findGroup(e.V2)

		// if parents are not the same, then this edge can be added to MST
		if v1Parent != v2Parent {
			edgesTaken++
			g.MST.PushBack(e)
			g.union(e.V1, e.V2)
		}
	}

	// sort the MST before returning
	// sort by name
	timsort.TimSort[Edge](g.MST, func(a, b Edge) int {
		e1Name := a.SortedEdgeName()
		e2Name := b.SortedEdgeName()
		return strings.Compare(e1Name, e2Name)
	})
	return g.MST
}
