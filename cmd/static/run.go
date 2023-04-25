package static

import (
	"errors"

	"github.com/mulmuri/grassflow/cmd/form"
)

var WorkflowName = "static"



func Run(*form.Argument) error {
	return errors.New("static command is not available now")
}