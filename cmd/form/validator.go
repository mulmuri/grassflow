package form

import "fmt"



func (a *Argument) Contains(argName string) bool {
	_, ok := a.argMap[argName]
	return ok
}



func (a *Argument) ShouldExist(argName string) error {

	if !a.Contains(argName) {
		return fmt.Errorf("it should contain argument: %s", argName)
	}

	return nil
}



func (a *Argument) ShouldContainMany(argName string) ([]string, error) {

	if err := a.ShouldExist(argName); err != nil {
		return nil, err
	}

	argList := a.argMap[argName]

	if len(argList) == 0 {
		return nil, fmt.Errorf("it should contain at least one argument: %s", argName)
	}

	return argList, nil
}



func (a *Argument) ShouldContainOne(argName string) (string, error) {

	if err := a.ShouldExist(argName); err != nil {
		return "", err
	}

	argList := a.argMap[argName]

	if len(argList) != 1 {
		return "", fmt.Errorf("it should contain only one argument: %s", argName)
	}

	return argList[0], nil
}



func (a *Argument) ShouldContainNone(argName string) error {

	if err := a.ShouldExist(argName); err != nil {
		return err
	}

	argList := a.argMap[argName]

	if len(argList) != 0 {
		return fmt.Errorf("it should contain at least one argument: %s", argName)
	}

	return nil
}