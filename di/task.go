package di

type Task struct {
}

func NewTask() *Task {
	return &Task{}
}
func (a *Task) Run() error {
	return nil
}
