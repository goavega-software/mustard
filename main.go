package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	mustardcore "mustard/core"
	events "mustard/core/events"
	_ "mustard/jobs"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type requestPayload struct {
	Path string `json:"path"`
}

type widget struct {
	Component string      `json:"component"`
	Class     []string    `json:"class"`
	Props     interface{} `json:"props"`
	State     interface{} `json:"state"`
}

type dashboard struct {
	Id      string   `json:"id"`
	Widgets []widget `json:"widgets"`
}

type configuration struct {
	Dashboards []dashboard            `json:"dashboards"`
	Jobs       []mustardcore.Schedule `json:"jobs"`
}

func parseArgs() (bool, string) {
	gistIDPtr := flag.String("gist", "", "gist id to install")
	flag.Parse()
	fmt.Println("gist", *gistIDPtr)
	return len(*gistIDPtr) > 0, *gistIDPtr
}

func installGist(gistID string) {
	gist := mustardcore.Gist{ID: gistID}
	gist.Install()
}

func readConfig() []byte {
	jsonFile, err := os.Open("./config/config.json")
	if err != nil {
		log.Println(err)
		return []byte("[]")
	}
	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
		return []byte("[]")
	}
	return bytes
}

func main() {
	hasArgs, gistID := parseArgs()
	if hasArgs {
		installGist(gistID)
		return
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		ExposeHeaders:    []string{"*"},
	}))
	e.Static("/js", "js")
	e.Static("/css", "css")
	e.File("/dashboard/*", "views/index.html")
	eventsManager := mustardcore.GetEventsManager()
	eventsManager.Init(e)
	eventListener := events.EventListener{URL: os.Getenv("KAFKA_URL"), Topic: os.Getenv("KAFKA_TOPIC"), GroupID: "mustard"}
	eventListener.Start()
	// read config to get widgets
	config := readConfig()
	var result configuration
	json.Unmarshal([]byte(config), &result)
	mustardcore.GetFactory().Process(result.Jobs)

	defer eventsManager.Destroy()
	defer mustardcore.DestroyJobs()

	e.POST("/api/nudge", func(c echo.Context) error {
		mustardcore.FireImmediately()
		return c.String(http.StatusOK, "Hello mustard")
	})

	e.POST("/api/layout", func(c echo.Context) (err error) {
		request := new(requestPayload)
		if err = c.Bind(request); err != nil {
			return
		}
		for _, f := range result.Dashboards {
			if f.Id == request.Path {
				return c.JSON(http.StatusOK, result.Dashboards[0].Widgets)
			}
		}
		return
	})

	e.POST("/api/webhook", func(c echo.Context) (err error) {
		apiKey := os.Getenv("API_KEY")
		headerKey := c.Request().Header.Get("X_API_KEY")
		if apiKey == "" || headerKey != b64.StdEncoding.EncodeToString([]byte(apiKey)) {
			return c.String(http.StatusUnauthorized, "Invalid key")
		}
		eventMessage := new(mustardcore.EventData)
		if err = c.Bind(eventMessage); err != nil || eventMessage.Data == nil {
			return c.String(http.StatusUnprocessableEntity, "Invalid request body")
		}
		mustardcore.GetEventsManager().Notify(*eventMessage)
		return c.String(http.StatusAccepted, "Ok")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), e))
}
