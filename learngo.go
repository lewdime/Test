package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

// creating a variable with type Job which is a struct type with capacity of 10
var jobs = make(chan Job, 10)

// creating a variable with type Result which is a struct type with capacity of 10
var results = make(chan Result, 10)

// function below will add the value of digits in a multi digit number like a random number
// 532, will be added as 5+3+2 and return the sum
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	// Un remark the code line below to simulate a run time process for digits func
	// time.Sleep(100 * time.Millisecond)
	return sum
}
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	// line code below will tell the caller routine or func that this func is done
	// and subtruct one unit value in the sync.WaitGroup of the calling function
	// calling function should have a created sync.WaitGroup.Add(value int)
	wg.Done()
}
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		// line below adds one unit value to the group
		wg.Add(1)
		go worker(&wg)
	}
	// line code below will wait for the wg.Done of worker func
	wg.Wait()
	// line below will clouse the result channel of type Result struct
	close(results)
}
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		// enlarge the value of randomo to get a diffetent time execution output
		randomno := rand.Intn(999)
		// Job struct type is copied to job variable
		job := Job{i, randomno}
		// job variable is sent to the jobs channel of type Job struc
		jobs <- job
	}
	close(jobs)
}
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
func main() {
	startTime := time.Now()
	// enlarge the value of noOfJobs to get a different time execution output
	noOfJobs := 10
	// allocate is called to add jobs to the jobs channel
	go allocate(noOfJobs)
	// done channel is created and passed to the result Goroutine so that it can start printing
	// the output and notify once everything has been printed
	done := make(chan bool)
	go result(done)
	// a pool of  worker Goroutines are created by the call to createWorkerPool function
	// enlarge the value of noOfWorkers to get a diffetrent time execution output
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	// then main Goroutine waits on the done channel for all the results to tbe printed
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
