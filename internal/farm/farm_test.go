package farm

import (
	"reflect"
	"testing"

	structs "git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

func TestGenerateFarm(t *testing.T) {
	type args struct {
		data structs.GenerationData
	}
	tests := []struct {
		name string
		args args
		want structs.Farm
	}{
		{
			name: "pass",
			args: args{
				data: structs.GenerationData{
					NumberOfAnts: 3,
					Rooms:        []string{"1", "0", "2"},
					StartIndex:   0,
					EndIndex:     1,
				},
			},
			want: structs.Farm{
				Ants: []*structs.Ant{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
					{
						Id: 2,
					},
				},
				Rooms: []*structs.Room{
					{
						Name:    "1",
						IsStart: true,
						IsEnd:   false,
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
			},
		},
		{
			name: "empty",
			args: args{
				data: structs.GenerationData{},
			},
			want: structs.Farm{
				Ants:  []*structs.Ant{},
				Rooms: []*structs.Room{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// This function, right now, doesn't generate links so your test should assume that none exist
			if got := GenerateFarm(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateFarm() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestConnectRooms(t *testing.T) {
	type args struct {
		farm structs.Farm
		data structs.GenerationData
	}
	tests := []struct {
		name string
		args args
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
				},
				data: structs.GenerationData{
					Links: []string{"1-0", "1-2"},
				},
			},
		},
		{
			name: "empty",
			args: args{
				farm: structs.Farm{},
				data: structs.GenerationData{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConnectRooms(tt.args.farm, tt.args.data)
			if tt.name == "pass" {
				if reflect.DeepEqual(got, structs.Farm{}) {
					t.Errorf("ConnectRooms() got %+v, wanted filled farm", got)
				}
				for _, room := range got.Rooms {
					if room.Name == "1" && len(room.Links) != 2 {
						t.Errorf("ConnectRooms() gave room 1 %d links, want 2", len(room.Links))
					}
					if room.Name == "0" && len(room.Links) != 1 {
						t.Errorf("ConnectRooms() gave room 0 %d links, want 1", len(room.Links))
					}
				}
			}
			if tt.name == "empty" {
				if !reflect.DeepEqual(got, structs.Farm{}) {
					t.Errorf("ConnectRooms() got %+v, wanted %+v", got, structs.Farm{})
				}
			}
		})
	}
}
