package main

import (
	"os"
	"tokei/output"
	"tokei/reader"
)

func main() {
	dir := handleArgs()

	info:= reader.Start(dir)
	
	output.OutputBuilder(info)
}

func handleArgs() string {
	args := os.Args[1:]

	if len(args) == 0 {
		panic("Insert the directory you want to read")
	}

	dir := args[0]

	return dir
}
