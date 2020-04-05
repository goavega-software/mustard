package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	mustardcore "mustard/core"
	"net/http"
)

type placeData struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type tempData struct {
	Temp    float32 `json:"temp"`
	TempMin float32 `json:"temp_min"`
	TempMax float32 `json:"temp_max"`
}
type weatherData struct {
	Weather []placeData `json:"weather"`
	Main    tempData    `json:"main"`
}

type outputData struct {
	Temp        float32 `json:"temp"`
	Min         float32 `json:"min"`
	Max         float32 `json:"max"`
	Description string  `json:"desc"`
	Icon        string  `json:"icon"`
	Image       string  `json:"image"`
}

var weatherAPIURL = "http://api.openweathermap.org/data/2.5/weather?id=%d&appid=%s&units=metric"

func init() {
	mustardcore.AddJob("@every 1h", func() {
		placeID := mustardcore.GetEnvVariables().OpenWeatherPlaceID
		apiKey := mustardcore.GetEnvVariables().OpenWeatherKey
		resp, err := http.Get(fmt.Sprintf(weatherAPIURL, placeID, apiKey))
		if err != nil {
			log.Fatal(err)
			return
		}
		defer resp.Body.Close()
		text, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		weather := weatherData{}
		err = json.Unmarshal(text, &weather)
		if err != nil {
			log.Fatal(err)
			return
		}
		weatherData := outputData{Max: weather.Main.TempMax, Min: weather.Main.TempMin, Temp: weather.Main.Temp}
		weatherData.Description = weather.Weather[0].Description
		weatherData.Icon = fmt.Sprintf("http://openweathermap.org/img/wn/%s@2x.png", weather.Weather[0].Icon)
		// get the flickr data
		flickr := mustardcore.Flickr{Q: weatherData.Description}
		photos := flickr.Get()
		log.Println("photos are")
		log.Println(photos)
		weatherData.Image = mustardcore.TakeOne(photos).GetUrl()
		data := mustardcore.EventData{Event: "weather", Data: weatherData}
		mustardcore.SseNotify(data)

	})
}
