
package task

import "testing"

func TestAddTask(t *testing.T) {
	// Reset state for this test
	oldTasks, oldNextID := tasks, nextID
	tasks = nil
	nextID = 1

	defer func() {
		tasks, nextID = oldTasks, oldNextID
	}()

	err := AddTask("Learn Go testing")
	if err != nil {
		t.Fatalf("AddTask() returned an error: %v", err)
	}

	got := ListTasks()

	if len(got) != 1 {
		t.Fatalf("expected 1 task, got %d", len(got))
	}

	if got[0].Description != "Learn Go testing" {
		t.Errorf("expected description %q, got %q",
			"Learn Go testing", got[0].Description)
	}

	if got[0].Completed {
		t.Error("new task should not be completed")
	}
}

func TestCompleteTask(t *testing.T) {
	oldTasks, oldNextID := tasks, nextID
	tasks = []Task{
		{ID: 1, Description: "Learn Go", Completed: false},
	}
	nextID = 2

	defer func() {
		tasks, nextID = oldTasks, oldNextID
	}()

	err := CompleteTask(1)
	if err != nil {
		t.Fatalf("CompleteTask() returned an error: %v", err)
	}

	if !tasks[0].Completed {
		t.Error("expected task to be completed")
	}

	// Test an ID that does not exist
	err = CompleteTask(999)
	if err == nil {
		t.Error("expected an error for a nonexistent task ID")
	}
}

func TestDeleteTask(t *testing.T) {
	oldTasks, oldNextID := tasks, nextID
	tasks = []Task{
		{ID: 1, Description: "Learn Go"},
		{ID: 2, Description: "Build a project"},
	}
	nextID = 3

	defer func() {
		tasks, nextID = oldTasks, oldNextID
	}()

	err := DeleteTask(1)
	if err != nil {
		t.Fatalf("DeleteTask() returned an error: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("expected 1 remaining task, got %d", len(tasks))
	}

	if tasks[0].ID != 2 {
		t.Errorf("expected task ID 2 to remain, got %d", tasks[0].ID)
	}

	// Test an ID that does not exist
	err = DeleteTask(999)
	if err == nil {
		t.Error("expected an error for a nonexistent task ID")
	}
}

func TestAddTaskRejectsEmptyDescription(t *testing.T) {
	for _, description := range []string{"", "   ", "\t\n"} {
		err := AddTask(description)
		if err == nil {
			t.Errorf("expected error for description %q", description)
		}
	}
}
