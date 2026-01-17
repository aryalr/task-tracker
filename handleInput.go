package main

func handleAdd(args []string) string {
	if len(args) < 4 {
		return "Usage: mytask add <task-name> <priority>"
	}

	input := args[1]
	taskName := args[2]
	priority := args[3]
	value := "test: " + input + taskName + priority
	return value
}

func handleList(args []string) string {
	input := args[1]
	value := "test: " + input
	return value
}

func handleUpdate(args []string) string {
	if len(args) < 5 {
		return "Usage: mytask update <id> <task-name> <priority>"
	}

	input := args[1]
	id := args[2]
	taskName := args[3]
	priority := args[4]
	value := "test: " + input + id + taskName + priority
	return value
}

func delete(args []string) string {
	if len(args) < 3 {
		return "Usage: mytask delete <id>"
	}

	input := args[1]
	id := args[2]
	value := "test: " + input + id
	return value
}
