package git

import (
	"context"
	"fmt"
	"strings"

	"github.com/mulmuri/grassflow/internal/infra/shell"
)



func CheckBranchExist(ctx context.Context, branch string) (bool, error) {

	cmd := fmt.Sprintf("git branch | grep -w %s", branch)	
	out, err := shell.Output(ctx, cmd)

	if err != nil {
		return false, err
	}

	if out == "" {
		return false, nil
	}

	if strings.Split(out, " ")[1] != branch {
		cmd := fmt.Sprintf("matched branch is not same with branch %s", branch)
		return false, fmt.Errorf(cmd)
	}

	return true, nil
}



func CreateBranch(ctx context.Context, branch string) error {
	cmd := fmt.Sprintf("git branch %s", branch)
	return shell.Run(ctx, cmd)
}



func CreateCheckout(ctx context.Context, branch string) error {
	cmd := fmt.Sprintf("git checkout -b %s", branch)
	return shell.Run(ctx, cmd)
}



func Checkout(ctx context.Context, branch string) error {
	cmd := fmt.Sprintf("git checkout %s", branch)
	return shell.Run(ctx, cmd)
}



func Merge(ctx context.Context, branch string, target string) error {
	var cmd string

	cmd = fmt.Sprintf("git branch %s", branch)
	if err := shell.Run(ctx, cmd); err != nil {
		return err
	}

	cmd = fmt.Sprintf("git merge %s", target)
	if err := shell.Run(ctx, cmd); err != nil {
		return err
	}

	return nil
} 