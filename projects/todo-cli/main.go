package main

import (
	"fmt"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

func main() {
	var tasks []Task
	var choice int

	for {
		fmt.Println("To-Do List CLI")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Complete")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			tasks = addTask(tasks)
		case 2:
			listTasks(tasks)
		case 3:
			tasks = completeTask(tasks)
		case 4:
			tasks = deleteTask(tasks)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addTask(tasks []Task) []Task {
	var title string
	fmt.Print("Enter task title: ")
	fmt.Scan(&title)

	task := Task{
		ID:    len(tasks) + 1,
		Title: title,
		Done:  false,
	}
	tasks = append(tasks, task)
	fmt.Println("Task added!")
	return tasks
}

func listTasks(tasks []Task) {
	fmt.Println("To-Do List:")
	for _, task := range tasks {
		status := "Incomplete"
		if task.Done {
			status = "Complete"
		}
		fmt.Printf("ID: %d | Title: %s | Status: %s\n", task.ID, task.Title, status)
	}
}

func completeTask(tasks []Task) []Task {
	var id int
	fmt.Print("Enter task ID to mark as complete: ")
	fmt.Scan(&id)

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Done = true
			fmt.Println("Task marked as complete!")
			return tasks
		}
	}

	fmt.Println("Task not found.")
	return tasks
}

func deleteTask(tasks []Task) []Task {
	var id int
	fmt.Print("Enter task ID to delete: ")
	fmt.Scan(&id)

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted!")
			return tasks
		}
	}

	fmt.Println("Task not found.")
	return tasks
}
