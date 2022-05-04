package ants

import (
	"reflect"
	"testing"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

func TestMoveAnts(t *testing.T) {
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
		{
			Id: 4,
		},
	}

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

	type args struct {
		ants []*structs.Ant
		path *structs.PathStruct
	}
	tests := []struct {
		name string
		args args
		want []*structs.Ant
	}{
		{
			name: "pass - easy",
			args: args{
				ants: []*structs.Ant{
					&ants[0],
					&ants[1],
					&ants[2],
					&ants[3],
				},
				path: &structs.PathStruct{
					Path: []*structs.Room{
						&rooms[0],
						&rooms[3],
						&rooms[2],
						&rooms[1],
					},
				},
			},
			want: []*structs.Ant{
				{
					Id: 1,
					CurrentRoom: &structs.Room{
						Name: "1",
					},
					RoomsPassed: 3,
				},
				{
					Id: 2,
					CurrentRoom: &structs.Room{
						Name: "1",
					},
					RoomsPassed: 3,
				},
				{
					Id: 3,
					CurrentRoom: &structs.Room{
						Name: "1",
					},
					RoomsPassed: 3,
				},
				{
					Id: 4,
					CurrentRoom: &structs.Room{
						Name: "1",
					},
					RoomsPassed: 3,
				},
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
