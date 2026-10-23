
package task

import (
    "errors"
    "strings"
)

type Task struct {
    ID          int
    Description string
    Completed   bool
}

var tasks []Task
var nextID = 1



func AddTask(description string) error {
	
	description = strings.TrimSpace(description)
	if description == "" {
		return errors.New("task description cannot be empty")
	}

	newTask := Task{
		ID:          nextID,
		Description: description,
		Completed:   false,
	}

	tasks = append(tasks, newTask)
	nextID++


	if err := saveTasksToFile(); err != nil {
	
		tasks = tasks[:len(tasks)-1]
		nextID--
		return err
	}

	return nil
}


func ListTasks() []Task {
    return tasks
}


func CompleteTask(id int) error {
	for i := range tasks {
		if tasks[i].ID == id {
		
			wasCompleted := tasks[i].Completed

			tasks[i].Completed = true

			
			if err := saveTasksToFile(); err != nil {
			
		tasks[i].Completed = wasCompleted
				return err
			}

			return nil
		}
	}

	return errors.New("task not found")
}




func DeleteTask(id int) error {
	for i := range tasks {
		if tasks[i].ID == id {
			
			originalTasks := tasks

			
			tasks = append(tasks[:i], tasks[i+1:]...)

		
			if err := saveTasksToFile(); err != nil {
				
				tasks = originalTasks
				return err
			}

			return nil
		}
	}

	return errors.New("task not found")
}
