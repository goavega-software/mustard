package jobs

import (
	mustardcore "mustard/core"
)

type slide struct {
	url string
}
type slideshow struct {
	items []slide
}

func init() {
	mustardcore.AddJob("@every 1h", func() {
		flickrUser := mustardcore.GetEnvVariables().FlickrUserID
		var items []slide
		flickr := mustardcore.Flickr{UserID: flickrUser}
		slideshow := slideshow{}
		for _, item := range flickr.Get() {
			items = append(items, slide{url: item.GetUrl()})
		}
		slideshow.items = items
		data := mustardcore.EventData{Event: "slideshow", Data: slideshow}
		mustardcore.SseNotify(data)

	})
}
