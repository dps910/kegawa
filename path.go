package main

import (
	"io/ioutil"
	"log"
	"os"
)

func cli() {
	if len(os.Args) < 2 {
		log.Println("You must provide a path")
	}

	data, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Println("Could not read path", err)
	}
	log.Printf("%s", data)
}

func main() {
	cli()
}
