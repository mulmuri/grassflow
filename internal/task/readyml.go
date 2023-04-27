package task

import "errors"



type ReadYmlTask struct {}

func (t *ReadYmlTask) ConditionCheck(c Condition) error {

	if _, ok := c.(*string); !ok {
		return errors.New("")
	}

	return nil
}