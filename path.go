package main

import (
	"bufio"
	"crypto/md5"
	"io"
	"log"
	"os"
)

func main() {
	// Read text file
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Println("There was an error reading the file", err)
	}

	// Defer closing of file, because now the md5 checksum needs to be checked
	defer file.Close()

	// Read file contents
	s := bufio.NewScanner(file)

	// Reads text from source file "file.txt"
	scantext := s.Scan()

	for scantext {
		// Reads text from the most recent call generated to Scan()
		// From my understanding, s.Scan() stores the text in a buffer
		// and now it can be read by calling s.Text()
		log.Printf("Here is the text from the file: %s", s.Text())

		// Return to stop it printing a million times
		return
	}

	// Return new hash.Hash to compute the md5 checksum of file
	hash := md5.New()

	// Copy "file" from src (reader) to "hash" dst (writer)
	_, err = io.Copy(hash, file)
	if err != nil {
		log.Println("Could not compute md5 checksum, error copying", err)
	}

	// Get md5 checksum of the file data
	log.Printf("md5 checksum: %x", hash.Sum(nil))
}
