package workers

type worker struct {
}

func (w worker) Execute(id uint, task Task) error {
	return task(id)
}
