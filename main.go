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
		fileCount := <-fileCountChannel
		elapsed := time.Since(start)
		fmt.Printf("Found %d files in %s\n", fileCount, elapsed)
	} else {
		fmt.Printf("Invalid directory: %s\n", directory)
	}
}

//-----------------------------------------------------------------------------
// Function to count the files in a directory and all sub-directories using
// recursion. Demonstrates concurrent go routines and channel synchronisation.
//-----------------------------------------------------------------------------
func countFilesInSubDirectory(directory string, parentChannel chan int) {

	localChannel := make(chan int)
	directories, files := getDirectoryItems(directory)
	fileCount := len(files)

	// Now process the sub-directories as new go-routines
	for _, name := range directories {
		subDirectory := directory + "\\" + name
		fmt.Printf("Found directory %s\n", subDirectory)
		go countFilesInSubDirectory(subDirectory, localChannel)
	}

	// Wait for all the file count messages
	for i := 0; i < len(directories); i++ {
		fileCount += <-localChannel
	}

	// Send total file count up to the parent routine
	parentChannel <- fileCount
}

//-----------------------------------------------------------------------------
// Function to return the list of sub-directories and files within a directory.
// Demonstrates returning multiple values and adding items to a slice.
//-----------------------------------------------------------------------------
func getDirectoryItems(path string) ([]string, []string) {
	directories := make([]string, 0)
	files := make([]string, 0)
	fileInfos, _ := ioutil.ReadDir(path)
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			directories = append(directories, fileInfo.Name())
		} else {
			files = append(files, fileInfo.Name())
		}
	}
	return directories, files
}
