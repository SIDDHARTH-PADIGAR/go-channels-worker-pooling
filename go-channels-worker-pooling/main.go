package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // simulate time-consuming task
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // pretend we're doubling the job value
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	//Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	//Collecting results
	for a := 1; a <= numJobs; a++ {
		fmt.Println("Results", <-results)
	}
}
