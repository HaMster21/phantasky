package main

import (
	"encoding/json"
	"bufio"
	"os"
	"log"
)

type Task map[string]interface{}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	rawTask, err := stdin.ReadBytes('\n')
	if err != nil {
		log.Println("Error during task reading")
		log.Fatal(err)
	}

	var preTask interface{}
	if err := json.Unmarshal(rawTask, &preTask); err != nil {
		log.Println("Error during json parsing")
		log.Fatal(err)
	}

	var task Task = preTask.(map[string]interface{})

	log.Printf("Description: %s", task["description"])
	log.Printf("%v\n", task)
}
