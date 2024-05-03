package main

import (
	"os"

	"github.intuit.com/gdanko/netspeed2/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	return
}
