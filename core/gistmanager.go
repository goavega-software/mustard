package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// GistFilename represents filename on a gist.
type GistFilename string

// Gist represents a GitHub's gist.
type Gist struct {
	ID    string                    `json:"id,omitempty"`
	Files map[GistFilename]GistFile `json:"files,omitempty"`
}

// GistFile represents a file on a gist.
type GistFile struct {
	Filename string `json:"filename,omitempty"`
	Language string `json:"language,omitempty"`
	Type     string `json:"type,omitempty"`
	Content  string `json:"content,omitempty"`
}

var (
	currentPath  string
	componentDir string
	jobsDir      string
)

func init() {
	currentPath, _ = os.Getwd()
	componentDir = fmt.Sprintf("%s/client/src/components/", currentPath)
	jobsDir = fmt.Sprintf("%s/jobs/", currentPath)
}

// Install installs a gist by moving .Vue files to
// ./client/src/components & .go files to ./jobs
func (gist Gist) Install() {
	if gist.ID == "" {
		log.Fatal("Invalid gist Id provided")
		return
	}
	gistURL := fmt.Sprintf("https://api.github.com/gists/%s", gist.ID)
	resp, err := http.Get(gistURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	text, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	v := Gist{}
	e := json.Unmarshal(text, &v)
	if e != nil {
		log.Fatal("Could not deserialize json", e)
		return
	}
	for _, item := range v.Files {
		log.Println("filename is", item.Filename)
		item.install()
	}
}

func (file GistFile) install() {
	fileExt := filepath.Ext(file.Filename)
	var createPath string
	log.Println(fileExt, componentDir)
	if fileExt == ".vue" {
		// vue file
		createPath = componentDir
	}
	if fileExt == ".go" {
		createPath = jobsDir
	}
	if fileExt == ".sh" {
		createPath = currentPath + "/js/"
	}
	createPath += file.Filename
	fileBuf, _ := os.Create(createPath)
	defer fileBuf.Close()
	fileBuf.WriteString(file.Content)
	fileBuf.Sync()
}
