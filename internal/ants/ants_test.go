package ants

import (
	"reflect"
	"testing"

	"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

func TestCreateAnts(t *testing.T) {
	type args struct {
		paths [][]*structs.Room
	}
	tests := []struct {
		name string
		args args
		want []structs.Ant
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateAnts(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAnts() = %v, want %v", got, tt.want)
			}
		})
	}
}
