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

// func LoadData(fileName string) [][]byte {
// 	data, err := os.ReadFile(os.Args[1])

// 	if err != nil {
// 		log.Fatalf("failed to open: %s", fileName)
// 	}

// 	separate := []byte{13, 10}
// 	changeData := bytes.Split(data, separate)

// 	return changeData
// }

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

// func ReadData(data [][]byte) structs.GenerationData {
// 	var result structs.GenerationData

// 	var err error
// 	structs.ANTCOUNTER, err = strconv.Atoi(string(data[0]))
// 	utils.CheckError(err)

// 	if structs.ANTCOUNTER <= 0 {
// 		log.Fatal("Invalid number of Ants!")
// 	}

// 	var start, end bool
// 	var comments int = 1
// 	for i := 1; i < len(data); i++ {
// 		if strings.Contains(string(data[i]), "##") {
// 			if string(data[i]) == "##start" {
// 				start = true
// 				result.StartIndex = i - comments
// 			} else if string(data[i]) == "##end" {
// 				end = true
// 				result.EndIndex = i - comments
// 			} else {
// 				log.Fatal("Invalid start or end data format!")
// 			}
// 			comments++
// 			continue
// 		} else if strings.Contains(string(data[i]), "#") {
// 			comments++
// 			continue
// 		}

// 		if strings.Count(string(data[i]), " ") == 2 {
// 			result.Rooms = append(result.Rooms, string(data[i]))
// 		} else if strings.Count(string(data[i]), "-") == 1 {
// 			result.Links = append(result.Links, string(data[i]))
// 		} else {
// 			log.Fatal("Invalid link data format!")
// 		}
// 	}

// 	if !start || !end {
// 		log.Fatal("Invalid data format, no start or end room found")
// 	}

// 	return result
// }
