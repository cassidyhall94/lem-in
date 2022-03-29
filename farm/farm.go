package farm

import (
	// "fmt"
	structs "lem-in/structs"
	utils "lem-in/utils"
	"log"
	"strconv"
	"strings"
)

// Generate farm based on generation data
func GenerateFarm(data structs.GenerationData) {
	var err error
	// var ANTCOUNTER int
	// var STARTROOMID int
	// var ENDROOMID int
	// var FARM []structs.Room
	for i := 0; i < len(data.Rooms); i++ {
		var roomToAdd structs.Room
		if i == data.StartIndex {
			roomToAdd.IsStart = true
			structs.STARTROOMID = i
			roomToAdd.Ants = structs.ANTCOUNTER
		} else if i == data.EndIndex {
			roomToAdd.IsEnd = true
			structs.ENDROOMID = i
		}
		splitData := strings.Split(data.Rooms[i], " ")

		roomToAdd.Name = splitData[0]
		roomToAdd.X_pos, err = strconv.Atoi(splitData[1])
		utils.CheckError(err)
		roomToAdd.Y_pos, err = strconv.Atoi(splitData[2])
		utils.CheckError(err)

		structs.FARM = append(structs.FARM, roomToAdd)
	}
	ConnectRooms(data.Links)
}

// Connect all rooms based on links from the file
func ConnectRooms(links []string) {
	// var FARM []structs.Room
	for i := 0; i < len(links); i++ {
		splitData := strings.Split(links[i], "-")
		for j := 0; j < len(structs.FARM); j++ {
			if structs.FARM[j].Name == splitData[0] {
				for k := 0; k < len(structs.FARM); k++ {
					if structs.FARM[k].Name == splitData[1] {
						if structs.FARM[k].Name == structs.FARM[j].Name {
							log.Fatal("Invalid data format")
						}
						structs.FARM[j].Links = append(structs.FARM[j].Links, &structs.FARM[k])
						structs.FARM[k].Links = append(structs.FARM[k].Links, &structs.FARM[j])
						break
					}
					if k == len(structs.FARM)-1 {
						log.Fatal("Invalid data format. Room not found")
					}

				}
				break
			}
			if j == len(structs.FARM)-1 {
				log.Fatal("Invalid data format. Room not found")
			}
		}
	}
}
