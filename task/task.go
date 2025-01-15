package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type (
	Task struct {
		// Id          int
		Description string
		Status      string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
	Tasks []Task
)

const filename = "task.json"

func NewTask(description string, date time.Time) {
	tasks := ParseTasksFile()
	task := Task{
		// Id:          1,
		Description: description,
		Status:      "todo",
		CreatedAt:   date,
		UpdatedAt:   date,
	}
	tasks = append(tasks, task)
	WriteTasks(tasks)
	fmt.Printf("Added task: %s | Time: %s\n", task.Description, task.CreatedAt)
}

func ParseTasksFile() Tasks {
	var tasks Tasks
	content, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		// file doesn't exist return an empty array
		tasks = make([]Task, 0)
		return tasks
	}
	// if err != nil {
	// 	log.Fatalf("Error opening file: %v ", err)
	// }

	err = json.Unmarshal(content, &tasks)
	if err != nil {
		log.Fatal("Error Unmarshalling file")
	}
	return tasks
}

func WriteTasks(tasks Tasks) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("Error Marshalling: %s", err)
	}
	fmt.Println(string(bytes))
	err = os.WriteFile(filename, bytes, 0644)
	if err != nil {
		log.Fatalf("Could not write to file: err")
	}
}
