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

// Connect all rooms based on links from the file
// func ConnectRooms(links []string) {
// 	for i := 0; i < len(links); i++ {
// 		splitData := strings.Split(links[i], "-")
// 		for j := 0; j < len(FARM); j++ {
// 			if FARM[j].Name == splitData[0] {
// 				for k := 0; k < len(FARM); k++ {
// 					if FARM[k].Name == splitData[1] {
// 						if FARM[k].Name == FARM[j].Name {
// 							log.Fatal("Invalid data format")
// 						}
// 						FARM[j].Links = append(FARM[j].Links, &FARM[k])
// 						FARM[k].Links = append(FARM[k].Links, &FARM[j])
// 						break
// 					}
// 					if k == len(FARM)-1 {
// 						log.Fatal("Invalid data format. Room not found")
// 					}

// 				}
// 				break
// 			}
// 			if j == len(FARM)-1 {
// 				log.Fatal("Invalid data format. Room not found")
// 			}
// 		}
// 	}
// }
