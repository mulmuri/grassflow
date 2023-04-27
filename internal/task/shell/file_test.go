package shell_test

import (
	"context"
	"testing"

	"github.com/mulmuri/grassflow/internal/task/shell"
)

func TestChangeDir(t *testing.T) {
	type args struct {
		ctx  context.Context
		from string
		to   string
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
			if err := shell.ChangeDir(tt.args.ctx, tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("ChangeDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCopyDir(t *testing.T) {
	type args struct {
		ctx  context.Context
		from string
		to   string
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
			if err := shell.CopyDir(tt.args.ctx, tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("CopyDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteDir(t *testing.T) {
	type args struct {
		ctx  context.Context
		file string
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
			if err := shell.DeleteDir(tt.args.ctx, tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("DeleteDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
