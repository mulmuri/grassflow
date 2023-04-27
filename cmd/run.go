package main

import (
	"errors"
	"log"
	"os"

	"github.com/mulmuri/grassflow/cmd/execute"
	"github.com/mulmuri/grassflow/cmd/help"
	"github.com/mulmuri/grassflow/cmd/rollback"
	"github.com/mulmuri/grassflow/cmd/static"
	"github.com/mulmuri/grassflow/cmd/version"
	"github.com/mulmuri/grassflow/cmd/form"
)



func main() {

	var err error

	workflowName := os.Args[1]
	newArgs := os.Args[2:]

	if (newArgs[0] == "-v" || newArgs[0] == "--version" || newArgs[0] == "v") {
		workflowName = "version"
		newArgs = []string{}
	}

	argument := form.NewWithParsing(newArgs)

	switch workflowName {

	case execute.WorkflowName:
		err = execute.Run(argument)

	case rollback.WorkflowName:
		err = rollback.Run(argument)
	
	case static.WorkflowName:
		err = static.Run(argument)

	case help.WorkflowName:
		err = help.Run()

	case version.WorkflowName:
		version.Run()

	default:
		err = errors.New("unknown parameter")
	}

	if err != nil {
		log.Fatal(err)
	}
}
