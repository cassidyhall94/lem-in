package paths

import (
	"fmt"
	"os"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// make an ordered slice of rooms in PathStruct, such that the order is from first room to last room with the fewest hops (hop is moving from one linked room to another)
// func for finding all the paths from start to end using farm.room, room.Links(take in farm.rooms and loop over it and its links), and gives you []*structs.PathStruct (use make)

// https://edu.anarcho-copy.org/Algorithm/grokking-algorithms-illustrated-programmers-curious.pdf
// https://en.wikipedia.org/wiki/Edmonds%E2%80%93Karp_algorithm
// https://medium.com/@jamierobertdawson/lem-in-finding-all-the-paths-and-deciding-which-are-worth-it-2503dffb893
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Practical_optimizations_and_infinite_graphs
// https://kalkicode.com/print-all-the-paths-from-the-source-to-destination-in-directed-graph-in-go
// https://github.com/e-tinkers/shortest-path-algorithm/blob/master/shortest.go

// see if rooms has been visited
//  node = room
// edge = link
// append rooms to the var if they are not the end room
// we have an undirected graph

// Find all paths from start to end
func findAllPaths(farm structs.Farm) []*structs.PathStruct {
	pathStruct := []*structs.PathStruct{}

	for _, farmRoom := range farm.Rooms {
		if farmRoom.IsStart {
			_, paths, _, _ := findValidPath(farmRoom, []string{}, [][]*structs.Room{}, []*structs.Room{})
			for _, p := range paths {
				pathStruct = append(pathStruct, &structs.PathStruct{Path: p})
			}
		}
	}
	return pathStruct
}

var counter int

// find valid path looks in a room for its links, it then asks itself about those links returning a slice of visited rooms and a slice of paths
// setting the returned bool indicates that we're winding up

// we pass in a room
// if it hasn't been visited we need to look at its links
// we want to ignore links that lead to a dead end
// we want to add a new path to paths and return from everything when we find the end
// when adding to visited or path we need to make copies of them to make sure we don't overwrite the underlying path/visited slice
// we want to call findValidPath on each link that isn't a dead end
// eventually this will return paths that is a slice of each path
func findValidPath(room *structs.Room, visited []string, paths [][]*structs.Room, path []*structs.Room) ([]string, [][]*structs.Room, []*structs.Room, bool) {
	counter++
	fmt.Printf("room.Name %s; len(room.Links) %d, room.Links[0] %s, visited %+v\n", room.Name, len(room.Links), room.Links[0].Name, visited)
	if counter > 20 {
		os.Exit(1)
	}

	path = append(path, room)
	fmt.Printf("path: %s\n", getAllNamesFromSliceOfRooms(path))

	visited = append(visited, room.Name)
	// fmt.Printf("visited: %s\n", visited)

	if room.IsEnd {
		return nil, append(paths, path), nil, true
	}

	rollup := false

	for _, l := range room.Links {
		// find out if this is a dead end
		if len(l.Links) == 1 && l.Links[0].Name == room.Name && !l.IsEnd {
			continue
		} else if visitedRoom(visited, l) {
			continue
		}

		newVisited := make([]string, len(visited))
		copy(newVisited, visited)
		newPath := make([]*structs.Room, len(path))
		copy(newPath, path)

		_, ps, _, r := findValidPath(l, newVisited, paths, newPath)
		if r {
			rollup = r
			paths = append(paths, ps...)
		}
	}
	return visited, paths, path, rollup
}

func getAllNamesFromSliceOfRooms(rooms []*structs.Room) []string {
	ret := []string{}
	for _, r := range rooms {
		ret = append(ret, r.Name)
	}
	return ret
}

func visitedRoom(visited []string, room *structs.Room) bool {
	for _, roomsVisited := range visited {
		if roomsVisited == room.Name {
			return true
		}
	}
	return false
}

// Find all paths from start to end
func FindPaths(path []*structs.Room, currentRoom structs.Room, step int, paths *[][]*structs.Room, previousRoom *structs.Room) {
	if currentRoom.IsEnd {
		var skip bool
		for i := 0; i < len(path); i++ {
			if path[i].IsStart {
				skip = true
				break
			}
		}

		if len(*paths) == 0 {
			*paths = append((*paths), nil)
		} else if (*paths)[len(*paths)-1] != nil {
			*paths = append((*paths), nil)
		}

		for i := 0; i < len(path); i++ {
			if !skip {
				(*paths)[len(*paths)-1] = append((*paths)[len(*paths)-1], path[i])
			} else {
				break
			}
		}
	}

	for i := 0; i < len(currentRoom.Links); i++ {
		var carryOn bool

		for j := 0; j < len(path); j++ {
			if path[j].Name == currentRoom.Links[i].Name {
				carryOn = true
				break
			}
		}

		if !carryOn {
			pathToGo := path
			pathToGo = append(pathToGo, currentRoom.Links[i])
			FindPaths(pathToGo, *currentRoom.Links[i], step+1, paths, &currentRoom)
			pathToGo = path
		}
	}

	for i := 0; i < len(*paths); i++ {
		if (*paths)[i] == nil {
			*paths = append((*paths)[:i], (*paths)[i+1:]...)
		}
	}
}

// Find all possible combinations
func SearchCombinations(paths [][]*structs.Room) [][][]*structs.Room {
	var result [][][]*structs.Room

	for i := 0; i < len(paths); i++ {
		var findComb [][]*structs.Room
		findComb = append(findComb, paths[i])
		for j := i + 1; j < len(paths); j++ {
			if !Divide(paths[i][:len(paths[i])-1], paths[j][:len(paths[j])-1]) &&
				!DivideWithinCombinations(paths[j][:len(paths[j])-1], findComb) {
				result = append(result, findComb)
				findComb = append(findComb, paths[j])
			}
		}

		result = append(result, findComb)
	}

	return result
}

// Check division between two paths
func Divide(currentPath, pathToCheck []*structs.Room) bool {
	for i := 0; i < len(currentPath); i++ {
		for j := i + 1; j < len(pathToCheck); j++ {
			if currentPath[i].Name == pathToCheck[j].Name {
				return true
			}
		}
	}
	return false
}

// Check division between existing SearchCombinations and paths
func DivideWithinCombinations(path []*structs.Room, findComb [][]*structs.Room) bool {
	for i := 0; i < len(findComb); i++ {
		for j := 0; j < len(path); j++ {
			for k := 0; k < len(findComb[i]); k++ {
				if path[j] == findComb[i][k] {
					return true
				}
			}
		}
	}
	return false
}

// Find best combination and calculate best path for each ant
func SearchBestCombination(r [][][]*structs.Room) [][]*structs.Room {
	var topScore int
	var topPath [][]*structs.Room

	var ANTCOUNTER int
	for _, paths := range r {
		var pathCombination [][]*structs.Room
		antPosition := make([]int, len(paths))
		currentIndex := 0
		var nextPathId int
		var updateNextPathId bool = true
		for i := 0; i < ANTCOUNTER; i++ {
			if i == 0 {
				pathCombination = append(pathCombination, paths[0])
				currentIndex = 0
				antPosition[currentIndex]++
				continue
			}
			for {
				if updateNextPathId {
					if len(paths) == currentIndex+1 {
						nextPathId = 0
					} else {
						nextPathId = currentIndex + 1
					}
					updateNextPathId = false
				}
				if len(paths) == 1 || paths[currentIndex][0].IsEnd {
					pathCombination = append(pathCombination, paths[currentIndex])
					antPosition[currentIndex]++
					break
				}
				if antPosition[currentIndex]+len(paths[currentIndex]) <= len(paths[nextPathId])+antPosition[nextPathId] {
					pathCombination = append(pathCombination, paths[currentIndex])
					antPosition[currentIndex]++
					break
				} else {
					pathCombination = append(pathCombination, paths[nextPathId])
					antPosition[nextPathId]++
					currentIndex = nextPathId
					updateNextPathId = true
					break
				}
			}
		}

		var currentScore int
		for i := 0; i < len(paths); i++ {
			temp := antPosition[i] + len(paths[i])
			if currentScore == 0 || temp > currentScore {
				currentScore = temp
			}
		}
		if topScore == 0 || currentScore < topScore {
			topScore = currentScore
			topPath = pathCombination
		}
	}

	return topPath
}
