package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func main() {
	noOfJobs := 100

	go allocate(noOfJobs)

	_, ok := <-jobs

	fmt.Println("The jobs chan if open? ", ok)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 15
	createWorkerPool(noOfWorkers)

	fmt.Println("The jobs chan if open? ", ok)

	<-done
}

func createWorkerPool(noWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noWorkers; i++ {
		wg.Add(1)
		go worker(&wg)

	}
	wg.Wait()

	close(results)

}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randno := rand.Intn(999)
		job := Job{i, randno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum if digits %d \n", result.job.id, result.job.randno, result.sum)
	}
	done <- true
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randno)}
		results <- output
	}

	wg.Done()
}

func digits(number int) int {
	sum := 0
	no := number

	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}

	time.Sleep(500 * time.Millisecond)

	return sum

}

type Job struct {
	id     int
	randno int
}

type Result struct {
	job Job
	sum int
}
