package utils

import (
	"encoding/json"
	"os"
)

func SaveResult(file string, result interface{}) {
	content, _ := json.MarshalIndent(result, "", "  ")
	os.WriteFile(file, content, 0644)
}
