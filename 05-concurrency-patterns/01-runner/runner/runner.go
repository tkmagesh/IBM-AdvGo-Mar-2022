/* implement the runner */
package runner

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"time"
)

type Task func(int)

type Runner struct {
	t         time.Duration
	tasks     []Task
	complete  chan error
	timeout   <-chan time.Time
	interrupt chan os.Signal
}

var ErrTimeout = errors.New("timeout occured")
var ErrInterrupt = errors.New("interrupt occured")

func New(timeout time.Duration) *Runner {
	return &Runner{
		t:         timeout,
		tasks:     make([]Task, 0),
		complete:  make(chan error),
		interrupt: make(chan os.Signal),
	}
}

func (r *Runner) Add(task Task) {
	r.tasks = append(r.tasks, task)
}

func (r *Runner) Start() error {
	r.timeout = time.After(r.t)
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.interrupt:
		return ErrInterrupt
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		fmt.Println("Interrupt received.. exiting")
		return true
	default:
		return false
	}
}
