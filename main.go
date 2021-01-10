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
	job := jobmanager.NewJob()
	err := job.Do(func() {
		fmt.Println("Hello world!")
	})
	if err != nil {
		panic(err)
	}

	jobManager.RunJobAndWait(job)

	job2, err := jobManager.RunAndWait(func() { fmt.Printf("Hello world!") })

	fmt.Println(job2.Status)
}
