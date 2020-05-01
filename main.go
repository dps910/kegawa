package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var client = http.Client{
	// to comply with e621 rate limits of 2 reqs per second
	Timeout: 5 * time.Second,
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

	// ioutil used to read the response :D
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s", body)
}

func main() {
	GetByMD5("d91986beca5ca13d88b200109a412d24")
}
