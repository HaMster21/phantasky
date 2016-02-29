package tasks

import (
	"encoding/json"
	"bufio"
	"os"
	"log"
)

type Task map[string]interface{}

func FromFile(path string) (Task, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	f := bufio.NewReader(file)
	rawTask, err := f.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	var preTask interface{}
	if err := json.Unmarshal(rawTask, &preTask); err != nil {
		return nil, err
	}
	return preTask.(map[string]interface{}), nil
}
