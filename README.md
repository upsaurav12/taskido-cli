# Taskido-CLI - Task Management CLI Tool

Taskido-CLI is a simple command-line tool to help you manage your tasks efficiently. With Taskido-CLI, you can add, list, complete, and delete tasks from your terminal while keeping track of important deadlines, priorities, and more. It’s built in **Go** and uses **SQLite** for local storage, but can be extended to the cloud for syncing tasks across devices.

---

## Features

- **Task Management**  
  - Add tasks with a title and due date
  - List all tasks and filter by status, priority, and due date
  - Mark tasks as completed
  - Delete tasks

- **Priorities & Due Dates**  
  - Assign priority levels to tasks (low, medium, high)
  - Set due dates and track tasks that are overdue

- **Search and Filter**  
  - Search tasks by keywords
  - Filter tasks by completion status or priority

- **Data Storage**  
  - Uses **SQLite** for storing tasks locally on your machine (easy to backup and portable)

- **Backup & Restore**  
  - Export tasks to a JSON file and restore tasks from backup

- **Cloud Sync (Optional)**  
  - Sync tasks between devices with a backend such as Appwrite (coming soon)

---

## Installation

### Prerequisites

- [Go](https://golang.org/dl/) (version 1.18 or higher)
- SQLite (for local storage)

### Install via `go install`

You can install **Taskido-CLI** directly from GitHub by running the following command:

```bash
go install github.com/upsaurav12/taskido-cli@latest
````

## Usage

### Add a New Task

To add a new task, simply use the `add` command:

```bash
task add "Finish writing blog post"
```


