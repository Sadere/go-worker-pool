package pool

import (
	"errors"
	"sync"
)

var ErrInvalidWorker = errors.New("worker with provided ID doesn't exist")

type WorkerPool struct {
	mu           sync.Mutex
	lastWorkerID int
	workers      map[int]*Worker
	taskQueue    chan string
}

func NewWorkerPool(workers int) *WorkerPool {
	taskQueue := make(chan string)

	workerPool := &WorkerPool{
		mu:        sync.Mutex{},
		workers:   make(map[int]*Worker),
		taskQueue: taskQueue,
	}

	// Add workers
	for i := 0; i < workers; i++ {
		workerPool.AddWorker()
	}

	return workerPool
}

func (wp *WorkerPool) TaskQueue() chan string {
	return wp.taskQueue
}

// Adds new worker to worker pool
func (wp *WorkerPool) AddWorker() {
	wp.mu.Lock()

	defer wp.mu.Unlock()

	// Generate worker ID
	id := wp.lastWorkerID + 1

	// Create worker instance
	worker := NewWorker(id, wp.taskQueue)

	// Save worker to pool
	wp.workers[id] = worker

	// Update last index
	wp.lastWorkerID = id

	// Run worker
	go worker.Run()
}

// Remove worker with provided id
func (wp *WorkerPool) RemoveWorker(workerID int) error {
	wp.mu.Lock()

	defer wp.mu.Unlock()

	// Try to retrieve worker from map
	w, ok := wp.workers[workerID]
	if !ok {
		return ErrInvalidWorker
	}

	// Notify worker it's time to quit
	w.QuitChan() <- struct{}{}

	// Remove worker from map
	delete(wp.workers, workerID)

	return nil
}

// Run tasks
func (wp *WorkerPool) RunTasks(tasks []string) {
	for _, task := range tasks {
		wp.taskQueue <- task
	}
}
