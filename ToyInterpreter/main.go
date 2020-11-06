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

		// Checks whether the line is an addition
		if strings.Contains(stripped, assignment) && strings.Contains(stripped, addition) {

			// Substrings of each of the non operator characters in every line
			a := stripped[strings.Index(stripped, assignment)+1 : strings.Index(stripped, addition)]
			b := stripped[strings.Index(stripped, addition)+1 : len(stripped)]
			c := stripped[0:strings.Index(stripped, assignment)]

			aInt, err0 := strconv.Atoi(a)
			bInt, err1 := strconv.Atoi(b)

			// uses the errors to determine whether an input substring is a string or an int
			if err1 != nil && err0 == nil {
				variableMap[c] = aInt + variableMap[b]
			}

			if err0 != nil && err1 == nil {
				variableMap[c] = variableMap[a] + bInt
			}

			if err0 == nil && err1 == nil {
				variableMap[c] = aInt + bInt
			}

			if err0 != nil && err1 != nil {
				variableMap[c] = variableMap[a] + variableMap[b]
			}

			// checks whether the line is an assignment
		} else if strings.Contains(stripped, assignment) {

			a := stripped[0:strings.Index(stripped, assignment)]
			b := stripped[strings.Index(stripped, assignment)+1 : len(stripped)]

			bAsInt, err := strconv.Atoi(b)
			if err != nil {
				log.Fatal(err)
			}

			variableMap[a] = bAsInt

			// Here if the line calls for a return
		} else {

			c := string(stripped)

			fmt.Print("The answer is ")
			fmt.Print(variableMap[c])
		}
	}
}

// main method runs the program
func main() {
	file := os.Args[1]
	reader(openFile(file))
}
