package git

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/mulmuri/grassflow/internal/infra/shell"
	"github.com/mulmuri/grassflow/internal/model"
)




func GitAdd(ctx context.Context, path ...string) error {
	cmd := "git add " + strings.Join(path, " ")
	return shell.Run(ctx, cmd)
}


func GitCommit(ctx context.Context, message string) error {
	cmd := fmt.Sprintf("git commit -m \"%s\"", message)
	return shell.Run(ctx, cmd)
}


func GitCommitWithDate(ctx context.Context, message string, date time.Time) error {
	cmd := fmt.Sprintf("git commit -m \"%s\" --date %s", message, date.Format(model.TimeFormat))
	return shell.Run(ctx, cmd)
}