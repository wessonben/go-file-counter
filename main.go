package main

import (
	"fmt"
	"io/ioutil"
	"time"
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
		go countFilesInSubDirectory(directory, fileCountChannel)
		fileCount := <- fileCountChannel
		elapsed := time.Since(start)
		fmt.Printf("Found %d files in %s\n", fileCount, elapsed)
	} else {
		fmt.Printf("Invalid directory: %s\n", directory)
	}
}

func countFilesInSubDirectory(directory string, parentChannel chan int) {
	
	fileCount := 0
	directoryCount := 0
	localChannel := make(chan int)

	files, _ := ioutil.ReadDir(directory)

	// Get the number of sub-directories and files in the current directory.
	// Need to know the number of directories in advance of launching the
	// recursive go-routine so that we know how many messages to wait for on 
	// the file count channel
	for _, file := range files {
		if file.IsDir() {
			directoryCount++
		} else {
			fileCount++
		}
	}

	// Now process the sub-directories as new go-routines
	for _, file := range files {
		if file.IsDir() {
			subDirectory := directory + "\\" + file.Name()
			fmt.Printf("Found directory %s\n", subDirectory)
			go countFilesInSubDirectory(subDirectory, localChannel)
		}
	}

	// Wait for all the file count messages
	for i := 1; i <= directoryCount; i++ {
        fileCount += <- localChannel
	}
	
	// Send total file count up to the parent routine
	parentChannel <- fileCount
}