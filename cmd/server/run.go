package main

import (
	"os"

	"github.com/ProgrammingLab/prolab-accounts/app"
)

func main() {
	os.Exit(run())
}

func run() int {
	err := app.Run()
	if err != nil {
		return 1
	}
	return 0
}
