package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var instructionArray []string
var variableArray []string

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

	variableMap := make(map[string]int)

	for i := 0; i < len(array); i++ {
		stripped := strings.ReplaceAll(array[i], " ", "")
		if strings.Contains(stripped, assignment) && strings.Contains(stripped, addition) {
			a := string(stripped[strings.Index(stripped, addition)-1])
			b := string(stripped[strings.Index(stripped, addition)+1])
			c := string(stripped[strings.Index(stripped, assignment)-1])
			instructionArray = append(instructionArray, c, a, b, addition, assignment)
			variableMap[c] = variableMap[a] + variableMap[b]

		} else if strings.Contains(stripped, assignment) {
			// fmt.Println("this is an assignment")
			a := string(stripped[strings.Index(stripped, assignment)-1])
			b := string(stripped[strings.Index(stripped, assignment)+1])
			instructionArray = append(instructionArray, a, b, assignment)
			variableArray = append(variableArray, a, b)

			bAsInt, err := strconv.Atoi(b)
			if err != nil {
				log.Fatal(err)
			}

			variableMap[a] = bAsInt

		} else {
			// fmt.Println("This is a return statement")
			c := string(stripped)
			instructionArray = append(instructionArray, c)
			fmt.Print("The answer is ")
			fmt.Print(variableMap[c])
		}
	}
}

func main() {
	file := os.Args[1]
	reader(openFile(file))
}
