package git

import (
	"context"
	"testing"
	"time"
)

func TestGitAdd(t *testing.T) {
	type args struct {
		ctx  context.Context
		path []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GitAdd(tt.args.ctx, tt.args.path...); (err != nil) != tt.wantErr {
				t.Errorf("GitAdd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGitCommit(t *testing.T) {
	type args struct {
		ctx     context.Context
		message string
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
			if err := GitCommit(tt.args.ctx, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("GitCommit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGitCommitWithDate(t *testing.T) {
	type args struct {
		ctx     context.Context
		message string
		date    time.Time
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
			if err := GitCommitWithDate(tt.args.ctx, tt.args.message, tt.args.date); (err != nil) != tt.wantErr {
				t.Errorf("GitCommitWithDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
