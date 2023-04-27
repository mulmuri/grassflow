package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {

	if err := os.Chdir("../"); err != nil {
		panic(err)
	}

	cmd := exec.Command("go", "build", "-o", "grassflow", "cmd/run.go")
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	code := m.Run()

	cmd = exec.Command("rm", "grassflow")
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	os.Exit(code)
}


func Test_main(t *testing.T) {

	tests := []struct {
		name  string
		args[]string
		wantErr bool
	}{
		{
			name: "grassflow execute",
			args: []string{"./grassflow", "execute", "-f" ,"test/config.yml"},
			wantErr: false,
		},
		{
			name: "grassflow execute",
			args: []string{"./grassflow", "execute", "-f" ,"."},
			wantErr: true,
		},
		{
			name: "grassflow rollback",
			args: []string{"./grassflow", "rollback"},
			wantErr: false,
		},
		{
			name: "grassflow statistic",
			args: []string{"./grassflow", "statistic"},
			wantErr: false,
		},
		{
			name: "grassflow help",
			args: []string{"./grassflow", "help"},
			wantErr: false,
		},
		{
			name: "grassflow --version",
			args: []string{"./grassflow", "version"},
			wantErr: false,
		},
		{
			name: "grassflow -v",
			args: []string{"./grassflow", "version"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command(tt.args[0], tt.args[1:]...)

			var stderr bytes.Buffer
			cmd.Stderr = &stderr
		
			if err := cmd.Run(); (err != nil) != tt.wantErr {
				t.Errorf("failed to run %s : %v", strings.Join(tt.args, " "), stderr.String())
			}
		})
	}
}
