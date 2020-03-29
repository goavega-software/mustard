package core

import (
	"encoding/json"
	"fmt"

	"github.com/alexandrevicenzi/go-sse"
	"github.com/labstack/echo"
)

var server *sse.Server

type EventData struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

/*
SseInit Initializes the sse server
*/
func SseInit(e *echo.Echo) {
	server = sse.NewServer(nil)
	e.Any("/events/:channel", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		server.ServeHTTP(res, req)
		return nil
	})
}

/*
SseNotify does nothing (for now)
*/
func SseNotify(payload EventData) {
	if server == nil {
		return
	}
	go func() {
		payload, _ := json.Marshal(payload)
		fmt.Println(string(payload))
		server.SendMessage(fmt.Sprintf("/events/%s", "dunno"), sse.SimpleMessage(string(payload)))
	}()
}

/*
SseDestroy is supposed to clean up
*/
func SseDestroy() {
	if server == nil {
		return
	}
	server.Shutdown()
}
