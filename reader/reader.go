package reader

import (
	"os"
	"path/filepath"
	"strings"
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
	"ts":   "Typescript",
	"js":   "Javascript",
	"json": "Json",
	"yaml": "Yaml",
	"go": "Go",
	"mod" : "Go",
	"md": "Markdown",
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
			if includes(markedPath, file.Name()) || strings.HasPrefix(file.Name(), ".") {
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

func includes(pathArray []string, word  string) bool {
	for _, path := range pathArray {
		if path == word {
			return true
		}
	}

	return false
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
