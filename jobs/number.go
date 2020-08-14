package jobs

import (
	"fmt"
	"io/ioutil"
	"log"
	mustardcore "mustard/core"
	"time"

	"net/http"
)

type number struct {
	Trivia string `json:"trivia"`
}

var numberURL = "http://numbersapi.com/%d"

func init() {
	mustardcore.GetFactory().Advertise("number", func() {
		now := time.Now()
		resp, err := http.Get(fmt.Sprintf(numberURL, now.Day()))
		if err != nil {
			log.Println(err)
			return
		}
		defer resp.Body.Close()
		text, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return
		}
		data := mustardcore.EventData{Event: "clockWidget", Data: number{Trivia: string(text)}}
		mustardcore.GetEventsManager().Notify(data)
	})
}
