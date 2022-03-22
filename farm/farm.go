package farm

import (
	structs "lem-in/structs"
	utils "lem-in/utils"
	"strconv"
	"strings"
)

// Generate farm based on generation data
func GenerateFarm(data structs.GenerationData) {
	var err error
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
}

// Connect all rooms based on links from the file
