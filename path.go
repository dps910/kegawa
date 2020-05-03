package main

import (
	"bufio"
	"crypto/md5"
	"io"
	"log"
	"os"
)

// List of species
var species = []string{"lucario", "fox", "pokemon"}

func main() {
	// Read text file
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Println("There was an error reading the file", err)
	}

	// Defer closing of file, because now the md5 checksum needs to be checked
	defer file.Close()

	// Return new hash.Hash to compute the md5 checksum of file
	hash := md5.New()

	// Copy "file" from src (reader) to "hash" dst (writer)
	_, err = io.Copy(hash, file)
	if err != nil {
		log.Println("Could not compute md5 checksum, error copying", err)
	}

	// Get md5 checksum of the file data
	log.Printf("md5 checksum: %x", hash.Sum(nil))

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Println("oof", err)
	}

	// Read file contents
	s := bufio.NewScanner(file)

	for s.Scan() {
		// Reads text from the most recent call generated to Scan()
		// From my understanding, s.Scan() stores the text in a buffer
		// and now it can be read by calling s.Text()
		log.Printf("Here is the text from the file: %s", s.Text())

		// Return to stop it printing a million times
		return
	}

	// How to use
	// flag.String("path", "specify path to file", "./path.go -path=<PATH>")
	// flag.Parse()
}
