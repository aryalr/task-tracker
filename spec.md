# Task-Tracker CLI Specification

A lightweight command-line interface for managing tasks efficiently, built with Go.

## Overview

Task-Tracker CLI (`mytask`) allows users to manage their daily tasks directly from the terminal. It supports adding, updating, deleting, listing, and marking tasks as done. Data is persisted locally in a JSON file.

## Getting Started

### Prerequisites

*   Go (version 1.20+)
*   Make (optional, for build automation)

### Installation & Build

1.  Clone the repository:
    ```bash
    git clone <repository-url>
    cd task-tracker
    ```

2.  Build the binary:
    ```bash
    make build
    ```
    The binary will be created in `build/task-tracker-cli`.

### Running

You can run the application directly using Go:

```bash
go run cmd/task-tracker/main.go [command]
```

Or execute the built binary:

```bash
./build/task-tracker-cli [command]
```

## Command Reference

The general syntax is:

```bash
mytask <command> [arguments]
```

### 1. Add Task

Create a new task with a name and priority level.

*   **Usage:** `add <task_name> <priority>`
*   **Priorities:** `p1` (High), `p2` (Medium), `p3` (Low)

**Example:**
```bash
mytask add "create user input" p1
```
**Output:**
`successfully added task "create user input" id => 1 | priority => high`

### 2. List Tasks

Display all current tasks.

*   **Usage:** `list`

**Example:**
```bash
mytask list
```
**Output:**
```text
task list:

id | Name                  | Priority
---|-----------------------|----------
1  | add update user input | Medium
2  | create database port  | High
```

### 3. Update Task

Modify an existing task's name or priority.

*   **Usage:** `update <id> <new_name> <new_priority>`

**Example:**
```bash
mytask update 1 "add update user input" p2
```
**Output:**
`successfully update task "add update user input" id => 1 | priority => medium`

### 4. Mark as Done

Mark a task as completed.

*   **Usage:** `done <id>`

**Example:**
```bash
mytask done 1
```
**Output:**
`task marked done id => 1`

### 5. Delete Task

Remove a task permanently.

*   **Usage:** `delete <id>`

**Example:**
```bash
mytask delete 1
```
**Output:**
`successfully delete task "add update user input" id => 1`

## Data Model

The application stores tasks in a JSON file named `mytask.json` created in the current working directory.

### Priority Mapping

| Code | Level  |
| :--- | :----- |
| `p1` | High   |
| `p2` | Medium |
| `p3` | Low    |

## Contribution Guidelines

We welcome contributions! Please follow these steps to contribute:

1.  **Fork** the repository.
2.  **Clone** your fork locally.
3.  **Create a branch** for your feature or bug fix (`git checkout -b feature/amazing-feature`).
4.  **Implement** your changes.
5.  **Test** your changes to ensure functionality.
6.  **Commit** your changes with clear messages.
7.  **Push** to your branch and submit a **Pull Request**.

### Code Structure

*   `cmd/task-tracker/`: Entry point of the application (`main.go`).
*   `internal/input/`: Handles CLI input parsing.
*   `internal/route/`: Routing logic for commands.
*   `internal/model/`: Defines data structures (e.g., `Task`).
*   `internal/repo/`: Handles data storage and retrieval (JSON file operations).