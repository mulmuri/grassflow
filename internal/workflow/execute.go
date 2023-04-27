package workflow

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mulmuri/grassflow/internal/io"
	"github.com/mulmuri/grassflow/internal/model"
	"github.com/mulmuri/grassflow/internal/shell"
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

		if err := shell.CreateCheckout(ctx, branch); err != nil {
			return err
		}

		for _, file := range fileList {

			var msg string
			msg = strings.Replace(msgForm, "{d}", file.Parent, 1)
			msg = strings.Replace(msg, "{f}", file.Name, 1)
			msg = strings.Replace(msg, "{f-ext}", strings.TrimSuffix(file.Name, filepath.Ext(file.Name)), 1)

			if err := shell.GitAdd(ctx, file.Path); err != nil {
				return err
			}

			if err := shell.GitCommitWithDate(ctx, msg, file.ModTime); err != nil {
				return err
			}			
		}

		if err := shell.Merge(ctx, mainBranch, branch); err != nil {
			return err
		}
	}

	rollbackRequired = false
	return nil
}