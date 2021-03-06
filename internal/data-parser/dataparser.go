package dataparser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	structs "git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// LoadData loads data from the file and saves it into a variable
func LoadData(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("file open error: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ret := []string{}

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("file scanner error: %w", err)
	}

	return ret, nil
}

// ReadData checks data from LoadData to make GenerationData for future farm
func ReadData(fileLines []string) (structs.GenerationData, error) {
	ants, rooms, links, start, end := 0, []string{}, []string{}, 0, 0
	foundStart, foundEnd := false, false
	for i, fileLine := range fileLines {
		// Ignore blank lines
		if len(fileLine) == 0 {
			continue
		}

		// Do nothing with comments (#) unless its start/end
		if string(fileLine[0]) == "#" {
			if fileLine == "##start" {
				foundStart = true
			} else if fileLine == "##end" {
				foundEnd = true
			}
			continue
		}

		a, err := strconv.Atoi(fileLine)
		if i == 0 && a == 0 {
			fmt.Println("ERROR: invalid number of ants")
		}
		if err == nil {
			ants = a
			continue
		} else if err != nil && i == 0 {
			return structs.GenerationData{}, fmt.Errorf("ERROR: parsing ants, got %s wanted int: %w", fileLine, err)
		}

		maybeRoom := strings.Split(fileLine, " ")
		if len(maybeRoom) == 3 {
			rooms = append(rooms, maybeRoom[0])
			if foundStart {
				start = len(rooms) - 1
				foundStart = false
			}
			if foundEnd {
				end = len(rooms) - 1
				foundEnd = false
			}
			continue
		}
		maybeLink := strings.Contains(fileLine, "-")
		if maybeLink {
			links = append(links, fileLine)
		}
	}
	return structs.GenerationData{
		NumberOfAnts: ants,
		Rooms:        rooms,
		Links:        links,
		StartIndex:   start,
		EndIndex:     end,
	}, nil
}
