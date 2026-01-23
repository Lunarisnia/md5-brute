package workers

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Task func(id uint) error

type Workers interface {
	SetWorkerCount(count uint) Workers
	SetTask(task Task) Workers
	Run(ctx context.Context) error
}

type workers struct {
	task        Task
	workerCount uint
}

func New() Workers {
	return &workers{}
}

func (w workers) SetWorkerCount(count uint) Workers {
	w.workerCount = count
	return w
}

func (w workers) SetTask(task Task) Workers {
	w.task = task
	return w
}

func (w workers) Run(ctx context.Context) error {
	eg, ctx := errgroup.WithContext(ctx)
	for i := range w.workerCount {
		eg.Go(func() error {
			wk := worker{}
			err := wk.Execute(i, w.task)
			return err
		})
	}
	return eg.Wait()
}
