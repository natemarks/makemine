package input

import (
	"reflect"
	"testing"

	"github.com/natemarks/makemine/input"
)

func TestInputFromUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    input.MakeMineInput
		wantErr bool
	}{
		{name: "valid",
			args: args{
				url: "https://gist.githubusercontent.com/natemarks/3620f565dc4e775143647d62c151658e/raw/3612b76c0ffcb86800de32d20714bc62cfe3e902/makemine.json"},
			want: input.MakeMineInput{
				Name:      "Nathan Marks",
				LocalUser: "nmarks",
				Email:     "npmarks@gmail.com",
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := input.FromUrl(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInputFromFile(t *testing.T) {
	type args struct {
		mFile string
	}
	tests := []struct {
		name    string
		args    args
		want    input.MakeMineInput
		wantErr bool
	}{
		{name: "vvalid",
			args: args{
				mFile: "../testdata/FromFile/vvalid/input.json"},
			want: input.MakeMineInput{
				Name:      "Nathan Marks",
				LocalUser: "nmarks",
				Email:     "npmarks@gmail.com",
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := input.FromFile(tt.args.mFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
