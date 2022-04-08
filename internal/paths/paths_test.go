package paths

import (
	"reflect"
	"testing"

	dataparser "git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/data-parser"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/farm"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

func Test_findAllPaths(t *testing.T) {
	tests := []struct {
		name    string
		fixture string
		want    []*structs.PathStruct
	}{
		{
			name: "pass - easy",
			fixture: "../data-parser/fixtures/test4.txt",
			want: []*structs.PathStruct{
				{
					Path: []*structs.Room{
						{
							Name:    "0",
							IsStart: true,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "3",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "1",
							IsEnd: true,
						},
					},
				},
			},
		// },
		// {
		// 	name: "pass - more paths",
		// 	fixture: "../data-parser/fixtures/test5.txt",
		// 	want: []*structs.PathStruct{
		// 		{
		// 			Path: []*structs.Room{
		// 				{
		// 					Name:    "0",
		// 					IsStart: true,
		// 				},
		// 				{
		// 					Name:    "2",
		// 					IsStart: false,
		// 					IsEnd:   false,
		// 				},
		// 				{
		// 					Name:    "3",
		// 					IsStart: false,
		// 					IsEnd:   false,
		// 				},
		// 				{
		// 					Name:  "1",
		// 					IsEnd: true,
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Path: []*structs.Room{
		// 				{
		// 					Name:    "0",
		// 					IsStart: true,
		// 				},
		// 				{
		// 					Name:    "2",
		// 					IsStart: false,
		// 					IsEnd:   false,
		// 				},
		// 				{
		// 					Name:    "3",
		// 					IsStart: false,
		// 					IsEnd:   false,
		// 				},
		// 				{
		// 					Name:    "4",
		// 					IsStart: false,
		// 					IsEnd:   false,
		// 				},
		// 				{
		// 					Name:  "1",
		// 					IsEnd: true,
		// 				},
		// 			},
		// 		},
		// 	},
		// },
		// {
		// 	name: "empty",
		// 	fixture: "../data-parser/fixtures/test3.txt",
		// 	want: []*structs.PathStruct{},
		// },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := dataparser.LoadData(tt.fixture)
			if err != nil {
				t.Errorf("dataparser.LoadData error: %+v", err)
			}
			generationData := dataparser.ReadData(data)
			filledFarm := farm.GenerateFarm(generationData)
			linkedFarm := farm.ConnectRooms(filledFarm, generationData)

			got := findAllPaths(linkedFarm)
			sliceOfAllGotRooms := [][]string{}
			for _, gotRooms := range got {
				sliceOfAllGotRooms = append(sliceOfAllGotRooms, getSliceOfRoomNames(gotRooms.Path))
			}
			if len(got) < len(tt.want) {
				t.Fatalf("findAllPaths() returned %d paths, wanted %d, here's all the paths we got: %+v", len(got), len(tt.want), sliceOfAllGotRooms)
			}
			for _, gotRooms := range got {
				foundMatchingPath := false
				gotRoomNames := getSliceOfRoomNames(gotRooms.Path)
				for _, wantRooms := range tt.want {
					wantRoomNames := getSliceOfRoomNames(wantRooms.Path)
					if !reflect.DeepEqual(gotRoomNames, wantRoomNames) {
						t.Logf("wanted %+v, here's all the paths we got: %+v", wantRoomNames, sliceOfAllGotRooms)
					} else {
						foundMatchingPath = true
						break
					}
				}
				if !foundMatchingPath {
					t.Error("unable to find matching path for gotPath, see log lines.")
				}
			}
		})
	}
}

func getSliceOfRoomNames(rooms []*structs.Room) []string {
	ret := make([]string, 0, len(rooms))
	for _, r := range rooms {
		ret = append(ret, r.Name)
	}
	return ret
}
