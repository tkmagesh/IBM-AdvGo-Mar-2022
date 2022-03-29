package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

type Pool struct {
	factory   func() (io.Closer, error)
	resources chan io.Closer
	mutex     *sync.Mutex
	closed    bool
}

var ErrInvalidPoolSize = errors.New("invalid pool size")
var ErrPoolClosed = errors.New("pool closed")

func New(factory func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, ErrInvalidPoolSize
	}
	return &Pool{
		factory:   factory,
		resources: make(chan io.Closer, size),
		mutex:     &sync.Mutex{},
		closed:    false,
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	select {
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("Acquire : From Pool")
		return r, nil
	default:
		fmt.Println("Acquire : From Factory")
		return p.factory()

	}
}

func (p *Pool) Release(r io.Closer) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		r.Close()
		return ErrPoolClosed
	}

	select {
	case p.resources <- r:
		fmt.Println("Release : In pool")
		return nil
	default:
		fmt.Println("Release : Close & discard the resource")
		return r.Close()
	}

}

func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}
