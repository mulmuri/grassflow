package git

import (
	"context"
	"os"
	"os/exec"
	"testing"

	"github.com/mulmuri/grassflow/internal/task/io"
)


func TestMain(m *testing.M) {
	
	if err := os.Chdir("../../../test"); err != nil {
		panic(err)
	}

	cmd := exec.Command("git", "init")
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	io.Write("hello.txt", "hello world", ".")

	code := m.Run()
	os.Exit(code)
}



func TestCheckBranchExist(t *testing.T) {
	type args struct {
		ctx    context.Context
		branch string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckBranchExist(tt.args.ctx, tt.args.branch)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckBranchExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckBranchExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateBranch(t *testing.T) {
	type args struct {
		ctx    context.Context
		branch string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateBranch(tt.args.ctx, tt.args.branch); (err != nil) != tt.wantErr {
				t.Errorf("CreateBranch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateCheckout(t *testing.T) {
	type args struct {
		ctx    context.Context
		branch string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateCheckout(tt.args.ctx, tt.args.branch); (err != nil) != tt.wantErr {
				t.Errorf("CreateCheckout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckout(t *testing.T) {
	type args struct {
		ctx    context.Context
		branch string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Checkout(tt.args.ctx, tt.args.branch); (err != nil) != tt.wantErr {
				t.Errorf("Checkout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args struct {
		ctx    context.Context
		branch string
		target string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Merge(tt.args.ctx, tt.args.branch, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("Merge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
