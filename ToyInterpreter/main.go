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

	// put out a bunch of instructions
	instructionArray := make([]string, 0)

	for i := 0; i < len(array); i++ {
		stripped := strings.ReplaceAll(array[i], " ", "")
		if strings.Contains(stripped, assignment) && strings.Contains(stripped, addition) {
			fmt.Println("This is an addition")

		} else if strings.Contains(stripped, assignment) {
			fmt.Println("this is an assignment")
			a := string(stripped[strings.Index(stripped, assignment)-1])
			b := string(stripped[strings.Index(stripped, assignment)+1])
			instructionArray = append(instructionArray, a)
			instructionArray = append(instructionArray, assignment)
			instructionArray = append(instructionArray, b)

		} else if strings.Contains(stripped, addition) {
			fmt.Println("This is an addition")

		} else {
			fmt.Println("This is a return statement")
		}
	}
}

func addition(a int, b int) int {
	return a + b
}

func assignment(a int) int {
	return a
}

func main() {
	file := os.Args[1]
	reader(openFile(file))
}
