package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	assignment := "="
	addition := "+"
	for i := 0; i < len(array); i++ {
		if strings.Contains(array[i], assignment) && strings.Contains(array[i], addition) {
			fmt.Println("This is an addition")

		} else if strings.Contains(array[i], assignment) {
			fmt.Println("this is an assignment")

		} else if strings.Contains(array[i], addition) {
			fmt.Println("This is an addition")

		} else {
			fmt.Println("This is a return statement")
		}
	}
}

func main() {
	file := os.Args[1]
	reader(openFile(file))
}
