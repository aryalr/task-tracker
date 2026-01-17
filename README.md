# 📝 Task Tracker CLI

![Go Version](https://img.shields.io/badge/Go-1.25%2B-00ADD8?style=flat&logo=go)
![Status](https://img.shields.io/badge/Status-Active_Development-orange)

**Task Tracker CLI** is a lightweight, efficient, and intuitive command-line tool built with **Go** to help you manage your daily tasks without leaving your terminal.

> 🚀 **Fast**, **Simple**, and **Persisted** locally in JSON.

---

## ✨ Features

- 📌 **Add Tasks** with priority levels (High, Medium, Low).
- 📋 **List** all your pending and completed tasks.
- ✏️ **Update** task details and priorities on the fly.
- ✅ **Mark Done** to track your progress.
- 🗑️ **Delete** unwanted tasks.
- 💾 **Local Storage** saves your data to `mytask.json`.

---

## 🚀 Getting Started

### Prerequisites

- **Go** (version 1.25 or higher)
- **Make** (optional, for easy build commands)

### 📥 Installation

1.  **Clone the repository**
    ```bash
    git clone https://github.com/yourusername/task-tracker.git
    cd task-tracker
    ```

2.  **Build the application**
    ```bash
    make build
    # Or manually: go build -o build/task-tracker-cli cmd/task-tracker/main.go
    ```

---

## 📖 Usage

You can run the tool directly using `go run` or via the built binary.

### 1. Add a New Task
```bash
# Syntax: add <task_name> <priority: p1|p2|p3>
./build/task-tracker-cli add "Finish the report" p1
```
> **Priorities:** `p1` (High) 🔴, `p2` (Medium) 🟡, `p3` (Low) 🟢

### 2. List All Tasks
```bash
./build/task-tracker-cli list
```

### 3. Update a Task
```bash
# Syntax: update <id> <new_name> <new_priority>
./build/task-tracker-cli update 1 "Finish the FINAL report" p1
```

### 4. Mark Task as Done
```bash
# Syntax: done <id>
./build/task-tracker-cli done 1
```

### 5. Delete a Task
```bash
# Syntax: delete <id>
./build/task-tracker-cli delete 1
```

---

## 📂 Project Structure

```
task-tracker/
├── cmd/
│   └── task-tracker/    # Main application entry point
├── internal/
│   ├── input/           # Input parsing logic
│   ├── model/           # Data structures (Task structs)
│   ├── repo/            # Data persistence (JSON handling)
│   └── route/           # Command routing logic
├── build/               # Compiled binaries
├── mytask.json          # Local database (auto-generated)
└── Makefile             # Build automation
```

## 🛠️ Development

To run the project in development mode:

```bash
make run
# equivalent to: go run cmd/task-tracker/main.go
```

## 🤝 Contributing

Contributions are welcome! Feel free to submit a Pull Request.

---

Made with ❤️ in Go.