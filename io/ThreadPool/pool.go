package ThreadPool

import (
	"errors"
	"sync/atomic"
)

const (
	RUNNING = 0
	STOPED  = 1
)

type Pool struct {
	capacity     uint64
	runningTasks uint64
	state        int64
	taskQueue    chan *Task
	closeC       chan bool
	PaninHandler func(interface{})
}

func (pool *Pool) incrRunningT() uint64 {
	return atomic.AddUint64(&pool.runningTasks, 1)
}

func (pool *Pool) decRunningT() uint64 {
	return atomic.AddUint64(&pool.runningTasks, ^uint64(0))
}

func (pool *Pool) getWorkers() uint64 {
	return atomic.LoadUint64(&pool.runningTasks)
}

func (p *Pool) GetCap() uint64 {
	return atomic.LoadUint64(&p.capacity)
}

func newPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, errors.New("invaild pool size")
	}
	return &Pool{
		capacity:  capacity,
		state:     RUNNING,
		taskQueue: make(chan *Task, capacity),
		closeC:    make(chan bool),
	}, nil
}

func (pool *Pool) run() {
	pool.incrRunningT()

	go func() {
		defer func() {
			pool.decRunningT()
		}()
		for {
			select {
			case task, ok := <-pool.taskQueue:
				if !ok {
					return
				}
				task.handler(task.param)
			case <-pool.closeC:
				return
			}
		}
	}()
}

func (pool *Pool) submit(task *Task) error {
	if pool.state == STOPED {
		return errors.New("pool is stopped")
	}
	if pool.getWorkers() < pool.capacity {
		pool.run()
		return nil
	}
	pool.taskQueue <- task
	return nil
}

func (pool *Pool) Close() {
	pool.state = STOPED

	for len(pool.taskQueue) > 0 {

	}

	pool.closeC <- true
	close(pool.taskQueue)
}
