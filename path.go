package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type A struct {
	P []P `json:"posts"`
}

type P struct {
	ID int `json:"id"`
}

var (
	URL = "https://e621.net/posts.json?tags=md5"
	c   = &http.Client{}
)

func hash() {
	// Check if there are enough args
	if len(os.Args) < 2 {
		log.Println("You didn't provide enough args")
	}
	f, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Println("You didn't provide the file name")
	}

	md5hash := md5.Sum(f)
	log.Printf("md5 checksum: %x", md5hash)

	// Send GET request to URL and hash
	// If request is successful, return URL+hash
	// If not successful, error
	urlAndMd5 := fmt.Sprintf("%s:%x", URL, md5hash)
	req, err := http.NewRequest("GET", urlAndMd5, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Add("User-Agent", "kegawa")

	response, err := c.Do(req)
	if err != nil {
		log.Println("Couldn't send HTTP request. This is not an md5 from e621")
	}
	// Defer closing of response body so it can still be read
	defer response.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Could not read response body", err)
	}

	abc := A{}
	json.Unmarshal(body, &abc)

	// Print this epic data
	if len(abc.P) == 0 {
		log.Println("request did not succeed because not e621 hash. Or maybe the server is down and the furries made afucky wucky UWU")
	} else {
		log.Printf("%s:%x", URL, md5hash)
	}
}

func main() {
	hash()
}
