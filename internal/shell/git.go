package shell

import (
	"context"
	"errors"
	"fmt"
	"strings"
)



func CheckBranchExist(ctx context.Context, branch string) (bool, error) {

	cmd := fmt.Sprintf("git branch | grep -w %s", branch)	
	out, err := output(ctx, cmd)

	if err != nil {
		return false, err
	}

	if out == "" {
		return false, nil
	}

	if strings.Split(out, " ")[1] != branch {
		cmd := fmt.Sprintf("matched branch is not same with branch %s", branch)
		return false, errors.New(cmd)
	}

	return true, nil
}



func CreateBranch(ctx context.Context, branch string) error {
	cmd := fmt.Sprintf("git branch %s", branch)
	return run(ctx, cmd)
}



func Checkout(ctx context.Context, branch string) error {
	cmd := fmt.Sprintf("git checkout %s", branch)
	return run(ctx, cmd)
}



func Merge(ctx context.Context, branch string, target string) error {
	var cmd string

	cmd = fmt.Sprintf("git branch %s", branch)
	if err := run(ctx, cmd); err != nil {
		return err
	}

	cmd = fmt.Sprintf("git merge %s", target)
	if err := run(ctx, cmd); err != nil {
		return err
	}

	return nil
} 