package dataparser

import (
	"fmt"
	structs "lem-in/structs"
	"log"
	"os"
	"strings"
)

// Loads data from the file and saves it into variable
func LoadData(fileName string) []byte {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open: %s", fileName)
	}

	// sep := []byte{13, 10}
	// transformedData := bytes.Split(data, sep)

	return data
}

// Reads and checks data from loaded data to make generation data for future farm
func ReadData(data []byte) structs.GenerationData {
	var result structs.GenerationData
	var err error

	stringdata := string(data)
	fmt.Println(stringdata)

	for _, slicedata := range stringdata {
		stringslicedata := string(slicedata)
		// splitstring := strings.SplitAfter(stringslicedata, "\n")
		datarune := stringslicedata[0]
		fmt.Print(string(datarune))
		structs.ANTCOUNTER = string(datarune)
		// fmt.Print(string(structs.ANTCOUNTER))
	}

	if err != nil {
		fmt.Printf("ANTCOUNTER error: %+v", err)
		// utils.CheckError(err)
	}

	// fmt.Println(structs.ANTCOUNTER)

	if structs.ANTCOUNTER == "" {
		log.Fatal("Invalid number of Ants!")
	}

	var startFound, endFound bool
	var commentsCounter int = 1
	for i := 1; i < len(data); i++ {
		if strings.Contains(string(data[i]), "##") {
			if string(data[i]) == "##start" {
				startFound = true
				// result.StartIndex = i - commentsCounter
			} else if string(data[i]) == "##end" {
				endFound = true
				// result.EndIndex = i - commentsCounter
			} else {
				log.Fatal("Invalid start or end data format!")
			}
			commentsCounter++
			continue
		} else if strings.Contains(string(data[i]), "#") {
			commentsCounter++
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

	if !startFound || !endFound {
		log.Fatal("Invalid data, no start or end room found")
	}

	return result
}
