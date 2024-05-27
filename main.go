package main

import (
	"Yomu/args"
	"Yomu/output"
	"Yomu/reader"
)

func main() {
	dir := args.HandleArgs()

	info := reader.Start(dir)
		
	output.OutputBuilder(info)
}
