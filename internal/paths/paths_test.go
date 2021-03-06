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
			name:    "pass - hard",
			fixture: "../data-parser/fixtures/test0.txt",
			want: []*structs.PathStruct{
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "4",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "3",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "3",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "4",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "3",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "4",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "4",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
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
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
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
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
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
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "4",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
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
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "3",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "4",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "3",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
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
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
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
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "3",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "4",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "3",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "4",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
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
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "7",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "4",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
							IsEnd: true,
						},
					},
				},
				{
					Path: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "4",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "3",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "5",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:    "6",
							IsStart: false,
							IsEnd:   false,
						},
						{
							Name:  "0",
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
			generationData, err := dataparser.ReadData(data)
			if err != nil {
				t.Errorf("ERROR: invalid data format, %v", err)
			}
			filledFarm := farm.GenerateFarm(generationData)
			linkedFarm := farm.ConnectRooms(filledFarm, generationData)

			got := FindAllPaths(linkedFarm)
			sliceOfAllGotRooms := [][]string{}
			for _, gotRooms := range got {
				sliceOfAllGotRooms = append(sliceOfAllGotRooms, GetSliceOfRoomNames(gotRooms.Path))
			}
			if len(got) < len(tt.want) {
				t.Fatalf("findAllPaths() returned %d paths, wanted %d, here's all the paths we got: %+v", len(got), len(tt.want), sliceOfAllGotRooms)
			}

			for _, wantPath := range tt.want {
				foundMatchingPath := false
				wantPathNames := GetSliceOfRoomNames(wantPath.Path)
				gotPathNamesTotal := [][]string{}
				for _, gotPath := range got {
					gotPathNames := GetSliceOfRoomNames(gotPath.Path)
					gotPathNamesTotal = append(gotPathNamesTotal, gotPathNames)
					if reflect.DeepEqual(wantPathNames, gotPathNames) {
						foundMatchingPath = true
						break
					}
				}
				if !foundMatchingPath {
					t.Errorf("unable to find path %+v in got: %+v", wantPathNames, gotPathNamesTotal)
				}
			}
		})
	}
}

