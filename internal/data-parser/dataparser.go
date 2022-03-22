package dataparser

import (
	"bufio"
	"fmt"
	structs "lem-in/structs"
	"os"
	"strconv"
	"strings"
)

// Loads data from the file and saves it into variable
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

// Reads and checks data from loaded data to make generation data for future farm
func ReadData(fileLines []string) structs.GenerationData {
	ants, rooms, links, start, end := 0, []string{}, []string{}, 0, 0
	foundStart, foundEnd := false, false
	for _, fileLine := range fileLines {
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
		if err == nil {
			ants = a
			continue
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

		// is link
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
	}
}
