package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// GetFileReader opens file and returns io.Reader
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

// GetFileContents reads the file and returns its contents
func GetFileContents(filePath string) ([]byte, error) {
	reader, err := GetFileReader(filePath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(reader)
}

// fileExists checks if file exists
func fileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
