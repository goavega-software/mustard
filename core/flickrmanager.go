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

var flickrURL = "https://www.flickr.com/services/rest/?per_page=15&method=flickr.photos.search&format=json&sort=interestingness-desc&api_key=%s&nojsoncallback=1"

type Flickr struct {
	Q      string
	UserID string
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

func (f Flickr) Get() []ApiPhotoResponse {
	apiKey := GetEnvVariables().FlickrKey
	filter := ""
	if f.UserID != "" {
		filter += "&user_id=" + f.UserID
	}
	if f.Q != "" {
		filter += "&text=" + url.QueryEscape(f.Q)
	}
	query := fmt.Sprintf(flickrURL, apiKey) + filter
	log.Println("flickr query", query)
	resp, err := http.Get(query)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	apiResponse := apiSearchResponse{}
	json.Unmarshal(text, &apiResponse)
	return apiResponse.Photos.Photo
}

func TakeOne(photos []ApiPhotoResponse) ApiPhotoResponse {
	return photos[getRand(len(photos)-1)]
}

func (p ApiPhotoResponse) GetUrl() string {
	return fmt.Sprintf("https://farm%d.staticflickr.com/%s/%s_%s_b.jpg", p.Farm, p.Server, p.ID, p.Secret)
}

func getRand(max int) int {
	return rand.Intn(max)
}
