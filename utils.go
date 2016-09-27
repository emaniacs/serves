package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// checking for exists directory
func existsDir(path string) bool{
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir();
}

// parse header filename
// header format is Key = Value
func parseHeader(filename string) []Header {
	var headers []Header

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return headers
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := string(scanner.Text());
		array := strings.Split(line, "=")
		if len(array) == 2 {
			key := strings.Trim(array[0], " ")
			value := strings.Trim(array[0], " ")
			headers = append(headers, Header{key, value})
		}
	}

	return headers
}
