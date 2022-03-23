package paths

import (
	"lem-in/structs"
)

// Finds all paths from start to end
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
