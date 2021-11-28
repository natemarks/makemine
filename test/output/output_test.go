package output

import (
	"testing"

	"github.com/natemarks/makemine/input"
	"github.com/natemarks/makemine/output"
)

func TestToJsonFile(t *testing.T) {
	mm := input.MakeMineInput{
		Name:      "Some P. User",
		LocalUser: "suser",
		Email:     "spuser@gmail.com",
	}
	type args struct {
		data  input.MakeMineInput
		oFile string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "valid",
			args: args{
				data:  mm,
				oFile: "/tmp/gg.json",
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := output.ToJsonFile(tt.args.data, tt.args.oFile); (err != nil) != tt.wantErr {
				t.Errorf("ToJsonFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			read, err := input.FromFile(tt.args.oFile)
			if err != nil {
				t.Errorf("Failed to open generated file %v", tt.args.oFile)
			}
			if read != mm {
				t.Errorf("Incorrect data in %v", tt.args.oFile)
			}
		})
	}
}
