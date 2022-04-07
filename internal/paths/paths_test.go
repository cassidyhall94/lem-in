package paths

import (
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

	data1, err := dataparser.LoadData("../data-parser/fixtures/test5.txt")
	if err != nil {
		t.Errorf("dataparser.LoadData error: %+v", err)
	}
	generationData1 := dataparser.ReadData(data1)
	filledFarm1 := farm.GenerateFarm(generationData1)
	linkedFarm1 := farm.ConnectRooms(filledFarm1, generationData1)

	type args struct {
		farm structs.Farm
	}
	tests := []struct {
		name string
		args args
		want []*structs.PathStruct
	}{
		{
			name: "pass - easy",
			args: args{
				farm: linkedFarm,
			},
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
			name: "pass - more paths",
			args: args{
				farm: linkedFarm1,
			},
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
			name: "empty",
			args: args{
				farm: structs.Farm{},
			},
			want: []*structs.PathStruct{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findAllPaths(tt.args.farm)
			for _, gotType := range got {
				for i, gotPath := range gotType.Path {
					t.Errorf("findAllPaths() testname: %s, index: %d, name: %+v, isStart: %+v, isEnd: %+v, want: %+v\n", tt.name, i, gotPath.Name, gotPath.IsStart, gotPath.IsEnd, tt.want)
				}
			}
		})
	}
}
