package core

import (
	"encoding/json"
	"log"

	"github.com/goavega-software/restful"
)

type JobFactory struct {
}

type Api struct {
	Url           string `json:"url"`
	Method        string `json:"method"`
	Transform     string `json:"transform"`
	Authorization string `json:"authorization"`
	Body          string `json:"body"`
}
type Schedule struct {
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
	Api      Api    `json:"api"`
}

var factory *JobFactory
var advertisedJobs map[string]func() = make(map[string]func())

/*
Process parses the json and only schedules enabled jobs
*/
func (jf JobFactory) Process(schedules []Schedule) {
	for _, item := range schedules {
		if item.Api.Url != "" {
			localItem := item
			AddJob(item.Schedule, func() {
				log.Printf("executing %s at %s", localItem.Name, localItem.Api.Url)
				options := restful.Options{}
				options.Method = localItem.Api.Method
				options.Headers = make(map[string]string)
				options.Headers["Content-Type"] = "application/json"
				if localItem.Api.Authorization != "" {
					options.Headers["Authorization"] = localItem.Api.Authorization
				}
				options.Transformer = localItem.Api.Transform
				options.Payload = localItem.Api.Body
				message, _ := restful.Call(localItem.Api.Url, &options)
				event := EventData{}
				json.Unmarshal([]byte(message), &event)
				GetEventsManager().Notify(event)
			})
			continue
		}
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
