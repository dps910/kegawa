package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// e621 API
type API struct {
	Posts []Posts `json:"posts"`
}

type Posts struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	File      File   `json:"file"`
}

type File struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Ext    string `json:"ext"`
	Size   int    `json:"size"`
	Md5    string `json:"md5"`
	URL    string `json:"url"`
}

// HTTP client
var client = http.Client{
	// time out of 2 seconds
	Timeout: time.Second * 2,
}

func GetByMD5(hash string) {
	// What the request will be
	request, err := http.NewRequest("GET", "https://e621.net/posts.json?tags=md5:"+hash, nil)
	if err != nil {
		log.Println("Invalid md5 hash", err)
	}

	request.Header.Set("User-Agent", "kegawa v0.1")

	// Now to actually make the request, where a response will be returned
	response, err := client.Do(request)
	if err != nil {
		log.Println("Couldn't make http request", err)
	}
	defer response.Body.Close()

	// ioutil.ReadAll used to read the response :D
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Status: ", response.Status)

	// Get data from the "Posts" struct
	a := API{}
	err = json.Unmarshal(body, &a)
	if len(a.Posts) == 0 {
		log.Println(err)
	}

	imgUrl := a.Posts[0].File.URL
	md5Hash := a.Posts[0].File.Md5
	log.Println("md5 checksum: ", md5Hash)
	log.Println("Image URL: ", imgUrl)
}

func httpserver() {
	// HTTP server
	s := &http.Server{Addr: ":8080", Handler: nil}
	http.Handle("/", http.FileServer(http.Dir("./http")))
	log.Println("Listening on port 8080")
	log.Fatal(s.ListenAndServe())
}

func main() {
	GetByMD5("d91986beca5ca13d88b200109a412d24")
	httpserver()
}
