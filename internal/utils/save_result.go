package utils

import (
	"encoding/json"
	"os"
)

func SaveResult(file string, result interface{}) {
	content, _ := json.MarshalIndent(result, "", "  ")
	err := os.WriteFile(file, append(content, '\n'), 0644)
	if err != nil {
		panic(err)
	}
}
