package jobs

import (
	mustardcore "mustard/core"
	"regexp"

	"github.com/mmcdole/gofeed"
)

type blogItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"url"`
}

type blogData struct {
	Title string     `json:"title"`
	Items []blogItem `json:"items"`
}

var blogURL = "https://www.goavega.com/feed"

func init() {
	mustardcore.AddJob("@every 1h", func() {
		re := regexp.MustCompile(`<[^>]*>`)
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(blogURL)
		totalItems := len(feed.Items)
		takeItems := totalItems
		if totalItems > 15 {
			takeItems = 15
		}
		slicedItems := feed.Items[1:takeItems]
		var (
			blog  blogData
			items []blogItem
		)

		blog = blogData{Title: feed.Title}
		for _, item := range slicedItems {
			items = append(items, blogItem{Title: item.Title, Description: re.ReplaceAllString(item.Description, ""),
				Link: mustardcore.QRAsDataURI(item.Link)})
		}
		blog.Items = items
		data := mustardcore.EventData{Event: "blogRoll", Data: blog}
		mustardcore.SseNotify(data)
	})
}
