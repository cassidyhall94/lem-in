package ants

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

var rooms = []structs.Room{
	{
		Name:    "0",
		IsStart: true,
		IsEnd:   false,
	},
	{
		Name:    "1",
		IsStart: false,
		IsEnd:   true,
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
}

var ants = []structs.Ant{
	{
		Id: 1,
	},
	{
		Id: 2,
	},
	{
		Id: 3,
	},
	{
		Id: 4,
	},
	{
		Id: 5,
	},
	{
		Id: 6,
	},
	{
		Id: 7,
	},
	{
		Id: 8,
	},
	{
		Id: 9,
	},
	{
		Id: 10,
	},
}

var room = []structs.Room{
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

func TestMoveAnts(t *testing.T) {
	type args struct {
		ants []*structs.Ant
		path *structs.PathStruct
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "pass - easy",
			args: args{
				ants: []*structs.Ant{
					&ants[0],
					&ants[1],
					&ants[2],
				},
				path: &structs.PathStruct{
					Path: []*structs.Room{
						&rooms[0],
						&rooms[2],
						&rooms[3],
						&rooms[1],
					},
				},
			},
			want: []string{
				"L1-2",
				"L1-3 L2-2",
				"L1-1 L2-3 L3-2",
				"L2-1 L3-3",
				"L3-1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MoveAnts(tt.args.ants, tt.args.path)
			fmt.Printf("got: %+v, want: %+v", len(got), len(tt.want))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveAnts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAssignAnts(t *testing.T) {
	type args struct {
		ants  []*structs.Ant
		paths []*structs.PathStruct
	}
	tests := []struct {
		name string
		args args
		want []*structs.PathStruct
	}{
		{
			name: "pass - easy",
			args: args{
				ants: []*structs.Ant{
					&ants[0],
					&ants[1],
					&ants[2],
					&ants[3],
					&ants[4],
					&ants[5],
				},
				paths: []*structs.PathStruct{
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
			want: []*structs.PathStruct{
				{
					Path: []*structs.Room{
						&rooms[0],
						&rooms[1],
					},
					Ants: []*structs.Ant{
						&ants[0],
						&ants[1],
						&ants[2],
						&ants[4],
					},
				},
				{
					Path: []*structs.Room{
						&rooms[0],
						&rooms[2],
						&rooms[3],
						&rooms[1],
					},
					Ants: []*structs.Ant{
						&ants[3],
						&ants[5],
					},
				},
			},
		},
		{
			name: "pass - hard",
			args: args{
				ants: []*structs.Ant{
					&ants[0],
					&ants[1],
					&ants[2],
					&ants[3],
					&ants[4],
					&ants[5],
					&ants[6],
					&ants[7],
					&ants[8],
					&ants[9],
				},
				paths: []*structs.PathStruct{
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
			want: []*structs.PathStruct{
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
					Ants: []*structs.Ant{
						&ants[0],
						&ants[3],
						&ants[6],
						&ants[9],
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
					Ants: []*structs.Ant{
						&ants[1],
						&ants[4],
						&ants[7],
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
					Ants: []*structs.Ant{
						&ants[2],
						&ants[5],
						&ants[8],
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, gotPaths := range AssignAnts(tt.args.ants, tt.args.paths) {
				for j, wantPaths := range tt.want {
					if i != j {
						continue
					}
					if len(gotPaths.Ants) != len(wantPaths.Ants) {
						t.Errorf("AssignAnts(Ants) = %v, want %v", len(gotPaths.Ants), len(wantPaths.Ants))
					}
					if !reflect.DeepEqual(gotPaths.Path, wantPaths.Path) {
						t.Errorf("AssignAnts(Paths): %v\n want: %+v", GetSliceOfRoomNames(gotPaths.Path), GetSliceOfRoomNames(wantPaths.Path))
					}
				}
			}
		})
	}
}

func GetSliceOfPathNames(paths []*structs.PathStruct) []string {
	ret := []string{}
	for _, path := range paths {
		for _, pathInfo := range path.Path {
			ret = append(ret, pathInfo.Name)
		}
	}
	return ret
}

func GetSliceOfAntNames(ants []*structs.Ant) []string {
	ret := []string{}
	for _, ant := range ants {
		ret = append(ret, strconv.Itoa(ant.Id))
	}
	return ret
}

func GetSliceOfRoomNames(rooms []*structs.Room) []string {
	ret := []string{}
	for _, r := range rooms {
		ret = append(ret, r.Name)
	}
	return ret
}
