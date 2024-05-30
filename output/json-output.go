package output

import (
	"Yomu/reader"
	"encoding/json"
	"os"
)

type Jsonoutput = struct {
	Files int
	TotalLines int
	CommentedLines int
	CodeLines int
	BlankLines int
}

func JsonOutput (information []reader.CodeInformation){

	jsonoutput := make(map[string]Jsonoutput)

	for _, language := range information {
		jsonOutputEntry := Jsonoutput{ Files: language.Files, 
			TotalLines: language.TotalLines, 
			CommentedLines: language.CommentedLines, 
			CodeLines: language.CodeLines,
			BlankLines: language.BlankLines,
		}

		jsonoutput[language.Extension] = jsonOutputEntry
	}

	jsonBytes, err := json.Marshal(jsonoutput)

	if err != nil {
		panic(err)
	}

	writeErr := os.WriteFile("Yomu-output.json", jsonBytes, 0644)

	if writeErr != nil {
		panic(writeErr)
	}
}
