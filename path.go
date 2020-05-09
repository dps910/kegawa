package main

import (
	"crypto/md5"
	"io/ioutil"
	"log"
	"os"
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
}

func main() {
	hash()
}
