package model

import (
	"reflect"
	"testing"

	"github.com/natemarks/makemine/model"
)

func TestMyDataFromFilePath(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    model.MyData
		wantErr bool
	}{
		{name: "valid",
			args:    args{filePath: "../testdata/FromFile/vvalid/input.json"},
			wantErr: false,
			want: model.MyData{
				FullName:  "Nathan Marks",
				LocalUser: "nmarks",
				Email:     "npmarks@gmail.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.MyDataFromFilePath(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("MyDataFromFilePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyDataFromFilePath() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDataFromURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    model.MyData
		wantErr bool
	}{
		{
			name:    "valid",
			args:    args{url: "https://gist.githubusercontent.com/natemarks/3620f565dc4e775143647d62c151658e/raw/54559a95b3ba58c0cc2ca8889ea12c72adbf7adb/makemine.json"},
			wantErr: false,
			want: model.MyData{
				FullName:  "Nathan Marks",
				LocalUser: "nmarks",
				Email:     "npmarks@gmail.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.MyDataFromURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("MyDataFromURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyDataFromURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
