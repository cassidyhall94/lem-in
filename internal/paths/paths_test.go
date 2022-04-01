package paths

import (
	"reflect"
	"testing"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// func for finding all the paths from start to end using farm.room, room.Links(take in farm.rooms and loop over it and its links), and gives you [][]*Room

func Test_findAllPaths(t *testing.T) {
	type args struct {
		farm structs.Farm
	}
	tests := []struct {
		name string
		args args
		want [][]*structs.Room
	}{
		{
			name: "pass",
			args: args{
				farm: structs.Farm{
					Rooms: []*structs.Room{
						{
							Name:    "1",
							IsStart: true,
							IsEnd:   false,
							Links: []*structs.Room{
								
							},
						},
						{
							Name:    "0",
							IsStart: false,
							IsEnd:   true,
						},
						{
							Name:    "2",
							IsStart: false,
							IsEnd:   false,
						},
					},
				}
			}
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllPaths(tt.args.farm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
