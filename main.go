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

const (
	// noPathsLongerThan controls run time and memory usage for lots of paths
	// increasing noPathsLongerThan can improve the correctness of the output
	// if allPaths is really long >~30 and this is too high your computer could crash
	noPathsLongerThan      = 8
	// maxPermutations controls run time for lots of paths
	// increasing maxPermutations can improve the correctness of the output
	// at the cost of time
	maxPermutations        = 4000000
	// permutationChunkLength controls the size of batches of goroutines when finding all moves through a set of paths
	// increasing permuationChunkLength can improve the speed of the program but needs tuning depending on the computer's cpu
	// if the program is too slow to run, try increasing this first
	permutationChunkLength = 50
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
	generationData, err := dataparser.ReadData(data)
	if err != nil {
		log.Fatalf("ERROR: invalid data format (%s)\n", err.Error())
	}
	filledFarm := farm.GenerateFarm(generationData)
	connectedFarm := farm.ConnectRooms(filledFarm, generationData)
	allPaths := paths.FindAllPaths(connectedFarm)
	if len(allPaths) == 0 {
		log.Fatalf("ERROR: farm has no valid paths")
	}
	allPaths = reduceSlice(allPaths, noPathsLongerThan)

	pathsFlat := [][]*structs.Room{}
	for _, p := range allPaths {
		pathsFlat = append(pathsFlat, p.Path)
	}
	lookupMap, sortString := paths.MapPathsToStrings(pathsFlat)
	permedStrings := paths.PermutationsIter(sortString, maxPermutations)

	moveStore := moveData{}
	m := sync.Mutex{}
	wg := sync.WaitGroup{}

	// increasing the int in chunkslice will make the program run faster and use more RAM
	chunks := chunkSlice(permedStrings, permutationChunkLength)

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

// getLongestMoveset gets the number of turns for a set of moves
func getLongestMoveset(allMoves [][]string) int {
	l := 0
	for _, moveset := range allMoves {
		if len(moveset) > l {
			l = len(moveset)
		}
	}
	return l
}

// chunkSlice splits a gives slice of strings into slices of slices of size chunkSize
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

// reduceSlice takes a slice of paths and will reduce them
// such that the result has a reasonably sized set of
// the shortest paths possible
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
