package ants

import (
	"reflect"
	"testing"
	//"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

func TestCreateAnts(t *testing.T) {
	type args struct {
		ants *Ant
	}
	tests := []struct {
		name string
		args args
		want *Ant
	}{
		{
			name: "Ant1",
			args: args{
				ants: &Ant{
					Id:          1,
					Path:        []string{"1", "0"},
					RoomsPassed: 2,
				},
			},
		},
		/*{
			name: "Ant2",
			args: args{
				ants: &Ant{
					Id:          2,
					Path:        []string{"1", "0"},
					RoomsPassed: 2,
				},
			},
		},
		{
			name: "Ant3",
			args: args{
				ants: &Ant{
					Id:          3,
					Path:        []string{"1", "0"},
					RoomsPassed: 2,
				},
			},
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateAnts(); !reflect.DeepEqual(got, tt.args.ants) {
				t.Errorf("CreateAnts() = %v, want %v", got, tt.args.ants)

			}
		})
	}
}
