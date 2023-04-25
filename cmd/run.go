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
)



func main() {

	workflowName := os.Args[1]
	args := os.Args[2:]

	argMap := make(map[string][]string)

	if len(args) != 0 && args[0][0] != '-' {
		err := errors.New("unknown parameter")
		log.Fatal(err)
		return
	}

	var prevArg []string

	for _, arg := range args {
		if arg[0] == '-' {
			argMap[arg] = make([]string, 0)
			prevArg = argMap[arg]
		} else {
			prevArg = append(prevArg, arg)
		}
	}

	switch workflowName {

	case execute.WorkflowName:
		execute.Run(argMap)

	case rollback.WorkflowName:
		rollback.Run(argMap)
	
	case static.WorkflowName:
		static.Run(argMap)

	case help.WorkflowName:
		help.Run()

	case version.WorkflowName:
		version.Run()

	default:
		err := errors.New("unknown parameter")
		log.Fatal(err)
		return
	}	
}
