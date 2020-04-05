package jobs

import (
	mustardcore "mustard/core"
)

type slide struct {
	URL string `json:"url"`
}
type slideshow struct {
	Items []slide `json:"items"`
}

func init() {
	mustardcore.AddJob("@every 1h", func() {
		flickrUser := mustardcore.GetEnvVariables().FlickrUserID
		var items []slide
		flickr := mustardcore.Flickr{UserID: flickrUser}
		slideshow := slideshow{}
		for _, item := range flickr.Get() {
			items = append(items, slide{URL: item.GetUrl()})
		}
		slideshow.Items = items
		data := mustardcore.EventData{Event: "slideshow", Data: slideshow}
		mustardcore.SseNotify(data)

	})
}
