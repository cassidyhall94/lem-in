package main

import (
	"net/http"
	"text/template"
)

// Graph represents an adjaceny list graph
/*type Graph struct {
	vertices []*Vertex
}

// Vertex represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a Vertex to the Graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not aded because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
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

// getVertex returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v:", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
	}
}

func main() {
	test := &Graph{}

	antsList := ants.SpawnAnts(bestCombination)
	ants.MakeStep(antsList)

	// test := &Graph{}

	// for i := 0; i < 5; i++ {
	// 	test.AddVertex(i)
	// }
	// test.AddEdge(1, 2)
	// test.Print()
}*/

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}
