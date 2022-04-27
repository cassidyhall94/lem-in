package ants

import (
	"reflect"
	"testing"
)

func TestCreateAnts(t *testing.T) {
	tests := []struct {
		name  string
		want  []string
		want1 []int
	}{
		{
			name:  "test",
			want:  []string{"3"},
			want1: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CreateAnts()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAnts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CreateAnts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
