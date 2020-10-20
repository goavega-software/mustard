package core

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

var jobWrapper *cron.Cron
var jobs []func()

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
FireImmediately fires all jobs added immediately without waiting for cron timer
the reason being cron only fires jobs *after* the time is elapsed which means
that first run will end up having blank data till job is run on schedule time.
It is assumed that all jobs will persist their data and that would be used to
send SSEs on cient connect.
*/
func FireImmediately() {
	for _, f := range jobs {
		go f()
	}
}

/*
AddJob adds a job
*/
func AddJob(interval string, f func()) {
	jobs = append(jobs, f)
	jobWrapper.AddFunc(interval, f)
}

/*
DestroyJobs closes the job server
*/
func DestroyJobs() {
	jobs = nil
	jobWrapper.Stop()
}
