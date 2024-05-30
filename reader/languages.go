package reader

import (
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

func ReadLanguages() {

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

func BuildExtensionMap() map[string]string {

	extensionMap :=  make(map[string]string)

	jsonFile, err := os.ReadFile("languages.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(jsonFile, &supportedLanguages)

	if err != nil {
		panic(err)
	}
	
	for _, language := range supportedLanguages {
		for _, extemnsions := range language.Extensions {
			extensionMap[extemnsions] = language.Name
		}
	}

	
	return extensionMap
} 