func TestSortPaths(t *testing.T) {
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
		want []int
	}{
		{
			name: "pass - easy",
			args: args{
				allPaths: []*structs.PathStruct{
					{
						Path: []*structs.Room{
							&rooms[0],
							&rooms[2],
							&rooms[3],
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
			want: []int{
				2,
				4,
			},
		},
		{
			name: "pass - hard",
			args: args{
				allPaths: []*structs.PathStruct{
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[3],
							&rooms[4],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[3],
							&rooms[4],
							&rooms[2],
							&rooms[5],
							&rooms[6],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[3],
							&rooms[4],
							&rooms[2],
							&rooms[7],
							&rooms[6],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[3],
							&rooms[4],
							&rooms[7],
							&rooms[6],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[3],
							&rooms[4],
							&rooms[7],
							&rooms[2],
							&rooms[5],
							&rooms[6],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[3],
							&rooms[5],
							&rooms[4],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[3],
							&rooms[5],
							&rooms[2],
							&rooms[4],
							&rooms[7],
							&rooms[6],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[3],
							&rooms[5],
							&rooms[2],
							&rooms[7],
							&rooms[6],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[3],
							&rooms[5],
							&rooms[2],
							&rooms[7],
							&rooms[4],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[3],
							&rooms[5],
							&rooms[6],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[3],
							&rooms[5],
							&rooms[6],
							&rooms[7],
							&rooms[2],
							&rooms[4],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[3],
							&rooms[5],
							&rooms[6],
							&rooms[7],
							&rooms[4],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[2],
							&rooms[5],
							&rooms[3],
							&rooms[4],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[2],
							&rooms[5],
							&rooms[3],
							&rooms[4],
							&rooms[7],
							&rooms[6],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[2],
							&rooms[5],
							&rooms[6],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[2],
							&rooms[5],
							&rooms[6],
							&rooms[7],
							&rooms[4],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[2],
							&rooms[4],
							&rooms[0],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[1],
							&rooms[2],
							&rooms[4],
							&rooms[3],
							&rooms[5],
							&rooms[6],
							&rooms[0],
						},
					},
				},
			},
			want: []int{
				4,
				4,
				5,
				5,
				5,
				6,
				6,
				7,
				7,
				7,
				7,
				7,
				7,
				7,
				8,
				8,
				8,
				8,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, gotPath := range SortPaths(tt.args.allPaths) {
				if len(gotPath.Path) != tt.want[i] {
					t.Errorf("SortPaths() = %v, want %v", len(gotPath.Path), tt.want[i])
				}
			}
		})
	}
}

func TestTrimPaths(t *testing.T) {

	rooms := []structs.Room{
		{
			Name:    "0",
			IsStart: true,
		},
		{
			Name:  "1",
			IsEnd: true,
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
	}
	room := []structs.Room{
		{
			Name:    "start",
			IsStart: true,
		},
		{
			Name: "h",
		},
		{
			Name: "n",
		},
		{
			Name: "e",
		},
		{
			Name:  "end",
			IsEnd: true,
		},
		{
			Name: "m",
		},
		{
			Name: "t",
		},
		{
			Name: "E",
		},
		{
			Name: "a",
		},
		{
			Name: "0",
		},
		{
			Name: "o",
		},
		{
			Name: "A",
		},
		{
			Name: "c",
		},
		{
			Name: "k",
		},
	}
	type args struct {
		allPaths []*structs.PathStruct
	}
	tests := []struct {
		name string
		args args
		want []*structs.PathStruct
	}{
		{
			name: "pass - easy",
			args: args{
				allPaths: []*structs.PathStruct{
					{
						Path: []*structs.Room{
							&rooms[0],
							&rooms[1],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[0],
							&rooms[2],
							&rooms[3],
							&rooms[1],
						},
					},
					{
						Path: []*structs.Room{
							&rooms[0],
							&rooms[2],
							&rooms[3],
							&rooms[4],
							&rooms[1],
						},
					},
				},
			},
			want: []*structs.PathStruct{
				{
					Path: []*structs.Room{
						&rooms[0],
						&rooms[1],
					},
				},
				{
					Path: []*structs.Room{
						&rooms[0],
						&rooms[2],
						&rooms[3],
						&rooms[1],
					},
				},
			},
		},
		{
			name: "pass - hard",
			args: args{
				allPaths: []*structs.PathStruct{
					{
						//[start h n e end]
						Path: []*structs.Room{
							&room[0],
							&room[1],
							&room[2],
							&room[3],
							&room[4],
						},
					},
					{
						//[start h n m end]
						Path: []*structs.Room{
							&room[0],
							&room[1],
							&room[2],
							&room[5],
							&room[4],
						},
					},
					{
						//[start t E a m end]
						Path: []*structs.Room{
							&room[0],
							&room[6],
							&room[7],
							&room[8],
							&room[5],
							&room[4],
						},
					},
					{
						//[start 0 o n m end]
						Path: []*structs.Room{
							&room[0],
							&room[9],
							&room[10],
							&room[2],
							&room[5],
							&room[4],
						},
					},
					{
						//[start h A c k end]
						Path: []*structs.Room{
							&room[0],
							&room[1],
							&room[11],
							&room[12],
							&room[13],
							&room[4],
						},
					},
					{
						//[start 0 o n e end]
						Path: []*structs.Room{
							&room[0],
							&room[9],
							&room[10],
							&room[2],
							&room[3],
							&room[4],
						},
					},
					{
						//[start t E a m n e end]
						Path: []*structs.Room{
							&room[0],
							&room[6],
							&room[7],
							&room[8],
							&room[5],
							&room[2],
							&room[3],
							&room[4],
						},
					},
					{
						//[start 0 o n h A c k end]
						Path: []*structs.Room{
							&room[0],
							&room[9],
							&room[10],
							&room[2],
							&room[1],
							&room[11],
							&room[12],
							&room[13],
							&room[4],
						},
					},
					{
						//[start t E a m n h A c k end]
						Path: []*structs.Room{
							&room[0],
							&room[6],
							&room[7],
							&room[8],
							&room[5],
							&room[2],
							&room[1],
							&room[11],
							&room[12],
							&room[13],
							&room[4],
						},
					},
				},
			},
			want: []*structs.PathStruct{
				{
					////[start t E a m end]
					Path: []*structs.Room{
						&room[0],
						&room[6],
						&room[7],
						&room[8],
						&room[5],
						&room[4],
					},
				},
				{
					//[start h A c k end]
					Path: []*structs.Room{
						&room[0],
						&room[1],
						&room[11],
						&room[12],
						&room[13],
						&room[4],
					},
				},
				{
					//[start 0 o n e end]
					Path: []*structs.Room{
						&room[0],
						&room[9],
						&room[10],
						&room[2],
						&room[3],
						&room[4],
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimPaths(tt.args.allPaths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimPaths(): %v, want: %v", GetSliceOfPathNames(got), GetSliceOfPathNames(tt.want))
			}
		})
	}
}
