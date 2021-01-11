package main

import (
	"github.com/ivanape/jobmanager"
	"log"
	"time"
)

var manager *jobmanager.JobsManager

func init() {
	manager = jobmanager.NewJobManager(10)
}

func main() {
	var job *jobmanager.Job
	var err error

	f := func(name string) (string, error) {
		log.Printf("Hello %s!\n", name)
		time.Sleep(1 * time.Second)

		return "any value", nil
	}

	// Run job as function async
	job, err = manager.Run(f, "Job")
	if err != nil {
		panic(err)
	}
	// Wait until job ends
	manager.WaitForJobs(job)
	log.Printf("Return value is: %s\n", job.Result.Value.(string))
	log.Printf("Error value is: %v\n", job.Result.Err)

	// Run and wait a function
	job, err = manager.RunAndWait(f, "Job")
	if err != nil {
		panic(err)
	}

	// Define a job as a function
	job, err = manager.NewJob(f, "Job")
	if err != nil {
		panic(err)
	}
	// Same flow to execute
	manager.RunJob(job)
	manager.WaitForJobs(job)
	manager.RunJobAndWait(job)

	// Batch job execution
	job1, _ := manager.NewJob(f, "Job 1")
	job2, _ := manager.NewJob(f, "Job 2")
	// Run both jobs in Sequentially
	manager.RunJobsInSerial(job1, job2)
	job3, _ := manager.NewJob(f, "Job 3")
	job4, _ := manager.NewJob(f, "Job 4")
	// Run both jobs in Parallel
	manager.RunJobsInParallel(job3, job4)

	// Logical job aggregation
	job5, _ := manager.NewJob(f, "Job 5")
	job6, _ := manager.NewJob(f, "Job 6")
	err = manager.CreateGroup("groupName", job5, job6)
	manager.RunJobsInSerial(job5, job6)
	jobs, err := manager.GetJobsByGroup("groupName")
	for _, j := range jobs {
		log.Printf("Job is %s status is %d\n", j.ID, j.Status)
	}
}
