package paths

import (
	"reflect"
	"testing"

	dataparser "git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/data-parser"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/farm"
	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

func Test_findAllPaths(t *testing.T) {
	data, err := dataparser.LoadData("../data-parser/fixtures/test4.txt")
	if err != nil {
		t.Errorf("dataparser.LoadData error: %+v", err)
	}
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
				{
					{
						Name:    "0",
						IsStart: true,
						Links: []*structs.Room{
							{
								Name:    "2",
								IsEnd:   false,
								IsStart: false,
							},
						},
					},
				},
				{
					{
						Name:    "2",
						IsStart: false,
						IsEnd:   false,
						Links: []*structs.Room{
							{
								Name:    "0",
								IsStart: true,
							},
						},
					},
				},
				{
					{
						Name:    "2",
						IsStart: false,
						IsEnd:   false,
						Links: []*structs.Room{
							{
								Name:    "3",
								IsEnd:   false,
								IsStart: false,
							},
						},
					},
				},
				{
					{
						Name:    "3",
						IsStart: false,
						IsEnd:   false,
						Links: []*structs.Room{
							{
								Name:    "2",
								IsEnd:   false,
								IsStart: false,
							},
						},
					},
				},
				{
					{
						Name:    "3",
						IsStart: false,
						IsEnd:   false,
						Links: []*structs.Room{
							{
								Name:  "1",
								IsEnd: true,
							},
						},
					},
				},
				{
					{
						Name:  "1",
						IsEnd: true,
						Links: []*structs.Room{
							{
								Name:    "3",
								IsEnd:   false,
								IsStart: false,
							},
						},
					},
				},
			},
		},
		{
			name: "empty",
			args: args{
				farm: linkedFarm,
			},
			want: [][]*structs.Room{},
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
