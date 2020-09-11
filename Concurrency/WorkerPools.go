//Start 3 worker threads , assign job to them and collect the
//Final result in results
package main

import (
	"fmt"
	"time"
)

//Once a worker finishes its current job , then only it picks another one
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//It will start 3 worker threads , but they will remain blocked
	// as there wont be any jobs
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		fmt.Println("Assigning j to jobs", j)
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
