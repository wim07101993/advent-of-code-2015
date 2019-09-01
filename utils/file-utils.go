package utils

import (
	"bufio"
	"io/ioutil"
	"os"
)

func ReadFileAsBytes(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}

func ReadFileAsString(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	return string(bytes), err
}

func ReadFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	return content, scanner.Err()
}
