package paths

import (
	"reflect"
	"testing"

	dataparser "git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/data-parser"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/farm"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

// func for finding all the paths from start to end using farm.room, room.Links(take in farm.rooms and loop over it and its links), and gives you [][]*Room

func Test_findAllPaths(t *testing.T) {
	data, _ := dataparser.LoadData("fixtures/test4.txt")
	generationData := dataparser.ReadData(data)
	filledFarm := farm.GenerateFarm(generationData)
	linkedFarm := farm.ConnectRooms(filledFarm, generationData)

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
				farm: linkedFarm,
			},
			want: [][]*structs.Room{
				
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAllPaths(tt.args.farm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
