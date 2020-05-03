package main

import (
	"crypto/md5"
	"io"
	"log"
	"os"
)

func main(path string) {
	// Read the file, if file doesn't exist throw error
	file, err := os.Open("./1c53f0a1a717e2462c29c1170911927d.png")
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

}
