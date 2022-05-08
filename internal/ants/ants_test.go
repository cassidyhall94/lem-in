package ants

import (
	"fmt"
	"reflect"
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
	// {
	// 	Name: "4",
	// },
	// {
	// 	Name: "5",
	// },
	// {
	// 	Name: "6",
	// },
	// {
	// 	Name: "7",
	// },
	// {
	// 	Name: "8",
	// },
	// {
	// 	Name: "9",
	// },
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
		ants []*structs.Ant
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AssignAnts(tt.args.ants, tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssignAnts() = %v, want %v", got, tt.want)
			}
		})
	}
}
