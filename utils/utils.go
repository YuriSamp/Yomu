package utils

func Includes(pathArray []string, word  string) bool {
	for _, path := range pathArray {
		if path == word {
			return true
		}
	}

	return false
}
