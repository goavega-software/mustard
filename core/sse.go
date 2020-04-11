package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

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
SseNotify sends SSE with payload
*/
func SseNotify(payload EventData) {
	if server == nil {
		return
	}
	go func() {
		payload, _ := jsonMarshal(payload, true)
		fmt.Println(string(payload))
		message := strings.ReplaceAll(string(payload), "%", "%%")
		server.SendMessage(fmt.Sprintf("/events/%s", "dashboard"), sse.SimpleMessage(message))
	}()
}

func jsonMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
	b, err := json.Marshal(v)

	if safeEncoding {
		b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
		b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
		b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	}
	return b, err
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
