package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	result       string
	resultWorker string
	mu           sync.Mutex
	tmpl         = template.Must(template.ParseFiles("./static/index.html"))
)

func generatestrings() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 5)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func funcForTask(data any, WId int) {
	temp := data.([]string)
	taskResult := fmt.Sprintf("Worker %d выполнил task %s: %s\n", WId, temp[0], temp[1])
	mu.Lock()
	result += taskResult
	mu.Unlock()
	time.Sleep(300 * time.Millisecond)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		result = ""
		resultWorker = ""
		r.ParseForm()

		numWorkers, err := strconv.Atoi(r.FormValue("counter"))
		if err != nil {
			fmt.Print("Error in main:handelr:numWorkers")
		}

		numTasks, err := strconv.Atoi(r.PostFormValue("taskCount"))
		if err != nil {
			fmt.Print("Error in main:handelr:taskCount")
		}

		startWorkerPool(numWorkers, numTasks)
	}
	mu.Lock()
	defer mu.Unlock()
	tmpl.Execute(w, []string{result, resultWorker})
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8080", nil)
}

func startWorkerPool(numWorkers int, numTasks int) {
	resultWorker += fmt.Sprintf("Количество воркеров: %d \nКоличетво задач: %d \n\n", numWorkers, numTasks)
	allTask := make(chan *Task, 20)

	go func() {
		for i := 1; i <= numTasks; i++ {
			allTask <- NewTask(funcForTask, []string{strconv.Itoa(i), generatestrings()})
			fmt.Printf("Task %d добавлен в канал\n", i)
		}
		close(allTask)
	}()

	runWorkers(allTask, numWorkers)
}
