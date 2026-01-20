package workers

type worker struct {
}

func (w worker) Execute(task Task) error {
	return task()
}
