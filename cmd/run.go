package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mulmuri/grassflow/cmd/execute"
	"github.com/mulmuri/grassflow/cmd/form"
	"github.com/mulmuri/grassflow/cmd/help"
	"github.com/mulmuri/grassflow/cmd/rollback"
	"github.com/mulmuri/grassflow/cmd/statistic"
	"github.com/mulmuri/grassflow/cmd/version"
)



func main() {

	var err error

	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("panic recovered: %v", err)
		}
	}()

	workflowName := os.Args[1]

	var newArgs = []string{}
	if len(os.Args) > 2 {
		newArgs = os.Args[2:]
	}

	if (workflowName == "-v" || workflowName == "--version" || workflowName == "v") {
		workflowName = "version"
		newArgs = []string{}
	}

	argument := form.NewWithParsing(newArgs)

	switch workflowName {

	case execute.WorkflowName:
		err = execute.Run(argument)

	case rollback.WorkflowName:
		err = rollback.Run(argument)
	
	case statistic.WorkflowName:
		err = statistic.Run(argument)

	case help.WorkflowName:
		err = help.Run()

	case version.WorkflowName:
		version.Run()

	default:
		err = fmt.Errorf("unknown parameter")
	}

	if err != nil {
		log.Fatalf("failed to run workflow: %v", err)
	}
}
