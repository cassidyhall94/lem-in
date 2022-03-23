package utils

import (
	"fmt"
	structs "lem-in/structs"
)

func CheckError(err error) {
	if err != nil {
		fmt.Printf("Invalid Data Format err:%+v", err)
	}
}

// Reset farm to defaults
func ResetFarm() {
	for i := 0; i < len(structs.FARM); i++ {
		if (structs.FARM)[i].IsStart {
			(structs.FARM)[i].Ants = (structs.ANTCOUNTER)
		} else {
			(structs.FARM)[i].Ants = 0
		}
	}
}

// Sort paths by length
func SortPathsByLength(paths *[][]*structs.Room) {
	l := len(*paths)
	for i := 0; i < l; i++ {
		min := i
		for j := i; j < l; j++ {
			if len((*paths)[j]) < len((*paths)[min]) {
				min = j
			}
		}
		(*paths)[i], (*paths)[min] = (*paths)[min], (*paths)[i]
	}
}

// Decodes paths list into room names
func DecodePaths(paths [][]*structs.Room) {
	fmt.Println()
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths[i]); j++ {
			fmt.Print(paths[i][j].Name)
			if j != len(paths[i])-1 {
				fmt.Print(" -> ")
			} else {
				fmt.Println()
			}
		}
	}
}
