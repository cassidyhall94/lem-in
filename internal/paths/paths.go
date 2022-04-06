package paths

import (
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// make an ordered slice of rooms in PathStruct, such that the order is from first room to last room with the fewest hops (hop is moving from one linked room to another)
// func for finding all the paths from start to end using farm.room, room.Links(take in farm.rooms and loop over it and its links), and gives you []*structs.PathStruct (use make)

// https://libgen.rocks/ads.php?md5=1a699911f1094229b4d6c5df601a09ad
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

// func for finding the shortest path from all valid paths
func findAllPaths(farm structs.Farm) []*structs.PathStruct {
	pathStruct := []*structs.PathStruct{}
	for _, farmRooms := range farm.Rooms {
		for _, farmLinks := range farmRooms.Links {
			// fmt.Printf("farmLinks: %+v\n", farmLinks.Name)
			if farmLinks.IsStart {
				for _, path := range pathStruct {
					for _, pathRoom := range path.Path {
						pathRoom.Name = farmLinks.Name
					}
				}
			}
		}
	}
	return pathStruct
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
