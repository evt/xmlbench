package file

import (
	"fmt"
	"io"
	"os"
)

// GetFileReader opens XML hotels file and returns reader
func GetFileReader(filePath string) (io.Reader, error) {
	filePath = os.Getenv("GOPATH") + "/src/xmlbench/file/" + filePath
	if !fileExists(filePath) {
		return nil, fmt.Errorf("File %s doesn't exist", filePath)
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// fileExists checks if file exists
func fileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
