package execute

import (
	"github.com/mulmuri/grassflow/cmd/form"
	"github.com/mulmuri/grassflow/internal/model"
	"github.com/mulmuri/grassflow/internal/workflow"
)


var (
	WorkflowName = "execute"
	fileFlag = "-f"
	workingDirFlag = "-d"
	targetBranchFlag = "-b"
)



func Run(args *form.Argument) error {

	var (
		config YamlConfig
		targetBranch string = "main"
		workingDir string = "."
		taskList []model.Task
	)

	configFileName, err := args.ShouldContainOne(fileFlag)
	if err != nil {
		return err
	}


	if err := ReadYML(configFileName, &config); err != nil {
		return err
	}

	if args.Contains(targetBranchFlag) {
		targetBranch, err = args.ShouldContainOne(targetBranchFlag)
		if err != nil {
			return err
		}
	}

	if args.Contains(workingDirFlag) {
		workingDir, err = args.ShouldContainOne(workingDirFlag)
		if err != nil {
			return err
		}
	}

	taskList = make([]model.Task, len(config.TaskConfig))

	for idx := range taskList {
		taskList[idx] = model.Task{
			Message: config.TaskConfig[idx].Message,
			Branch: config.TaskConfig[idx].Branch,
			Dir: config.TaskConfig[idx].Dir,
			Match: config.TaskConfig[idx].Match,
		}
	}

	return workflow.Execute(targetBranch, workingDir, taskList)
}