package main

import (
	"os"

	"github.intuit.com/gdanko/netspeed/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	return
}
