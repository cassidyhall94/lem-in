package ants

import (
	"fmt"
	"os"

	dataparser "git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/data-parser"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/farm"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/paths"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// Call graph → SortAnts to sort ants between paths. Sort the paths for BFS structures.
//    Utilise Jamie Dawson's algorithm to find the most optimal paths;
//    I.e we have to make sure that we are sending all the ants from start to end in the most optimized way possible.
//    Send the first ant on the shortest first journey. Next send the next Ant on it's journey, etc.
//    Until all Ants have reached the End room.

// 6. Call graph → MoveAnts which moves ants between rooms from start room to the end room.
//    There exists a queue for active ants that have not reached the end.
//    An Ant can only move into the next room (on it's path) if that room is empty.
//    The Ants are moved using the Paths found using the BFS (Breadth First Search) method.
//    This Iteration is repeated until all Ants are in the End Room.

// func CreateAnts assigns a path to each ant
func CreateAnts(Ants structs.Farm, allPaths []*structs.PathStruct) []*structs.Ant {
	data, _ := dataparser.LoadData(os.Args[1])
	generationData := dataparser.ReadData(data)
	filledFarm := farm.GenerateFarm(generationData)
	connectedFarm := farm.ConnectRooms(filledFarm, generationData)
	paths := paths.FindAllPaths(connectedFarm)
	// shortestPath := paths.FindShortestPath(paths)

	noOfPaths := lenPathStruct(paths)

	antsPath := antsPerPath(paths)

	fmt.Println(antsPath)

	for p := 0; p < noOfPaths; p++ {
		var i int
		for i = p + 1; i <= 20-p; {
			var antToAdd *structs.Ant = new(structs.Ant)
			noOfRooms := 0

			antToAdd.Id = i
			antToAdd.Path = paths[p].Path
			antToAdd.RoomsPassed = noOfRooms

			fmt.Println(antToAdd)

			if i < 19 {
				i = i + 2
			} else {
				i = i + 1
			}
		}
	}

	return []*structs.Ant{}
}

func lenPathStruct(paths []*structs.PathStruct) int {
	lenPathCounter := 0

	for _, path := range paths {
		for _, pathRoom := range path.Path {
			if pathRoom.Name != "" {
				lenPathCounter = lenPathCounter + 1
			}
		}
	}
	return lenPathCounter
}

func diffRoomsInPaths(paths []*structs.PathStruct) []int {
	difPaths := []int{}
	noOfRoomsPath := []int{}

	for n := 0; n < lenPathStruct(paths); n++ {
		noOfRooms := len(paths[n].Path)
		noOfRoomsPath = append(noOfRoomsPath, noOfRooms)
	}

	for d := 1; d < len(noOfRoomsPath); d++ {
		difRooms := noOfRoomsPath[d] - noOfRoomsPath[d-1]
		difPaths = append(difPaths, difRooms)
	}

	return difPaths
}

func antsPerPath(paths []*structs.PathStruct) []int {
	noOfPaths := lenPathStruct(paths)
	difPaths := diffRoomsInPaths(paths)
	antsPerPath := []int{}

	antsPerPath = append(antsPerPath, 20/noOfPaths+difPaths[0]-1)
	for dr := 1; dr < noOfPaths; dr++ {
		antCounter := 20 - antsPerPath[dr-1]
		if dr == len(difPaths) {
			antsPerPath = append(antsPerPath, antCounter/len(difPaths))
		} else {
			antsPerPath = append(antsPerPath, antCounter/len(difPaths)+difPaths[dr]-1)
		}
	}

	return antsPerPath
}

