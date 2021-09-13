package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded    = errors.New("errors limit exceeded")
	ErrErrorsUnexpectedCount  = errors.New("errors unexpected count")
	ErrWorkersUnexpectedCount = errors.New("workers unexpected count")
	ErrTasksEmpty             = errors.New("no tasks")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	taskCh := make(chan Task)
	wg := sync.WaitGroup{}
	var errCount int32

	if m <= 0 {
		return ErrErrorsUnexpectedCount
	}
	if n <= 0 {
		return ErrWorkersUnexpectedCount
	}
	if len(tasks) == 0 {
		return ErrTasksEmpty
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskCh {
				if err := task(); err != nil {
					atomic.AddInt32(&errCount, 1)
				}
			}
		}()
	}

	for _, task := range tasks {
		if int32(m) <= atomic.LoadInt32(&errCount) {
			break
		}
		taskCh <- task
	}
	close(taskCh)

	wg.Wait()

	if int32(m) <= errCount {
		return ErrErrorsLimitExceeded
	}

	return nil
}
