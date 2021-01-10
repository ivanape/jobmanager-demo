package main

import (
	"fmt"
	"github.com/ivanape/jobmanager"
)

var jobManager *jobmanager.JobsManager

func init() {
	jobManager = jobmanager.NewJobManager(1)
}

func main() {
	job, err := jobmanager.NewJob(func() {
		fmt.Println("Hello world!")
	})
	if err != nil {
		panic(err)
	}

	jobManager.RunJobAndWait(job)

	job2, err := jobManager.RunAndWait(func() { fmt.Println("Hello world!") })

	jobManager.RunJobsInSequence(job, job2)
	jobManager.RunJobsInParallel(job, job2)

	fmt.Println(job2.Status)
}
