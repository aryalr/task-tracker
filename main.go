package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input command: add, list, update, atau delete")
		return
	}

	// Input Routing
	command := os.Args[1]
	switch command {
	case "add":
		hasil := handleAdd(os.Args)
		fmt.Println(hasil)
	case "list":
		hasil := handleList(os.Args)
		fmt.Println(hasil)
	}
}
