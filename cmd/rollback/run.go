package rollback

import (
	"errors"

	"github.com/mulmuri/grassflow/cmd/form"
)

var WorkflowName = "rollback"




func Run(*form.Argument) error {
	return errors.New("rollback command is not available now")
}