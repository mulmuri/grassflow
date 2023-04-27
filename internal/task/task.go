package task



type Condition interface {}


type Task interface {
	Execute(Condition) Condition
	Rollback(Condition) Condition
	ConditionCheck(Condition) error
}

type TaskStack struct {
	taskList []Task
}

func (t *TaskStack) AppendTask(task Task) *TaskStack {
	t.taskList = append(t.taskList, task)
	return t
}

func (t *TaskStack) RunTask(task Task) {
	
}