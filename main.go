package main

import (
	"fmt"
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

	fmt.Println(args)

	dir := args[0]

	return dir
}
