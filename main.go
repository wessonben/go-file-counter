package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	count := 0
	directory := ""

	fmt.Print("Enter directory: ")
    fmt.Scanln(&directory)

	fmt.Printf("Counting files in %s\n", directory)
	countFilesInDirectory(directory, &count)
	fmt.Printf("Found %d files in %s\n", count, directory)
}

func countFilesInDirectory(directory string, count *int) {
	files, err := ioutil.ReadDir(directory)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				subDirectory := directory + "\\" + file.Name()
				fmt.Printf("Found directory %s\n", subDirectory)
				countFilesInDirectory(subDirectory, count)
			} else {
				*count++
			}
		}
	}
}