package main

import (
	"flag"
	"fmt"
	"log"
	mustardcore "mustard/core"
	_ "mustard/jobs"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

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
	e.Static("/dashboard", "views")

	mustardcore.SseInit(e)
	mustardcore.InitJobs()

	defer mustardcore.SseDestroy()
	defer mustardcore.DestroyJobs()
	port := os.Getenv("PORT")
	e.POST("/api/nudge", func(c echo.Context) error {
		mustardcore.FireImmediately()
		return c.HTML(200, "<p>Hello</p>")
	})
	if port == "" {
		port = "8090"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), e))
}
