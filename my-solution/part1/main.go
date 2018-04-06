// Solution to part 1 of the Whispering Gophers code lab.
// This program reads from standard input and writes JSON-encoded messages to
// standard output. For example, this input line:
//	Hello!
// Produces this output:
//	{"Body":"Hello!"}
//
package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Body string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	encoder := json.NewEncoder(os.Stdout)
	for scanner.Scan() {
		text := scanner.Text()
		message := Message{Body: text}
		err := encoder.Encode(message)

		if err != nil {
			log.Fatal(err)
		}
	}
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
