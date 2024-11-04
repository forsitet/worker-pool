package main

type Task struct {
	Data  interface{}
	funct func(data any, WId int)
}

func NewTask(f func(data any, WId int), data interface{}) *Task {
	return &Task{funct: f, Data: data}
}

func runTask(workerID int, task *Task) {
	task.funct(task.Data, workerID)
}
