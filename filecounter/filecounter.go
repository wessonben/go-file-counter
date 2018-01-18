package filecounter

import (
	"fmt"
	"io/ioutil"
)

// CountFilesInDirectory counts the files in a directory and all
// sub-directories using recursion. Demonstrates concurrent go routines
// and channel synchronisation.
func CountFilesInDirectory(directory string, parentChannel chan int) {

	localChannel := make(chan int)
	directories, files := getDirectoryItems(directory)
	fileCount := len(files)

	// Now process the sub-directories as new go-routines
	for _, name := range directories {
		subDirectory := directory + "\\" + name
		fmt.Printf("Found directory %s\n", subDirectory)
		go CountFilesInDirectory(subDirectory, localChannel)
	}

	// Wait for all the file count messages
	for i := 0; i < len(directories); i++ {
		fileCount += <-localChannel
	}

	// Send total file count up to the parent routine
	parentChannel <- fileCount
}

// GetDirectoryItems returns the list of sub-directories and files within a
// directory.  Demonstrates returning multiple values from a function and
// adding items to a slice.
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
