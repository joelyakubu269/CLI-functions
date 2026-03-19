package main

import (
	"fmt"

	"time"
)

func AddTask(description string) {
	tasks := loadTasks()

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	id := maxID + 1

	task := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, task)
	saveTasks(tasks)
	fmt.Printf("Task with ID %d added successfully\n", id)
}

func ListTasks() {
	tasks := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks available")
		return
	}
	for _, r := range tasks {
		fmt.Printf(
			"ID: %d | %s | status: %s | created: %s\n",
			r.ID,
			r.Description,
			r.Status,
			r.CreatedAt.Format("2006-01-02 15:04:05"),
		)
	}
}
func deleteTask(id int) {
	tasks := loadTasks()
	var updated []Task
	found := false
	for _, r := range tasks {
		if r.ID != id {
			updated = append(updated, r)
		} else {
			found = true
		}
	}
	if !found {
		fmt.Printf("ID of %d not found\n", id)
		return
	}
	saveTasks(updated)
	fmt.Printf("Task with ID %d deleted successfully\n", id)
}
func updateTask(id int, newDescription string) {
	tasks := loadTasks()
	found := false
	for i, r := range tasks {
		if r.ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Task with id %d not found\n", id)
		return
	}
	saveTasks(tasks)
	fmt.Printf("Tasks with ID %d successfully updated\n", id)
}
func markTaskDone(id int) {
	tasks := loadTasks()
	found := false
	for i, r := range tasks {
		if r.ID == id {
			tasks[i].Status = "Done"
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("The task with ID of %d does not exist", id)
		return
	}
	saveTasks(tasks)

	fmt.Printf("Task with ID %d marked as done\n", id)

}
func markTaskInprogress(id int) {
	tasks := loadTasks()
	found := false
	for i, r := range tasks {
		if r.ID == id {
			tasks[i].Status = "in-Progress"
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Task with ID of %d is not found\n", id)
		return
	}
	saveTasks(tasks)
}
func ListByStatus(status string) {
	tasks := loadTasks()
	found := false
	for _, r := range tasks {
		if r.Status == status {
			fmt.Printf("ID: %d | %s | status: %s\n", r.ID, r.Description, r.Status)
			found = true
		}
	}
	if !found {
		fmt.Printf("No tasks with status: %s\n", status)
		return
	}
}
