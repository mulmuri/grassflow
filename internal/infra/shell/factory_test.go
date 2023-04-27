package shell

import (
	"context"
	"os"
	"os/exec"
	"testing"
)



func TestMain(m *testing.M) {

	if err := os.Chdir("../../../test"); err != nil {
		panic(err)
	}

	cmd := exec.Command("git", "init")
	cmd.Run()
	
	code := m.Run()
	os.Exit(code)
}


func TestRun(t *testing.T) {

	type args struct {
		ctx     context.Context
		command string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ls",
			args: args{context.TODO(), "ls -a"},
			wantErr: false,
		},
		{
			name: "python is not installed",
			args: args{context.TODO(), "python --version"},
			wantErr: true,
		},
		{
			name: "copy dir",
			args: args{context.TODO(), "cp -r .git .git.swp"},
			wantErr: false,
		},
		{
			name: "delete dir",
			args: args{context.TODO(), "rm -rf .git.swp"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Run(tt.args.ctx, tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOutput(t *testing.T) {

	type args struct {
		ctx     context.Context
		command string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "git status",
			args: args{context.TODO(), "git status"},
			want: "On branch main\n\nNo commits yet\n\nnothing to commit (create/copy files and use \"git add\" to track)\n",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Output(tt.args.ctx, tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("output() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("output() = %v, want %v", got, tt.want)
			}
		})
	}
}
