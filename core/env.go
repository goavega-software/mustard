package core

import (
	"log"
	"os"
	"strconv"
)

type EnvVariables struct {
	FlickrKey          string
	OpenWeatherKey     string
	OpenWeatherPlaceID int
	FlickrUserID       string
}

var envVars *EnvVariables

func (e *EnvVariables) Load() {
	log.Println(" in load method")
}

func GetEnvVariables() EnvVariables {
	if envVars == nil {
		envVars = &EnvVariables{}
	}
	envVars.FlickrKey = os.Getenv("FLICKR_API_KEY")
	log.Println(envVars.FlickrKey)
	envVars.OpenWeatherKey = os.Getenv("OPENWEATHER_API_KEY")
	envVars.OpenWeatherPlaceID, _ = strconv.Atoi(os.Getenv("OPENWEATHER_PLACE_ID"))
	envVars.FlickrUserID = os.Getenv("FLICKR_USER_ID")
	return *envVars
}
