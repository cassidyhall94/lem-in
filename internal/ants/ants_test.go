package ants

import (
	"reflect"
	"testing"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

func TestCreateAnts(t *testing.T) {
	type args struct {
		Ants     structs.Farm
		allPaths []*structs.PathStruct
	}
	tests := []struct {
		name string
		args args
		want []*structs.Ant
	}{
		{
			name: "pass - easy",
			args: args{
				Ants: structs.Farm{
					Ants: []*structs.Ant{
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
					},
				},
				allPaths: []*structs.PathStruct{
					{
						Path: []*structs.Room{
							{
								Name: "0",
							},
							{
								Name: "3",
							},
							{
								Name: "2",
							},
							{
								Name: "1",
							},
						},
					},
				},
			},
			want: []*structs.Ant{
				{
					Id: 1,
					Path: []*structs.Room{
						{
							Name: "0",
						},
						{
							Name: "3",
						},
						{
							Name: "2",
						},
						{
							Name: "1",
						},
					},
					CurrentRoom: &structs.Room{
						Name: "1",
					},
					RoomsPassed: 3,
				},
				{
					Id: 2,
					Path: []*structs.Room{
						{
							Name: "0",
						},
						{
							Name: "3",
						},
						{
							Name: "2",
						},
						{
							Name: "1",
						},
					},
					CurrentRoom: &structs.Room{
						Name: "1",
					},
					RoomsPassed: 3,
				},
				{
					Id: 3,
					Path: []*structs.Room{
						{
							Name: "0",
						},
						{
							Name: "3",
						},
						{
							Name: "2",
						},
						{
							Name: "1",
						},
					},
					CurrentRoom: &structs.Room{
						Name: "1",
					},
					RoomsPassed: 3,
				},
				{
					Id: 4,
					Path: []*structs.Room{
						{
							Name: "0",
						},
						{
							Name: "3",
						},
						{
							Name: "2",
						},
						{
							Name: "1",
						},
					},
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
			if got := CreateAnts(tt.args.Ants, tt.args.allPaths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAnts() = %v, want %v", got, tt.want)
			}
		})
	}
}
