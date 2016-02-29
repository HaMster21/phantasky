package main

import (
	t "github.com/hamster21/phantasky/tasks"
	"os"
	"log"
)

func main() {
	for _, file := range os.Args {
		task, err := t.FromFile(file)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("Description: %s", task["description"])
		log.Printf("%v\n", task)
	}
}
