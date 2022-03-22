package utils

import (
	"fmt"
	structs "lem-in/structs"
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal("Invalid data format!")
	}
}

// Decodes paths list into room names
func DecodePaths(paths [][]*structs.Room) {
	fmt.Println()
	for b := 0; b < len(paths); b++ {
		for c := 0; c < len(paths[b]); c++ {
			fmt.Print(paths[b][c].Name)
			if c != len(paths[b])-1 {
				fmt.Print(" -> ")
			} else {
				fmt.Println()
			}
		}
	}
}
