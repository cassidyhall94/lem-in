package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/ants"
	dataparser "git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/data-parser"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/farm"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/paths"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

type moveData map[int][][]string

func main() {
	if len(os.Args) != 2 {
		log.Fatal("ERROR: File name missing")
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

	allPaths = reduceSlice(allPaths, 8)

	pathsFlat := [][]*structs.Room{}
	for _, p := range allPaths {
		pathsFlat = append(pathsFlat, p.Path)
	}
	lookupMap, sortString := paths.MapPathsToStrings(pathsFlat)
	permedStrings := paths.PermutationsIter(sortString)

	moveStore := moveData{}
	m := sync.Mutex{}
	wg := sync.WaitGroup{}

	// increasing the int in chunkslice will make the program run faster and use more RAM
	chunks := chunkSlice(permedStrings, 50)

	for _, c := range chunks {
		for _, s := range c {
			wg.Add(1)
			go func(s string) {
				defer wg.Done()
				allPathsFlat := paths.GetPathsFromStrings(lookupMap, s)
				allPaths := []*structs.PathStruct{}
				for _, p := range allPathsFlat {
					allPaths = append(allPaths, &structs.PathStruct{Path: p})
				}
				trimmedPaths := paths.TrimPaths(allPaths)
				assignedAnts := ants.AssignAnts(connectedFarm.Ants, trimmedPaths)

				allMoves := [][]string{}
				for _, ps := range assignedAnts {
					antsMoved := ants.MoveAnts(ps.Ants, ps)
					allMoves = append(allMoves, antsMoved)
				}

				longestMoveset := getLongestMoveset(allMoves)

				m.Lock()
				if _, ok := moveStore[longestMoveset]; !ok {
					moveStore[longestMoveset] = allMoves
				}
				m.Unlock()
			}(s)
		}
		wg.Wait()
	}

	allMoves, numMoves := [][]string{}, 0
	for n, m := range moveStore {
		if numMoves == 0 {
			allMoves = m
			numMoves = n
		} else if n < numMoves {
			allMoves = m
			numMoves = n
		}
	}

	printMoves(allMoves, numMoves)
}

func printMoves(allMoves [][]string, longestMoveset int) {
	out := make([]string, longestMoveset)
	for _, moveset := range allMoves {
		for i, move := range moveset {
			out[i] = fmt.Sprintf("%s %s", out[i], move)
		}
	}

	for _, o := range out {
		fmt.Printf("%s\n", o)
	}
}

// longestMoveset gets the number of turns (needs verifying)
func getLongestMoveset(allMoves [][]string) int {
	l := 0
	for _, moveset := range allMoves {
		if len(moveset) > l {
			l = len(moveset)
		}
	}
	return l
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func reduceSlice(allPaths []*structs.PathStruct, reduceTo int) []*structs.PathStruct {
	for i, ps := range allPaths {
		if !(i+1 > len(allPaths)-1) {
			if len(ps.Path) > reduceTo {
				allPaths = append(allPaths[:i], allPaths[i+1:]...)
			}
		}
	}
	countAbove, countBelow := 0, 0
	for _, ps := range allPaths {
		if len(ps.Path) > reduceTo {
			countAbove++
		} else if len(ps.Path) < reduceTo {
			countBelow++
		}
	}
	if len(allPaths) > 10 && countAbove-countBelow > 0 {
		return reduceSlice(allPaths, reduceTo-1)
	}
	return allPaths
}
