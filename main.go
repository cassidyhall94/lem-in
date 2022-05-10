package main

import (
	"fmt"
	"log"
	"os"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/ants"
	dataparser "git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/data-parser"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/farm"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/paths"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("The file name is missing")
		os.Exit(1)
	}

	data, err := dataparser.LoadData(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	generationData := dataparser.ReadData(data)
	filledFarm := farm.GenerateFarm(generationData)
	connectedFarm := farm.ConnectRooms(filledFarm, generationData)
	allPaths := paths.FindAllPaths(connectedFarm)

	// instead of making a proper algorithm lets use randomness and brute force the problem
	// 1. using a loop, randomise the order of allpaths n times (n is the number of possible ways to randomise len(allpaths) paths), collect all of the random path orders into a slice
	// 2. in a loop 
	// 2a. run trimpaths
	// 2b. run sortpaths
	// 2c. run antAssignments
	// 2d. run moveants
	// 2e. get the longestMoveset, append to a helper (this needs to be able to refer back to the allMoves that it came from)
	// 3. printMoves() the shortest allMoves (determined by the smallest longestMoveset in the helper)


	sortedPaths := paths.SortPaths(allPaths)
	trimmedPaths := paths.TrimPaths(sortedPaths)
	assignedAnts := ants.AssignAnts(connectedFarm.Ants, trimmedPaths)

	allMoves := [][]string{}
	for _, ps := range assignedAnts {
		antsMoved := ants.MoveAnts(ps.Ants, ps)
		allMoves = append(allMoves, antsMoved)
	}

	// longestMoveset is the number of turns (needs verifying)
	longestMoveset := func() int {
		l := 0
		for _, moveset := range allMoves {
			if len(moveset) > l {
				l = len(moveset)
			}
		}
		return l
	}()

	printMoves(allMoves, longestMoveset)
}

func printMoves(allMoves [][]string, longestMoveset int) {
	out := make([]string, longestMoveset)
	for _, moveset := range allMoves {
		for i, move := range moveset {
			out[i] = fmt.Sprintf("%s %s", out[i], move)
		}
	}

	for _, o := range out {
		fmt.Println(o)
	}
}