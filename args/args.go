package args

import (
	"Yomu/output"
	"Yomu/reader"
	"Yomu/utils"
	"encoding/json"
	"fmt"
	"os"
)

type Languages map[string] Language

type Language struct {
	Name               string   `json:"name"`
  LineComment       []string `json:"line_comment,omitempty"`
  MultiLineComments []string `json:"multi_line_comments,omitempty"`
  Extensions        []string `json:"extensions,omitempty"`
}

var supportedLanguages Languages

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
		readLanguages()
		os.Exit(0)
	}

	if utils.Includes(args, "-o") || utils.Includes(args, "--output") {
		// readLanguages()
		// os.Exit(0)
	}

	dir := args[0]

	info := reader.Start(dir)
		
	output.OutputBuilder(info)
}


func readLanguages() {

	jsonFile, err := os.ReadFile("languages.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonFile, &supportedLanguages)

	
	if err != nil {
		panic(err)
	}
	
	for _, language := range supportedLanguages {
		fmt.Println(language.Name)
	}
}
