package main

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// A API struct
type A struct {
	P []P `json:"posts"`
}

// P Posts struct
type P struct {
	File F    `json:"file"`
	Tags Tags `json:"tags"`
}

// F File struct
type F struct {
	Md5 string `json:"md5"`
}

// Tags struct
type Tags struct {
	Artist  []string `json:"artist"`
	Species []string `json:"species"`
}

var (
	// URL variable (why do I need to comment this?)
	URL = "https://e621.net/posts.json?tags=md5"
	c   = &http.Client{}
)

func check(text string, err error) {
	if err != nil {
		log.Fatal(text, err)
	}
}

func hash() {
	// Check if there are enough args
	if len(os.Args) < 2 {
		log.Println("You didn't provide enough args")
	}
	f, err := ioutil.ReadFile(os.Args[1])
	check("You didn't provide the file name:", err)

	md5hash := md5.Sum(f)
	log.Printf("md5 checksum: %x", md5hash)

	// Send GET request to URL and hash
	// If request is successful, return URL+hash
	// If not successful, error
	urlAndMd5 := fmt.Sprintf("%s:%x", URL, md5hash)
	req, err := http.NewRequest("GET", urlAndMd5, nil)
	check("Couldn't create http request:", err)

	req.Header.Add("User-Agent", "kegawa UwU")

	response, err := c.Do(req)
	check("Couldn't send HTTP request:", err)
	
	// Defer closing of response body so it can still be read
	defer response.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(response.Body)
	check("Could not read response body:", err)

	abc := A{}
	json.Unmarshal(body, &abc)

	// Print this epic data
	if len(abc.P) == 0 {
		log.Println("request did not succeed because not e621 hash.")
	} else {
		artist := abc.P[0].Tags.Artist
		species := abc.P[0].Tags.Species
		log.Printf("%s:%x", URL, md5hash)
		log.Printf("Artist: %s", strings.Join(artist, ""))
		log.Printf("Species: %s", strings.Join(species, ", "))
	}

	// Database
	db, err := sql.Open("sqlite3", "./e621.db")
	check("Couldn't find database:", err)

	// Check if database is alive (and not dead) :)
	ctx, timeout := context.WithTimeout(context.Background(), 30*time.Second)
	defer timeout()

	connection := "Success"
	if err := db.PingContext(ctx); err != nil {
		connection = "Nope"
	}
	log.Println(connection)

	// Create a table in sqlite3 :D
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS e621 (id INTEGER PRIMARY KEY, artist TEXT NOT NULL, urlMd5 TEXT NOT NULL, species TEXT NOT NULL)")
	check("Couldn't create table:", err)

	// Insert data :3
	_, err = db.Exec(
		"INSERT INTO e621 (artist, urlMd5, species) VALUES ($1, $2, $3)",
		strings.Join(abc.P[0].Tags.Artist, ""),
		urlAndMd5,
		strings.Join(abc.P[0].Tags.Species, ", "),
	)
}

func main() {
	hash()
}
