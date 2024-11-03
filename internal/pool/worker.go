package pool

import "fmt"

type Worker struct {
	id        int
	taskQueue <-chan string
	quitChan  chan struct{}
}

func NewWorker(id int, taskQueue <-chan string) *Worker {
	quitChan := make(chan struct{})

	return &Worker{
		id:        id,
		taskQueue: taskQueue,
		quitChan:  quitChan,
	}
}

func (w *Worker) QuitChan() chan struct{} {
	return w.quitChan
}

// Main worker method
func (w *Worker) Run() {
	fmt.Printf("Worker #%d started\n", w.id)

	for {
		select {
		case task := <-w.taskQueue:
			fmt.Printf("Worker #%d has done task = %s\n", w.id, task)
		case <-w.quitChan:
			fmt.Printf("Worker #%d quitting\n", w.id)
			return
		}
	}
}
