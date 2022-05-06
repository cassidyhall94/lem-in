package paths

import (
	"sort"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// Find all paths from start to end
func FindAllPaths(farm structs.Farm) []*structs.PathStruct {
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

// func findValidPath looks in a room for its links, it then asks itself about those links returning a slice of visited rooms and a slice of paths
// setting the returned bool indicates that we're rolling back up the tree or graph
func findValidPath(room *structs.Room, visited []string, paths [][]*structs.Room, path []*structs.Room) ([]string, [][]*structs.Room, []*structs.Room, bool) {
	path = append(path, room)
	visited = append(visited, room.Name)

	if room.IsEnd {
		return nil, append(paths, path), nil, true
	}

	rollup := false

	for _, l := range room.Links {
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
			paths = ps
		}
	}

	return visited, paths, path, rollup
}

func getSliceOfRoomNames(rooms []*structs.Room) []string {
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

func FindShortestPath(allPaths []*structs.PathStruct) *structs.PathStruct {
	p := -1
	shortestPath := &structs.PathStruct{}
	for _, path := range allPaths {
		if p == -1 {
			p = len(path.Path)
		} else if len(path.Path) < p {
			p = len(path.Path)
			shortestPath = path
		}
	}
	return shortestPath
}

func SortPaths(allPaths []*structs.PathStruct) []*structs.PathStruct {
	sort.Slice(allPaths, func(i, j int) bool {
		shortToLong := len(allPaths[i].Path) < len(allPaths[j].Path)
		return shortToLong
	})
	return allPaths
}

func TrimPaths(allPaths []*structs.PathStruct) []*structs.PathStruct {
	return []*structs.PathStruct{}
}
