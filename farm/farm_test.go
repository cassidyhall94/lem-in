package farm

import (
	structs "lem-in/structs"
	"reflect"
	"testing"
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
				// In this struct, room 0 is the start and room 1 is the end
				data: structs.GenerationData{
					NumberOfAnts: 3,
					Rooms:        []string{"1", "0"},
					Links:        []string{"0-1"},
					StartIndex:   0,
					EndIndex:     1,
				},
				// In this struct, room bar is the start and room foo is the end
				// data: structs.GenerationData{
				// 	NumberOfAnts: 3,
				// 	Rooms:        []string{"foo", "zip", "bar"},
				// 	Links:        []string{"foo-zip", "zip-bar"},
				// 	StartIndex:   1,
				// 	EndIndex:     0,
				// },
			},
			want: structs.Farm{
				// What about the Ants: []*Ant?
				Rooms: []*structs.Room{
					// How many rooms do you need to define?
					{
						Name: "0",
						// How many ants can a room have? What is the significance of this value?
						Ants:    3,
						X_pos:   0,
						Y_pos:   1,
						// Is the room defined here the start room?
						IsStart: true,
						// What about it being the end room?
						IsEnd:   false,
						// Links: []*structs.Room {

						// },

					},
				},
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
