package main

import (
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
)

func main() {
	// create an app
	app := cli.App("cp", "Copy files around")

	// Here's what differentiates mow.cli from other CLI libraries:
	// This line is not just for help message generation.
	// It also validates the call to reject calls with less than 2 arguments
	// and split the arguments between SRC or DST
	app.Spec = "[-r] SRC... DST"

	var (
		// declare the -r flag as a boolean flag
		recursive = app.BoolOpt("r recursive", false, "Copy files recursively")
		// declare the SRC argument as a multi-string argument
		src = app.StringsArg("SRC", nil, "Source files to copy")
		// declare the DST argument as a single string (string slice) arguments
		dst = app.StringArg("DST", "", "Destination where to copy files to")
	)

	// Specify the action to execute when the app is invoked correctly
	app.Action = func() {
		fmt.Printf("Copying %v to %s [recursively: %v]\n", *src, *dst, *recursive)
	}

	// Invoke the app passing in os.Args
	app.Run(os.Args)
}
