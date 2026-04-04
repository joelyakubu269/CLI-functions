package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("There are not enough arguements")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)

		description := addCmd.String("description", "", "add a description of the task")
		addCmd.Parse(os.Args[2:])
		AddTask(*description)
	case "list":
		ListTasks()
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("usage: deelete <Id>")
			return
		}

		id := parseID(os.Args[2])
		deleteTask(id)
	case "update":
		if len(os.Args) != 4 {
			fmt.Println("usage: update <id> <description>")
			return
		}
		id := parseID(os.Args[2])
		description := os.Args[3]
		updateTask(id, description)
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: done <id>")
			return
		}
		id := parseID(os.Args[2])
		markTaskDone(id)
	case "in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: done <id>")
			return
		}
		id := parseID(os.Args[2])
		markTaskInprogress(id)
	case "done-list":
		ListByStatus("done")

	case "progress-list":
		ListByStatus("in-progress")

	case "todo-list":
		ListByStatus("todo")
	default:
		fmt.Println("Unknown command:", command)

	}

}
