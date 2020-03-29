package core

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

var jobWrapper *cron.Cron

func init() {
	fmt.Println("init of job manager called")
	jobWrapper = cron.New()
}

/*
InitJobs initializes the job server
*/
func InitJobs() {
	jobWrapper.Start()
}

/*
AddJob adds a job
*/
func AddJob(interval string, f func()) {
	jobWrapper.AddFunc(interval, f)
}

/*
DestroyJobs closes the job server
*/
func DestroyJobs() {
	jobWrapper.Stop()
}
