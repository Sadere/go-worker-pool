package main

import (
	"math/rand"
	"time"

	"github.com/Sadere/go-worker-pool/internal/pool"
)

func generateTasks(taskNum int) []string {
	tasks := make([]string, 0, taskNum)

	for i := 0; i < taskNum; i++ {
		tasks = append(tasks, randomString())
	}

	return tasks
}

func randomString() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	length := 10

	str := make([]byte, length)

	for i := 0; i < length; i++ {
		str[i] = charset[rand.Intn(len(charset))]
	}

	return string(str)

}

func main() {
	// Generate first tasks
	tasks := generateTasks(10)

	// Create new worker pool
	wp := pool.NewWorkerPool(5)

	// Run first tasks
	wp.RunTasks(tasks)

	// Wait for tasks to complete
	time.Sleep(time.Second)

	// Remove some workers
	wp.RemoveWorker(1)
	wp.RemoveWorker(3)

	// Add new worker
	wp.AddWorker()

	// Generate new tasks
	tasks = generateTasks(20)

	// Run second batch of tasks
	wp.RunTasks(tasks)

	// Wait for completion
	time.Sleep(time.Second)
}
