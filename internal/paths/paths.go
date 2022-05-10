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

func GetSliceOfRoomNames(rooms []*structs.Room) []string {
	ret := []string{}
	for _, r := range rooms {
		ret = append(ret, r.Name)
	}
	return ret
}

func GetSliceOfPathNames(paths []*structs.PathStruct) []string {
	ret := []string{}
	for _, path := range paths {
		for _, pathInfo := range path.Path {
			ret = append(ret, pathInfo.Name)
		}
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

func SortPaths(allPaths []*structs.PathStruct) []*structs.PathStruct {
	sort.Slice(allPaths, func(i, j int) bool {
		shortToLong := len(allPaths[i].Path) < len(allPaths[j].Path)
		return shortToLong
	})
	return allPaths
	// allPaths:
	// [start h n e end]
	// [start h n m end]
	// [start t E a m end]
	// [start 0 o n m end]
	// [start h A c k end]
	// [start 0 o n e end]
	// [start t E a m n e end]
	// [start 0 o n h A c k end]
	// [start t E a m n h A c k end]
}

// func to remove overlapping paths (except start/end rooms)
// loop over allPaths for each path, store path in helper variable and then in next loop compare path with helper variable.
// if helper variable does match path, don't return path
func TrimPaths(allPaths []*structs.PathStruct) []*structs.PathStruct {
	trimmedPaths := []*structs.PathStruct{}
	helper := []*structs.Room{}
	med := calcMedianPathLen(allPaths)
	allPaths = rearrangePathsOnTheMedian(allPaths, med)
	for _, path := range allPaths {
		pathToAppend := true
		var doesContain bool
		if doesContain, helper = contains(helper, path.Path); doesContain {
			pathToAppend = false
		}
		if pathToAppend {
			trimmedPaths = append(trimmedPaths, path)
			continue

		}
	}
	return trimmedPaths
}

func contains(s []*structs.Room, path []*structs.Room) (bool, []*structs.Room) {
	for _, v := range s {
		for _, r := range path {
			if v.Name == r.Name {
				return true, s
			}
		}
	}

	return false, addPathToHelper(s, path)
}

func addPathToHelper(h []*structs.Room, path []*structs.Room) []*structs.Room {
	for _, r := range path {
		if !r.IsEnd && !r.IsStart {
			h = append(h, r)
		}
	}
	return h
}

func calcMedianPathLen(allPaths []*structs.PathStruct) int {
	totalLengths := 0
	for _, ps := range allPaths {
		totalLengths = totalLengths + len(ps.Path)
	}
	if totalLengths == 0 || len(allPaths) == 0 {
		return 0
	}

	return totalLengths / len(allPaths)
}

func rearrangePathsOnTheMedian(paths []*structs.PathStruct, med int) []*structs.PathStruct {
	ret := []*structs.PathStruct{}
	didAThing := false
	for i, ps := range paths {
		if len(ps.Path) == med {
			didAThing = true
			ret = append(ret, paths[i:]...)
			ret = append(ret, paths[:i]...)
			break
		}
	}
	if didAThing == false {
		return paths
	}
	return ret
}
