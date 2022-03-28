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
						Name: "1",
						IsStart: true,
						IsEnd: false,
					},
					{
						Name: "0",
						IsStart: false,
						IsEnd: true,
					},
					{
						Name: "2",
						IsStart: false,
						IsEnd: false,
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
