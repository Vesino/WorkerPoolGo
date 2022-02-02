package main

import "fmt"

// https://gobyexample.com/worker-pools

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}
	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fibonnacci with id %d\n", id, job)
		fib := Fibonnaci(job)
		fmt.Printf("Worker with id %d, job %d, and fibonacci %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonnaci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonnaci(n-1) + Fibonnaci(n-2)
}
