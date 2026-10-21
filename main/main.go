
package main

import (
    "fmt"

    "Project3_cli/Internal/Task"
)

func main() {
    err := task.AddTask("Learn Go")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    err = task.AddTask("Build CLI")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("My tasks:")

    for _, t := range task.ListTasks() {
        fmt.Printf("%d. %s | Completed: %t\n",
            t.ID, t.Description, t.Completed)
    }
}
