package dataparser

import (
	"bytes"
	"lem-in/structs"
	"lem-in/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

// Load data from the file and save it into var
func LoadData(fileName string) [][]byte {
	data, err := os.ReadFile(os.Args[1])

	if err != nil {
		log.Fatalf("failed to open: %s", fileName)
	}

	separate := []byte{13, 10}
	changeData := bytes.Split(data, separate)

	return changeData
}

// Read and check data from loaded data
func ReadData(data [][]byte) structs.GenerationData {
	var result structs.GenerationData

	var err error
	structs.ANTCOUNTER, err = strconv.Atoi(string(data[0]))
	utils.CheckError(err)

	if structs.ANTCOUNTER <= 0 {
		log.Fatal("Invalid number of Ants!")
	}

	var start, end bool
	var comments int = 1
	for i := 1; i < len(data); i++ {
		if strings.Contains(string(data[i]), "##") {
			if string(data[i]) == "##start" {
				start = true
				result.StartIndex = i - comments
			} else if string(data[i]) == "##end" {
				end = true
				result.EndIndex = i - comments
			} else {
				log.Fatal("Invalid start or end data format!")
			}
			comments++
			continue
		} else if strings.Contains(string(data[i]), "#") {
			comments++
			continue
		}

		if strings.Count(string(data[i]), " ") == 2 {
			result.Rooms = append(result.Rooms, string(data[i]))
		} else if strings.Count(string(data[i]), "-") == 1 {
			result.Links = append(result.Links, string(data[i]))
		} else {
			log.Fatal("Invalid link data format!")
		}
	}

	if !start || !end {
		log.Fatal("Invalid data format, no start or end room found")
	}

	return result
}
