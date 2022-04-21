package paths

import (
	"reflect"
	"testing"

	dataparser "git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/data-parser"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/farm"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

func TestFindAllPaths(t *testing.T) {
	tests := []struct {
		name    string
		fixture string
		want    []*structs.PathStruct
	}{
		{
			name:    "pass - easy",
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
		},
		{
			name:    "pass - more paths",
			fixture: "../data-parser/fixtures/test5.txt",
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
							Name:    "4",
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
		},
		{
			name:    "empty",
			fixture: "../data-parser/fixtures/test3.txt",
			want:    []*structs.PathStruct{},
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

			got := FindAllPaths(linkedFarm)
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

func TestFindShortestPath(t *testing.T) {
	rooms := []structs.Room{
		{
			Name: "0",
		},
		{
			Name: "1",
		},
		{
			Name: "2",
		},
		{
			Name: "3",
		},
		{
			Name: "4",
		},
		{
			Name: "5",
		},
		{
			Name: "6",
		},
		{
			Name: "7",
		},
		{
			Name: "8",
		},
		{
			Name: "9",
		},
	}
	type args struct {
		allPaths []*structs.PathStruct
	}
	tests := []struct {
		name string
		args args
		want *structs.PathStruct
	}{
		{
			name: "pass - easy",
			args: args{
				allPaths: []*structs.PathStruct{
					{
						Path: []*structs.Room{
							&rooms[0],
							&rooms[3],
							&rooms[2],
							&rooms[1],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[0],
							&rooms[1],
						},
					},
				},
			},
			want: &structs.PathStruct{
				Path: []*structs.Room{
					&rooms[0],
					&rooms[1],
				},
			},
		},
		{
			name: "pass - hard",
			args: args{
				allPaths: []*structs.PathStruct{
					{
						Path: []*structs.Room{
							&rooms[2],
							&rooms[4],
							&rooms[3],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[2],
							&rooms[4],
							&rooms[5],
							&rooms[6],
							&rooms[7],
							&rooms[3],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[2],
							&rooms[4],
							&rooms[5],
							&rooms[6],
							&rooms[7],
							&rooms[8],
							&rooms[9],
							&rooms[3],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[2],
							&rooms[3],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[2],
							&rooms[4],
							&rooms[3],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[2],
							&rooms[4],
							&rooms[5],
							&rooms[6],
							&rooms[7],
							&rooms[8],
							&rooms[9],
							&rooms[0],
							&rooms[3],
						},
					},
				},
			},
			want: &structs.PathStruct{
				Path: []*structs.Room{
					&rooms[2],
					&rooms[3],
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindShortestPath(tt.args.allPaths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findShortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
