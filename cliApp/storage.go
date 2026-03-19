package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func loadTasks() []Task {
	file, err := os.ReadFile("tasks.json")
	if os.IsNotExist(err) {
		err = os.WriteFile("tasks.json", []byte("[]"), 0644)
		if err != nil {
			log.Fatal(err)
		}

		file = []byte("[]")
	} else if err != nil {
		log.Fatal(err)

	}
	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("There was an error parsing into tasks")
		log.Fatal(err)

	}
	return tasks

}
func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
