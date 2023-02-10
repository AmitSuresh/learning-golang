package main

import (
	"fmt"
	"log"
	"os"
)

type Job struct {
	Command string
	*log.Logger
}

func NewJob(command string, logger *log.Logger) *Job {
	return &Job{command, logger}
}

func (job *Job) Printf(format string, args ...interface{}) {
	job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}

func main() {
	job := &Job{"adfg", log.New(os.Stderr, "Job: ", log.Ldate)}
	fmt.Println(job)
	job.Logger.Print()

	fmt.Println(NewJob("something", log.New(os.Stderr, "Job: ", log.Ldate)))
	job2 := NewJob("something", log.New(os.Stderr, "Job: ", log.Ldate))
	job2.Logger.Print()
}
