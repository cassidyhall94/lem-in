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

// func GenerateFarm(data structs.GenerationData) {
// 	var err error
// 	// var ANTCOUNTER int
// 	// var STARTROOMID int
// 	// var ENDROOMID int
// 	// var FARM []structs.Room
// 	for i := 0; i < len(data.Rooms); i++ {
// 		var roomToAdd structs.Room
// 		if i == data.StartIndex {
// 			roomToAdd.IsStart = true
// 			structs.STARTROOMID = i
// 			roomToAdd.Ants = structs.ANTCOUNTER
// 		} else if i == data.EndIndex {
// 			roomToAdd.IsEnd = true
// 			structs.ENDROOMID = i
// 		}
// 		splitData := strings.Split(data.Rooms[i], " ")

// 		roomToAdd.Name = splitData[0]
// 		roomToAdd.X_pos, err = strconv.Atoi(splitData[1])
// 		utils.CheckError(err)
// 		roomToAdd.Y_pos, err = strconv.Atoi(splitData[2])
// 		utils.CheckError(err)

// 		structs.FARM = append(structs.FARM, roomToAdd)
// 	}
// 	ConnectRooms(data.Links)
// }

// Using farm(rooms) found from generateFarm func, connect the farms(rooms) together using data(links) structs.GenerationData
// range over farm to get room one at a time
// match Farm.Rooms.Names with GenerationData.Rooms
// if matched, store GenerationData.Links from matched room name

func ConnectRooms(farm structs.Farm, data structs.GenerationData) structs.Farm {

	return structs.Farm{}
}

// func ConnectRooms(links []string) {
// 	// var FARM []structs.Room
// 	for i := 0; i < len(links); i++ {
// 		splitData := strings.Split(links[i], "-")
// 		for j := 0; j < len(structs.FARM); j++ {
// 			if structs.FARM[j].Name == splitData[0] {
// 				for k := 0; k < len(structs.FARM); k++ {
// 					if structs.FARM[k].Name == splitData[1] {
// 						if structs.FARM[k].Name == structs.FARM[j].Name {
// 							log.Fatal("Invalid data format")
// 						}
// 						structs.FARM[j].Links = append(structs.FARM[j].Links, &structs.FARM[k])
// 						structs.FARM[k].Links = append(structs.FARM[k].Links, &structs.FARM[j])
// 						break
// 					}
// 					if k == len(structs.FARM)-1 {
// 						log.Fatal("Invalid data format. Room not found")
// 					}

// 				}
// 				break
// 			}
// 			if j == len(structs.FARM)-1 {
// 				log.Fatal("Invalid data format. Room not found")
// 			}
// 		}
// 	}
// }
