package io

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	if err := os.Chdir("../../../test"); err != nil {
		panic(err)
	}

	code := m.Run()
	os.Exit(code)
}


func TestWrite(t *testing.T) {

	type args struct {
		name    string
		content string
		path    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create hello world file",
			args: args{name: "hello.txt", content: "Hello World!", path: "."},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Write(tt.args.name, tt.args.content, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}

			_, err := os.Open(filepath.Join(tt.args.path, tt.args.name))
			if err != nil {
				t.Errorf("Error opening file: %v", err)
			}
		})
	}
}
