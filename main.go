package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	directory := ""
	

	fmt.Print("Enter directory or leave blank for C:\\Temp: ")
	fmt.Scanln(&directory)
	
	if directory == "" {
		directory = "C:\\Temp"
	}

	_, err := ioutil.ReadDir(directory)
	
	if err == nil {
		fmt.Printf("Counting files in %s\n", directory)
		fileCountChannel := make(chan int)
		go countFilesInFirstDirectory(directory, fileCountChannel)
		fileCount := <- fileCountChannel
		fmt.Printf("Found %d files in %s\n", fileCount, directory)
	} else {
		fmt.Printf("Invalid directory: %s\n", directory)
	}
}

func countFilesInFirstDirectory(directory string, fileCountChannel chan int) {

	count := 0
	files, _ := ioutil.ReadDir(directory)

	for _, file := range files {
		if file.IsDir() {
			subDirectory := directory + "\\" + file.Name()
			fmt.Printf("Found directory %s\n", subDirectory)
			countFilesInSubDirectory(subDirectory, &count)
		} else {
			count++
		}
	}

	fileCountChannel <- count
}

func countFilesInSubDirectory(directory string, count *int) {
	files, _ := ioutil.ReadDir(directory)
	for _, file := range files {
		if file.IsDir() {
			subDirectory := directory + "\\" + file.Name()
			fmt.Printf("Found directory %s\n", subDirectory)
			countFilesInSubDirectory(subDirectory, count)
		} else {
			*count++
		}
	}
}