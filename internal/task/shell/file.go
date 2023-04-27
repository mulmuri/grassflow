package shell

import (
	"context"
	"fmt"

	"github.com/mulmuri/grassflow/internal/infra/shell"
)



func ChangeDir(ctx context.Context, from string, to string) error {
	cmd := fmt.Sprintf("mv %s %s", from, to)
	return shell.Run(ctx, cmd)
}



func CopyDir(ctx context.Context, from string, to string) error {
	cmd := fmt.Sprintf("cp -r %s %s", from, to)
	return shell.Run(ctx, cmd)
}


func DeleteDir(ctx context.Context, file string) error {
	cmd := fmt.Sprintf("rm -rf %s", file)
	return shell.Run(ctx, cmd)
}