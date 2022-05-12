package dataparser

import (
	"reflect"
	"testing"

	structs "git.learn.01founders.co/Cassidy.Hall94/lem-in/internal/structs"
)

func TestLoadData(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "read all lines",
			args: args{
				fileName: "fixtures/test1.txt",
			},
			want:    []string{"2-1", "7-6", "7-2", "7-4", "6-5"},
			wantErr: false,
		},
		{
			name: "file does not exist",
			args: args{
				fileName: "doesnotexis.txt",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "empty file name",
			args: args{
				fileName: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadData(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadData(t *testing.T) {
	tests := []struct {
		name          string
		fileLinesFile string
		want          structs.GenerationData
	}{
		{
			name:          "basic file",
			fileLinesFile: "fixtures/test2.txt",
			want: structs.GenerationData{
				NumberOfAnts: 3,
				Rooms:        []string{"1", "0"},
				Links:        []string{"0-1"},
				StartIndex:   0,
				EndIndex:     1,
			},
		},
		{
			name:          "empty file",
			fileLinesFile: "fixtures/test3.txt",
			want: structs.GenerationData{
				NumberOfAnts: 0,
				Rooms:        []string{},
				Links:        []string{},
				StartIndex:   0,
				EndIndex:     0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileLines, err := LoadData(tt.fileLinesFile)
			if err != nil {
				t.Errorf("unable to set up ReadData test: %+v", err)
			}
			got, _ := ReadData(fileLines)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadData() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
