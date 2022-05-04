package main

import (
	// "fmt"

	"log"
	"os"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/ants"
	dataparser "git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/data-parser"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/farm"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/paths"
)

// // Graph represents an adjaceny list graph
// type Graph struct {
// 	vertices []*Vertex
// }

// // Vertex represents a graph vertex
// type Vertex struct {
// 	// room is the vertex identifer (room name/number in text file)
// 	room string
// 	// adjacent is the rooms that are connected by an edge/tunnel/link
// 	adjacent []*Vertex
// }

// AddVertex adds a Vertex to the Graph
// func (g *Graph) AddVertex(k string) {
// 	if contains(g.vertices, k) {
// 		err := fmt.Errorf("Vertex %v not added because it is an existing key", k)
// 		fmt.Println(err.Error())
// 	} else {
// 		g.vertices = append(g.vertices, &Vertex{room: k})
// 	}
// }

// // AddEdge adds an edge to the graph
// func (g *Graph) AddEdge(from, to string) {
// 	// get vertex
// 	fromVertex := g.getVertex(from)
// 	toVertex := g.getVertex(to)
// 	// check error
// 	if fromVertex == nil || toVertex == nil {
// 		err := fmt.Errorf("Invalid edge (%v-->%v)", from, to)
// 		fmt.Println(err.Error())
// 	} else if contains(fromVertex.adjacent, to) {
// 		err := fmt.Errorf("Existing edge (%v-->%v)", from, to)
// 		fmt.Println(err.Error())
// 	} else {
// 		// add edge
// 		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
// 	}
// }

// // getVertex returns a pointer to the Vertex with a key integer
// func (g *Graph) getVertex(k string) *Vertex {
// 	for i, v := range g.vertices {
// 		if v.room == k {
// 			return g.vertices[i]
// 		}
// 	}
// 	return nil
// }

// // contains
// func contains(s []*Vertex, k string) bool {
// 	for _, v := range s {
// 		if k == v.room {
// 			return true
// 		}
// 	}
// 	return false
// }

// // Print will print the adjacent list for each vertex of the graph
// func (g *Graph) Print() {
// 	for _, v := range g.vertices {
// 		fmt.Printf("\nVertex %v:", v.room)
// 		for _, v := range v.adjacent {
// 			fmt.Printf("%v", v.room)
// 		}
// 	}
// }

func main() {
	if len(os.Args) != 2 {
		log.Fatal("The file name is missing")
		os.Exit(1)
	}

	data, _ := dataparser.LoadData(os.Args[1])
	generationData := dataparser.ReadData(data)
	filledFarm := farm.GenerateFarm(generationData)
	connectedFarm := farm.ConnectRooms(filledFarm, generationData)
	allPaths := paths.FindAllPaths(connectedFarm)
	shortestPath := paths.FindShortestPath(allPaths)

	ants.MoveAnts(connectedFarm.Ants, shortestPath)

	// var allPaths [][]*structs.Room
	// paths.FindPaths(make([]*structs.Room, 0), structs.FARM[structs.STARTROOMID], 0, &allPaths, &structs.FARM[structs.STARTROOMID])
	// utils.SortPathsByLength(&allPaths)

	// differentCombinations := paths.SearchCombinations(allPaths)
	// bestCombination := paths.SearchBestCombination(differentCombinations)

	// antsList := ants.CreateAnts(bestCombination)
	// ants.CreateStep(antsList)
}
