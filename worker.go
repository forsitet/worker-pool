package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	ID       int
	taskChan chan *Task
}

func NewWorker(channel chan *Task, ID int) *Worker {
	return &Worker{
		ID:       ID,
		taskChan: channel,
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	mu.Lock()
	resultWorker += fmt.Sprintf("Starting worker %d\n", w.ID)
	mu.Unlock()

	for task := range w.taskChan {
		runTask(w.ID, task)
	}

	mu.Lock()
	resultWorker += fmt.Sprintf("Stoping worker %d\n", w.ID)
	mu.Unlock()

	wg.Done()
}

func runWorkers(Task chan *Task, numWorkers int) {
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		worker := NewWorker(Task, i)
		wg.Add(1)
		go worker.Start(&wg)
	}
	wg.Wait()
}
