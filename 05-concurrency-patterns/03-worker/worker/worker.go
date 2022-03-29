package worker

import (
	"fmt"
	"sync"
)

type Work interface {
	Task()
}

type Worker struct {
	work chan Work
	wg   sync.WaitGroup
}

func New(maxTasks int) *Worker {
	worker := &Worker{
		work: make(chan Work),
	}
	worker.wg.Add(maxTasks)
	for idx := 0; idx < maxTasks; idx++ {
		go func() {
			for w := range worker.work {
				w.Task()
			}
			worker.wg.Done()
		}()
	}
	return worker
}

func (worker *Worker) Run(w Work) {
	worker.work <- w
}

func (worker *Worker) Shutdown() {
	close(worker.work)
	worker.wg.Wait()
	fmt.Println("worker shutdown complete")
}
