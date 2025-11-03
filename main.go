package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`
		----- Task Tracker Cli -----
		available command: add, update, delete, mark-in-progress, mark-done
		list, list done, list todo, list in-progress
		----------------------------
		`)
		return
	}

}
