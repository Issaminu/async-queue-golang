package main

import (
	"log"
	"net/http"
	"sync"
)

// Job represents a unit of work to be processed by a worker.
type Job struct {
	r *http.Request       // HTTP request to be processed
	w http.ResponseWriter // Response writer to send the result
}

// Queue manages a list of jobs to be processed by workers.
type Queue struct {
	jobs []*Job     // List of jobs in the queue
	mu   sync.Mutex // Mutex to synchronize access to the queue
	cond *sync.Cond // Condition variable for signaling
}

var q Queue // Global instance of the queue

// Push adds a job to the queue.
func (q *Queue) Push(j *Job) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.jobs = append(q.jobs, j)
	q.cond.Signal() // Signal a waiting worker that a job is available
	log.Println("Job added to queue")
}

// Pop retrieves and removes a job from the queue.
func (q *Queue) Pop() (*Job, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.jobs) == 0 {
		q.cond.Wait() // If the queue is empty, wait for a signal
	}
	job := q.jobs[0]
	q.jobs = q.jobs[1:]
	log.Println("Job removed from queue")
	return job, true
}

// handler adds a job to the queue.
func handler(w http.ResponseWriter, r *http.Request) {
	// Create a job with the request and response.
	job := &Job{r, w}

	// Push the job onto the queue.
	q.Push(job)
	log.Println("Received request and added job to queue")
}

// init initializes the condition variable and starts worker goroutines.
func init() {
	q.cond = sync.NewCond(&q.mu)
	for i := 0; i < 4; i++ {
		go worker()
	}
}

// worker processes jobs from the queue.
func worker() {
	for {
		job, ok := q.Pop()
		if ok {
			log.Println("Worker processing job")
			doWork(job)
		}
	}
}

// doWork simulates processing a job and sends a response.
func doWork(job *Job) {
	// Extract the "Name" parameter from the request query.
	name := job.r.URL.Query().Get("Name")

	// Check if the name is not empty.
	if name != "" {
		// Real work is done here.
		flusher := job.w.(http.Flusher)
		// Send the name as the response.
		flusher.Flush()
		_, err := job.w.Write([]byte("Hello, " + name))
		flusher.Flush()
		if err != nil {
			log.Println("Error writing response:", err)
		}
		log.Println("Response sent: Hello,", name)
	} else {
		// If the "Name" parameter is missing or empty, send an error response.
		http.Error(job.w, "Name parameter is missing or empty", http.StatusBadRequest)
		log.Println("Error: Name parameter is missing or empty")
	}
}

func main() {
	http.HandleFunc("/addJob", handler)
	log.Println("Server started and listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
