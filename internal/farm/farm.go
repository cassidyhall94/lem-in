package farm

import (
	"strings"

	structs "git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// Generate farm based on generation data

func GenerateFarm(data structs.GenerationData) structs.Farm {
	rooms := []*structs.Room{}
	for i := 0; i < len(data.Rooms); i++ {
		var roomToAdd structs.Room
		if i == data.StartIndex {
			roomToAdd.IsStart = true
		} else if i == data.EndIndex {
			roomToAdd.IsEnd = true
		}
		roomToAdd.Name = data.Rooms[i]
		rooms = append(rooms, &roomToAdd)
	}

	ants := []*structs.Ant{}
	for i := 0; i < data.NumberOfAnts; i++ {
		ants = append(ants, &structs.Ant{
			Id: i,
		})
	}

	return structs.Farm{
		Ants:  ants,
		Rooms: rooms,
	}
}

// Using farm(rooms) found from generateFarm func, connect the farms(rooms) together using data(links) structs.GenerationData

// range over data.Links for each link
// find corresponding rooms in farm.rooms, save the links in a helper variable before appending
// append &RoomB to RoomA.Links and &RoomsA to RoomB.Links
// add links back into the farm.Room after appending

func ConnectRooms(farm structs.Farm, data structs.GenerationData) structs.Farm {
	for _, dataLink := range data.Links {
		for _, farmRoomA := range farm.Rooms {
			splitDataLink := strings.Split(dataLink, "-")
			if farmRoomA.Name == splitDataLink[0] {
				for _, farmRoomB := range farm.Rooms {
					if farmRoomB.Name == splitDataLink[1] {
						roomAVar := farmRoomA.Links
						roomAVar = append(roomAVar, farmRoomB)
						farmRoomA.Links = roomAVar
						roomBVar := farmRoomB.Links
						roomBVar = append(roomBVar, farmRoomA)
						farmRoomB.Links = roomBVar
					}
				}
			}
		}
	}
	return farm
}
