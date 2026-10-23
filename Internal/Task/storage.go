
package task

import (
	"encoding/json"
	"os"
)

const storageFile = "tasks.json"

func saveTasksToFile() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(storageFile, data, 0644)
}
func loadTasksFromFile() error {
	data, err := os.ReadFile(storageFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil 
		}
		return err
	}

	var loadedTasks []Task
	if err := json.Unmarshal(data, &loadedTasks); err != nil {
		return err
	}

	tasks = loadedTasks

	
	nextID = 1
	for _, t := range tasks {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}

	return nil
}


func LoadTasks() error {
    return loadTasksFromFile()
}