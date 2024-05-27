package reader

import (
	"os"
	"path/filepath"
	"strings"
	"Yomu/utils"
)

type CodeInformation struct {
	Extension string
	Files int
	TotalLines int
	CommentedLines int
	CodeLines int
	BlankLines int
}

var extensionMap = map[string]string{
	"ts": "Typescript",
	"tsx": "Typescript",
	"js":  "Javascript",
	"cjs": "Javascript",
	"mjs": "Javascript",
	"json": "Json",
	"yaml": "Yaml",
	"yml" : "Yaml",
	"go": "Go",
	"mod" : "Go",
	"sum" : "Go",
	"md": "Markdown",
	"css" : "CSS",
	"svg" : "SVG",
	"html" : "HTML",
	"Dockerfile" : "Dockerfile", // handle unique name file ?
}

func Start(dir  string) []CodeInformation {
	EmptyCodeInformationList := []CodeInformation{}
	markedPath := []string{}

	CodeInformationList, _ := readCurrDir(dir, EmptyCodeInformationList, markedPath)

	return reduceInformation(CodeInformationList)
}

func readCurrDir(dir string, CodeInformationList []CodeInformation, markedPath []string) ([]CodeInformation,[]string ) {
		files, err := os.ReadDir(dir)

		if err != nil {
			panic(err)
		}

		for _ ,file := range files {
			if isIgnoredFile(markedPath, file.Name()) {
				continue
			}

			filePath := filepath.Join(dir, file.Name())

			if file.IsDir(){
				markedPath = append(markedPath, dir)
				CodeInformationList, markedPath = readCurrDir(filePath, CodeInformationList, markedPath)
			}

			if !file.IsDir() {
				info := processFile(filePath)
        CodeInformationList = append(CodeInformationList, info)
			}
		}

		markedPath = append(markedPath, dir)
		return CodeInformationList, markedPath
}

func reduceInformation(informationList []CodeInformation) []CodeInformation {

	patchedInfo := []CodeInformation{}
	summedData := make(map[string]CodeInformation)

	
	for _, entry := range informationList {
		if val, ok := summedData[entry.Extension]; ok {
			val.TotalLines += entry.TotalLines
			val.Files += entry.Files
			val.CommentedLines += entry.CommentedLines
			val.CodeLines += entry.CodeLines
			val.BlankLines += entry.BlankLines
			summedData[entry.Extension] = val
			} else {
				summedData[entry.Extension] = entry
			}
		}
		
		for _, info := range summedData {
			patchedInfo = append(patchedInfo, info)
		}
		
		return patchedInfo
}

func isIgnoredFile(markedPath []string, word  string) bool {

	ignoredFiles := []string{"node_modules", "dist"}

	if utils.Includes(markedPath, word) || strings.HasPrefix(word, ".") || utils.Includes(ignoredFiles, word) {
		return true
	}

	return false
}
