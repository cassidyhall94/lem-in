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
}

// func to remove overlapping paths (except start/end rooms)
// loop over allPaths for each path, store path in helper variable and then in next loop compare path with helper variable.
// if helper variable does match path, don't return path
func TrimPaths(allPaths []*structs.PathStruct) []*structs.PathStruct {
	trimmedPaths := []*structs.PathStruct{}
	helper := []*structs.Room{}
	for _, path := range allPaths {
		if len(path.Path) == 2 {
			trimmedPaths = append(trimmedPaths, path)
			continue
		}
		pathToAppend := true
		for _, room := range path.Path {
			if contains(helper, room) {
				pathToAppend = false
				break
			}
			if !room.IsStart && !room.IsEnd {
				helper = append(helper, room)
			}
		}
		if pathToAppend {
			trimmedPaths = append(trimmedPaths, path)
		}
	}
	return trimmedPaths
}

func contains(s []*structs.Room, str *structs.Room) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func addPathToHelper(h []*structs.Room, path []*structs.Room) []*structs.Room {
	for _, r := range path {
		if !r.IsEnd && !r.IsStart {
			h = append(h, r)
		}
	}
	return h
}

type mapOfStringsToPaths map[string][]*structs.Room

func join(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}

func PermutationsRec(testStr string) []string {
	var n func(testStr []rune, p []string) []string
	n = func(testStr []rune, p []string) []string {
		if len(testStr) == 0 {
			return p
		} else {
			result := []string{}
			for _, e := range p {
				result = append(result, join([]rune(e), testStr[0])...)
			}
			return n(testStr[1:], result)
		}
	}

	output := []rune(testStr)
	return n(output[1:], []string{string(output[0])})
}

func MapPathsToStrings(allPaths [][]*structs.Room) (map[string][]*structs.Room, string) {
	lookupMap, sortString := map[string][]*structs.Room{}, ""
	letter := func(i int) string {
		return string(rune(i + 33))
	}
	for i, path := range allPaths {
		lookupMap[letter(i)] = path
		sortString = sortString + letter(i)
	}
	return lookupMap, sortString
}

func GetPathsFromStrings(lookupMap map[string][]*structs.Room, sortString string) [][]*structs.Room {
	allPaths := [][]*structs.Room{}
	for _, r := range sortString {
		allPaths = append(allPaths, lookupMap[string(r)])
	}
	return allPaths
}

func PermutationsIter(str string) []string {
	var count = 0
	ret := []string{}
	temp := []rune(str)
	if len(temp) == 1 {
		return []string{str}
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})
	index, lowest := 0, 0
	for {
		// 4 million permutions should be enough for anyone
		if count > 4000000 {
			break
		}
		count++
		for i := 0; i < len(temp)-1; i++ {
			if temp[i] < temp[i+1] {
				lowest = i
			}
		}
		index = lowest
		j := findCeiling(temp, index)
		if j == index {
			break
		}
		swap(temp, index, j)
		b := temp[index+1:]
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		a := string(temp[0:index+1]) + string(b)
		temp = []rune(a)
		ret = append(ret, string(temp))
	}
	return ret
}

func swap(r []rune, i, j int) string {
	r[i], r[j] = r[j], r[i]
	return string(r)
}

func findCeiling(temp []rune, index int) int {
	k := index
	test := temp[index]
	for k < len(temp)-1 {
		if temp[index] < temp[k+1] {
			index = k + 1
			break
		}
		k++
	}
	k = index
	for k < len(temp)-1 {
		if (temp[index] > temp[k+1]) && (temp[k+1] > test) {
			index = k + 1
		}
		k++
	}
	return index
}
