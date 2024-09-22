package args

import (
	"Yomu/output"
	"Yomu/reader"
	"Yomu/utils"
	"fmt"
	"os"
)

func HandleArgs() {
	version := "0.0.1"

	args := os.Args[1:]

	if len(args) == 0 {
		panic("Insert at least one argument")
	}

	if utils.Includes(args, "-v") || utils.Includes(args, "--version") {
		fmt.Printf("Yomu %s \n", version)
		os.Exit(0)
	}

	if utils.Includes(args, "-l") || utils.Includes(args, "--languages") {
		reader.ReadLanguages()
		os.Exit(0)
	}

	dir := args[0]

	info := reader.Start(dir)

	if utils.Includes(args, "-o") || utils.Includes(args, "--output") {
		output.JsonOutput(info)
		os.Exit(0)
	}

	output.TableOutputBuilder(info)
}
