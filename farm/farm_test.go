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
				data: structs.GenerationData{
					NumberOfAnts: 3,
					Rooms:        []string{"1", "0"},
					Links:        []string{"0-1"},
					StartIndex:   0,
					EndIndex:     1,
				},
			},
			want: structs.Farm{
				Rooms: []*structs.Room{
					{
						Name:  "0",
						Ants:  3,
						X_pos: 0,
						// Y_pos:   1,
						IsStart: true,
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
			if got := GenerateFarm(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateFarm() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
