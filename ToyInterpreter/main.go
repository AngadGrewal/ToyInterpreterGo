package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// bufio is used to read a file line by line this implements buffered
// io. Wraps an io.Reader or io.Writer Object, creating another object
// that also implements the interface but provides buffering and some
// help for textual I/O.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
