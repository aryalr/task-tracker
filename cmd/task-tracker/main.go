package main

import (
	"fmt"
	"os"
	"task-tracker/internal/route"
)

func main() {
	hasil := route.Route(os.Args)
	fmt.Println(hasil)
}
