package farm

import (
	structs "lem-in/structs"
	"reflect"
	"testing"
)

func TestGenerateFarm(t *testing.T) {
	tests := []struct {
		name string

		want  structs.GenerationData
		want1 int
		want2 int
		want3 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3 := GenerateFarm(tt.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateFarm() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GenerateFarm() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("GenerateFarm() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("GenerateFarm() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
