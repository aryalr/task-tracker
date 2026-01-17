package route

import "task-tracker/internal/input"

func Route(args []string) string {
	if len(args) > 2 {
		return "Please input command: add, list, update, atau delete"
	}

	// Command Routing Handler
	command := args[1]
	switch command {
	case "add":
		hasil := input.Add(args)
		return hasil

	case "list":
		hasil := input.List(args)
		return hasil

	case "update":
		hasil := input.Update(args)
		return hasil

	case "delete":
		hasil := input.Delete(args)
		return hasil

	case "help":
		return "Display the mytask help page."

	default:
		return "invalid argument: " + command
	}
}
