package feed

import (
	"sync"
)

// Task encapsulates a work item that should go in a work
// pool.
type Task struct {
	f func()
}

// Run runs a Task and does appropriate accounting via a
// given sync.WorkGroup.
func (t *Task) run(wg *sync.WaitGroup) {
	t.f()
	wg.Done()
}

// Pool is a worker group that runs a number of tasks at a
// configured concurrency.
type Pool struct {
	Tasks []*Task

	sync.RWMutex
	concurrency int
	tasksChan   chan *Task
	wg          sync.WaitGroup
}

// NewPool initializes a new pool with the given tasks and
// at the given concurrency.
func NewPool(concurrency int) *Pool {
	p := &Pool{
		Tasks:       []*Task{},
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}

	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}

	return p
}

// Size number of tasks added
func (p *Pool) Size() int {
	return len(p.Tasks)
}

// Add task ...
func (p *Pool) Add(fn func()) {
	p.Lock()
	defer p.Unlock()
	p.Tasks = append(p.Tasks, &Task{f: fn})
}

// Run runs all work within the pool and blocks until it's
// finished.
func (p *Pool) Run() {

	p.wg.Add(len(p.Tasks))
	for _, task := range p.Tasks {
		p.tasksChan <- task
	}
	//clear tasks
	p.Tasks = []*Task{}

	p.wg.Wait()
}

// Quit ...
func (p *Pool) Quit() {
	// all workers return
	close(p.tasksChan)
}

// The work loop for any single goroutine.
func (p *Pool) work() {
	for task := range p.tasksChan {
		task.run(&p.wg)
	}
}
