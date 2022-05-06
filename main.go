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

	data, _ := dataparser.LoadData(os.Args[1])
	generationData := dataparser.ReadData(data)
	filledFarm := farm.GenerateFarm(generationData)
	connectedFarm := farm.ConnectRooms(filledFarm, generationData)
	allPaths := paths.FindAllPaths(connectedFarm)
	shortestPath := paths.FindShortestPath(allPaths)
	antsMoved := ants.MoveAnts(connectedFarm.Ants, shortestPath)
	for _, move := range antsMoved {
		fmt.Println(move)
	}
}
