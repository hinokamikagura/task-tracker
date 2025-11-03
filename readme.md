Absolutely! Here’s a **professional README** template for your Go CLI task tracker project. I’ve written it in a way that looks polished, structured, and developer-friendly. You can adjust the details (like GitHub username, project path) as needed.

---

# Task Tracker CLI

A simple and professional **CLI task manager** written in Go.
Supports adding, updating, deleting, and listing tasks stored in a JSON file.

---

## Table of Contents

* [Features](#features)
* [Installation](#installation)
* [Usage](#usage)
* [Commands](#commands)
* [JSON Storage](#json-storage)
* [Project Structure](#project-structure)

---

## Features

* Add tasks with title, description, priority, and due date.
* Update task titles and descriptions.
* Delete tasks by ID.
* List all tasks in **pretty JSON format**.
* No external dependencies — uses **Go standard library only**.
* Portable and lightweight — runs on Windows, macOS, Linux.

---

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/hinokamikagura/task-tracker.git
cd task-tracker
```

### 2. Build the CLI

```bash
go build -o task-cli main.go
```

This will create an executable `task-cli` (or `task-cli.exe` on Windows).

---

## Usage

Run commands in the terminal:

```bash
./task-cli <command> [flags]
```

Example:

```bash
./task-cli add --title="Buy groceries" --desc="Cook dinner" 
```

---

## Commands

| Command            | Flags                                       | Description                            |
| -------------------|-------------------------------------------- | -------------------------------------- |
| `add`              | `--title` `--desc`                          | Add a new task                         |
| `update`           | `--taskId` `--title` `[--desc]`             | Update task by ID                      |
| `delete`           | `--taskId`                                  | Delete task by ID                      |
| `mark-in-progress` | `--taskId`                                  | Mark in progress by ID                 |
| `mark-done`        | `--taskId`                                  | Mark done by ID                        |
| `list`             | -                                           | List all tasks in JSON format          |
| `list done`        | -                                           | List done tasks in JSON format         |
| `list todo`        | -                                           | List todo tasks in JSON format         |
| `list in-progress` | -                                           | List in progress tasks in JSON format  |

### Example: Add a task

```bash
./task-cli add --title="task-tracker" --desc="CLI task manager" 
```

### Example: Update a task

```bash
./task-cli update --taskId=1 --title="Updated task title"
```

### Example: Delete a task

```bash
./task-cli delete --taskId=1
```

### Example: List all tasks

```bash
./task-cli list
```

Output:

```json
{
    "Id": 4,
    "Title": "task-tracker",
    "Description": "This is the updated version of add task command",
    "Status": "todo",
    "CreatedAt": "2025-11-03T12:08:36",
    "UpdatedAt": "2025-11-03T12:08:36"
}
```

---

## JSON Storage

Tasks are stored in a local JSON file:

```
db/task.json
```

* Automatically created if not present.
* Handles empty files safely.
* Each task has the following fields:

```json
{
    "Id": 1,
    "Title": "Task title",
    "Description": "Task description",
    "Status": "todo",
    "CreatedAt": "YYYY-MM-DDTHH:MM:SS",
    "UpdatedAt": "YYYY-MM-DDTHH:MM:SS"
}
```

---

## Project Structure

```
task-tracker/
├── main.go          # CLI entry point
├── command/         # Command handlers
│   ├── addTask.go
│   ├── updateTask.go
│   ├── deleteTask.go
│   ├── listTask.go
│   ├── command.go
│   ├── helper.go
│   ├── markDone.go
│   └── markProgress.go
├── schemas/         # Struct definitions
│   └── task.go
├── db/              # JSON storage
│   └── task.json
└── config/          # Logger and utilities
    └── config.go
    └── database.go
    └── logger.go
```

---

### Optional Enhancements

* Add **interactive mode** (REPL) in future.
* Add **search/filter by status or priority**.
* Add **command history** for faster input.
* Support **sorting tasks by due date**.
