package main

import (
	"fmt"
	"io/ioutil"
	"time"

	fc "go-file-counter/filecounter"
)

func main() {

	directory := ""

	fmt.Print("Enter directory or leave blank for C:\\Temp: ")
	fmt.Scanln(&directory)

	if directory == "" {
		directory = "C:\\Temp"
	}

	// verify the directory exists
	_, err := ioutil.ReadDir(directory)
	if err == nil {
		start := time.Now()
		fmt.Printf("Counting files in %s\n", directory)
		fileCountChannel := make(chan int)
		go fc.CountFilesInDirectory(directory, fileCountChannel)
		fileCount := <-fileCountChannel
		elapsed := time.Since(start)
		fmt.Printf("Found %d files in %s\n", fileCount, elapsed)
	} else {
		fmt.Printf("Invalid directory: %s\n", directory)
	}
}
