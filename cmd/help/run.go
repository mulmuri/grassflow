package help

import "errors"

var WorkflowName = "help"




func Run() error {
	return errors.New("help command is not available now")
}