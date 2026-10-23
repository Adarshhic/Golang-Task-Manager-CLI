# Go CLI Task Manager

A lightweight command-line task manager built with **Go**, **Cobra**, and **JSON-based persistence**. Create, view, complete, and delete tasks directly from your terminal, with task data preserved between executions.

## Features

- **Task management:** Add, list, complete, and delete tasks.
- **Persistent storage:** Store tasks locally in a JSON file.
- **Input validation:** Reject empty task descriptions and invalid task IDs.
- **CLI interface:** Use Cobra to organize commands and provide built-in help.
- **Automated testing:** Verify core task operations with Go's testing package.

## Tech Stack

| Technology | Purpose |
|---|---|
| Go | Application logic and task management |
| Cobra | Command-line interface and command routing |
| JSON | Local data persistence |
| Go testing | Unit tests for task operations |

## Prerequisites

- [Go](https://go.dev/dl/) installed on your machine.
- Git, if you want to clone the repository.

Verify your Go installation:

```bash
go version
```

## Getting Started

### 1. Clone the repository

Replace the placeholder URL with your repository URL.

```bash
git clone <your-repository-url>
cd Project3_cli
```

Alternatively, open your existing project directory.

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Build the application

```bash
go build -o taskmanager.exe .
```

On Windows, run the executable with:

```powershell
.\taskmanager.exe --help
```

You can also run the application directly using `go run .` without building an executable.

## Usage

### Add a task

```bash
go run . add "Learn Go interfaces"
```

Example output:

```text
Task added successfully: Learn Go interfaces
```

### List all tasks

```bash
go run . list
```

Example output:

```text
1. Learn Go interfaces | Completed: false
2. Build a REST API | Completed: true
```

### Complete a task

Use the task ID to mark a task as completed.

```bash
go run . complete 1
```

### Delete a task

Delete a task using its ID.

```bash
go run . delete 2
```

### View help

Display available commands and their descriptions.

```bash
go run . --help
go run . add --help
go run . complete --help
```

## Project Structure

```text
Project3_cli/
├── main.go
├── go.mod
├── go.sum
├── tasks.json
├── cmd/
│   ├── root.go
│   ├── add.go
│   ├── list.go
│   ├── complete.go
│   └── delete.go
└── internal/
    └── task/
        ├── task.go
        ├── storage.go
        └── task_test.go
```

- `main.go` — Application entry point.
- `cmd/` — Cobra command definitions and CLI setup.
- `internal/task/task.go` — Task model and task operations.
- `internal/task/storage.go` — JSON loading and saving.
- `internal/task/task_test.go` — Unit tests for task operations.
- `tasks.json` — Local task data, generated when tasks are saved.

## Data Storage

Tasks are stored locally in `tasks.json`.

Each task contains three fields:

| Field | Description |
|---|---|
| `ID` | Unique identifier assigned to the task |
| `Description` | Human-readable task description |
| `Completed` | Boolean indicating whether the task is complete |

Example JSON:

```json
[
  {
    "ID": 1,
    "Description": "Learn Go interfaces",
    "Completed": false
  }
]
```

**Note:** Keep `tasks.json` in the application's working directory when using the current storage implementation. Back up this file if you need to preserve your tasks.

## Testing

Run all tests across the project:

```bash
go test ./...
```

Display detailed test output:

```bash
go test -v ./...
```

## Engineering Concepts Demonstrated

This project provides hands-on practice with:

- Structs, slices, functions, and error handling in Go.
- Package organization and Go module imports.
- CLI design using Cobra.
- JSON serialization and deserialization.
- File I/O and persistence across application runs.
- Unit testing and test assertions.
- Rollback logic when persistence operations fail.

## Roadmap

Potential improvements for future versions:

- [ ] Atomic JSON writes to reduce the risk of corrupted data.
- [ ] More comprehensive tests with isolated temporary storage.
- [ ] Task priorities, due dates, and filtering.
- [ ] An update command for editing existing tasks.
- [ ] A storage interface to decouple business logic from persistence.
- [ ] SQLite support for database-backed storage.
- [ ] CI workflows to run tests automatically.

## License

Choose a license before distributing the project publicly. For an open-source project, the [MIT License](https://opensource.org/license/mit/) is one common option.

---

Built as a hands-on project while learning Go and backend engineering.
