package farm

import (
	structs "lem-in/structs"
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
// range over farm to get room one at a time
// match Farm.Rooms.Names with GenerationData.Rooms
// if matched, store GenerationData.Links from matched room name

func ConnectRooms(farm structs.Farm, data structs.GenerationData) structs.Farm {

	return structs.Farm{}
}
