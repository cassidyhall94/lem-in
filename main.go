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
	sortedPaths := paths.SortPaths(allPaths)
	trimmedPaths := paths.TrimPaths(sortedPaths)
	assignedAnts := ants.AssignAnts(connectedFarm.Ants, trimmedPaths)

	for _, ps := range trimmedPaths {
		fmt.Printf("%+v\n", paths.GetSliceOfRoomNames(ps.Path))
	}

	allMoves := [][]string{}
	for _, ps := range assignedAnts {
		antsMoved := ants.MoveAnts(ps.Ants, ps)
		allMoves = append(allMoves, antsMoved)
	}

	longestMoveset := func() int {
		l := 0
		for _, moveset := range allMoves {
			if len(moveset) > l {
				l = len(moveset)
			}
		}
		return l
	}()

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
