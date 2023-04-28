package task

import (
	"fmt"
)



type ReadYmlTask struct {}

func (t *ReadYmlTask) ConditionCheck(c Condition) error {

	if _, ok := c.(*string); !ok {
		return fmt.Errorf("")
	}

	return nil
}