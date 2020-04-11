package jobs

import (
	mustardcore "mustard/core"
	"os"
)

func init() {
	mustardcore.AddJob("@every 4h", func() {
		var items []slide
		s3Images := mustardcore.S3Images{Bucket: os.Getenv("AWS_S3_BUCKET"), Region: os.Getenv("AWS_S3_REGION"), Prefix: os.Getenv("")}
		slideshow := slideshow{}
		for _, item := range s3Images.Get() {
			items = append(items, slide{URL: item})
		}
		slideshow.Items = items
		data := mustardcore.EventData{Event: "slideshow", Data: slideshow}
		mustardcore.SseNotify(data)

	})
}
