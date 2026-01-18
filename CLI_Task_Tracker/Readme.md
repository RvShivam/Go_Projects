# CLI Task Tracker

A simple command-line task tracker built in Go using the Cobra library.  
Tasks are stored locally in a JSON file, making the tool lightweight, fast, and easy to use without any external dependencies.

---

## Features

- Add new tasks
- Update task titles
- Delete tasks
- Mark tasks as `todo`, `in-progress`, or `done`
- List all tasks
- Persistent local storage using a JSON file

---

## Requirements

- Go 1.20 or later

---

## Installation

### Option 1: Build locally

Clone the repository:

```bash
git clone https://github.com/RvShivam/Go_Projects.git
cd Go_Projects/CLI_Task_Tracker
```

Build the binary:

```bash
go build -o task
```

Move the binary to a directory in your PATH (Linux / macOS):

```bash
sudo mv task /usr/local/bin/
```

For Windows:

```powershell
Move-Item task.exe $env:USERPROFILE\bin\
```
```powershell
setx PATH "$env:PATH;$env:USERPROFILE\bin"
```

You can now run the CLI from anywhere:

```
task add "My first task"
```

### Option 2: Go install

```bash
go install github.com/RvShivam/Go_Projects/CLI_Task_Tracker@latest
```

Ensure $GOPATH/bin is in your PATH, then run:

```bash
task list
```

---

## Usage

### Add a Task

```
task add "Learn Cobra"
```

### List all Tasks

```
task list
```

Example output:

```
ID   TITLE                     STATUS        CREATED AT
1    Learn Cobra               todo          2026-01-09 14:30
```

### Update a task

```
task update 1 "Learn Cobra deeply"
```

### Delete a task

```
task delete 1
```

### Update task status

```
task updatestatus 1 in-progress
task updatestatus 1 done
task updatestatus 1 todo
```

---

## Data Storage

All tasks are stored locally in a tasks.json file in the project directory.

- No database
- No network calls
- No background services

The file is created automatically when the first task is added.

---

## Project Structure

```
CLI_Task_Tracker/
├── cmd/              # Cobra commands (add, list, update, delete, mark)
├── internal/         # Task model and JSON storage logic
├── main.go           # Application entry point
├── tasks.json        # Runtime task data (not committed to Git)
└── README.md
```

---
