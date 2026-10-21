
package task

import "errors"

type Task struct {
    ID          int
    Description string
    Completed   bool
}

var tasks []Task
var nextID = 1

func AddTask(description string) error {
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

    return nil
}

func ListTasks() []Task {
    return tasks
}
