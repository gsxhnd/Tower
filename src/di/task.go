package di

type Task struct {
	// engine *gin.Engine
}

func NewTask() *Task {
	return &Task{}
}
func (a *Task) Run() error {
	return nil
}
