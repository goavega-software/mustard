package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"net/http"
	"net/url"
)

var flickrURL = "https://www.flickr.com/services/rest/?per_page=15&method=flickr.photos.search&format=json&text=%s&sort=interestingness-desc&api_key=%s&nojsoncallback=1"

type Flickr struct {
}

type apiPhotosResponse struct {
	Photo []ApiPhotoResponse `json:"photo"`
	Pages int                `json:"pages"`
}

type ApiPhotoResponse struct {
	ID     string `json:"id"`
	Secret string `json:"secret"`
	Server string `json:"server"`
	Farm   int    `json:"farm"`
}

type apiSearchResponse struct {
	Photos apiPhotosResponse `json:"photos"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (f Flickr) Get(q string) ApiPhotoResponse {
	apiKey := GetEnvVariables().FlickrKey
	response := ApiPhotoResponse{}
	resp, err := http.Get(fmt.Sprintf(flickrURL, url.QueryEscape(q), apiKey))
	if err != nil {
		log.Fatal(err)
		return response
	}
	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return response
	}
	apiResponse := apiSearchResponse{}
	json.Unmarshal(text, &apiResponse)
	return apiResponse.Photos.Photo[getRand(len(apiResponse.Photos.Photo)-1)]
}

func (p ApiPhotoResponse) GetUrl() string {
	return fmt.Sprintf("https://farm%d.staticflickr.com/%s/%s_%s_b.jpg", p.Farm, p.Server, p.ID, p.Secret)
}

func getRand(max int) int {
	return rand.Intn(max)
}
