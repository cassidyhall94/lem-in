package farm

import (
	"fmt"
	structs "lem-in/structs"
	utils "lem-in/utils"
	"log"
	"strconv"
	"strings"
)

var FARM []structs.Room

// Generate farm based on generation data

func GenerateFarm(data structs.GenerationData) (structs.Farm) {
	var err error
	var ANTCOUNTER int // Amount of ants to spawn
	var STARTROOMID int
	var ENDROOMID int
	for i := 0; i < len(data.Rooms); i++ {
		var roomToAdd structs.Room
		if i == data.StartIndex {
			roomToAdd.IsStart = true
			i = STARTROOMID
			roomToAdd.Ants = ANTCOUNTER
		} else if i == data.EndIndex {
			roomToAdd.IsEnd = true
			i = ENDROOMID
		}
		splitData := strings.Split(data.Rooms[i], " ")

		roomToAdd.Name = splitData[0]
		roomToAdd.X_pos, err = strconv.Atoi(splitData[1])
		if err != nil {
			fmt.Printf("Generate Farm(roomsToAdd.X_pos error: %+v", err)
		}
		utils.CheckError(err)
		roomToAdd.Y_pos, err = strconv.Atoi(splitData[2])
		if err != nil {
			fmt.Printf("Generate Farm(roomsToAdd.Y_pos error: %+v", err)
		}
		utils.CheckError(err)

		FARM = append(FARM, roomToAdd)
	}
	return structs.Farm{}
}

// Connect all rooms based on links from the file
func ConnectRooms(links []string) {
	for i := 0; i < len(links); i++ {
		splitData := strings.Split(links[i], "-")
		for j := 0; j < len(FARM); j++ {
			if FARM[j].Name == splitData[0] {
				for k := 0; k < len(FARM); k++ {
					if FARM[k].Name == splitData[1] {
						if FARM[k].Name == FARM[j].Name {
							log.Fatal("Invalid data format")
						}
						FARM[j].Links = append(FARM[j].Links, &FARM[k])
						FARM[k].Links = append(FARM[k].Links, &FARM[j])
						break
					}
					if k == len(FARM)-1 {
						log.Fatal("Invalid data format. Room not found")
					}

				}
				break
			}
			if j == len(FARM)-1 {
				log.Fatal("Invalid data format. Room not found")
			}
		}
	}
}
