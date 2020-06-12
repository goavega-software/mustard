package jobs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	mustardcore "mustard/core"

	"net/http"
)

type comparer struct {
	NewVal int `json:"newVal"`
	OldVal int `json:"oldVal"`
}

var dataURL = "https://pomber.github.io/covid19/timeseries.json"
var country = "India"

func init() {
	mustardcore.AddJob("@every 6h", func() {
		resp, err := http.Get(dataURL)
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
		var v interface{}
		json.Unmarshal(text, &v)
		data := v.(map[string]interface{})[country]
		dataInterface := data.([]interface{})
		totalItems := len(dataInterface)
		newData := dataInterface[totalItems-1].(map[string]interface{})
		oldValue := 0
		if totalItems > 1 {
			oldValue = int(dataInterface[totalItems-2].(map[string]interface{})["confirmed"].(float64))
		}
		newValue := int(newData["confirmed"].(float64))
		eventData := mustardcore.EventData{Event: "comparisonWidget", Data: comparer{OldVal: oldValue, NewVal: newValue}}
		mustardcore.GetEventsManager().Notify(eventData)
	})
}
