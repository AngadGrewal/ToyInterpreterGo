package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// openfile method - takes runtime argument of txt file
func openFile(arg string) []string {
	file, err := os.Open(arg)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fileToArray := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileToArray = append(fileToArray, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return fileToArray
}

// Takes array and decides what every line is
// assignment, function or return
func reader(array []string) {
	for i := 0; i < len(array); i++ {
		if array
	}
}


func main() {
	file := os.Args[1]
	reader(openFile(file))
}
