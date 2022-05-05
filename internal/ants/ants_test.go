package ants

import (
	"reflect"
	"testing"
	//"git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

func TestCreateAnts(t *testing.T) {

	tests := []struct {
		name string
		want []string
	}{
		{
			name: "pass - easy",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateAnts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAnts() = %v, want %v", got, tt.want)
			}
		})
	}
}
