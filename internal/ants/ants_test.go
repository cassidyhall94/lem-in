package ants

import (
	"reflect"
	"testing"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

func TestMoveAnts(t *testing.T) {

	rooms := []structs.Room{
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

	ants := []structs.Ant{
		{
			Id: 1,
		},
		{
			Id: 2,
		},
		{
			Id: 3,
		},
	}

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
			if got := MoveAnts(tt.args.ants, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MoveAnts() = %v, want %v", got, tt.want)
			}
		})
	}
}
