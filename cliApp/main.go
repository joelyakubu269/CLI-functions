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
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		listCmd.Parse(os.Args[2:])
		ListTasks()
	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", 0, "id of task to be deleted")
		deleteCmd.Parse(os.Args[2:])

		deleteTask(*id)
	case "update":
		upCmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := upCmd.Int("id", 0, "id of the task you want to update")
		description := upCmd.String("description", "", "description of the task you want to update")
		upCmd.Parse(os.Args[2:])

		updateTask(*id, *description)
	case "done":
		doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
		id := doneCmd.Int("id", 0, "id of the task to be marked done")
		doneCmd.Parse(os.Args[2:])
		markTaskDone(*id)
	case "in-progress":
		inProgCmd := flag.NewFlagSet("inProgress", flag.ExitOnError)
		id := inProgCmd.Int("id", 0, "id which status is to be checked")
		inProgCmd.Parse(os.Args[2:])
		markTaskInprogress(*id)
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
