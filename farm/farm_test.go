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
			name: "foo",
			args: args{
				data: structs.GenerationData{
					
				},
			},
			want: structs.Farm{
				Rooms: []*structs.Room{
					{
						IsStart: false,
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
