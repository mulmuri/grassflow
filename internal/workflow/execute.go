package workflow

import (
	"context"
	"os"
	"time"

	"github.com/mulmuri/grassflow/internal/model"
	"github.com/mulmuri/grassflow/internal/task/format"
	"github.com/mulmuri/grassflow/internal/task/git"
	"github.com/mulmuri/grassflow/internal/task/io"
	"github.com/mulmuri/grassflow/internal/task/shell"
)




func Execute(mainBranch string, workingDir string, taskList []model.Task) error {
	
	if err := os.Chdir(workingDir); err != nil {
		return err
	}

	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(10 * time.Second))
	
	var rollbackRequired bool = true
	defer func() {
		cancelFunc()

		if rollbackRequired {
			if err := shell.Rollback(context.Background()); err != nil {
				panic(err)
			}
		}
	}()

	if err := shell.Backup(ctx); err != nil {
		return err
	}

	for _, task := range taskList {

		branch := task.Branch
		msgForm := task.Message

		fileList := make([]model.File, 0)

		for _, dir := range task.Dir {
			list, err := io.ReadFile(dir)
			
			if err != nil {
				return err
			}

			fileList = append(fileList, list...)
		}

		if err := git.CreateCheckout(ctx, branch); err != nil {
			return err
		}

		for _, file := range fileList {

			msg := msgForm
			msg = format.ReplaceD(msg, file.Parent)
			msg = format.ReplaceF(msg, file.Name)
			msg = format.ReplaceN(msg, file.Name)

			if err := git.GitAdd(ctx, file.Path); err != nil {
				return err
			}

			if err := git.GitCommitWithDate(ctx, msg, file.ModTime); err != nil {
				return err
			}			
		}

		if err := git.Merge(ctx, mainBranch, branch); err != nil {
			return err
		}
	}

	rollbackRequired = false
	return nil
}