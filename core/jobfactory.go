package core

import (
	"log"
)

type JobFactory struct {
}

type Schedule struct {
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
}

var factory *JobFactory
var advertisedJobs map[string]func() = make(map[string]func())

/*
Process parses the json and only schedules enabled jobs
*/
func (jf JobFactory) Process(schedules []Schedule) {
	for _, item := range schedules {
		value, ok := advertisedJobs[item.Name]
		if ok {
			log.Printf("%s, %s", item.Schedule, item.Name)
			AddJob(item.Schedule, value)
		}
	}
	InitJobs()
}

/*
Advertise All jobs should "advertise" themselves with a name and execute function
***/
func (jf JobFactory) Advertise(name string, executor func()) {
	advertisedJobs[name] = executor
}

/*
GetFactory a non-thread safe way of getting the factory. It's not called make since it mimics a singleton
*/
func GetFactory() *JobFactory {
	if factory == nil {
		factory = &JobFactory{}
	}
	return factory
}
