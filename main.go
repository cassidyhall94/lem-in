package main

import (
	"fmt"
)

// Graph represents an adjaceny list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex represents a graph vertex
type Vertex struct {
	// room is the vertex identifer (room name/number in text file)
	room     string
	// adjacent is the rooms that are connected by an edge/tunnel/link
	adjacent []*Vertex
}

// AddVertex adds a Vertex to the Graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it is an existing room", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{room: k})
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex returns a pointer to the Vertex with a room integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.room == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.room {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v:", v.room)
		for _, v := range v.adjacent {
			fmt.Printf("%v", v.room)
		}
	}
}

func main() {
	// read the arg
	// open the file
	// read the file lines into []string
	// parse lines
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	test.AddEdge(1, 2)
	test.Print()
}
